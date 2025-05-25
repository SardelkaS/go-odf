package ods

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var chsNameIter = atomic.Uint64{}

// ChartStyle contains styling configurations for charts
// All fields are private to enforce encapsulation
type ChartStyle struct {
	name              string
	title             string   // Chart title text
	titleFontSize     int      // Title font size in points
	titleColor        string   // Title color in hex format
	backgroundColor   string   // Chart background color
	legendPosition    string   // Where to place the legend
	showGrid          bool     // Whether to display grid lines
	gridColor         string   // Color for grid lines
	axisTitleFontSize int      // Font size for axis labels
	axisColor         string   // Color for axis lines and labels
	seriesColors      []string // Colors for each data series
	borderColor       string   // Chart border color
	borderWidth       int      // Chart border width in pixels
}

// NewChartStyle creates a new ChartStyle with default values:
//   - 12pt black title
//   - White background
//   - Legend at bottom
//   - Gray grid lines visible
//   - Standard Google colors for series
func NewChartStyle() *ChartStyle {
	iter := chsNameIter.Load()
	chsNameIter.Add(1)
	if iter == 0 {
		iter = 1
		chsNameIter.Add(1)
	}

	return &ChartStyle{
		name:              fmt.Sprintf("chs%s", strconv.FormatUint(iter, 10)),
		titleFontSize:     12,
		titleColor:        "#000000",
		backgroundColor:   "#FFFFFF",
		legendPosition:    LegendBottom,
		showGrid:          true,
		gridColor:         "#CCCCCC",
		axisTitleFontSize: 10,
		axisColor:         "#000000",
		seriesColors:      []string{"#4285F4", "#EA4335", "#FBBC05", "#34A853"},
		borderColor:       "#000000",
		borderWidth:       1,
	}
}

// WithTitle sets the chart's title text
// Returns self for method chaining
func (cs *ChartStyle) WithTitle(title string) *ChartStyle {
	cs.title = title
	return cs
}

// WithTitleFontSize sets the font size for chart title
// Size is in points (pt)
func (cs *ChartStyle) WithTitleFontSize(size int) *ChartStyle {
	cs.titleFontSize = size
	return cs
}

// WithTitleColor sets the color for chart title
// Color should be in hex format (e.g., "#FF5733")
func (cs *ChartStyle) WithTitleColor(color string) *ChartStyle {
	cs.titleColor = color
	return cs
}

// WithBackgroundColor sets the background color for the chart
func (cs *ChartStyle) WithBackgroundColor(color string) *ChartStyle {
	cs.backgroundColor = color
	return cs
}

// WithLegendPosition configures where the legend appears
// Use constants: LegendTop, LegendBottom, LegendLeft, LegendRight, LegendNone
func (cs *ChartStyle) WithLegendPosition(position string) *ChartStyle {
	cs.legendPosition = position
	return cs
}

// WithShowGrid toggles visibility of grid lines
func (cs *ChartStyle) WithShowGrid(show bool) *ChartStyle {
	cs.showGrid = show
	return cs
}

// WithGridColor sets the color for grid lines
func (cs *ChartStyle) WithGridColor(color string) *ChartStyle {
	cs.gridColor = color
	return cs
}

// WithAxisTitleFontSize sets font size for axis labels
func (cs *ChartStyle) WithAxisTitleFontSize(size int) *ChartStyle {
	cs.axisTitleFontSize = size
	return cs
}

// WithAxisColor sets color for both x and y axes
func (cs *ChartStyle) WithAxisColor(color string) *ChartStyle {
	cs.axisColor = color
	return cs
}

// WithSeriesColors defines colors for each data series
// Colors are applied to series in order
func (cs *ChartStyle) WithSeriesColors(colors []string) *ChartStyle {
	cs.seriesColors = colors
	return cs
}

// WithBorderColor sets the border color around the chart
func (cs *ChartStyle) WithBorderColor(color string) *ChartStyle {
	cs.borderColor = color
	return cs
}

// WithBorderWidth sets the border width in pixels
func (cs *ChartStyle) WithBorderWidth(width int) *ChartStyle {
	cs.borderWidth = width
	return cs
}

// generate creates XML style definition for ODF format
func (cs *ChartStyle) generate() string {
	var buf bytes.Buffer

	// Open style definition
	buf.WriteString("<style:style style:name=\"chart-style\" style:family=\"chart\">\n")
	buf.WriteString("  <style:chart-properties>\n")

	// Title configuration
	if cs.title != "" {
		buf.WriteString(fmt.Sprintf("    <chart:title text=\"%s\"/>\n", escapeXML(cs.title)))
		buf.WriteString(fmt.Sprintf("    <style:text-properties fo:color=\"%s\" fo:font-size=\"%dpt\"/>\n",
			cs.titleColor, cs.titleFontSize))
	}

	// Background and border
	buf.WriteString(fmt.Sprintf("    <style:graphic-properties draw:fill-color=\"%s\" draw:stroke=\"solid\" draw:stroke-color=\"%s\" draw:stroke-width=\"%d\"/>\n",
		cs.backgroundColor, cs.borderColor, cs.borderWidth))

	// Legend position
	if cs.legendPosition != LegendNone {
		buf.WriteString(fmt.Sprintf("    <chart:legend position=\"%s\"/>\n", cs.legendPosition))
	}

	// Grid lines
	if cs.showGrid {
		buf.WriteString("    <chart:wall>\n")
		buf.WriteString(fmt.Sprintf("      <style:graphic-properties chart:stroke-color=\"%s\" chart:stroke-width=\"0.5pt\"/>\n", cs.gridColor))
		buf.WriteString("    </chart:wall>\n")
	}

	// Axis styling
	buf.WriteString("    <chart:axis>\n")
	buf.WriteString(fmt.Sprintf("      <style:text-properties fo:color=\"%s\" fo:font-size=\"%dpt\"/>\n",
		cs.axisColor, cs.axisTitleFontSize))
	buf.WriteString("    </chart:axis>\n")

	// Series colors
	for i, color := range cs.seriesColors {
		buf.WriteString(fmt.Sprintf("    <chart:series style:data-series-number=\"%d\">\n", i))
		buf.WriteString(fmt.Sprintf("      <style:graphic-properties draw:fill-color=\"%s\" draw:stroke=\"none\"/>\n", color))
		buf.WriteString("    </chart:series>\n")
	}

	// Close elements
	buf.WriteString("  </style:chart-properties>\n")
	buf.WriteString("</style:style>")

	return buf.String()
}
