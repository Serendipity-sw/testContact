package main

import (
	"fmt"
	"os"
)

func main() {
	fs := os.Getpid()
	fmt.Println(fs)
	ff, err := os.FindProcess(fs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ff.Kill())
}
