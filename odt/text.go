package odt

import (
	"fmt"
	"github.com/SardelkaS/go-odf/helpers"
)

// text simple text with style
type text struct {
	style *Style
	text  string
	tag   string
}

// newText creates new text block
// param t contains the actual text to be displayed
// param style defines formatting rules (nil for default formatting)
func newText(t string, style *Style) *text {
	return &text{
		style: style,
		text:  t,
		tag:   _textTagSpan,
	}
}

func (t *text) withTag(tag string) *text {
	t.tag = tag
	return t
}

// setText set text content. Content contains the actual text to be displayed
// Special characters will be automatically XML-escaped
func (t *text) setText(text string) {
	t.text = text
}

// GetElementType returns element type
func (t *text) getElementType() string {
	return _textElement
}

// generateStyles returns text style
func (t *text) generateStyles() string {
	return t.style.generate()
}

// generate generates xml code
func (t *text) generate() string {
	return fmt.Sprintf(`<text:%s text:style-name="%s">%s</text:%s>`,
		t.tag, t.style.getName(), helpers.EscapeXML(t.text), t.tag)
}
