package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("1")
	}()
	fmt.Println("nihao")
	defer func() {
		fmt.Println("2")
	}()
}
