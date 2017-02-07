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

	model=append(model,&nameObj{
		Name:"nihao",
	})
	fmt.Println(model[0].Name)


	jsonasdf,_=json.Marshal(model)
	fmt.Println(string(jsonasdf))
}
