package configext

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

type FileConfig struct {
	FilePath string `mapstructure:"file_path"`
}

func (s *FileConfig) GetValue() string {
	apiKey, err := os.ReadFile(s.FilePath)
	if err != nil {
		panic(fmt.Sprintf("Error to read file in path %v with error: %v", s.FilePath, err))
	}
	return strings.TrimSpace(string(apiKey))
}

type Config struct {
	internalDefaultConf []byte
	configPath          string
	serviceName         string

	HTTP     SectionHttp     `yaml:"http"`
	Core     SectionCore     `yaml:"core"`
	Postgres SectionPostgres `yaml:"postgres"`
	Log      SectionLog      `yaml:"log"`
}

type SectionCore struct {
	ServiceName string `mapstructure:"service_name"`
	Mode        string `yaml:"mode"`
}

type SectionPostgres struct {
	Connection FileConfig `yaml:"connection"`
	Host       string     `yaml:"host"`
	Port       int        `yaml:"port"`
	DB         string     `yaml:"db"`
	User       string     `yaml:"user"`
	Pass       string     `yaml:"pass"`
}

type SectionLog struct {
	Level   string `yaml:"level"`
	LogFile string `mapstructure:"log_file"`
	LogSize int    `mapstructure:"log_size"` // megabytes
	LogAge  int    `mapstructure:"log_age"`  // days
}
type SectionHttp struct {
	Address             string `yaml:"address"`
	Port                uint16 `yaml:"port"`
	MultipleAccountAuth bool   `mapstructure:"multi_account_auth"`
	User                string `yaml:"user"`
	Pass                string `yaml:"pass"`
}

func (c *Config) configureViper() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()                                            // read in environment variables that match
	viper.SetEnvPrefix(strings.ReplaceAll(c.serviceName, "-", "_")) // will be uppercase automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath("/etc/" + c.serviceName + "/")
	viper.AddConfigPath("$HOME/." + c.serviceName)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
}

func (c *Config) loadConf() error {
	c.configureViper()

	if err := c.readConf(); err != nil {
		return err
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("unable to decode into config struct, ", err)
		return err
	}

	return nil
}

func (c *Config) readConfFromFile() error {
	if c.configPath != "" {
		content, err := ioutil.ReadFile(c.configPath)
		if err != nil {
			log.Errorf("File does not exist : %s", c.configPath)
			return err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return err
		}
	} else {
		if err := viper.MergeInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			fmt.Println("Config file not found.")
		}
	}
	return nil
}

func (c *Config) readConf() error {
	// load default config
	if err := viper.ReadConfig(bytes.NewBuffer(c.internalDefaultConf)); err != nil {
		return err
	}
	if err := c.readConfFromFile(); err != nil {
		return err
	}
	return nil
}

func NewConfig(path string, serviceName string, internalDefaultConfig []byte) *Config {
	conf := Config{
		internalDefaultConf: internalDefaultConfig,
		configPath:          path,
		serviceName:         serviceName,
	}
	err := conf.loadConf()
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return nil
	}
	return &conf
}
