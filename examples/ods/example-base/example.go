package main

import (
	"github.com/SardelkaS/go-odf/ods"
)

func main() {
	odsFile := ods.New()

	// create styles
	cellStyle1 := ods.NewCellStyle().WithBackgroundColor("#FF0000")
	rowStyle1 := ods.NewRowStyle().WithBackgroundColor("#00AA00")
	cellStyle2 := ods.NewCellStyle().WithBackgroundColor("#0000AA")

	// create new sheet
	sheet1 := ods.NewSheet("Sheet1")
	odsFile.Sheet(sheet1)

	// add data to sheet
	sheet1.SetRowStyle(2, rowStyle1)
	sheet1.SetColumnDefaultCellStyle(2, cellStyle2)
	_ = sheet1.SetCellStyle("A1", cellStyle1)
	_ = sheet1.SetCellValue("A1", "test")
	_ = sheet1.SetCellValueType("B1", ods.Float)
	_ = sheet1.SetCellValue("B1", "0.95")
	_ = sheet1.SetCellValue("D3", "test test")

	// change metadata
	odsFile.Meta.SetInitialCreator("Hi it's me")
	odsFile.Meta.SetCreator("It's me too")
	odsFile.Meta.SetSubject("just test odt file")
	odsFile.Meta.SetDescription("just test odt file")

	// save generated file
	err := odsFile.SaveToFile("./examples/ods/example-base/example.ods")
	if err != nil {
		panic(err)
	}
}
