package main

import "fmt"

func main() {
	var array []string
	array = append(array, "123")
	array = append(array, "1234")
	array = append(array, "1235")
	fmt.Println(array[:2])
}
