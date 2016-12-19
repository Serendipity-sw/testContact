package main

import (
	"github.com/guotie/config"
	"encoding/json"
	"fmt"
)

type objectStructType struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
}

var (
	objectStructList []objectStructType
)

func main() {
	config.ReadCfg("./config.json")
	objectInter:=config.Get("objectStruct")
	objectByte,err:=json.Marshal(objectInter)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(objectByte))
}