package main

import "fmt"

func main() {
	fileStr:=[]string{"1","2","3"}
	for _, value := range fileStr {
		fmt.Println(value)
		return
	}
	fmt.Println("2")
}
