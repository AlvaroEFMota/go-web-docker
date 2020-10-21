package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

//GetConexao retorna uma conexão ativa com o banco de dados
func GetConexao() *sql.DB {
		
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/go_web_docker")
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println("\033[32m"+"Connected to database!"+"\033[0m")
	}
	return db
}