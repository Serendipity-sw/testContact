package main

import (
	"encoding/json"
	"fmt"
)

type tableMap struct {
	Mob string
	Sex string
}

func main() {
	var (
		modelList         map[string]tableMap          = make(map[string]tableMap)
		itemList          map[string]tableMap          = make(map[string]tableMap)
		resultMessageCode map[string]map[string]string = make(map[string]map[string]string)
	)

	fjsisfkl := "{ \"eucpResultCode\":{ \"0\":\"成功\", \"DELIVRD\":\"成功\", \"EM:000\":\"内容为空失败\", \"EM:001\":\"手机号和内容相同失败，频繁发送\", \"EM:101\":\"黑名单失败\", \"EM:102\":\"签名未报备，失败\", \"EM:103\":\"黑字典失败\", \"EM:104\":\"移动通道无签名失败\", \"EM:105\":\"联通通道/电信通道无签名失败\", \"EM:2000\":\"无签名失败\", \"EM:2001\":\"双签名，失败\", \"MOBFAIL\":\"运营商黑名单\", \"CJ:0007\":\"签名被运营商屏蔽\" } }"

	err := json.Unmarshal([]byte(fjsisfkl), &resultMessageCode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resultMessageCode["eucpResultCode"]["0"])

	modelList["邵炜"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	modelList["邵炜1"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	modelList["邵炜2"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	modelList["邵炜3"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}

	itemList["1邵炜"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	itemList["2邵炜1"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	itemList["3邵炜2"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	itemList["4邵炜3"] = tableMap{
		Mob: "",
		Sex: "asdf",
	}
	fmt.Println(len(itemList))
	updateMap(&itemList)
	fmt.Println(len(itemList))
	fmt.Println(itemList["123"].Mob)
}

func updateMap(model *map[string]tableMap) {
	models := *model
	models["123"] = tableMap{
		Mob: "18551734732",
		Sex: "asdf",
	}
}
