package main

import "fmt"

func main() {
	var (
		haveValueMap map[string]string=make(map[string]string)
		noValueMap map[string]string=make(map[string]string)
	)
	haveValueMap["1"]="2"
	fmt.Println(len(haveValueMap))
	fmt.Println(len(noValueMap))

	number:=4
	chushu:=3
	fmt.Println(number%chushu)

	modelList["nihao"]=stateType  {
		Name :"123",
		PilotShowPv :123,
		PilotShowUV :123,
		PilotClickPV :123,
		PilotClickUV :123,
		PilotClosePV :123,
		PilotClostUV :123,
		DateTime :"123",
		ApiInvokePV :123,
	}
	model,_:=modelList["nihao"]
	model.ApiInvokePV+=50
	modelList["nihao"]=model

	fmt.Println(modelList["nihao"].ApiInvokePV)
}




var modelList map[string]stateType=make(map[string]stateType)

type stateType struct {
	Name string
	PilotShowPv int
	PilotShowUV int
	PilotClickPV int
	PilotClickUV int
	PilotClosePV int
	PilotClostUV int
	DateTime string
	ApiInvokePV int
}