package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(1)
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return errors.New("nihao")
	})
	fmt.Println(2)
}
