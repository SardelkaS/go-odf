package main

import (
	"github.com/SardelkaS/go-odf/odt/content/style"
	"github.com/SardelkaS/go-odf/odt/content/style/types"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()

	style1 := style.New()
	_ = style1.Props.Text.SetFontSize(20)
	style1.Props.Text.SetFontName(types.FontName_LiberationSans)

	odtFile.Text("Hello, world!", style1)

	style2 := style.New()
	_ = style2.Props.Text.SetFontSize(12)
	style2.Props.Text.SetFontName(types.FontName_NSimSun)

	odtFile.Text("Hello world 1111", style1)
	odtFile.Text("Nice to meet you", style2)

	err := odtFile.SaveToFile("./odt/examples/example1.odt")
	if err != nil {
		panic(err)
	}
}
