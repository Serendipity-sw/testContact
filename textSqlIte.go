package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./bcy.s3db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stmt, err := db.Prepare("insert into bcyTable(urlPath,zanNumber) values(?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res, err := stmt.Exec("http://www.baidu.com", 33)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(id)

	rows, err := db.Query("select urlPath,zanNumber from bcyTable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	type bcyTable struct {
		UrlPath   string
		ZanNumber int
	}
	var (
		modelList []bcyTable
		urlPath   string
		zanNumber int
	)
	for rows.Next() {
		err = rows.Scan(&urlPath, &zanNumber)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		modelList = append(modelList, bcyTable{
			UrlPath:   urlPath,
			ZanNumber: zanNumber,
		})
	}
	fmt.Println(modelList)
}
