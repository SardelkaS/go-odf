package main

import "github.com/SardelkaS/go-odf/odt"

func main() {
	odtFile := odt.New()

	// set page style
	ps := odt.NewPageStyle()
	ps.SetPrintOrientation("landscape")
	ps.SetBackgroundColor("#510034")
	odtFile.PageStyle(ps)

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt").WithColor("#FF0000")
	style2 := odt.NewTextStyle().WithFontSize("14pt").WithBold()

	// add header
	odtFile.Header("Header", style1, 1)

	// add paragraph with text
	pr := odt.NewParagraph()
	pr.AddText("Some text", style2)
	odtFile.Paragraph(pr)

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-page-style/example.odt")
	if err != nil {
		panic(err)
	}
}
