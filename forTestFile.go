package main

import "fmt"

func main() {
	//fileStr:=[]string{"1","2","3"}
	//for _, value := range fileStr {
	//	fmt.Println(value)
	//	break
	//}
	//fmt.Println("2")

	var list map[int]string=make(map[int]string)

	list[1]="1"
	list[2]="1"
	list[3]="1"
	list[4]="1"
	list[5]="1"
	list[6]="1"
	list[7]="1"
	list[8]="1"
	list[9]="1"

	for index, _ := range list {
		if index<5 {
			fmt.Println(index)
			list[index+10]="2"
		}
		//fmt.Println(index)
	}

	for index, _ := range list {
		fmt.Println(index)
	}
	fmt.Println(len(list))
}
