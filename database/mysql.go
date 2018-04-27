package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root@/test?charset=utf8")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
