package app

import (
	"database/sql"
	"time"
)

func ConnectDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/hocgolang?parseTime=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return DB
}
