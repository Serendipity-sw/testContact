package main

import "fmt"

func main() {
	var (
		haveValueMap map[string]string=make(map[string]string)
		noValueMap map[string]string=make(map[string]string)
	)
	haveValueMap["1"]="2"
	fmt.Println(len(haveValueMap))
	fmt.Println(len(noValueMap))
}
