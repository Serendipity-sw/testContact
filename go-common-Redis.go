package main

import (
	"fmt"
	"github.com/swgloomy/go-common"
)

func main() {
	err := common.OpenRedis("192.200.154.109:6379", 5)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err.Error())
		return
	}
	err = common.SetRedisCache("gloomysw", "hello world!", 863000)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err.Error())
		return
	}
	result, err := common.GetRedisCache("gloomysw")
	if err != nil {
		fmt.Println(3)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
