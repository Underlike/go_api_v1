package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/url-rewrite")

    if err != nil {
        panic(err.Error())
	}
	return db
}