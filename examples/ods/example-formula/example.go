package main

import "github.com/SardelkaS/go-odf/ods"

func main() {
	odsFile := ods.New()

	// create new sheet
	sheet1 := ods.NewSheet("Sheet1")
	odsFile.Sheet(sheet1)

	// add data to sheet
	_ = sheet1.SetCellValueType("A1", ods.Float)
	_ = sheet1.SetCellValue("A1", "1")
	_ = sheet1.SetCellValueType("A2", ods.Float)
	_ = sheet1.SetCellValue("A2", "3")
	_ = sheet1.SetCellFormula("B1", "=A1+A2")

	// save generated file
	err := odsFile.SaveToFile("./examples/ods/example-formula/example.ods")
	if err != nil {
		panic(err)
	}
}
