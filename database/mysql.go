package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123@abc@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("Chuck", "技术中心", "2018-04-27")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET created=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("2018-04-28", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	fmt.Println("Delete affect ", affect)

	// 关闭连接
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
