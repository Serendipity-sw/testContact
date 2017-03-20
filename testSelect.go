package main

import (
	"fmt"
	//"time"
)

func main() {
	ch1 := make(chan bool, 1)
	//ch2 := make (chan int, 1)

	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	//case <-ch2:
	//	fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}

	//var outLine =make(chan bool,1)
	//go func() {
	//	time.Sleep(1*time.Second)
	//}()
	//select {
	//case <-outLine:
	//	fmt.Println(2)
	//	break
	//}
	//fmt.Println(1)
}
