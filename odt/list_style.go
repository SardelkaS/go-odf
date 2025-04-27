package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var listStyleIter = atomic.Uint64{}

// listLevelStyle defines formatting for a specific list level
type listLevelStyle struct {
	level         int
	bullet        bool
	numFormat     string
	bulletChar    string
	prefix        string
	suffix        string
	displayLevels int
	marginLeft    string
}

// listStyle defines a complete list style
type listStyle struct {
	name   string
	levels map[int]*listLevelStyle
}

func newListStyle() *listStyle {
	iter := listStyleIter.Load()
	listStyleIter.Add(1)
	if iter == 0 {
		iter = 1
		listStyleIter.Add(1)
	}

	return &listStyle{
		name: fmt.Sprintf("L%s", strconv.FormatUint(iter, 10)),
		levels: map[int]*listLevelStyle{
			1: {
				level:      1,
				bullet:     true,
				bulletChar: BulletStyleDisc,
				marginLeft: "1cm",
			},
			2: {
				level:      2,
				bullet:     true,
				bulletChar: BulletStyleDisc,
				marginLeft: "2cm",
			},
			3: {
				level:      3,
				bullet:     true,
				bulletChar: BulletStyleDisc,
				marginLeft: "3cm",
			},
			4: {
				level:      4,
				bullet:     true,
				bulletChar: BulletStyleDisc,
				marginLeft: "4cm",
			},
			5: {
				level:      5,
				bullet:     true,
				bulletChar: BulletStyleDisc,
				marginLeft: "5cm",
			},
		},
	}
}

func (ls *listStyle) generate() string {
	var buf bytes.Buffer

	// Start list-style element with name attribute
	buf.WriteString(`<text:list-style style:name="`)
	buf.WriteString(ls.name)
	buf.WriteString(`">`)

	// Generate numbered level styles
	for _, level := range ls.levels {
		if level.bullet {
			buf.WriteString(`<text:list-level-style-bullet`)
			buf.WriteString(` text:level="`)
			buf.WriteString(strconv.Itoa(level.level))
			buf.WriteString(`"`)

			buf.WriteString(` text:bullet-char="`)
			buf.WriteString(escapeXML(level.bulletChar))
			buf.WriteString(`"`)

			buf.WriteString(`>`)
			buf.WriteString(fmt.Sprintf(`<style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab"
                        text:list-tab-stop-position="%s" fo:text-indent="-0.635cm"
                        fo:margin-left="%s" />
                </style:list-level-properties>
				<style:text-properties fo:font-family="OpenSymbol" />`, level.marginLeft, level.marginLeft))
			buf.WriteString(`</text:list-level-style-bullet>`)
		} else {
			buf.WriteString(`<text:list-level-style-number`)
			buf.WriteString(` text:level="`)
			buf.WriteString(strconv.Itoa(level.level))
			buf.WriteString(`"`)

			buf.WriteString(` style:num-format="`)
			buf.WriteString(level.numFormat)
			buf.WriteString(`"`)

			if level.prefix != "" {
				buf.WriteString(` style:num-prefix="`)
				buf.WriteString(escapeXML(level.prefix))
				buf.WriteString(`"`)
			}

			if level.suffix != "" {
				buf.WriteString(` style:num-suffix="`)
				buf.WriteString(escapeXML(level.suffix))
				buf.WriteString(`"`)
			}

			if level.displayLevels > 0 {
				buf.WriteString(` text:display-levels="`)
				buf.WriteString(strconv.Itoa(level.displayLevels))
				buf.WriteString(`"`)
			}

			buf.WriteString(`>`)
			buf.WriteString(fmt.Sprintf(`<style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab"
                        text:list-tab-stop-position="%s" fo:text-indent="-0.635cm"
                        fo:margin-left="%s" />
                </style:list-level-properties>`, level.marginLeft, level.marginLeft))
			buf.WriteString(`</text:list-level-style-number>`)
		}
	}

	// Close list-style element
	buf.WriteString(`</text:list-style>`)

	return buf.String()
}
