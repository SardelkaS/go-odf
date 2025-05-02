package odt

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

var hNameIter = atomic.Uint64{}

// Header represents header
type Header struct {
	style *Style
	text  string
	level int64
	ref   string
}

// NewHeader creates empty header
func NewHeader() *Header {
	iter := hNameIter.Load()
	hNameIter.Add(1)
	if iter == 0 {
		iter = 1
		hNameIter.Add(1)
	}

	return &Header{
		level: 1,
		ref:   fmt.Sprintf("__RefHeading__%s", strconv.FormatUint(iter, 10)),
	}
}

// WithText sets header text
func (h *Header) WithText(t string) *Header {
	h.text = t
	return h
}

// WithLevel sets header level. Default: 1
func (h *Header) WithLevel(level int64) *Header {
	h.level = level
	return h
}

// WithStyle sets header style
func (h *Header) WithStyle(s *Style) *Header {
	h.style = s.copy().asHeaderStyle()
	return h
}

// getRef return Header ref
func (h *Header) getRef() string {
	return h.ref
}

func (h *Header) getFilesInfo() []fileInfo {
	return []fileInfo{}
}

// generateStyles returns Header style
func (h *Header) generateStyles() string {
	return h.style.generate()
}

// generate generates xml code
func (h *Header) generate() string {
	return fmt.Sprintf(`<text:h text:style-name="%s" text:outline-level="%s"><text:bookmark-start text:name="%s" />%s<text:bookmark-end text:name="%s" /></text:h>`,
		h.style.getName(), strconv.FormatInt(h.level, 10), h.ref, h.text, h.ref)
}
