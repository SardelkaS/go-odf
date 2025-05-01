package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var psNameIter = atomic.Uint64{}

// PageStyle represents page style
type PageStyle struct {
	name             string
	width            string
	height           string
	margins          pageMargins
	backgroundColor  string
	headerFooter     pageHeaderFooter
	writingMode      string
	printOrientation string
}

type pageMargins struct {
	top    string
	bottom string
	left   string
	right  string
}

type pageHeaderFooter struct {
	headerDistance string
	footerDistance string
}

// NewPageStyle creates new PageStyle with default settings
//
//	 Default settings:
//		width:  "21cm"
//		height: "29.7cm"
//		margins:
//			top:    "2cm"
//			bottom: "2cm"
//			left:   "2cm"
//			right:  "2cm"
//		orientation:      "portrait"
//		writingMode:      "lr-tb"
//		printOrientation: "none"
//		headerFooter:
//			headerDistance: "0.5cm"
//			footerDistance: "0.5cm"
func NewPageStyle() *PageStyle {
	iter := psNameIter.Load()
	psNameIter.Add(1)
	if iter == 0 {
		iter = 1
		psNameIter.Add(1)
	}

	return &PageStyle{
		name:   fmt.Sprintf("PgS%s", strconv.FormatUint(iter, 10)),
		width:  "21cm",
		height: "29.7cm",
		margins: pageMargins{
			top:    "2cm",
			bottom: "2cm",
			left:   "2cm",
			right:  "2cm",
		},
		writingMode:      "lr-tb",
		printOrientation: "none",
		headerFooter: pageHeaderFooter{
			headerDistance: "0.5cm",
			footerDistance: "0.5cm",
		},
	}
}

// SetWidth sets the page width in valid CSS units (e.g., "21cm", "8.5in").
// Common values for A4: "21cm", Letter: "8.5in".
// Must be positive value with unit.
func (ps *PageStyle) SetWidth(width string) {
	ps.width = width
}

// SetHeight sets the page height in valid CSS units (e.g., "29.7cm", "11in").
// Common values for A4: "29.7cm", Letter: "11in".
// Must be positive value with unit.
func (ps *PageStyle) SetHeight(height string) {
	ps.height = height
}

// SetMargins sets all page margins at once.
// Parameters should be in valid CSS units (e.g., "2cm", "0.5in").
// Sets margins in order: top, bottom, left, right.
func (ps *PageStyle) SetMargins(top, bottom, left, right string) {
	ps.margins = pageMargins{
		top:    top,
		bottom: bottom,
		left:   left,
		right:  right,
	}
}

// SetMarginTop sets only the top margin in CSS units.
// Typically larger than other margins for title pages.
func (ps *PageStyle) SetMarginTop(margin string) {
	ps.margins.top = margin
}

// SetMarginBottom sets only the bottom margin in CSS units.
// Often contains footer space.
func (ps *PageStyle) SetMarginBottom(margin string) {
	ps.margins.bottom = margin
}

// SetMarginLeft sets only the left margin in CSS units.
// Binding side in printed documents may need larger margin.
func (ps *PageStyle) SetMarginLeft(margin string) {
	ps.margins.left = margin
}

// SetMarginRight sets only the right margin in CSS units.
// Typically matches left margin unless for binding.
func (ps *PageStyle) SetMarginRight(margin string) {
	ps.margins.right = margin
}

// SetBackgroundColor sets the page background color.
// Format: hexadecimal "#RRGGBB" or named colors "white", "lightblue".
// Empty string means transparent background.
func (ps *PageStyle) SetBackgroundColor(color string) {
	ps.backgroundColor = color
}

// SetHeaderDistance sets the distance from page top to Header content.
// Value in CSS units. Affects Header positioning.
func (ps *PageStyle) SetHeaderDistance(distance string) {
	ps.headerFooter.headerDistance = distance
}

// SetFooterDistance sets the distance from page bottom to footer content.
// Value in CSS units. Affects footer positioning.
func (ps *PageStyle) SetFooterDistance(distance string) {
	ps.headerFooter.footerDistance = distance
}

// SetWritingMode sets the text flow direction.
// Valid values: "lr-tb" (left-right, top-bottom, default),
// "tb-rl" (top-bottom, right-left for vertical scripts).
func (ps *PageStyle) SetWritingMode(mode string) {
	switch mode {
	case "lr-tb", "tb-rl":
		ps.writingMode = mode
	default:
		ps.writingMode = "lr-tb"
	}
}

// SetPrintOrientation sets the forced print orientation.
// Valid values: "none" (default), "portrait", "landscape".
// Overrides automatic orientation detection when printing.
func (ps *PageStyle) SetPrintOrientation(orientation string) {
	switch orientation {
	case "none", "portrait", "landscape":
		if orientation != ps.printOrientation {
			ps.width, ps.height = ps.height, ps.width
		}
		ps.printOrientation = orientation
	default:
		if ps.printOrientation == "landscape" {
			ps.width, ps.height = ps.height, ps.width
		}

		ps.printOrientation = "none"
	}
}

func (ps *PageStyle) generate() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`<style:page-layout style:name="%s">`, ps.name))
	buf.WriteString(`<style:page-layout-properties`)

	buf.WriteString(fmt.Sprintf(` fo:page-width="%s"`, ps.width))
	buf.WriteString(fmt.Sprintf(` fo:page-height="%s"`, ps.height))
	buf.WriteString(fmt.Sprintf(` style:writing-mode="%s"`, ps.writingMode))
	buf.WriteString(fmt.Sprintf(` style:print-orientation="%s"`, ps.printOrientation))

	buf.WriteString(fmt.Sprintf(` fo:margin-top="%s"`, ps.margins.top))
	buf.WriteString(fmt.Sprintf(` fo:margin-bottom="%s"`, ps.margins.bottom))
	buf.WriteString(fmt.Sprintf(` fo:margin-left="%s"`, ps.margins.left))
	buf.WriteString(fmt.Sprintf(` fo:margin-right="%s"`, ps.margins.right))

	if ps.backgroundColor != "" {
		buf.WriteString(fmt.Sprintf(` fo:background-color="%s"`, ps.backgroundColor))
	}

	buf.WriteString(`>`)

	buf.WriteString(fmt.Sprintf(`<style:Header-style><style:Header-footer-properties fo:min-height="0.5cm" fo:margin-left="0cm" fo:margin-right="0cm" fo:margin-bottom="%s"/></style:Header-style>`, ps.headerFooter.headerDistance))
	buf.WriteString(fmt.Sprintf(`<style:footer-style><style:Header-footer-properties fo:min-height="0.5cm" fo:margin-left="0cm" fo:margin-right="0cm" fo:margin-top="%s"/></style:footer-style>`, ps.headerFooter.footerDistance))

	buf.WriteString(`</style:page-layout-properties>`)
	buf.WriteString(`</style:page-layout>`)

	return buf.String()
}
