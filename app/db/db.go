package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func New() *sql.DB {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		DBName:    "moneyhook",
		User:      "moneyhook",
		Passwd:    "password",
		Addr:      "sql:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
