package main

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func main() {
	c, err := ftp.DialTimeout("10.10.141.71:21", 5*time.Second)
	if err != nil {
		fmt.Println(1, " ", err.Error())
		return
	}
	err = c.Login("shaowei", "shaowei@123")
	if err != nil {
		fmt.Println(2, " ", err.Error())
		return
	}
	err = c.NoOp()
	if err != nil {
		fmt.Println(3, " ", err.Error())
		return
	}

	nameList, err := c.NameList("")
	if err != nil {
		fmt.Println(12)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(nameList)
	return

	data := bytes.NewBufferString("nihaobuhaoa")
	err = c.Stor("gloomyswtestFile13", data)
	if err != nil {
		fmt.Println("4", " ", err.Error())
		return
	}
	fmt.Println("successfully")
	err = c.Stor("shaowei23", bytes.NewBufferString("fsasdfsadf"))
	if err != nil {
		fmt.Println("5", " ", err.Error())
		return
	}
	fmt.Println("123123fullay")
}
