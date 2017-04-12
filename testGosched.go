package main

import (
	"fmt"
	"os"
	"time"
	//"log"
	"sync"
)

func fib(x int) int {
	if x < 2 {
		fmt.Println(3)
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

var tokens = make(chan struct{}, 20)

func main() {
	var (
		threadLock sync.WaitGroup
		ch         = make(chan []string, 3)
	)
	threadLock.Add(3)
	//threadLock.Add(3)
	go func(ch chan []string) {
		time.Sleep(3 * time.Second)
		ch <- []string{"1", "3"}
		threadLock.Done()
		//threadLock.Done()
	}(ch)
	go func(ch chan []string) {
		time.Sleep(12 * time.Second)
		ch <- []string{"4", "3"}
		threadLock.Done()
		//threadLock.Done()
	}(ch)
	go func(ch chan []string) {
		time.Sleep(7 * time.Second)
		ch <- []string{"7", "3"}
		threadLock.Done()
		//threadLock.Done()
	}(ch)
	threadLock.Wait()
	close(ch)
	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println(55)

	return
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		fmt.Println(list)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- []string{"1"}
				}(link)
			}
		}
	}
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) string {
	return hostname
}
