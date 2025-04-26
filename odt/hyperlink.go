package odt

import "fmt"

// Hyperlink represents hyperlink element
type Hyperlink struct {
	text         string
	href         string
	style        *Style
	visitedStyle *Style
}

// NewHyperLink creates new empty Hyperlink
func NewHyperLink() *Hyperlink {
	return &Hyperlink{}
}

// WithText sets text of hyperlink
func (h *Hyperlink) WithText(t string) *Hyperlink {
	h.text = t
	return h
}

// WithLink sets destination link
func (h *Hyperlink) WithLink(l string) *Hyperlink {
	h.href = l
	return h
}

// WithStyle sets style of hyperlink element
func (h *Hyperlink) WithStyle(s *Style) *Hyperlink {
	h.style = s
	return h
}

// WithVisitedStyle sets style of hyperlink element after click
func (h *Hyperlink) WithVisitedStyle(s *Style) *Hyperlink {
	h.visitedStyle = s
	return h
}

func (h *Hyperlink) getElementType() string {
	return _linkElement
}

func (h *Hyperlink) generateStyles() string {
	return h.style.generate() + h.visitedStyle.generate()
}

func (h *Hyperlink) generate() string {
	return fmt.Sprintf(`<text:a xlink:type="simple" xlink:href="%s" text:style-name="%s" text:visited-style-name="%s">%s</text:a>`,
		h.href, h.style.getName(), h.visitedStyle.getName(), h.text)
}
