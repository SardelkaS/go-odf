package ods

const (
	_defaultCellStyleName = "Default"
)

// document constants
const (
	_metaFileName     = "meta.xml"
	_settingsFileName = "settings.xml"
	_stylesFileName   = "styles.xml"
	_contentFileName  = "content.xml"
	_mimeTypeFileName = "mimetype"
	_manifestFileName = "META_INF/manifest.xml"
)

// Cell value types
const (
	Float      = "float"
	Percentage = "percentage"
	Date       = "date"
	Time       = "time"
	Boolean    = "boolean"
	String     = "string"
	Void       = "void"
	Currency   = "currency"
	Formula    = "formula"
)

// Chart constants
const (
	Bar     string = "bar"     // Vertical bar chart
	Line    string = "line"    // Line chart showing trends
	Pie     string = "pie"     // Pie chart showing proportions
	Scatter string = "scatter" // Scatter plot for correlations
	Area    string = "area"    // Area chart emphasizing volume

	LegendTop    string = "top"    // Legend above chart
	LegendBottom string = "bottom" // Legend below chart
	LegendLeft   string = "left"   // Legend to the left
	LegendRight  string = "right"  // Legend to the right
	LegendNone   string = "none"   // No legend displayed
)
