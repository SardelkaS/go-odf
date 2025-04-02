package content

import (
	"bytes"
)

type Element interface {
	GetStyles() string
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

// Add - add new element
func (c *Content) Add(e Element) {
	c.Elements = append(c.Elements, e)
}

// Generate generates xml code
func (c *Content) Generate() string {
	var contentBuffer bytes.Buffer

	contentBuffer.WriteString(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-content xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" 
    xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" 
    xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" 
    xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0">
    <office:automatic-styles>
        <style:style style:name="P1" style:family="paragraph">
            <style:text-properties fo:font-size="12pt"/>
        </style:style>`)

	for _, e := range c.Elements {
		contentBuffer.WriteString(e.GetStyles())
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
