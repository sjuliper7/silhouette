package config

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
)

//Config is struct to access all config
type Config struct {
	DB *sql.DB
}

//LoadConfig is function to load config
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Env has been loaded successfully..")
	}

}

//InitConfig is function to init all config
func InitConfig() {

	var cf Config
	// cf.initDatabase()
}
