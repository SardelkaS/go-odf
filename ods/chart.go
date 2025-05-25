package ods

import "fmt"

// Chart represents a chart in an ODS document
type Chart struct {
	style     *ChartStyle
	zIndex    int
	width     string
	height    string
	posX      string
	posY      string
	dataRange string
}

// NewChart creates new empty chart
func NewChart() *Chart {
	return &Chart{}
}

// SetStyle sets chart style
func (c *Chart) SetStyle(cs *ChartStyle) {
	c.style = cs
}

// SetZIndex sets Z-index of chart
func (c *Chart) SetZIndex(z int) {
	if z < 0 {
		z = 0
	}

	c.zIndex = z
}

// SetWidth sets chart width in valid CSS units (e.g., "21cm", "8.5in")
func (c *Chart) SetWidth(w string) {
	c.width = w
}

// SetHeight sets chart height in valid CSS units (e.g., "21cm", "8.5in")
func (c *Chart) SetHeight(h string) {
	c.height = h
}

// SetPosX sets chart position X in valid CSS units (e.g., "21cm", "8.5in")
func (c *Chart) SetPosX(x string) {
	c.posX = x
}

// SetPosY sets chart position Y in valid CSS units (e.g., "21cm", "8.5in")
func (c *Chart) SetPosY(y string) {
	c.posY = y
}

// SetDataRange sets data range (e.g. "Sheet1.A1:Sheet1.A7")
func (c *Chart) SetDataRange(r string) {
	c.dataRange = r
}

func (c *Chart) generate() string {
	return fmt.Sprintf(`<draw:frame draw:z-index="%d" draw:style-name="%s" draw:text-style-name="P1"
                        svg:width="%s" svg:height="%s" svg:x="%s" svg:y="%s">
                        <draw:object draw:notify-on-update-of-ranges="%s">
                            <loext:p />
                        </draw:object>
                    </draw:frame>`,
		c.zIndex, c.style.name, c.width, c.height, c.posX, c.posY, c.dataRange)
}
