package ods

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var rsNameIter = atomic.Uint64{}

// RowStyle defines formatting properties for a table row in ODS.
// Example:
//
//	style := NewRowStyle("HeaderRow").
//			WithHeight(1.2). // 1.2cm
//			WithBackgroundColor("#EEEEEE").
//			WithBreakBefore("page")
type RowStyle struct {
	name            string  // Style name (referenced in table:table-row)
	height          float64 // Row height in centimeters (0 = auto)
	minHeight       float64 // Minimum height (cm)
	useOptimal      bool    // Auto-adjust height to fit content
	backgroundColor string  // Hex color ("#RRGGBB")
	breakBefore     string  // Break before row: "auto", "column", "page"
	breakAfter      string  // Break after row
	keepTogether    bool    // Prevent row splitting across pages
	borderTop       string  // Top border (e.g., "0.5pt solid #000000")
	borderBottom    string  // Bottom border
}

// NewRowStyle creates a new empty RowStyle.
func NewRowStyle() *RowStyle {
	iter := rsNameIter.Load()
	rsNameIter.Add(1)
	if iter == 0 {
		iter = 1
		rsNameIter.Add(1)
	}

	return &RowStyle{
		name: fmt.Sprintf("rs%s", strconv.FormatUint(iter, 10)),
	}
}

func (rs *RowStyle) copy() *RowStyle {
	iter := rsNameIter.Load()
	rsNameIter.Add(1)
	if iter == 0 {
		iter = 1
		rsNameIter.Add(1)
	}

	newStyle := *rs
	newStyle.name = fmt.Sprintf("rs%s", strconv.FormatUint(iter, 10))
	return &newStyle
}

// WithHeight sets the row height in centimeters.
// If height <= 0, the property is omitted (auto height).
func (rs *RowStyle) WithHeight(height float64) *RowStyle {
	rs.height = height
	return rs
}

// WithMinHeight sets the minimum row height in centimeters.
func (rs *RowStyle) WithMinHeight(height float64) *RowStyle {
	rs.minHeight = height
	return rs
}

// WithUseOptimal enables/disables automatic height adjustment.
func (rs *RowStyle) WithUseOptimal(enable bool) *RowStyle {
	rs.useOptimal = enable
	return rs
}

// WithBackgroundColor sets the background color in "#RRGGBB" format.
// Returns error if format is invalid.
func (rs *RowStyle) WithBackgroundColor(color string) *RowStyle {
	rs.backgroundColor = color
	return rs
}

// WithBreakBefore sets a page/column break before this row.
// Values: "auto" (default), "column", "page".
func (rs *RowStyle) WithBreakBefore(value string) *RowStyle {
	switch value {
	case "auto", "column", "page":
		rs.breakBefore = value
	}
	return rs
}

// WithKeepTogether prevents the row from splitting across pages.
func (rs *RowStyle) WithKeepTogether(keep bool) *RowStyle {
	rs.keepTogether = keep
	return rs
}

// WithBorderTop sets the top border style (e.g., "0.5pt solid #000000").
func (rs *RowStyle) WithBorderTop(style string) *RowStyle {
	rs.borderTop = style
	return rs
}

// WithBorderBottom sets the bottom border style.
func (rs *RowStyle) WithBorderBottom(style string) *RowStyle {
	rs.borderBottom = style
	return rs
}

func (rs *RowStyle) generate() string {
	var buf bytes.Buffer

	// Open style tag
	buf.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="table-row">`, rs.name))
	buf.WriteString(`<style:table-row-properties`)

	// Height
	if rs.height > 0 {
		buf.WriteString(fmt.Sprintf(` fo:height="%.2fcm"`, rs.height))
	}

	// Minimum height
	if rs.minHeight > 0 {
		buf.WriteString(fmt.Sprintf(` style:min-row-height="%.2fcm"`, rs.minHeight))
	}

	// Optimal height
	if rs.useOptimal {
		buf.WriteString(` style:use-optimal-row-height="true"`)
	}

	// Background color
	if rs.backgroundColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:background-color="%s"`, rs.backgroundColor))
	}

	// Breaks
	if rs.breakBefore != "" {
		buf.WriteString(fmt.Sprintf(` fo:break-before="%s"`, rs.breakBefore))
	}
	if rs.breakAfter != "" {
		buf.WriteString(fmt.Sprintf(` fo:break-after="%s"`, rs.breakAfter))
	}

	// Keep together
	if rs.keepTogether {
		buf.WriteString(` fo:keep-together="true"`)
	}

	// Borders
	if rs.borderTop != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-top="%s"`, rs.borderTop))
	}
	if rs.borderBottom != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-bottom="%s"`, rs.borderBottom))
	}

	// Close tags
	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
