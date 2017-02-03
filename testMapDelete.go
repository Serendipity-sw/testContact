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

	for key, _ := range modelList {
		for itemKey, _ := range itemList {
			fmt.Println(itemKey)
			delete(itemList,itemKey)
		}
		fmt.Println(key)
		delete(modelList,key)
	}
	fmt.Println(len(itemList))
	fmt.Println(len(modelList))
}