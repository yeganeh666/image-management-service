package gormext

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Configs struct {
	PostgresConnection string `env:"POSTGRES_CONNECTION,required,file"`
}

func Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: false})
}
