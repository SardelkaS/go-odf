package odt

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

var hNameIter = atomic.Uint64{}

type header struct {
	style *Style
	text  string
	level int64
	ref   string
}

func newHeader(t string, style *Style, level int64) *header {
	iter := hNameIter.Load()
	hNameIter.Add(1)
	if iter == 0 {
		iter = 1
		hNameIter.Add(1)
	}

	return &header{
		style: style.copy().asHeaderStyle(),
		text:  t,
		level: level,
		ref:   fmt.Sprintf("__RefHeading__%s", strconv.FormatUint(iter, 10)),
	}
}

// getRef return header ref
func (h *header) getRef() string {
	return h.ref
}

func (h *header) getFilesInfo() []FileInfo {
	return []FileInfo{}
}

// generateStyles returns header style
func (h *header) generateStyles() string {
	return h.style.generate()
}

// generate generates xml code
func (h *header) generate() string {
	return fmt.Sprintf(`<text:h text:style-name="%s" text:outline-level="%s"><text:bookmark-start text:name="%s" />%s<text:bookmark-end text:name="%s" /></text:h>`,
		h.style.getName(), strconv.FormatInt(h.level, 10), h.ref, h.text, h.ref)
}
