package table_style

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var cellStyleNameIter = atomic.Uint64{}

// CellStyle represents style properties for individual table cells in ODF documents.
// It controls visual appearance, padding, borders, and text alignment within cells.
type CellStyle struct {
	name            string // Unique identifier for this style
	backgroundColor string // Cell background color (hex or color name)
	border          string // border specification (e.g., "0.05pt solid #000000")
	borderTop       string // Top border override
	borderBottom    string // Bottom border override
	borderLeft      string // Left border override
	borderRight     string // Right border override
	padding         string // padding around cell content (e.g., "0.2cm")
	paddingTop      string // Top padding override
	paddingBottom   string // Bottom padding override
	paddingLeft     string // Left padding override
	paddingRight    string // Right padding override
	textAlign       string // Horizontal text alignment ("left", "center", "right")
	verticalAlign   string // Vertical text alignment ("top", "middle", "bottom")
	wrapOption      string // Text wrapping ("wrap", "no-wrap")
	rotationAngle   int    // Text rotation angle (0-360 degrees)
	shrinkToFit     bool   // Whether to shrink content to fit cell
	repeatContent   bool   // Whether to repeat content when spanning
}

func NewCellStyle() *CellStyle {
	iter := cellStyleNameIter.Load()
	cellStyleNameIter.Add(1)
	if iter == 0 {
		iter = 1
		cellStyleNameIter.Add(1)
	}

	return &CellStyle{
		name: fmt.Sprintf("TblCellS%s", strconv.FormatUint(iter, 10)),
	}
}

// WithBackground sets the cell background color.
// Accepts hex values (#RRGGBB) or color names.
// Example:
//
//	cellStyle.WithBackground("#E6E6E6")
func (c *CellStyle) WithBackground(color string) *CellStyle {
	c.backgroundColor = color
	return c
}

// WithBorder sets all cell borders uniformly.
// Format: "width style color" (e.g., "0.05pt solid #000000").
// Example:
//
//	cellStyle.WithBorder("0.5pt solid #CCCCCC")
func (c *CellStyle) WithBorder(border string) *CellStyle {
	c.border = border
	return c
}

// WithIndividualBorders sets borders independently for each side.
// Example:
//
//	cellStyle.WithIndividualBorders("1pt solid #000", "none", "1pt solid #000", "none")
func (c *CellStyle) WithIndividualBorders(top, bottom, left, right string) *CellStyle {
	c.borderTop = top
	c.borderBottom = bottom
	c.borderLeft = left
	c.borderRight = right
	return c
}

// WithPadding sets uniform padding for all sides.
// Example:
//
//	cellStyle.WithPadding("0.2cm")
func (c *CellStyle) WithPadding(padding string) *CellStyle {
	c.padding = padding
	return c
}

// WithIndividualPadding sets padding independently for each side.
// Example:
//
//	cellStyle.WithIndividualPadding("0.3cm", "0.1cm", "0.2cm", "0.2cm")
func (c *CellStyle) WithIndividualPadding(top, bottom, left, right string) *CellStyle {
	c.paddingTop = top
	c.paddingBottom = bottom
	c.paddingLeft = left
	c.paddingRight = right
	return c
}

// WithAlignment sets both horizontal and vertical text alignment.
// Example:
//
//	cellStyle.WithAlignment("center", "middle")
func (c *CellStyle) WithAlignment(horizontal, vertical string) *CellStyle {
	c.textAlign = horizontal
	c.verticalAlign = vertical
	return c
}

// WithTextWrap controls text wrapping behavior.
// Valid values: "wrap", "no-wrap".
// Example:
//
//	cellStyle.WithTextWrap("wrap")
func (c *CellStyle) WithTextWrap(wrapOption string) *CellStyle {
	c.wrapOption = wrapOption
	return c
}

// WithTextRotation sets the text rotation angle.
// Angle should be between 0 and 360 degrees.
// Example:
//
//	cellStyle.WithTextRotation(90)
func (c *CellStyle) WithTextRotation(angle int) *CellStyle {
	c.rotationAngle = angle
	return c
}

// WithShrinkToFit enables content shrinking to fit cell.
// Example:
//
//	cellStyle.WithShrinkToFit(true)
func (c *CellStyle) WithShrinkToFit(shrink bool) *CellStyle {
	c.shrinkToFit = shrink
	return c
}

// GetName returns style name
func (c *CellStyle) GetName() string {
	return c.name
}

// Generate produces the XML representation of the cell style using bytes.Buffer.
// The output follows ODF 1.2 specification for table-cell-properties.
func (c *CellStyle) Generate() string {
	var buf bytes.Buffer

	// Style declaration start
	buf.WriteString(`<style:style style:name="`)
	buf.WriteString(c.name)
	buf.WriteString(`" style:family="table-cell">`)
	buf.WriteString(`<style:table-cell-properties`)

	// Background color
	if c.backgroundColor != "" {
		buf.WriteString(` fo:background-color="`)
		buf.WriteString(c.backgroundColor)
		buf.WriteString(`"`)
	}

	// Border properties
	if c.border != "" {
		buf.WriteString(` fo:border="`)
		buf.WriteString(c.border)
		buf.WriteString(`"`)
	}
	if c.borderTop != "" {
		buf.WriteString(` fo:border-top="`)
		buf.WriteString(c.borderTop)
		buf.WriteString(`"`)
	}
	if c.borderBottom != "" {
		buf.WriteString(` fo:border-bottom="`)
		buf.WriteString(c.borderBottom)
		buf.WriteString(`"`)
	}
	if c.borderLeft != "" {
		buf.WriteString(` fo:border-left="`)
		buf.WriteString(c.borderLeft)
		buf.WriteString(`"`)
	}
	if c.borderRight != "" {
		buf.WriteString(` fo:border-right="`)
		buf.WriteString(c.borderRight)
		buf.WriteString(`"`)
	}

	// Padding properties
	if c.padding != "" {
		buf.WriteString(` fo:padding="`)
		buf.WriteString(c.padding)
		buf.WriteString(`"`)
	}
	if c.paddingTop != "" {
		buf.WriteString(` fo:padding-top="`)
		buf.WriteString(c.paddingTop)
		buf.WriteString(`"`)
	}
	if c.paddingBottom != "" {
		buf.WriteString(` fo:padding-bottom="`)
		buf.WriteString(c.paddingBottom)
		buf.WriteString(`"`)
	}
	if c.paddingLeft != "" {
		buf.WriteString(` fo:padding-left="`)
		buf.WriteString(c.paddingLeft)
		buf.WriteString(`"`)
	}
	if c.paddingRight != "" {
		buf.WriteString(` fo:padding-right="`)
		buf.WriteString(c.paddingRight)
		buf.WriteString(`"`)
	}

	// Text alignment
	if c.textAlign != "" {
		buf.WriteString(` fo:text-align="`)
		buf.WriteString(c.textAlign)
		buf.WriteString(`"`)
	}
	if c.verticalAlign != "" {
		buf.WriteString(` style:vertical-align="`)
		buf.WriteString(c.verticalAlign)
		buf.WriteString(`"`)
	}

	// Text wrapping
	if c.wrapOption != "" {
		buf.WriteString(` fo:wrap-option="`)
		buf.WriteString(c.wrapOption)
		buf.WriteString(`"`)
	}

	// Text rotation
	if c.rotationAngle != 0 {
		buf.WriteString(` style:rotation-angle="`)
		buf.WriteString(fmt.Sprintf("%d", c.rotationAngle))
		buf.WriteString(`"`)
	}

	// Content options
	if c.shrinkToFit {
		buf.WriteString(` style:shrink-to-fit="true"`)
	}
	if c.repeatContent {
		buf.WriteString(` style:repeat-content="true"`)
	}

	// Close tags
	buf.WriteString(`/>`)
	buf.WriteString(`</style:style>`)

	return buf.String()
}
