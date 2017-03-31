package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	blockObj.Add(1)
	go threadDownasdf(&blockObj)
	go threadDownasdf(&blockObj)
	time.Sleep(10 * time.Second)
	blockObj.Done()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	//信号等待
	<-c
}

func threadDownasdf(wg *sync.WaitGroup) {
	jsTmr := time.NewTimer(1 * time.Minute)
	go func() {
		wg.Wait()
		fmt.Println(5)
		jsTmr.Stop()
	}()
	for {
		fmt.Println(4)
		jsTmr.Reset(3 * time.Second)
		<-jsTmr.C
	}
}
