package main

import (
	"encoding/base64"
	"github.com/SardelkaS/go-odf/odt"
	"os"
)

func main() {
	odtFile := odt.New()

	// create styles
	style1 := odt.NewTextStyle().WithFontSize("30pt").WithBold()
	style2 := odt.NewTextStyle().WithFontSize("25pt").WithBold()
	captionStyle := odt.NewTextStyle().WithFontSize("20pt").WithUnderline().WithBold()

	// load images from files
	data, err := os.ReadFile("./examples/odt/example-image/1.jpg")
	if err != nil {
		panic(err)
	}
	img1Base64Str := base64.StdEncoding.EncodeToString(data)

	data, err = os.ReadFile("./examples/odt/example-image/2.jpg")
	if err != nil {
		panic(err)
	}
	img2Base64Str := base64.StdEncoding.EncodeToString(data)

	// create image for odt
	img1, err := odt.NewImage(img1Base64Str)
	if err != nil {
		panic(err)
	}

	img2, err := odt.NewImage(img2Base64Str)
	if err != nil {
		panic(err)
	}
	img2.SetHeight("8cm")
	img2.SetWidth("300px")
	img2.SetCaption("Some image name")
	img2.SetCaptionStyle(captionStyle)

	// add paragraphs
	pr := odt.NewParagraph()
	pr.AddText("Images", style1)
	odtFile.Paragraph(pr)

	pr = odt.NewParagraph()
	pr.AddImage(img1)
	odtFile.Paragraph(pr)

	pr = odt.NewParagraph()
	pr.AddText("Custom image", style2)
	odtFile.Paragraph(pr)

	pr = odt.NewParagraph()
	pr.AddImage(img2)
	odtFile.Paragraph(pr)

	// save generated file
	err = odtFile.SaveToFile("./examples/odt/example-image/example.odt")
	if err != nil {
		panic(err)
	}
}
