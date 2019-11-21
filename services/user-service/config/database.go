package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//InitDatabase is function to create connection to Database
func (cf *Config) initDatabase() {
	db, err := sqlx.Connect("mysql", populateStringConnection())

	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Println("Database connected successfully..")
	}

	cf.DB = db

}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME")
	stringConnection += fmt.Sprintf("?%s", "parseTime=true")

	return stringConnection
}
