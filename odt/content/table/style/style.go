package table_style

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var styleNameIter = atomic.Uint64{}

// Style contains styling for the entire table
type Style struct {
	name            string
	width           string
	align           string
	margin          string
	backgroundColor string
	borderModel     string
	border          string
}

func NewStyle() *Style {
	iter := styleNameIter.Load()
	styleNameIter.Add(1)
	if iter == 0 {
		iter = 1
		styleNameIter.Add(1)
	}

	return &Style{
		name: fmt.Sprintf("TblS%s", strconv.FormatUint(iter, 10)),
	}
}

// WithWidth sets table width
func (s *Style) WithWidth(w string) *Style {
	s.width = w
	return s
}

// WithAlign sets table align
func (s *Style) WithAlign(a string) *Style {
	s.align = a
	return s
}

// WithMargin sets table margin
func (s *Style) WithMargin(m string) *Style {
	s.margin = m
	return s
}

// WithBackgroundColor sets table background color
func (s *Style) WithBackgroundColor(c string) *Style {
	s.backgroundColor = c
	return s
}

// GetName returns style name
func (s *Style) GetName() string {
	return s.name
}

// WithBorderModel sets border model
//
// BorderModelCollapsing or BorderModelSeparating
func (s *Style) WithBorderModel(m string) *Style {
	s.borderModel = m
	return s
}

// WithBorder sets border properties
//
// example: "0.002cm solid #000000"
func (s *Style) WithBorder(border string) *Style {
	s.border = border
	return s
}

// Generate generates the XML representation of the table style
func (s *Style) Generate() string {
	var buf bytes.Buffer

	buf.WriteString(`<style:style style:name="`)
	buf.WriteString(s.name)
	buf.WriteString(`" style:family="table">`)
	buf.WriteString(`<style:table-properties`)

	if s.width != "" {
		buf.WriteString(` style:width="`)
		buf.WriteString(s.width)
		buf.WriteString(`"`)
	}

	if s.align != "" {
		buf.WriteString(` table:align="`)
		buf.WriteString(s.align)
		buf.WriteString(`"`)
	}

	if s.margin != "" {
		buf.WriteString(` fo:margin="`)
		buf.WriteString(s.margin)
		buf.WriteString(`"`)
	}

	if s.backgroundColor != "" {
		buf.WriteString(` fo:background-color="`)
		buf.WriteString(s.backgroundColor)
		buf.WriteString(`"`)
	}

	if s.borderModel != "" {
		buf.WriteString(` table:border-model="`)
		buf.WriteString(s.borderModel)
		buf.WriteString(`"`)
	}

	if s.border != "" {
		buf.WriteString(` fo:border="`)
		buf.WriteString(s.border)
		buf.WriteString(`"`)
	}

	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
