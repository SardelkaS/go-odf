package main

import (
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/text/style"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()

	// create styles
	style1 := style.New().WithFontSize("30pt").WithColor("#FF0000")
	style2 := style.New().WithFontSize("14pt").WithBold()
	style3 := style.New().WithFontSize("17pt").WithUnderline()

	// add paragraph with text
	pr := paragraph.New()
	pr.AddText("Header", style1)
	odtFile.Paragraph(pr)

	// you can use setters for create paragraph
	odtFile.Paragraph(paragraph.New().
		WithText("Some text 1.", style2).
		WithText("Some text 2.", style3))

	// change metadata
	odtFile.Meta.SetInitialCreator("Hi it's me")
	odtFile.Meta.SetCreator("It's me too")
	odtFile.Meta.SetSubject("just test odt file")
	odtFile.Meta.SetDescription("just test odt file")

	// save generated file
	err := odtFile.SaveToFile("./odt/examples/example-base/example.odt")
	if err != nil {
		panic(err)
	}
}
