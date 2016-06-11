package main

import (
	"database/sql"

	"fmt"

	_ "github.com/ziutek/mymysql/godrv"
)

const (
	DB_NAME = "test"
	DB_USER = "mysql"
	DB_PASS = ""
)

type User struct {
	id int `json:"id"`

	name string `json:"name"`

	sex int `json:"sex"`

	degree float64 `json:"degree"`
}

func OpenDB() *sql.DB {

	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))

	if err != nil {
		fmt.Println("err1:", err)
		panic(err)

	}

	return db

}

func UserById(id int) {

	db := OpenDB()
	defer db.Close()

	row, err := db.Query("select * from MyClass")
	if err != nil {
		fmt.Println("err2:", err)
		panic(err)

	}

	for row.Next() {
		user := new(User)
		row.Scan(&user.id, &user.name, &user.sex, &user.degree)
		fmt.Println(user)
	}

	// return user

}

func main() {
	fmt.Println("Hello world")
	UserById(0)
	// fmt.Println(user)
}
