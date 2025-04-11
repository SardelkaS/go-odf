package odt

import (
	"bytes"
	"strings"
)

type pElement interface {
	getElementType() string
	getStyle() string
	generate() string
}

// Paragraph represents a Paragraph containing text elements
type Paragraph struct {
	elements []pElement
}

func NewParagraph() *Paragraph {
	return &Paragraph{
		elements: []pElement{},
	}
}

// AddText add text with its style to Paragraph. All text in Paragraph will be written in one line
func (p *Paragraph) AddText(text string, style *Style) {
	p.elements = append(p.elements, newText(text, style))
}

// WithText setter analog of AddText
func (p *Paragraph) WithText(text string, style *Style) *Paragraph {
	p.AddText(text, style)
	return p
}

// AddImage add picture to Paragraph
func (p *Paragraph) AddImage(img *Image) {
	p.elements = append(p.elements, img)
}

// WithImage setter analog of AddImage
func (p *Paragraph) WithImage(img *Image) *Paragraph {
	p.AddImage(img)
	return p
}

// generateStyles generates XML representation of the Paragraph styles
func (p *Paragraph) generateStyles() string {
	var stylesBuffer bytes.Buffer
	for _, e := range p.elements {
		stylesBuffer.WriteString(e.getStyle())
	}

	return stylesBuffer.String()
}

// getFilesInfo returns information about additional files
func (p *Paragraph) getFilesInfo() []FileInfo {
	var result []FileInfo
	for _, e := range p.elements {
		if e.getElementType() == _imageElement {
			ie, ok := e.(*Image)
			if !ok {
				continue
			}

			info := ie.getFileInfo()
			if info.Valid() {
				result = append(result, info)
			}
		}
	}

	return result
}

// generate generates XML representation of the Paragraph
func (p *Paragraph) generate() string {
	var builder strings.Builder
	builder.WriteString(`<text:p text:style-name="P1">`)

	for _, element := range p.elements {
		builder.WriteString(element.generate())
	}

	builder.WriteString("</text:p>")
	return builder.String()
}
