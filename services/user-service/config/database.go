package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//InitDatabase is function to create connection to Database
func (cf *Config) initDatabase() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/practice")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected successfully..")
	}

	cf.DB = db

}
