package main

import "github.com/SardelkaS/go-odf/odt"

func main() {
	odtFile := odt.New()

	// title for the document
	style1 := odt.NewTextStyle().WithFontSize("25pt").WithBold()
	pr := odt.NewParagraph()
	pr.AddText("Charts", style1)
	odtFile.Paragraph(pr)

	// bar chart
	barChart := odt.NewChart(odt.ChartTypeBar)
	barChart.SetTitle("Sales by quarter")
	barChart.SetSize("14cm", "8cm")
	barChart.SetCategories([]string{"Q1", "Q2", "Q3", "Q4"})
	barChart.AddSeries("2024", []float64{12, 18, 9, 21})
	barChart.AddSeries("2025", []float64{15, 20, 14, 25})
	odtFile.Chart(barChart)

	// line chart
	lineChart := odt.NewChart(odt.ChartTypeLine)
	lineChart.SetTitle("Temperature, °C")
	lineChart.SetSize("14cm", "8cm")
	lineChart.SetCategories([]string{"Mon", "Tue", "Wed", "Thu", "Fri"})
	lineChart.AddSeries("Min", []float64{10, 12, 9, 11, 13})
	lineChart.AddSeries("Max", []float64{20, 22, 19, 21, 23})
	odtFile.Chart(lineChart)

	// pie chart
	pieChart := odt.NewChart(odt.ChartTypePie)
	pieChart.SetTitle("Market share")
	pieChart.SetSize("12cm", "8cm")
	pieChart.SetCategories([]string{"Product A", "Product B", "Product C"})
	pieChart.AddSeries("Share", []float64{45, 35, 20})
	odtFile.Chart(pieChart)

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-chart/example.odt")
	if err != nil {
		panic(err)
	}
}
