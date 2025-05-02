package ods

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var csNameIter = atomic.Uint64{}

// ColumnStyle defines formatting properties for a table column in ODS.
// Example:
//
//	style := NewColumnStyle("MyColumn").
//				SetWidth(3.5). // 3.5cm
//				SetBackgroundColor("#FF0000") // Red
type ColumnStyle struct {
	name            string  // Style name (required, referenced in table:table-column)
	width           float64 // Column width in centimeters (0 = auto)
	relWidth        int     // Relative width (e.g., "1*", "2*" for proportional columns)
	useOptimal      bool    // If true, auto-adjusts width to fit content
	backgroundColor string  // Hex color ("#RRGGBB")
	borderLeft      string  // Left border style (e.g., "0.5pt solid #000000")
	borderRight     string  // Right border style
	borderTop       string  // Top border style
	borderBottom    string  // Bottom border style
	breakBefore     string  // Page/column break before ("auto", "column", "page")
	breakAfter      string  // Page/column break after
}

// NewColumnStyle creates a new empty ColumnStyle.
func NewColumnStyle() *ColumnStyle {
	iter := csNameIter.Load()
	csNameIter.Add(1)
	if iter == 0 {
		iter = 1
		csNameIter.Add(1)
	}

	return &ColumnStyle{
		name: fmt.Sprintf("cs%s", strconv.FormatUint(iter, 10)),
	}
}

func (cs *ColumnStyle) copy() *ColumnStyle {
	iter := csNameIter.Load()
	csNameIter.Add(1)
	if iter == 0 {
		iter = 1
		csNameIter.Add(1)
	}

	newStyle := *cs
	newStyle.name = fmt.Sprintf("cs%s", strconv.FormatUint(iter, 10))
	return &newStyle
}

// WithWidth sets the column width in centimeters.
// If width <= 0, the property is omitted (auto width).
func (cs *ColumnStyle) WithWidth(width float64) *ColumnStyle {
	cs.width = width
	return cs
}

// WithRelWidth sets proportional width (e.g., "1*", "2*").
// Used in multi-column layouts. Overrides absolute width if > 0.
func (cs *ColumnStyle) WithRelWidth(relWidth int) *ColumnStyle {
	if relWidth > 0 {
		cs.relWidth = relWidth
	}
	return cs
}

// WithUseOptimal enables/disables automatic width adjustment.
func (cs *ColumnStyle) WithUseOptimal(enable bool) *ColumnStyle {
	cs.useOptimal = enable
	return cs
}

// WithBackgroundColor sets the background color in "#RRGGBB" format.
// Returns error if format is invalid.
func (cs *ColumnStyle) WithBackgroundColor(color string) *ColumnStyle {
	cs.backgroundColor = color
	return cs
}

// WithBorder sets all borders (left, right, top, bottom) to the same style.
// Example: "0.5pt solid #000000"
func (cs *ColumnStyle) WithBorder(style string) *ColumnStyle {
	cs.borderLeft = style
	cs.borderRight = style
	cs.borderTop = style
	cs.borderBottom = style
	return cs
}

// WithBreakBefore sets a page/column break before this column.
// Values: "auto" (default), "column", "page".
func (cs *ColumnStyle) WithBreakBefore(value string) *ColumnStyle {
	switch value {
	case "auto", "column", "page":
		cs.breakBefore = value
	}
	return cs
}

func (cs *ColumnStyle) generate() string {
	var buf bytes.Buffer

	// Open style tag
	buf.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="table-column">`, cs.name))
	buf.WriteString(`<style:table-column-properties`)

	// Width (absolute or relative)
	if cs.relWidth > 0 {
		buf.WriteString(fmt.Sprintf(` style:rel-width="%d*"`, cs.relWidth))
	} else if cs.width > 0 {
		buf.WriteString(fmt.Sprintf(` fo:width="%.2fcm"`, cs.width))
	}

	// Optimal width
	if cs.useOptimal {
		buf.WriteString(` style:use-optimal-column-width="true"`)
	}

	// Background color
	if cs.backgroundColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:background-color="%s"`, cs.backgroundColor))
	}

	// Borders
	for _, border := range []struct{ prop, value string }{
		{"fo:border-left", cs.borderLeft},
		{"fo:border-right", cs.borderRight},
		{"fo:border-top", cs.borderTop},
		{"fo:border-bottom", cs.borderBottom},
	} {
		if border.value != "" {
			buf.WriteString(fmt.Sprintf(` %s="%s"`, border.prop, border.value))
		}
	}

	// Page/column breaks
	if cs.breakBefore != "" {
		buf.WriteString(fmt.Sprintf(` fo:break-before="%s"`, cs.breakBefore))
	}
	if cs.breakAfter != "" {
		buf.WriteString(fmt.Sprintf(` fo:break-after="%s"`, cs.breakAfter))
	}

	// Close tags
	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
