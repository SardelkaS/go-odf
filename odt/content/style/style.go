package style

import (
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

var nameIter = atomic.Uint64{}

// Style represents text formatting options for ODT documents.
// It allows customization of font properties, weight, style, and color.
type Style struct {
	name      string
	fontName  string
	fontSize  string // e.g. "12pt", "1.5cm"
	bold      bool
	italic    bool
	underline bool
	color     string // hex color like "#000000"
}

// New creates new Style with default values
func New() *Style {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 2
		nameIter.Add(2)
	}

	return &Style{
		name: fmt.Sprintf("T%s", strconv.FormatUint(iter, 10)),
	}
}

// WithFontName set FontName. FontName specifies the font family (e.g., "Arial", "Times New Roman")
func (s *Style) WithFontName(fontName string) *Style {
	s.fontName = fontName
	return s
}

// WithFontSize set FontSize. FontSize specifies the text size with units (e.g., "12pt", "1.5cm")
// Common values: "10pt", "12pt", "14pt", "16pt", "18pt"
func (s *Style) WithFontSize(fontSize string) *Style {
	s.fontSize = fontSize
	return s
}

// WithBold enables bold text formatting
func (s *Style) WithBold() *Style {
	s.bold = true
	return s
}

// WithItalic enables italic text formatting
func (s *Style) WithItalic() *Style {
	s.italic = true
	return s
}

// WithUnderline enables text underlining
func (s *Style) WithUnderline() *Style {
	s.underline = true
	return s
}

// WithColor specifies text color in hex format (e.g., "#FF0000" for red)
// Supported formats: "#RGB", "#RRGGBB"
func (s *Style) WithColor(c string) *Style {
	s.color = c
	return s
}

// GetName returns style name
func (s *Style) GetName() string {
	return s.name
}

// Generate generates xml code
func (s *Style) Generate() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="text">`, s.name))
	builder.WriteString("<style:text-properties")

	if s.fontName != "" {
		builder.WriteString(fmt.Sprintf(` style:font-name="%s"`, s.fontName))
	}
	if s.fontSize != "" {
		builder.WriteString(fmt.Sprintf(` fo:font-size="%s"`, s.fontSize))
	}
	if s.bold {
		builder.WriteString(` fo:font-weight="bold"`)
	}
	if s.italic {
		builder.WriteString(` fo:font-style="italic"`)
	}
	if s.underline {
		builder.WriteString(` style:text-underline-style="solid" style:text-underline-width="auto" style:text-underline-color="font-color"`)
	}
	if s.color != "" {
		builder.WriteString(fmt.Sprintf(` fo:color="%s"`, s.color))
	}

	builder.WriteString("/></style:style>")
	return builder.String()
}
