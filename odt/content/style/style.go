package style

import (
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

var nameIter = atomic.Uint64{}

type Style struct {
	name      string
	FontName  string
	FontSize  string // e.g. "12pt", "1.5cm"
	Bold      bool
	Italic    bool
	Underline bool
	Color     string // hex color like "#000000"
}

// New creates new Style with default values
func New() Style {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 2
		nameIter.Add(2)
	}

	return Style{
		name: fmt.Sprintf("T%s", strconv.FormatUint(iter, 10)),
	}
}

// GetName returns style name
func (s Style) GetName() string {
	return s.name
}

// Generate generates xml code
func (s Style) Generate() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="text">`, s.name))
	builder.WriteString("<style:text-properties")

	if s.FontName != "" {
		builder.WriteString(fmt.Sprintf(` style:font-name="%s"`, s.FontName))
	}
	if s.FontSize != "" {
		builder.WriteString(fmt.Sprintf(` fo:font-size="%s"`, s.FontSize))
	}
	if s.Bold {
		builder.WriteString(` fo:font-weight="bold"`)
	}
	if s.Italic {
		builder.WriteString(` fo:font-style="italic"`)
	}
	if s.Underline {
		builder.WriteString(` style:text-underline-style="solid" style:text-underline-width="auto" style:text-underline-color="font-color"`)
	}
	if s.Color != "" {
		builder.WriteString(fmt.Sprintf(` fo:color="%s"`, s.Color))
	}

	builder.WriteString("/></style:style>")
	return builder.String()
}
