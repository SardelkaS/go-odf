package main

import (
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/text/style"
	"github.com/SardelkaS/go-odf/odt/content/table"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()

	// create styles
	style1 := style.New().WithFontSize("30pt")

	// add paragraph with text
	pr := paragraph.New()
	pr.AddText("Table", style1)
	odtFile.Paragraph(pr)

	// create table
	tbl := table.New(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tbl.SetValue(i, j, fmt.Sprintf("Cell %d:%d", i, j))
		}
	}
	odtFile.Table(tbl)

	// save generated file
	err := odtFile.SaveToFile("./odt/examples/example-table/example.odt")
	if err != nil {
		panic(err)
	}
}
