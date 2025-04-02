package paragraph

import (
	"bytes"
	"github.com/SardelkaS/go-odf/odt/content/style"
	content_text "github.com/SardelkaS/go-odf/odt/content/text"
	"strings"
)

// Paragraph represents a paragraph containing text elements
type Paragraph struct {
	elements []*content_text.Text
}

func New() *Paragraph {
	return &Paragraph{
		elements: []*content_text.Text{},
	}
}

// AddText add text with its style to paragraph. All text in paragraph will be written in one line
func (p *Paragraph) AddText(text string, style *style.Style) {
	p.elements = append(p.elements, content_text.New(text, style))
}

// GenerateStyles generates XML representation of the paragraph styles
func (p *Paragraph) GenerateStyles() string {
	var stylesBuffer bytes.Buffer
	for _, e := range p.elements {
		stylesBuffer.WriteString(e.GetStyle().Generate())
	}

	return stylesBuffer.String()
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
