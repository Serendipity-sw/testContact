package main

import (
	"github.com/Luxurioust/excelize"
	"fmt"
)

func main() {
	xlsx,_:=excelize.OpenFile("./ffffff.xlsx")
	fmt.Println(xlsx.SheetCount)
	name:=xlsx.GetCellValue(fmt.Sprintf("Sheet%d",1),"A2")
	fmt.Println(xlsx.GetRows(fmt.Sprintf("Sheet%d",1)))
	fmt.Println(name)
}
