package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

//Config is struct to access all config
type Config struct {
	DB *sqlx.DB
}

//LoadConfig is function to load config
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file")
	} else {
		logrus.Info("env has been loaded successfully..")
	}

}

//InitConfig is function to init all config
func InitConfig() {

	var cf Config
	cf.initDatabase()
	cf.initService()
}
