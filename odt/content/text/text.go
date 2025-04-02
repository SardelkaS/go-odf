package content_text

import (
	"fmt"
	"github.com/SardelkaS/go-odf/helpers"
	"github.com/SardelkaS/go-odf/odt/content/style"
)

// Text simple text with style
type Text struct {
	Style style.Style
	text  string
}

// New creates new text block
func New(text string, s style.Style) *Text {
	return &Text{
		Style: s,
		text:  text,
	}
}

// SetText set text data
func (t *Text) SetText(text string) {
	t.text = text
}

// GetStyle returns Text style
func (t *Text) GetStyle() style.Style {
	return t.Style
}

// Generate generates xml code
func (t *Text) Generate() string {
	return fmt.Sprintf(`<text:span text:style-name="%s">%s</text:span>`,
		t.Style.GetName(), helpers.EscapeXML(t.text))
}
