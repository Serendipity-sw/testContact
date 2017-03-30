package main

import (
	"database/sql"
	"fmt"
	"github.com/swgloomy/go-common"
)

func main() {
	marketDBModel := common.MySqlDBStruct{
		DbUser: "axon",
		DbHost: "10.10.141.79",
		DbPort: 3306,
		DbPass: "Axon.2017",
		DbName: "market",
	}
	dbs := common.MySqlSQlConntion(marketDBModel)
	result, err := common.MySqlSelect(dbs, marketDBModel, "select * from conf_syncinfo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	columnStr, err := result.Columns()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(columnStr)
	values := make([]sql.RawBytes, len(columnStr))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for result.Next() {
		err = result.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columnStr[i], ": ", value)
		}
	}
	common.MySqlClose(dbs)
}
