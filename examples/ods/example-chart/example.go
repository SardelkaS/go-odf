package main

import "github.com/SardelkaS/go-odf/ods"

func main() {
	odsFile := ods.New()

	// create new sheet with some source data
	sheet1 := ods.NewSheet("Sheet1")
	odsFile.Sheet(sheet1)

	_ = sheet1.SetCellValue("A1", "Fruit")
	_ = sheet1.SetCellValue("B1", "Count")
	_ = sheet1.SetCellValue("A2", "Apple")
	_ = sheet1.SetCellValueType("B2", ods.Float)
	_ = sheet1.SetCellValue("B2", "5")
	_ = sheet1.SetCellValue("A3", "Pear")
	_ = sheet1.SetCellValueType("B3", ods.Float)
	_ = sheet1.SetCellValue("B3", "2")
	_ = sheet1.SetCellValue("A4", "Plum")
	_ = sheet1.SetCellValueType("B4", ods.Float)
	_ = sheet1.SetCellValue("B4", "7")

	// bar chart placed on the sheet
	barChart := ods.NewChart(ods.ChartTypeBar)
	barChart.SetTitle("Fruit count")
	barChart.SetSize("14cm", "8cm")
	barChart.SetPosition("6cm", "0.5cm")
	barChart.SetCategories([]string{"Apple", "Pear", "Plum"})
	barChart.AddSeries("Count", []float64{5, 2, 7})
	sheet1.AddChart(barChart)

	// second sheet with a pie chart
	sheet2 := ods.NewSheet("Sheet2")
	odsFile.Sheet(sheet2)

	pieChart := ods.NewChart(ods.ChartTypePie)
	pieChart.SetTitle("Market share")
	pieChart.SetSize("12cm", "8cm")
	pieChart.SetCategories([]string{"Product A", "Product B", "Product C"})
	pieChart.AddSeries("Share", []float64{45, 35, 20})
	sheet2.AddChart(pieChart)

	// save generated file
	err := odsFile.SaveToFile("./examples/ods/example-chart/example.ods")
	if err != nil {
		panic(err)
	}
}
