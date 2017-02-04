package main

import (
	"github.com/gobwas/glob"
	"fmt"
)

func main() {
	var g glob.Glob
	g = glob.MustCompile("*.serv.*.cn")
	fmt.Println(g.Match("*.serv.wostore.*"))
}
