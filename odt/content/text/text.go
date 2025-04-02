package content_text

import (
	"fmt"
	"github.com/SardelkaS/go-odf/helpers"
	"github.com/SardelkaS/go-odf/odt/content/style"
)

// Text simple text with style
type Text struct {
	style *style.Style
	text  string
}

// New creates new text block
// param text contains the actual text to be displayed
// param style defines formatting rules (nil for default formatting)
func New(text string, style *style.Style) *Text {
	return &Text{
		style: style,
		text:  text,
	}
}

// SetText set text content. Content contains the actual text to be displayed
// Special characters will be automatically XML-escaped
func (t *Text) SetText(text string) {
	t.text = text
}

// GetStyle returns Text style
func (t *Text) GetStyle() *style.Style {
	return t.style
}

// Generate generates xml code
func (t *Text) Generate() string {
	return fmt.Sprintf(`<text:span text:style-name="%s">%s</text:span>`,
		t.style.GetName(), helpers.EscapeXML(t.text))
}
