package paragraph

import (
	"bytes"
	"github.com/SardelkaS/go-odf/odt/content/style"
	content_text "github.com/SardelkaS/go-odf/odt/content/text"
	"strings"
)

// Paragraph represents a paragraph containing text elements
type Paragraph struct {
	Elements []*content_text.Text
}

func New() *Paragraph {
	return &Paragraph{
		Elements: []*content_text.Text{},
	}
}

func (p *Paragraph) AddText(text string, style style.Style) {
	p.Elements = append(p.Elements, content_text.New(text, style))
}

func (p *Paragraph) GetStyles() string {
	var stylesBuffer bytes.Buffer
	for _, e := range p.Elements {
		stylesBuffer.WriteString(e.Style.Generate())
	}

	return stylesBuffer.String()
}

// Generate generates XML representation of the paragraph
func (p *Paragraph) Generate() string {
	var builder strings.Builder
	builder.WriteString(`<text:p text:style-name="P1">`)

	for _, element := range p.Elements {
		builder.WriteString(element.Generate())
	}

	builder.WriteString("</text:p>")
	return builder.String()
}
