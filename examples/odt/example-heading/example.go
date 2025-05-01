package main

import "github.com/SardelkaS/go-odf/odt"

func main() {
	odtFile := odt.New()

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt").WithColor("#FF0000")
	style2 := odt.NewTextStyle().WithFontSize("20pt").WithBold()
	style3 := odt.NewTextStyle().WithFontSize("30pt").WithUnderline()

	// create heading
	heading := odt.NewHeading()
	heading.SetHeader("My Heading")
	heading.SetHeaderStyle(style2)
	odtFile.Heading(heading)

	// create headers
	h1 := odt.NewHeader().
		WithText("Header 1").
		WithLevel(1).
		WithStyle(style1)
	h2 := odt.NewHeader().
		WithText("Header 2").
		WithLevel(2).
		WithStyle(style2)
	h3 := odt.NewHeader().
		WithText("Header 3").
		WithLevel(1).
		WithStyle(style3)

	odtFile.Header(h1)
	odtFile.Header(h2)
	odtFile.Header(h3)

	// add headers links to heading
	heading.AddLink(h1)
	heading.AddLink(h2)
	heading.AddLink(h3)

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-heading/example.odt")
	if err != nil {
		panic(err)
	}
}
