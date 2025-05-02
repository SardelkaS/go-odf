package ods

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var cellsNameIter = atomic.Uint64{}

// CellStyle defines formatting properties for a table cell in ODS.
// Example:
//
//	style := NewCellStyle("ResultCell").
//			WithBackgroundColor("#FFFF00"). // Yellow
//			WithFontName("Arial").
//			WithHorizontalAlign("center")
type CellStyle struct {
	name string // Style name (referenced in table:table-cell)

	// Table Cell Properties
	backgroundColor string // Hex color ("#RRGGBB")
	borderTop       string // Top border (e.g., "0.5pt solid #000000")
	borderBottom    string // Bottom border
	borderLeft      string // Left border
	borderRight     string // Right border
	padding         string // Padding (e.g., "0.1cm")
	wrapOption      string // "wrap", "no-wrap"

	// Text Properties
	fontName        string // Font family (e.g., "Liberation Sans")
	fontSize        string // Font size (e.g., "12pt")
	fontWeight      string // "normal", "bold"
	fontStyle       string // "normal", "italic"
	textColor       string // Text color ("#RRGGBB")
	horizontalAlign string // "left", "center", "right", "justify"
	verticalAlign   string // "top", "middle", "bottom"
	textRotation    int    // Rotation angle (0-360)
}

// NewCellStyle creates a new empty CellStyle.
func NewCellStyle() *CellStyle {
	iter := cellsNameIter.Load()
	cellsNameIter.Add(1)
	if iter == 0 {
		iter = 1
		cellsNameIter.Add(1)
	}

	return &CellStyle{
		name:            fmt.Sprintf("cellS%s", strconv.FormatUint(iter, 10)),
		wrapOption:      "wrap",   // Default
		horizontalAlign: "left",   // Default
		verticalAlign:   "middle", // Default
	}
}

func (cs *CellStyle) copy() *CellStyle {
	iter := cellsNameIter.Load()
	cellsNameIter.Add(1)
	if iter == 0 {
		iter = 1
		cellsNameIter.Add(1)
	}

	newStyle := *cs
	newStyle.name = fmt.Sprintf("cells%s", strconv.FormatUint(iter, 10))
	return &newStyle
}

// --- Table Cell Property Setters ---

// WithBackgroundColor sets the cell background color ("#RRGGBB").
func (cs *CellStyle) WithBackgroundColor(color string) *CellStyle {
	cs.backgroundColor = color
	return cs
}

// WithBorder sets all borders to the same style.
// Example: "0.5pt solid #000000"
func (cs *CellStyle) WithBorder(style string) *CellStyle {
	cs.borderTop = style
	cs.borderBottom = style
	cs.borderLeft = style
	cs.borderRight = style
	return cs
}

// WithBorderTop sets the top border style.
func (cs *CellStyle) WithBorderTop(style string) *CellStyle {
	cs.borderTop = style
	return cs
}

// WithPadding sets the cell padding (e.g., "0.1cm").
func (cs *CellStyle) WithPadding(padding string) *CellStyle {
	cs.padding = padding
	return cs
}

// WithWrapOption sets text wrapping ("wrap" or "no-wrap").
func (cs *CellStyle) WithWrapOption(option string) *CellStyle {
	switch option {
	case "wrap", "no-wrap":
		cs.wrapOption = option
	}
	return cs
}

// --- Text Property Setters ---

// WithFontName sets the font family (e.g., "Arial").
func (cs *CellStyle) WithFontName(name string) *CellStyle {
	cs.fontName = name
	return cs
}

// WithFontSize sets the font size (e.g., "12pt").
func (cs *CellStyle) WithFontSize(size string) *CellStyle {
	cs.fontSize = size
	return cs
}

// WithBold enables/disables bold text.
func (cs *CellStyle) WithBold(enable bool) *CellStyle {
	if enable {
		cs.fontWeight = "bold"
	} else {
		cs.fontWeight = "normal"
	}
	return cs
}

// WithItalic enables/disables italic text.
func (cs *CellStyle) WithItalic(enable bool) *CellStyle {
	if enable {
		cs.fontStyle = "italic"
	} else {
		cs.fontStyle = "normal"
	}
	return cs
}

// WithTextColor sets the text color ("#RRGGBB").
func (cs *CellStyle) WithTextColor(color string) *CellStyle {
	cs.textColor = color
	return cs
}

// WithHorizontalAlign sets horizontal alignment.
// Values: "left", "center", "right", "justify".
func (cs *CellStyle) WithHorizontalAlign(align string) *CellStyle {
	switch align {
	case "left", "center", "right", "justify":
		cs.horizontalAlign = align
	}
	return cs
}

// WithTextRotation sets text rotation angle (0-360 degrees).
func (cs *CellStyle) WithTextRotation(angle int) *CellStyle {
	if angle >= 0 && angle <= 360 {
		cs.textRotation = angle
	}
	return cs
}

func (cs *CellStyle) generate() string {
	var buf bytes.Buffer

	// Open style tag
	buf.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="table-cell">`, cs.name))
	buf.WriteString(`<style:table-cell-properties`)

	// Cell properties
	if cs.backgroundColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:background-color="%s"`, cs.backgroundColor))
	}
	if cs.borderTop != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-top="%s"`, cs.borderTop))
	}
	if cs.borderBottom != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-bottom="%s"`, cs.borderBottom))
	}
	if cs.borderLeft != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-left="%s"`, cs.borderLeft))
	}
	if cs.borderRight != "" {
		buf.WriteString(fmt.Sprintf(` fo:border-right="%s"`, cs.borderRight))
	}
	if cs.padding != "" {
		buf.WriteString(fmt.Sprintf(` fo:padding="%s"`, cs.padding))
	}
	if cs.wrapOption != "" {
		buf.WriteString(fmt.Sprintf(` fo:wrap-option="%s"`, cs.wrapOption))
	}

	// Close table-cell-properties and open text-properties
	buf.WriteString(`/><style:text-properties`)

	// Text properties
	if cs.fontName != "" {
		buf.WriteString(fmt.Sprintf(` fo:font-name="%s"`, cs.fontName))
	}
	if cs.fontSize != "" {
		buf.WriteString(fmt.Sprintf(` fo:font-size="%s"`, cs.fontSize))
	}
	if cs.fontWeight != "" {
		buf.WriteString(fmt.Sprintf(` fo:font-weight="%s"`, cs.fontWeight))
	}
	if cs.fontStyle != "" {
		buf.WriteString(fmt.Sprintf(` fo:font-style="%s"`, cs.fontStyle))
	}
	if cs.textColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:color="%s"`, cs.textColor))
	}
	if cs.horizontalAlign != "" {
		buf.WriteString(fmt.Sprintf(` fo:text-align="%s"`, cs.horizontalAlign))
	}
	if cs.verticalAlign != "" {
		buf.WriteString(fmt.Sprintf(` style:vertical-align="%s"`, cs.verticalAlign))
	}
	if cs.textRotation != 0 {
		buf.WriteString(fmt.Sprintf(` style:text-rotation-angle="%d"`, cs.textRotation))
	}

	// Close all tags
	buf.WriteString(`/></style:style>`)

	return buf.String()
}
