package content

import (
	"bytes"
	"github.com/SardelkaS/go-odf/odt/model"
)

type Element interface {
	GetFilesInfo() []model.FileInfo
	GenerateStyles() string
	Generate() string
}

type Content struct {
	Elements []Element
}

// New creates new empty Content
func New() *Content {
	return &Content{
		Elements: []Element{},
	}
}

// Add adds new element
func (c *Content) Add(e Element) {
	c.Elements = append(c.Elements, e)
}

// GetFilesInfo returns information about additional files
func (c *Content) GetFilesInfo() []model.FileInfo {
	var result []model.FileInfo
	for _, e := range c.Elements {
		info := e.GetFilesInfo()
		if len(info) > 0 {
			result = append(result, info[:]...)
		}
	}

	return result
}

// Generate generates xml code
func (c *Content) Generate() string {
	var contentBuffer bytes.Buffer

	contentBuffer.WriteString(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-content xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" 
    xmlns:ooo="http://openoffice.org/2004/office"
    xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0"
    xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
    xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0"
    xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0"
    xmlns:rpt="http://openoffice.org/2005/report"
    xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0"
    xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0"
    xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0"
    xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0"
    xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0"
    xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0"
    xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc"
    xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2"
    xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:tableooo="http://openoffice.org/2009/table"
    xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0"
    xmlns:drawooo="http://openoffice.org/2010/draw" xmlns:xhtml="http://www.w3.org/1999/xhtml"
    xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0"
    xmlns:field="urn:openoffice:names:experimental:ooo-ms-interop:xmlns:field:1.0"
    xmlns:math="http://www.w3.org/1998/Math/MathML"
    xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0"
    xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0"
    xmlns:formx="urn:openoffice:names:experimental:ooxml-odf-interop:xmlns:form:1.0"
    xmlns:dom="http://www.w3.org/2001/xml-events" xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:grddl="http://www.w3.org/2003/g/data-view#" xmlns:css3t="http://www.w3.org/TR/css3-text/"
    xmlns:officeooo="http://openoffice.org/2009/office" office:version="1.4">
    <office:automatic-styles>
        <style:style style:name="P1" style:family="paragraph">
            <style:text-properties fo:font-size="12pt"/>
        </style:style>`)

	for _, e := range c.Elements {
		contentBuffer.WriteString(e.GenerateStyles())
	}

	contentBuffer.WriteString(`</office:automatic-styles>
    <office:body>
        <office:text>`)

	for _, e := range c.Elements {
		contentBuffer.WriteString(e.Generate())
	}

	contentBuffer.WriteString(`</office:text>
    </office:body>
</office:document-content>`)

	return contentBuffer.String()
}
