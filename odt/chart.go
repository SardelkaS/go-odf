package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

// ChartType defines the kind of diagram rendered by a Chart.
//
// The value matches the ODF "chart:class" attribute.
type ChartType string

const (
	ChartTypeBar  ChartType = "chart:bar"    // Bar/column chart
	ChartTypeLine ChartType = "chart:line"   // Line chart
	ChartTypeArea ChartType = "chart:area"   // Area chart
	ChartTypePie  ChartType = "chart:circle" // Pie chart
)

// chartColors is the default color palette applied to chart series and,
// for pie charts, to individual data points (slices)
var chartColors = []string{
	"#004586", "#ff420e", "#ffd320", "#579d1c", "#7e0021",
	"#83caff", "#314004", "#aecf00", "#4b1f6f", "#ff950e",
}

// chartSeries represents a single named data series of a Chart
type chartSeries struct {
	name   string
	values []float64
}

var chartNameIter = atomic.Uint64{}

// Chart represents an embedded chart (diagram) object that can be added to a Document.
//
// A Chart is stored as a separate "Object N" sub-document inside the ODF package
// (the same way LibreOffice embeds charts), with its data placed in an internal
// table so the resulting document does not depend on any external data source.
type Chart struct {
	name      string // object folder name inside the package, e.g. "Object 1"
	styleName string // graphic style name used to draw the chart frame
	chartType ChartType
	title     string
	width     string
	height    string

	categories []string
	series     []chartSeries
}

// NewChart creates a new empty Chart of the given type.
//
// Default size is 12cm x 8cm, use SetSize to change it.
// Use SetCategories and AddSeries to fill the chart with data.
func NewChart(chartType ChartType) *Chart {
	iter := chartNameIter.Load()
	chartNameIter.Add(1)
	if iter == 0 {
		iter = 1
		chartNameIter.Add(1)
	}

	num := strconv.FormatUint(iter, 10)
	return &Chart{
		name:      fmt.Sprintf("Object %s", num),
		styleName: fmt.Sprintf("frChart%s", num),
		chartType: chartType,
		width:     "12cm",
		height:    "8cm",
	}
}

// SetTitle sets the chart title displayed above the plot area
func (c *Chart) SetTitle(title string) {
	c.title = title
}

// SetSize sets the chart frame width and height (e.g. "12cm", "8cm")
func (c *Chart) SetSize(width, height string) {
	c.width = width
	c.height = height
}

// SetCategories sets category labels used for the X-axis (or pie slice names)
func (c *Chart) SetCategories(categories []string) {
	c.categories = categories
}

// AddSeries adds a new named data series to the chart.
//
// values should have the same length as the categories set by SetCategories,
// missing values are treated as 0. For pie charts only the first series is used.
func (c *Chart) AddSeries(name string, values []float64) {
	c.series = append(c.series, chartSeries{name: name, values: values})
}

// getFilesInfo returns the additional package files (the embedded chart object)
// that must be stored alongside the document content
func (c *Chart) getFilesInfo() []fileInfo {
	return []fileInfo{
		{Path: c.name + "/content.xml", ContentType: "text/xml", Data: []byte(c.generateObjectContent())},
		{Path: c.name + "/styles.xml", ContentType: "text/xml", Data: []byte(_chartObjectStyles)},
		{Path: c.name + "/meta.xml", ContentType: "text/xml", Data: []byte(_chartObjectMeta)},
		// directory entry: declares the embedded object MIME-type in the manifest,
		// no actual file content is written for it
		{Path: c.name + "/", ContentType: "application/vnd.oasis.opendocument.chart", Data: []byte{}},
	}
}

// generateStyles returns the graphic style used to draw the chart frame
func (c *Chart) generateStyles() string {
	return fmt.Sprintf(`<style:style style:name="%s" style:family="graphic" style:parent-style-name="OLE"><style:graphic-properties style:horizontal-pos="center" style:horizontal-rel="paragraph" draw:ole-draw-aspect="1"/></style:style>`, c.styleName)
}

// generate returns the XML representation of the frame that embeds the chart object
// into the document body
func (c *Chart) generate() string {
	return fmt.Sprintf(`<text:p text:style-name="Standard"><draw:frame draw:style-name="%s" draw:name="%s" text:anchor-type="char" svg:width="%s" svg:height="%s" draw:z-index="0"><draw:object xlink:href="./%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/></draw:frame></text:p>`,
		c.styleName, c.name, c.width, c.height, c.name)
}

// chartObjectXMLDeclaration and namespaces shared by the generated chart sub-document
const _chartObjectNamespaces = `xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" ` +
	`xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" ` +
	`xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" ` +
	`xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" ` +
	`xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" ` +
	`xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" ` +
	`xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" ` +
	`xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" ` +
	`xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" ` +
	`xmlns:xlink="http://www.w3.org/1999/xlink"`

// _chartObjectMeta is the static content of the embedded chart object's meta.xml
const _chartObjectMeta = `<?xml version="1.0" encoding="UTF-8"?>
<office:document-meta xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" office:version="1.4"><office:meta><meta:generator>go-odf</meta:generator></office:meta></office:document-meta>`

// _chartObjectStyles is the static content of the embedded chart object's styles.xml
const _chartObjectStyles = `<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" office:version="1.4"><office:styles/></office:document-styles>`

// columnName converts a zero-based column index into a spreadsheet-style column
// name (0 -> "A", 1 -> "B", ..., 25 -> "Z", 26 -> "AA", ...)
func columnName(index int) string {
	name := ""
	index++
	for index > 0 {
		index--
		name = string(rune('A'+index%26)) + name
		index /= 26
	}
	return name
}

// generateObjectContent generates the XML of the embedded chart object's content.xml.
//
// The chart carries its own data in a "local-table" element, so it does not
// depend on any cell ranges of the parent document.
func (c *Chart) generateObjectContent() string {
	isPie := c.chartType == ChartTypePie
	numPoints := len(c.categories)

	var buf bytes.Buffer
	buf.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	buf.WriteString(fmt.Sprintf(`<office:document-content %s office:version="1.4">`, _chartObjectNamespaces))

	// --- automatic styles ---
	buf.WriteString(`<office:automatic-styles>`)
	buf.WriteString(`<number:number-style style:name="N0"><number:number number:min-integer-digits="1"/></number:number-style>`)
	buf.WriteString(`<style:style style:name="chArea" style:family="chart"><style:graphic-properties draw:stroke="none" draw:fill="none"/></style:style>`)

	if c.title != "" {
		buf.WriteString(`<style:style style:name="chTitle" style:family="chart"><style:text-properties fo:font-size="13pt" fo:font-weight="bold"/></style:style>`)
	}

	buf.WriteString(`<style:style style:name="chLegend" style:family="chart"><style:chart-properties chart:auto-position="true"/><style:text-properties fo:font-size="10pt"/></style:style>`)
	buf.WriteString(`<style:style style:name="chPlotArea" style:family="chart"><style:chart-properties chart:auto-position="true" chart:auto-size="true" chart:include-hidden-cells="false" chart:right-angled-axes="true"/></style:style>`)

	buf.WriteString(`<style:style style:name="chAxisX" style:family="chart" style:data-style-name="N0"><style:chart-properties chart:display-label="true" chart:axis-position="0"/><style:text-properties fo:font-size="9pt"/></style:style>`)
	if !isPie {
		buf.WriteString(`<style:style style:name="chAxisY" style:family="chart" style:data-style-name="N0"><style:chart-properties chart:display-label="true" chart:axis-position="0"/><style:text-properties fo:font-size="9pt"/></style:style>`)
		buf.WriteString(`<style:style style:name="chGrid" style:family="chart"><style:graphic-properties svg:stroke-color="#cccccc"/></style:style>`)
		buf.WriteString(`<style:style style:name="chWall" style:family="chart"><style:graphic-properties draw:stroke="none" draw:fill="none"/></style:style>`)
		buf.WriteString(`<style:style style:name="chFloor" style:family="chart"><style:graphic-properties draw:stroke="none" draw:fill="none"/></style:style>`)
	}

	// series styles
	for i, s := range c.series {
		_ = s
		color := chartColors[i%len(chartColors)]
		switch c.chartType {
		case ChartTypeLine:
			buf.WriteString(fmt.Sprintf(`<style:style style:name="chSeries%d" style:family="chart" style:data-style-name="N0"><style:chart-properties chart:symbol-type="automatic" chart:link-data-style-to-source="true"/><style:graphic-properties svg:stroke-width="0.05cm" svg:stroke-color="%s" draw:fill-color="%s"/></style:style>`, i+1, color, color))
		default:
			buf.WriteString(fmt.Sprintf(`<style:style style:name="chSeries%d" style:family="chart" style:data-style-name="N0"><style:graphic-properties draw:stroke="solid" draw:fill="solid" svg:stroke-color="%s" draw:fill-color="%s"/></style:style>`, i+1, color, color))
		}

		// for pie charts each category becomes its own slice with its own color
		if isPie {
			for j := range c.categories {
				pieColor := chartColors[j%len(chartColors)]
				buf.WriteString(fmt.Sprintf(`<style:style style:name="chPoint%d" style:family="chart"><style:graphic-properties draw:stroke="solid" svg:stroke-color="#ffffff" draw:fill="solid" draw:fill-color="%s"/></style:style>`, j+1, pieColor))
			}
			break // only the first series is used for pie charts
		}
	}
	buf.WriteString(`</office:automatic-styles>`)

	// --- chart body ---
	buf.WriteString(`<office:body><office:chart>`)
	buf.WriteString(fmt.Sprintf(`<chart:chart svg:width="%s" svg:height="%s" xlink:href="." xlink:type="simple" chart:class="%s" chart:style-name="chArea">`,
		c.width, c.height, string(c.chartType)))

	if c.title != "" {
		buf.WriteString(fmt.Sprintf(`<chart:title chart:style-name="chTitle"><text:p>%s</text:p></chart:title>`, escapeXML(c.title)))
	}
	buf.WriteString(`<chart:legend chart:legend-position="end" style:legend-expansion="high" chart:style-name="chLegend"/>`)

	buf.WriteString(`<chart:plot-area chart:style-name="chPlotArea">`)

	// X axis with categories (used by all chart types to label data points)
	buf.WriteString(`<chart:axis chart:dimension="x" chart:name="primary-x" chart:style-name="chAxisX">`)
	if numPoints > 0 {
		buf.WriteString(fmt.Sprintf(`<chart:categories table:cell-range-address="local-table.$A$2:.$A$%d"/>`, numPoints+1))
	}
	buf.WriteString(`</chart:axis>`)

	if !isPie {
		buf.WriteString(`<chart:axis chart:dimension="y" chart:name="primary-y" chart:style-name="chAxisY"><chart:grid chart:style-name="chGrid" chart:class="major"/></chart:axis>`)
	}

	for i, s := range c.series {
		_ = s
		col := columnName(i + 1)
		styleName := fmt.Sprintf("chSeries%d", i+1)

		if isPie {
			buf.WriteString(fmt.Sprintf(`<chart:series chart:style-name="%s" chart:values-cell-range-address="local-table.$%s$2:.$%s$%d" chart:label-cell-address="local-table.$%s$1" chart:class="%s">`,
				styleName, col, col, numPoints+1, col, string(c.chartType)))
			for j := 0; j < numPoints; j++ {
				buf.WriteString(fmt.Sprintf(`<chart:data-point chart:style-name="chPoint%d"/>`, j+1))
			}
			buf.WriteString(`</chart:series>`)
			break // only the first series is rendered for pie charts
		}

		buf.WriteString(fmt.Sprintf(`<chart:series chart:style-name="%s" chart:values-cell-range-address="local-table.$%s$2:.$%s$%d" chart:label-cell-address="local-table.$%s$1" chart:class="%s"><chart:data-point chart:repeated="%d"/></chart:series>`,
			styleName, col, col, numPoints+1, col, string(c.chartType), numPoints))
	}

	if !isPie {
		buf.WriteString(`<chart:wall chart:style-name="chWall"/><chart:floor chart:style-name="chFloor"/>`)
	}
	buf.WriteString(`</chart:plot-area>`)

	// --- embedded data table ---
	buf.WriteString(`<table:table table:name="local-table">`)
	buf.WriteString(`<table:table-header-columns><table:table-column/></table:table-header-columns>`)
	if len(c.series) > 0 {
		buf.WriteString(fmt.Sprintf(`<table:table-columns><table:table-column table:number-columns-repeated="%d"/></table:table-columns>`, len(c.series)))
	}

	buf.WriteString(`<table:table-header-rows><table:table-row><table:table-cell><text:p/></table:table-cell>`)
	for _, s := range c.series {
		buf.WriteString(fmt.Sprintf(`<table:table-cell office:value-type="string"><text:p>%s</text:p></table:table-cell>`, escapeXML(s.name)))
	}
	buf.WriteString(`</table:table-row></table:table-header-rows>`)

	buf.WriteString(`<table:table-rows>`)
	for i, category := range c.categories {
		buf.WriteString(`<table:table-row>`)
		buf.WriteString(fmt.Sprintf(`<table:table-cell office:value-type="string"><text:p>%s</text:p></table:table-cell>`, escapeXML(category)))
		for _, s := range c.series {
			var value float64
			if i < len(s.values) {
				value = s.values[i]
			}
			valStr := strconv.FormatFloat(value, 'f', -1, 64)
			buf.WriteString(fmt.Sprintf(`<table:table-cell office:value-type="float" office:value="%s"><text:p>%s</text:p></table:table-cell>`, valStr, valStr))
		}
		buf.WriteString(`</table:table-row>`)
	}
	buf.WriteString(`</table:table-rows>`)
	buf.WriteString(`</table:table>`)

	buf.WriteString(`</chart:chart></office:chart></office:body></office:document-content>`)

	return buf.String()
}
