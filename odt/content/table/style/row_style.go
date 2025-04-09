package table_style

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var rowStyleNameIter = atomic.Uint64{}

// RowStyle represents style properties for table rows in ODF documents.
// It controls visual appearance and layout of rows within a table.
type RowStyle struct {
	name             string // Unique identifier for this style
	height           string // Row height (e.g., "1.5cm", "0.5in")
	minHeight        string // Minimum row height
	backgroundColor  string // Row background color (hex or color name)
	breakBefore      string // Page break before row ("auto", "column", "page", "even-page", etc.)
	breakAfter       string // Page break after row
	keepTogether     bool   // Whether to keep row contents on same page
	useOptimalHeight bool   // Whether to automatically adjust row height
}

func NewRowStyle() *RowStyle {
	iter := rowStyleNameIter.Load()
	rowStyleNameIter.Add(1)
	if iter == 0 {
		iter = 1
		rowStyleNameIter.Add(1)
	}

	return &RowStyle{
		name: fmt.Sprintf("TblRS%s", strconv.FormatUint(iter, 10)),
	}
}

// WithHeight sets the row height with units.
// Supported units: cm, mm, in, pt, pc, px.
// Example:
//
//	rowStyle.WithHeight("0.8cm")
func (r *RowStyle) WithHeight(height string) *RowStyle {
	r.height = height
	return r
}

// WithMinHeight sets the minimum row height.
// Example:
//
//	rowStyle.WithMinHeight("0.5cm")
func (r *RowStyle) WithMinHeight(minHeight string) *RowStyle {
	r.minHeight = minHeight
	return r
}

// WithBackground sets the row background color.
// Accepts hex values (#RRGGBB) or color names.
// Example:
//
//	rowStyle.WithBackground("#F5F5F5")
func (r *RowStyle) WithBackground(color string) *RowStyle {
	r.backgroundColor = color
	return r
}

// WithBreakBefore controls page/column breaks before the row.
// Valid values: "auto", "column", "page", "even-page", "odd-page".
// Example:
//
//	rowStyle.WithBreakBefore("page")
func (r *RowStyle) WithBreakBefore(breakType string) *RowStyle {
	r.breakBefore = breakType
	return r
}

// WithKeepTogether ensures all row content stays on same page.
// Example:
//
//	rowStyle.WithKeepTogether(true)
func (r *RowStyle) WithKeepTogether(keep bool) *RowStyle {
	r.keepTogether = keep
	return r
}

// WithOptimalHeight enables automatic row height adjustment.
// Example:
//
//	rowStyle.WithOptimalHeight(true)
func (r *RowStyle) WithOptimalHeight(optimal bool) *RowStyle {
	r.useOptimalHeight = optimal
	return r
}

// GetName returns style name
func (r *RowStyle) GetName() string {
	return r.name
}

// Generate produces the XML representation of the row style using bytes.Buffer.
// The output follows ODF 1.2 specification for table-row-properties.
func (r *RowStyle) Generate() string {
	var buf bytes.Buffer

	// Style declaration start
	buf.WriteString(`<style:style style:name="`)
	buf.WriteString(r.name)
	buf.WriteString(`" style:family="table-row">`)
	buf.WriteString(`<style:table-row-properties`)

	// Height properties
	if r.height != "" {
		buf.WriteString(` style:row-height="`)
		buf.WriteString(r.height)
		buf.WriteString(`"`)
	}

	if r.minHeight != "" {
		buf.WriteString(` style:min-row-height="`)
		buf.WriteString(r.minHeight)
		buf.WriteString(`"`)
	}

	// Background color
	if r.backgroundColor != "" {
		buf.WriteString(` fo:background-color="`)
		buf.WriteString(r.backgroundColor)
		buf.WriteString(`"`)
	}

	// Page break controls
	if r.breakBefore != "" {
		buf.WriteString(` fo:break-before="`)
		buf.WriteString(r.breakBefore)
		buf.WriteString(`"`)
	}

	if r.breakAfter != "" {
		buf.WriteString(` fo:break-after="`)
		buf.WriteString(r.breakAfter)
		buf.WriteString(`"`)
	}

	// Content keeping
	if r.keepTogether {
		buf.WriteString(` fo:keep-together="true"`)
	}

	// Automatic height adjustment
	if r.useOptimalHeight {
		buf.WriteString(` style:use-optimal-row-height="true"`)
	}

	// Close tags
	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
