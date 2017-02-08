package main

import (
	"fmt"
	"encoding/json"
)

type nameObj struct {
	Name string
}

func main() {
	var model []*nameObj=make([]*nameObj,10)

	jsonasdf,_:=json.Marshal(model)
	fmt.Println(string(jsonasdf))

	jsonasdf,_=json.Marshal(model)
	fmt.Println(string(jsonasdf))
}
