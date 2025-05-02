package ods

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var ssNameIter = atomic.Uint64{}

// SheetStyle defines formatting properties for an entire spreadsheet in ODS.
// It controls global sheet settings like margins, print layout, and background.
// Example:
//
//	style := NewSheetStyle("MySheetStyle").
//			WithMarginLeft("1cm").
//			WithPrintOrientation("landscape")
type SheetStyle struct {
	name string // Style name (referenced in table:table)

	// Writing Mode
	writingMode string // Valid values: "lr-tb", "rl-tb", "tb-rl", "tb-lr", "page"

	// Page Layout
	marginLeft   string // Left margin (e.g., "1cm")
	marginRight  string // Right margin
	marginTop    string // Top margin
	marginBottom string // Bottom margin
	pageWidth    string // Page width (e.g., "21cm")
	pageHeight   string // Page height
	orientation  string // "portrait" or "landscape"
	printScale   int    // Print scaling percentage (10-100)
	displayGrid  bool   // Whether to show grid lines
	printGrid    bool   // Whether to print grid lines

	// Background
	backgroundColor string // Hex color ("#RRGGBB")
	backgroundImage string // Image URL (internal path in ODS package)

	// Headers/Footers
	headerLeft   string // Left header text
	footerCenter string // Center footer text
}

// NewSheetStyle creates a new empty SheetStyle.
func NewSheetStyle() *SheetStyle {
	iter := ssNameIter.Load()
	ssNameIter.Add(1)
	if iter == 0 {
		iter = 1
		ssNameIter.Add(1)
	}

	return &SheetStyle{
		name:        fmt.Sprintf("ss%s", strconv.FormatUint(iter, 10)),
		orientation: "portrait", // Default
		writingMode: "lr-tb",    // Default
		printScale:  100,        // Default
		displayGrid: true,       // Default
	}
}

func (ss *SheetStyle) copy() *SheetStyle {
	iter := ssNameIter.Load()
	ssNameIter.Add(1)
	if iter == 0 {
		iter = 1
		ssNameIter.Add(1)
	}

	newStyle := *ss
	newStyle.name = fmt.Sprintf("ss%s", strconv.FormatUint(iter, 10))
	return &newStyle
}

// --- Page Layout Setters ---

// WithMargins sets all margins to the same value.
// Example: WithMargins("1.5cm")
func (ss *SheetStyle) WithMargins(margin string) *SheetStyle {
	ss.marginLeft = margin
	ss.marginRight = margin
	ss.marginTop = margin
	ss.marginBottom = margin
	return ss
}

// WithPageSize sets the page dimensions.
// Standard sizes: "21cm x 29.7cm" (A4), "21.6cm x 27.9cm" (Letter)
func (ss *SheetStyle) WithPageSize(width, height string) *SheetStyle {
	ss.pageWidth = width
	ss.pageHeight = height
	return ss
}

// WithPrintOrientation sets the page orientation.
// Valid values: "portrait", "landscape"
func (ss *SheetStyle) WithPrintOrientation(orientation string) *SheetStyle {
	switch orientation {
	case "portrait", "landscape":
		ss.orientation = orientation
	}
	return ss
}

// WithPrintScale sets the print scaling percentage (10-100).
func (ss *SheetStyle) WithPrintScale(scale int) *SheetStyle {
	if scale >= 10 && scale <= 100 {
		ss.printScale = scale
	}
	return ss
}

// --- Background Setters ---

// WithBackgroundColor sets the sheet background color ("#RRGGBB").
func (ss *SheetStyle) WithBackgroundColor(color string) *SheetStyle {
	ss.backgroundColor = color
	return ss
}

// --- Grid Setters ---

// WithGridDisplay controls visibility of grid lines on screen.
func (ss *SheetStyle) WithGridDisplay(show bool) *SheetStyle {
	ss.displayGrid = show
	return ss
}

// WithGridPrint controls whether grid lines are printed.
func (ss *SheetStyle) WithGridPrint(print bool) *SheetStyle {
	ss.printGrid = print
	return ss
}

// WithWritingMode sets the text flow direction.
// Valid values:
//   - "lr-tb" : Left-to-right, top-to-bottom (default)
//   - "rl-tb" : Right-to-left, top-to-bottom (Hebrew/Arabic)
//   - "tb-rl" : Top-to-bottom, right-to-left (Asian vertical)
//   - "tb-lr" : Top-to-bottom, left-to-right (Mongolian)
//   - "page"  : Inherits from page layout
func (ss *SheetStyle) WithWritingMode(mode string) *SheetStyle {
	switch mode {
	case "lr-tb", "rl-tb", "tb-rl", "tb-lr", "page":
		ss.writingMode = mode
	default:
	}

	return ss
}

func (ss *SheetStyle) generate() string {
	var buf bytes.Buffer

	// Open style tag
	buf.WriteString(fmt.Sprintf(`<style:style style:name="%s" style:family="table" style:master-page-name="Default">`, ss.name))
	buf.WriteString(`<style:table-properties`)

	// Table properties
	buf.WriteString(fmt.Sprintf(` style:writing-mode="%s"`, ss.writingMode))

	if ss.displayGrid {
		buf.WriteString(` table:display="true"`)
	} else {
		buf.WriteString(` table:display="false"`)
	}

	// Background
	if ss.backgroundColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:background-color="%s"`, ss.backgroundColor))
	}

	// Close table-properties and open page layout properties
	buf.WriteString(`/><style:page-layout-properties`)

	// Page size
	if ss.pageWidth != "" && ss.pageHeight != "" {
		buf.WriteString(fmt.Sprintf(` fo:page-width="%s" fo:page-height="%s"`, ss.pageWidth, ss.pageHeight))
	}

	// Margins
	if ss.marginLeft != "" {
		buf.WriteString(fmt.Sprintf(` fo:margin-left="%s"`, ss.marginLeft))
	}
	if ss.marginRight != "" {
		buf.WriteString(fmt.Sprintf(` fo:margin-right="%s"`, ss.marginRight))
	}
	if ss.marginTop != "" {
		buf.WriteString(fmt.Sprintf(` fo:margin-top="%s"`, ss.marginTop))
	}
	if ss.marginBottom != "" {
		buf.WriteString(fmt.Sprintf(` fo:margin-bottom="%s"`, ss.marginBottom))
	}

	// Print settings
	buf.WriteString(fmt.Sprintf(` style:print-orientation="%s"`, ss.orientation))
	if ss.printScale != 100 {
		buf.WriteString(fmt.Sprintf(` style:scale-to="%d"`, ss.printScale))
	}
	if ss.printGrid {
		buf.WriteString(` style:print="true"`)
	}

	// Close all tags
	buf.WriteString(`/></style:style>`)

	return buf.String()
}
