# go-odf

[![Go Reference](https://pkg.go.dev/badge/github.com/SardelkaS/go-odf.svg)](https://pkg.go.dev/github.com/SardelkaS/go-odf)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub Issues](https://img.shields.io/github/issues/SardelkaS/go-odf)](https://github.com/SardelkaS/go-odf/issues)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/SardelkaS/go-odf/pulls)

**go-odf** is a Go library for working with **Open Document Format (ODF)** files (now only `.odt` and `.ods` is supported). 
It provides utilities to create ODF documents programmatically.

## **Features**
✔ Generate new ODF files from scratch  
✔ Modify created ODF documents (text, styles, metadata)  
✔ Lightweight and dependency-free (pure Go)

## **Installation**
```sh
go get github.com/SardelkaS/go-odf
```  

## **Quick Start**
### .odt
```go
package main

import (
	"github.com/SardelkaS/go-odf/odt"
)

func main() {
	odtFile := odt.New()
	
	// create style
	style1 := odt.NewTextStyle().WithFontSize("30pt").WithColor("#FF0000")

	// add paragraph with text
	pr := odt.NewParagraph()
	pr.AddText("Header", style1)
	odtFile.Paragraph(pr)

	// save generated file
	err := odtFile.SaveToFile("./example.odt")
	if err != nil {
		panic(err)
	}
}
```  

### .ods
```go
package main

import (
	"github.com/SardelkaS/go-odf/ods"
)

func main() {
	odsFile := ods.New()

	// create styles
	cellStyle1 := ods.NewCellStyle().WithBackgroundColor("#FF0000")

	// create new sheet
	sheet1 := ods.NewSheet("Sheet1")
	odsFile.Sheet(sheet1)

	// add data to sheet
	_ = sheet1.SetCellStyle("A1", cellStyle1)
	_ = sheet1.SetCellValue("A1", "test")
	_ = sheet1.SetCellValueType("B1", ods.Float)
	_ = sheet1.SetCellValue("B1", "0.95")
	_ = sheet1.SetCellValue("D3", "test test")

	// save generated file
	err := odsFile.SaveToFile("./example.ods")
	if err != nil {
		panic(err)
	}
}
```

## **Usage Examples**
See the [examples directory](https://github.com/SardelkaS/go-odf/tree/main/examples/odt) for practical implementations.
