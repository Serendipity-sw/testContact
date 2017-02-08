package main

import (
	"encoding/json"
	"fmt"
)

type interFaceMapObj struct {
	UrlGroupList map[int]UrlGroupMap
	UrlList map[int]UrlTableMap
}

/**
网址列表
 */
type UrlTableMap struct {
	UrlPath string
	UrlGroupId int
	Name string
}

/**
网址列表分组
 */
type UrlGroupMap struct {
	Name string
	Count int
	Id int
	CreateId int
}

func main() {
	var (
		urlList map[int]UrlTableMap=make(map[int]UrlTableMap)
		urlGroupList map[int]UrlGroupMap=make(map[int]UrlGroupMap)
	)
	urlList[1]=UrlTableMap{
		UrlGroupId:1,
	}
	urlGroupList[2]=UrlGroupMap{
		Name:"124",
	}
	asfasdf:=&urlList
	vadss:=&urlGroupList
	modelList :=interFaceMapObj{
		UrlGroupList:*vadss,
		UrlList:*asfasdf,
	}

	objStrByte,_:=json.Marshal(modelList)
	fmt.Println(string(objStrByte))


	option:=map[string]interface{}{
		"typ":"file",
	}
option["log"]="./logs"
	fmt.Println(option["typ"])
	fmt.Println(option["log"])
}
