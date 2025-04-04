package main

import (
	"encoding/base64"
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/image"
	"github.com/SardelkaS/go-odf/odt/content/style"
	"github.com/SardelkaS/go-odf/odt/document"
	"os"
)

func main() {
	odtFile := document.New()

	// create styles
	style1 := style.New().WithFontSize("30pt").WithBold()
	style2 := style.New().WithFontSize("25pt").WithBold()
	captionStyle := style.New().WithFontSize("20pt").WithUnderline().WithBold()

	// load images from files
	data, err := os.ReadFile("./odt/examples/example-image/1.jpg")
	if err != nil {
		panic(err)
	}
	img1Base64Str := base64.StdEncoding.EncodeToString(data)

	data, err = os.ReadFile("./odt/examples/example-image/2.jpg")
	if err != nil {
		panic(err)
	}
	img2Base64Str := base64.StdEncoding.EncodeToString(data)

	// create image for odt
	img1, err := image.New(img1Base64Str)
	if err != nil {
		panic(err)
	}

	img2, err := image.New(img2Base64Str)
	if err != nil {
		panic(err)
	}
	img2.SetHeight("8cm")
	img2.SetWidth("300px")
	img2.SetCaption("Some image name")
	img2.SetCaptionStyle(captionStyle)

	// add paragraphs
	pr := paragraph.New()
	pr.AddText("Images", style1)
	odtFile.Paragraph(pr)

	pr = paragraph.New()
	pr.AddImage(img1)
	odtFile.Paragraph(pr)

	pr = paragraph.New()
	pr.AddText("Custom image", style2)
	odtFile.Paragraph(pr)

	pr = paragraph.New()
	pr.AddImage(img2)
	odtFile.Paragraph(pr)

	// save generated file
	err = odtFile.SaveToFile("./odt/examples/example-image/example.odt")
	if err != nil {
		panic(err)
	}
}
