package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var blockObj sync.WaitGroup
	blockObj.Add(1)
	go func() {
		time.Sleep(time.Second * 10)
		blockObj.Done()
	}()
	blockObj.Wait()
	fmt.Println(123)
}
