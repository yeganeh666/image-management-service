package config

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	DefaultConfig = []byte(`
http:
  address: "localhost"
  port: "8080"
postgres:
  connection:
    file_path: "./config/postgres.secret"
  host: "localhost"
  port: 5432
  db: "images"
  user: "postgres"
  pass: "postgres"
image:
  source_path: "./data/links.txt"
  upload_path: "./images/user-content"
  download_path: "./images/"
`)
)

type Config struct {
	serviceName string
	HTTP        SectionHttp     `yaml:"http"`
	Core        SectionCore     `yaml:"core"`
	Log         SectionLog      `yaml:"log"`
	Image       SectionImage    `yaml:"postgres"`
	Postgres    SectionPostgres `yaml:"postgres"`
}

type SectionImage struct {
	UploadPath   string `mapstructure:"upload_path"`
	DownloadPath string `mapstructure:"download_path"`
	SourcePath   string `mapstructure:"source_path"`
}

type SectionCore struct {
	ServiceName string `mapstructure:"service_name"`
	Mode        string `yaml:"mode"`
}

type SectionPostgres struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
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

func LoadConfigs(defaultConfig bool) (*Config, error) {
	log.Infof("reding configs...")

	if defaultConfig {
		viper.SetConfigType("yaml")
		log.Infof("reading deafult configs")
		err := viper.ReadConfig(bytes.NewBuffer(DefaultConfig))
		if err != nil {
			log.WithError(err).Error("read from default configs failed")
			return nil, err
		}
	} else {
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Warnf("failed to read from env: %v", err)
			viper.AddConfigPath("./configs")  //path for docker compose configs
			viper.AddConfigPath("../configs") //path for local configs
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
			if err = viper.ReadInConfig(); err != nil {
				log.Warnf("failed to read from yaml: %v", err)
				localErr := viper.ReadConfig(bytes.NewBuffer(DefaultConfig))
				if localErr != nil {
					log.WithError(localErr).Error("read from default configs failed")
					return nil, localErr
				}
			}
		}
	}

	conf := new(Config)
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return conf, nil
}
