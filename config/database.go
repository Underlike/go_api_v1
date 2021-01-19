package config

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func InitializeDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/url-rewrite")

    if err != nil {
        panic(err.Error())
	}
	
	fmt.Printf(InfoColor, "Initialization of database - OK \n")
	return db
}