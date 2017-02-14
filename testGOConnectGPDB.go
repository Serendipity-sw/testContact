package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"axon"
)

func main() {
	db, err := sql.Open("postgres", "postgres://gpadmin:gpadmin@10.10.141.80:5432/pg_testdb?sslmode=disable")

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select phone from tag.bssofday_test limit 1")
	var phone string
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err=rows.Scan(&phone)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	rows.Close()
	phone=axon.AxonDecrypt(phone)
	fmt.Println(phone)
}