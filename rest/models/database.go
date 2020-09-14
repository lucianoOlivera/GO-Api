package models

import (
	"database/sql"
	"fmt"
	"log"

	"../config"

	_ "github.com/go-sql-driver/mysql"
)

//<username>:<password>@tcp(<host>:<port>)/<database>

var db *sql.DB

func CreatConnection() {
	url := config.GetUrl()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		fmt.Println("conecion exitosa ")
		db = connection
	}

}
func CloseConnection() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func existsTables(tableName string) bool {
	sql := fmt.Sprintf("show tables like '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func creatTable(tableName, schema string) {
	if !existsTables(tableName) {
		Exec(schema)
	} else {
		TruncateTable(tableName)
	}
}

func CreaTables() {
	creatTable("users", userShema)
}

//exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Panicln(err)
	}
	return result, err

}
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("truncate %s", tableName)
	Exec(sql)
}
