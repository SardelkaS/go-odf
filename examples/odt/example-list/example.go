package main

import "github.com/SardelkaS/go-odf/odt"

func main() {
	odtFile := odt.New()

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt")
	style2 := odt.NewTextStyle().WithFontSize("14pt")
	style3 := odt.NewTextStyle().WithFontSize("14pt").WithBold()

	// add paragraph with text
	pr := odt.NewParagraph()
	pr.AddText("List", style1)
	odtFile.Paragraph(pr)

	odtFile.Paragraph(odt.NewParagraph().WithText("Simple list", style3))
	vList := odt.NewList()
	vList.AddText("cucumber", style2)
	vList.AddText("tomato", style2)
	odtFile.List(vList)

	odtFile.Paragraph(odt.NewParagraph().WithText("Custom list", style3))
	fList := odt.NewList()
	fList.SetNumFormat(1, odt.NumberStyleLowerAlpha)
	fList.SetNumSuffix(1, ".")
	fList.SetNumFormat(2, odt.NumberStyleLowerAlpha)
	fList.SetNumSuffix(2, ".")
	fList.AddText("apple", style2)
	fList.AddText("banana", style2)
	fList.AddText("pineapple", style3)
	odtFile.List(fList)

	odtFile.Paragraph(odt.NewParagraph().WithText("Sublist", style3))

	sList := odt.NewList()
	sList.AddText("fruits", style2)
	sList.AddList(fList)
	sList.AddText("vegetables", style2)
	sList.AddList(vList)
	odtFile.List(sList)

	// save generated file
	err := odtFile.SaveToFile("./examples/odt/example-list/example.odt")
	if err != nil {
		panic(err)
	}
}
