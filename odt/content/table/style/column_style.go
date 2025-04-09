package table_style

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var columnStyleNameIter = atomic.Uint64{}

// ColumnStyle represents style properties for table columns in ODF documents.
// It controls visual appearance and layout of columns within a table.
type ColumnStyle struct {
	name            string // Unique identifier for this style
	width           string // Column width (e.g., "3cm", "2.5in")
	relWidth        string // Relative width as percentage (e.g., "25*")
	breakBefore     string // Column break before ("auto", "column", "page")
	breakAfter      string // Column break after
	backgroundColor string // Column background color
	border          string // Column border specification
	optimalWidth    bool   // Whether to use optimal column width
}

func NewColumnStyle() *ColumnStyle {
	iter := columnStyleNameIter.Load()
	columnStyleNameIter.Add(1)
	if iter == 0 {
		iter = 1
		columnStyleNameIter.Add(1)
	}

	return &ColumnStyle{
		name: fmt.Sprintf("TblCS%s", strconv.FormatUint(iter, 10)),
	}
}

// WithWidth sets the absolute column width with units.
// Supported units: cm, mm, in, pt, pc, px.
// Example:
//
//	colStyle.WithWidth("3.5cm")
func (c *ColumnStyle) WithWidth(width string) *ColumnStyle {
	c.width = width
	return c
}

// WithRelativeWidth sets the relative column width.
// Format: "N*" where N is the relative weight (e.g., "2*" for twice default width).
// Example:
//
//	colStyle.WithRelativeWidth("3*")
func (c *ColumnStyle) WithRelativeWidth(relWidth string) *ColumnStyle {
	c.relWidth = relWidth
	return c
}

// WithBackground sets the column background color.
// Accepts hex values (#RRGGBB) or color names.
// Example:
//
//	colStyle.WithBackground("#F8F8F8")
func (c *ColumnStyle) WithBackground(color string) *ColumnStyle {
	c.backgroundColor = color
	return c
}

// WithBorder sets the column border specification.
// Format: "width style color" (e.g., "0.5pt solid #000000").
// Example:
//
//	colStyle.WithBorder("0.05cm solid #CCCCCC")
func (c *ColumnStyle) WithBorder(border string) *ColumnStyle {
	c.border = border
	return c
}

// WithBreakBefore controls breaks before the column.
// Valid values: "auto", "column", "page".
// Example:
//
//	colStyle.WithBreakBefore("column")
func (c *ColumnStyle) WithBreakBefore(breakType string) *ColumnStyle {
	c.breakBefore = breakType
	return c
}

// WithOptimalWidth enables automatic column width adjustment.
// Example:
//
//	colStyle.WithOptimalWidth(true)
func (c *ColumnStyle) WithOptimalWidth(optimal bool) *ColumnStyle {
	c.optimalWidth = optimal
	return c
}

// GetName returns style name
func (c *ColumnStyle) GetName() string {
	return c.name
}

// Generate produces the XML representation of the column style using bytes.Buffer.
// The output follows ODF 1.2 specification for table-column-properties.
func (c *ColumnStyle) Generate() string {
	var buf bytes.Buffer

	// Style declaration start
	buf.WriteString(`<style:style style:name="`)
	buf.WriteString(c.name)
	buf.WriteString(`" style:family="table-column">`)
	buf.WriteString(`<style:table-column-properties`)

	// Width properties
	if c.width != "" {
		buf.WriteString(` style:column-width="`)
		buf.WriteString(c.width)
		buf.WriteString(`"`)
	}

	if c.relWidth != "" {
		buf.WriteString(` style:rel-column-width="`)
		buf.WriteString(c.relWidth)
		buf.WriteString(`"`)
	}

	// Background color
	if c.backgroundColor != "" {
		buf.WriteString(` fo:background-color="`)
		buf.WriteString(c.backgroundColor)
		buf.WriteString(`"`)
	}

	// Border specification
	if c.border != "" {
		buf.WriteString(` fo:border="`)
		buf.WriteString(c.border)
		buf.WriteString(`"`)
	}

	// Page/column break controls
	if c.breakBefore != "" {
		buf.WriteString(` fo:break-before="`)
		buf.WriteString(c.breakBefore)
		buf.WriteString(`"`)
	}

	if c.breakAfter != "" {
		buf.WriteString(` fo:break-after="`)
		buf.WriteString(c.breakAfter)
		buf.WriteString(`"`)
	}

	// Automatic width adjustment
	if c.optimalWidth {
		buf.WriteString(` style:use-optimal-column-width="true"`)
	}

	// Close tags
	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
