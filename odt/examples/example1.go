package main

import (
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/style"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()

	style1 := style.New()
	style1.FontSize = "20pt"
	style1.Color = "#FF0000"

	style2 := style.New()
	style2.FontSize = "9pt"
	style2.Bold = true

	style3 := style.New()
	style3.FontSize = "11pt"
	style3.Underline = true

	pr1 := paragraph.New()
	pr1.AddText("Header", style1)

	pr2 := paragraph.New()
	pr2.AddText("some text 1", style2)
	pr2.AddText("some test 2", style3)

	odtFile.Paragraph(pr1)
	odtFile.Paragraph(pr2)

	err := odtFile.SaveToFile("./odt/examples/example1.odt")
	if err != nil {
		panic(err)
	}
}
