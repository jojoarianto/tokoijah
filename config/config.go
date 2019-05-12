package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Config object config for database connection
type Config struct {
	DB *DBConfig
}

// DBConfig database gorm setup
type DBConfig struct {
	Dialeg string
	DBUri  string
}

// NewConfig method to set new db config
func NewConfig(dial string, uri string) *Config {

	return &Config{
		DB: &DBConfig{
			Dialeg: dial,
			DBUri:  uri,
		},
	}
}

// ConnectDB method to connect db with config
func (conf *Config) ConnectDB() (*gorm.DB, error) {

	db, err := gorm.Open(conf.DB.Dialeg, conf.DB.DBUri)
	if err != nil {
		log.Fatal("Fail to connect database", err.Error())
		return nil, err
	}

	return db, nil
}
