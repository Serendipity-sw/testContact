package main

import (
	"fmt"
	"github.com/guotie/deferinit"
)

func main() {
	deferinit.AddInit(runOne,runTwo,1)
	deferinit.AddInit(runThree,runFour,2)
	deferinit.InitAll()
	deferinit.FiniAll()
}

func runOne() {
	fmt.Println(1)
}
func runTwo() {
	fmt.Println(2)
}
func runThree() {
	fmt.Println(3)
}
func runFour() {
	fmt.Println(4)
}
