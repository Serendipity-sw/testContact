package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lockWait sync.WaitGroup

	lockWait.Add(1)
	go func() {
		fmt.Println(321)
		time.Sleep(3 * time.Second)
		lockWait.Done()
	}()
	lockWait.Wait()
	lockWait.Wait()
	fmt.Println(123)
}
