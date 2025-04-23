package main

import (
	"github.com/SardelkaS/go-odf/odt"
)

func main() {
	odtFile := odt.New()

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt").WithColor("#FF0000")
	style2 := odt.NewTextStyle().WithFontSize("14pt").WithBold()
	style3 := odt.NewTextStyle().WithFontSize("17pt").WithUnderline()

	// add header
	odtFile.Header("Header", style1, 1)

	// add paragraph with text
	pr := odt.NewParagraph()
	pr.AddText("Some text", style2)
	odtFile.Paragraph(pr)

	// you can use setters for create paragraph
	odtFile.Paragraph(odt.NewParagraph().
		WithText("Some text 1.", style2).
		WithText("Some text 2.", style3))

	// change metadata
	odtFile.Meta.SetInitialCreator("Hi it's me")
	odtFile.Meta.SetCreator("It's me too")
	odtFile.Meta.SetSubject("just test odt file")
	odtFile.Meta.SetDescription("just test odt file")

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-base/example.odt")
	if err != nil {
		panic(err)
	}
}
