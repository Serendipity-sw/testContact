package main

import "fmt"

type tableMap struct {
	Mob string
	Sex string
}


func main() {
	var (
		modelList map[string]tableMap=make(map[string]tableMap)
		itemList map[string]tableMap=make(map[string]tableMap)
	)
	modelList["邵炜"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	modelList["邵炜1"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	modelList["邵炜2"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	modelList["邵炜3"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}

	itemList["1邵炜"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	itemList["2邵炜1"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	itemList["3邵炜2"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
	itemList["4邵炜3"]=tableMap{
		Mob:"",
		Sex:"asdf",
	}
fmt.Println(len(itemList))
	updateMap(&itemList)
	fmt.Println(len(itemList))
	fmt.Println(itemList["123"].Mob)
}

func updateMap(model *map[string]tableMap) {
	models:=*model
	models["123"]=tableMap{
		Mob:"18551734732",
		Sex:"asdf",
	}
}