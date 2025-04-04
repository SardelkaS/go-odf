package paragraph

import (
	"bytes"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/image"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/text"
	"github.com/SardelkaS/go-odf/odt/content/style"
	"github.com/SardelkaS/go-odf/odt/model"
	"strings"
)

type element interface {
	GetElementType() string
	GetStyle() string
	Generate() string
}

// Paragraph represents a paragraph containing text elements
type Paragraph struct {
	elements []element
}

func New() *Paragraph {
	return &Paragraph{
		elements: []element{},
	}
}

// AddText add text with its style to paragraph. All text in paragraph will be written in one line
func (p *Paragraph) AddText(text string, style *style.Style) {
	p.elements = append(p.elements, content_text.New(text, style))
}

// AddImage add picture to paragraph
func (p *Paragraph) AddImage(img *image.Image) {
	p.elements = append(p.elements, img)
}

// GenerateStyles generates XML representation of the paragraph styles
func (p *Paragraph) GenerateStyles() string {
	var stylesBuffer bytes.Buffer
	for _, e := range p.elements {
		stylesBuffer.WriteString(e.GetStyle())
	}

	return stylesBuffer.String()
}

// GetFilesInfo returns information about additional files
func (p *Paragraph) GetFilesInfo() []model.FileInfo {
	var result []model.FileInfo
	for _, e := range p.elements {
		if e.GetElementType() == components.ImageElement {
			ie, ok := e.(*image.Image)
			if !ok {
				continue
			}

			info := ie.GetFileInfo()
			if info.Valid() {
				result = append(result, info)
			}
		}
	}

	return result
}

// Generate generates XML representation of the paragraph
func (p *Paragraph) Generate() string {
	var builder strings.Builder
	builder.WriteString(`<text:p text:style-name="P1">`)

	for _, element := range p.elements {
		builder.WriteString(element.Generate())
	}

	builder.WriteString("</text:p>")
	return builder.String()
}
