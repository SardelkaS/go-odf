# go-odf

[![Go Reference](https://pkg.go.dev/badge/github.com/SardelkaS/go-odf.svg)](https://pkg.go.dev/github.com/SardelkaS/go-odf)  
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)  
[![GitHub Issues](https://img.shields.io/github/issues/SardelkaS/go-odf)](https://github.com/SardelkaS/go-odf/issues)  
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/SardelkaS/go-odf/pulls)

**go-odf** is a Go library for working with **Open Document Format (ODF)** files (now only `.odt` is supported). It provides utilities to create ODF documents programmatically.

## **Features**
✔ Generate new ODF files from scratch  
✔ Modify created ODF documents (text, styles, metadata)  
✔ Lightweight and dependency-free (pure Go)

## **Installation**
```sh
go get github.com/SardelkaS/go-odf
```  

## **Quick Start**
```go
package main

import (
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/text/style"
	"github.com/SardelkaS/go-odf/odt/document"
)

func main() {
	odtFile := document.New()
	
	// create style
	style1 := style.New().WithFontSize("30pt").WithColor("#FF0000")

	// add paragraph with text
	pr := paragraph.New()
	pr.AddText("Header", style1)
	odtFile.Paragraph(pr)

	// save generated file
	err := odtFile.SaveToFile("./example.odt")
	if err != nil {
		panic(err)
	}
}
```  

## **Usage Examples**
See the [examples directory](https://github.com/SardelkaS/go-odf/tree/main/odt/examples) for practical implementations.
