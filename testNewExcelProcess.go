package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func main() {
	xlFile,err:=xlsx.OpenFile("./内容计费初始库（现网20161214）.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
				text, _ := row.Cells[0].String()
				fmt.Printf("%s\n", text)
		}

	}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	cell=row.AddCell()
	cell.Value="nihao"
	row=sheet.AddRow()
	cell=row.AddCell()
	cell.Value="asdfsadf"
	err = file.Save("./MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

}
