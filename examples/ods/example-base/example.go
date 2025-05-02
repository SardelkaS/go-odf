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
	sheet1.SetCellStyle(0, 0, cellStyle1)
	sheet1.SetCellValue(0, 0, "test", ods.String)
	sheet1.SetCellValue(0, 1, "0.95", ods.Float)
	sheet1.SetCellValue(3, 4, "test test", ods.String)

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
