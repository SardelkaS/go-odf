package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var hhNameIter = atomic.Uint64{}

type headingLevel struct {
	level int64
	style *headingStyle
}

// Heading represents heading
type Heading struct {
	name        string
	header      string
	headerStyle *Style
	elements    []*Header
	levels      map[int64]headingLevel
}

// NewHeading creates empty heading
func NewHeading() *Heading {
	iter := hhNameIter.Load()
	hhNameIter.Add(1)
	if iter == 0 {
		iter = 1
		hhNameIter.Add(1)
	}

	return &Heading{
		name:     fmt.Sprintf("Heading_%s", strconv.FormatUint(iter, 10)),
		header:   "Heading",
		elements: make([]*Header, 0),
		levels:   make(map[int64]headingLevel),
	}
}

// SetHeader sets heading header text
func (h *Heading) SetHeader(t string) {
	h.header = t
}

// SetHeaderStyle sets heading header style
func (h *Heading) SetHeaderStyle(s *Style) {
	h.headerStyle = s.copy().asHeaderStyle()
}

// AddLink adds new element to heading
func (h *Heading) AddLink(hh *Header) {
	if _, ok := h.levels[hh.level]; !ok {
		h.levels[hh.level] = headingLevel{
			level: hh.level,
			style: newHeadingStyle(hh.level),
		}
	}

	h.elements = append(h.elements, hh)
}

func (h *Heading) generateStyles() string {
	var stylesBuffer bytes.Buffer

	if h.headerStyle != nil {
		stylesBuffer.WriteString(h.headerStyle.generate())
	}

	for _, lvl := range h.levels {
		stylesBuffer.WriteString(lvl.style.generate())
	}

	return stylesBuffer.String()
}

func (h *Heading) getFilesInfo() []FileInfo {
	return []FileInfo{}
}

func (h *Heading) generate() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`<text:table-of-content text:style-name="%s" text:protected="true" text:name="%s"><text:table-of-content-source text:outline-level="%s">`,
		_sectStyleName, h.name, strconv.FormatInt(int64(len(h.levels)), 10)))
	buf.WriteString("<text:index-title-template ")
	if h.headerStyle != nil {
		buf.WriteString(fmt.Sprintf(`text:style-name="%s"`, h.headerStyle.getName()))
	}
	buf.WriteString(fmt.Sprintf(`>%s</text:index-title-template>`, h.header))

	for _, lvl := range h.levels {
		buf.WriteString(fmt.Sprintf(`<text:table-of-content-entry-template text:outline-level="%s"
                        text:style-name="%s">
                        <text:index-entry-link-start text:style-name="%s" />
                        <text:index-entry-chapter />
                        <text:index-entry-text />
                        <text:index-entry-tab-stop style:type="right" style:leader-char="." />
                        <text:index-entry-page-number />
                        <text:index-entry-link-end />
                    </text:table-of-content-entry-template>
`, strconv.FormatInt(lvl.level, 10), lvl.style.name, _index20StyleName))
	}

	buf.WriteString(`</text:table-of-content-source><text:index-body>`)
	buf.WriteString(fmt.Sprintf(`<text:index-title text:style-name="%s" text:name="%s_Head" text:protected="true"> <text:p `,
		_sectStyleName, h.name))
	if h.headerStyle != nil {
		buf.WriteString(fmt.Sprintf(`text:style-name="%s"`, h.headerStyle.getName()))
	}
	buf.WriteString(fmt.Sprintf(`>%s</text:p></text:index-title>`, h.header))

	for _, e := range h.elements {
		buf.WriteString(fmt.Sprintf(`<text:p text:style-name="%s">
                        <text:a xlink:type="simple" xlink:href="#%s"
                            text:style-name="%s" text:visited-style-name="%s">
                            %s</text:a>
						</text:p>`,
			h.levels[e.level].style.name, e.getRef(), _index20StyleName, _index20StyleName, e.text))
	}

	buf.WriteString(`</text:index-body></text:table-of-content>`)

	return buf.String()
}
