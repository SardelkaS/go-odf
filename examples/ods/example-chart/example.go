package main

import (
	"github.com/SardelkaS/go-odf/ods"
)

func main() {
	odsFile := ods.New()

	// create new sheet
	sheet1 := ods.NewSheet("Sheet1")
	odsFile.Sheet(sheet1)

	// fill data
	_ = sheet1.SetCellValueType("A1", ods.Float)
	_ = sheet1.SetCellValue("A1", "1")
	_ = sheet1.SetCellValueType("A2", ods.Float)
	_ = sheet1.SetCellValue("A2", "2")
	_ = sheet1.SetCellValueType("A3", ods.Float)
	_ = sheet1.SetCellValue("A3", "3")
	_ = sheet1.SetCellValueType("A4", ods.Float)
	_ = sheet1.SetCellValue("A4", "4")
	_ = sheet1.SetCellValueType("A5", ods.Float)
	_ = sheet1.SetCellValue("A5", "5")

	// create chart style
	chartStyle := ods.NewChartStyle()

	// create chart
	chart := ods.NewChart()
	chart.SetStyle(chartStyle)
	chart.SetWidth("15cm")
	chart.SetHeight("10cm")
	chart.SetPosX("20cm")
	chart.SetPosY("20cm")
	chart.SetDataRange("Sheet1.A1:Sheet1.A5")

	sheet1.AddChart(chart)

	// save generated file
	err := odsFile.SaveToFile("./examples/ods/example-chart/example.ods")
	if err != nil {
		panic(err)
	}
}
