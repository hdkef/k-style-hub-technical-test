package driver

import (
	"database/sql"
	"fmt"
	"kstylehub/app/config"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(config *config.AppConfig) *sql.DB {
	//make a connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USERNAME, config.DB_PASS,
		config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	//ping a connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
