package main

import (
	"fmt"
	"reflect"
)

func main() {
	If(true, fhusadf, asveawrtgewrt)

}

func fhusadf() {
	fmt.Println(123)
}

func asveawrtgewrt() {
	fmt.Println(321)
}

/**
构建三元运算
创建人:邵炜
创建时间:2017年2月8日18:51:36
输入参数:是否匹配 第一返回数 第二返回数
*/
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		fasdf := reflect.TypeOf(trueVal)
		fmt.Println(fasdf.Name())
		if fasdf.Kind() == reflect.Func { //通过Kind()来判断反射出的类型是否为需要的类型
			fmt.Println("err: type invalid!")
		}
		return trueVal
	}
	return falseVal
}
