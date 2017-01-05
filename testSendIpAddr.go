package main

import (
	"net/http"
	"fmt"
)

func main() {
	_,err:=http.Post("http://10.10.136.18:8754/nihao","",nil)
	if err != nil {
		fmt.Println("http run err! err: %s \n",err.Error())
		return
	}
	fmt.Println("http success!")
}
