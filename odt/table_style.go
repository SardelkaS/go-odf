package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var styleNameIter = atomic.Uint64{}

// TableStyle contains styling for the entire table
type TableStyle struct {
	name            string
	width           string
	align           string
	margin          string
	backgroundColor string
	borderModel     string
	border          string
}

// NewTableStyle creates empty table style
func NewTableStyle() *TableStyle {
	iter := styleNameIter.Load()
	styleNameIter.Add(1)
	if iter == 0 {
		iter = 1
		styleNameIter.Add(1)
	}

	return &TableStyle{
		name: fmt.Sprintf("TblS%s", strconv.FormatUint(iter, 10)),
	}
}

// WithWidth sets table width
func (s *TableStyle) WithWidth(w string) *TableStyle {
	s.width = w
	return s
}

// WithAlign sets table align
func (s *TableStyle) WithAlign(a string) *TableStyle {
	s.align = a
	return s
}

// WithMargin sets table margin
func (s *TableStyle) WithMargin(m string) *TableStyle {
	s.margin = m
	return s
}

// WithBackgroundColor sets table background color
func (s *TableStyle) WithBackgroundColor(c string) *TableStyle {
	s.backgroundColor = c
	return s
}

// getName returns style name
func (s *TableStyle) getName() string {
	return s.name
}

// WithBorderModel sets border model
//
// BorderModelCollapsing or BorderModelSeparating
func (s *TableStyle) WithBorderModel(m string) *TableStyle {
	s.borderModel = m
	return s
}

// WithBorder sets border properties
//
// example: "0.002cm solid #000000"
func (s *TableStyle) WithBorder(border string) *TableStyle {
	s.border = border
	return s
}

// generate generates the XML representation of the table style
func (s *TableStyle) generate() string {
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
