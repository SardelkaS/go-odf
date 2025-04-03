package main

import (
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/style"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()

	// create styles
	style1 := style.New().WithFontSize("30pt").WithColor("#FF0000")
	style2 := style.New().WithFontSize("14pt").WithBold()
	style3 := style.New().WithFontSize("17pt").WithUnderline()

	// add paragraph with text
	pr1 := paragraph.New()
	pr1.AddText("Header", style1)

	pr2 := paragraph.New()
	pr2.AddText("Some text 1.", style2)
	pr2.AddText("Some text 2.", style3)

	odtFile.Paragraph(pr1)
	odtFile.Paragraph(pr2)

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
