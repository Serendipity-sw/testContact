package main

import (
	//"sync"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go func() {
		for {
			fmt.Println(3)
			return
		}
		fmt.Println(4)
	}()

	fmt.Println(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	//信号等待
	<-c
}
