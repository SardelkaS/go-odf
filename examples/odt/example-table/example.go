package main

import (
	"fmt"
	"github.com/SardelkaS/go-odf/odt"
)

func main() {
	odtFile := odt.New()

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt")

	// add paragraph with text
	pr := odt.NewParagraph()
	pr.AddText("Table", style1)
	odtFile.Paragraph(pr)

	// create table
	tbl := odt.NewTable(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tbl.SetValue(i, j, fmt.Sprintf("cell %d:%d", i, j))
		}
	}
	odtFile.Table(tbl)

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-table/example.odt")
	if err != nil {
		panic(err)
	}
}
