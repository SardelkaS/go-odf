package odt

import (
	"bytes"
	"fmt"
)

type listElement interface {
	generateStyles() string
	generate() string
}

type listConfig struct {
	continueNumber bool // Continue numbering from previous list
	startValue     int  // Starting number for ordered lists
	indentLevel    int  // Nesting level (0 = top level)
}

// List represents a list in ODF document
type List struct {
	config listConfig
	style  *listStyle
	items  []listElement
}

// NewList creates empty List
func NewList() *List {
	return &List{
		config: listConfig{},
		style:  newListStyle(),
	}
}

// SetNumFormat sets num format for provided level
//
//	NumberStyleArabic
//	NumberStyleUpperAlpha
//	NumberStyleLowerAlpha
//	NumberStyleUpperRoman
//	NumberStyleLowerRoman
//	BulletStyleDisc
//	BulletStyleCircle
//	BulletStyleSquare
func (l *List) SetNumFormat(level int, nf string) {
	if lvl, ok := l.style.levels[level]; ok {
		lvl.bullet = false
		lvl.numFormat = nf
	} else {
		l.style.levels[level] = &listLevelStyle{
			level:      level,
			bullet:     false,
			numFormat:  nf,
			marginLeft: fmt.Sprintf("%dcm", level+1),
		}
	}
}

// SetBulletFormat sets bullet format for provided level
//
//	BulletStyleDisc
//	BulletStyleCircle
//	BulletStyleSquare
func (l *List) SetBulletFormat(level int, bf string) {
	if lvl, ok := l.style.levels[level]; ok {
		lvl.bullet = true
		lvl.bulletChar = bf
	} else {
		l.style.levels[level] = &listLevelStyle{
			level:      level,
			bullet:     true,
			bulletChar: bf,
			marginLeft: fmt.Sprintf("%dcm", level+1),
		}
	}
}

// SetNumSuffix sets num suffix (example, ".")
// without suffix:
//
//	1 record
//	2 record
//
// with suffix "."
//  1. record
//  2. record
func (l *List) SetNumSuffix(level int, ns string) {
	if lvl, ok := l.style.levels[level]; ok {
		lvl.suffix = ns
	} else {
		l.style.levels[level] = &listLevelStyle{
			level:      level,
			bullet:     false,
			numFormat:  NumberStyleArabic,
			suffix:     ns,
			marginLeft: fmt.Sprintf("%dcm", level+1),
		}
	}
}

// SetMarginLeft sets margin left for provided level
//
// default is 1cm * (level + 1)
func (l *List) SetMarginLeft(level int, ml string) {
	if lvl, ok := l.style.levels[level]; ok {
		lvl.marginLeft = ml
	} else {
		l.style.levels[level] = &listLevelStyle{
			level:      level,
			bullet:     true,
			numFormat:  BulletStyleDisc,
			marginLeft: ml,
		}
	}
}

// SetContinueNumber sets continue numbering from previous list flag
func (l *List) SetContinueNumber(fl bool) {
	l.config.continueNumber = fl
}

// SetStartValue sets starting number for ordered lists
func (l *List) SetStartValue(v int) {
	l.config.startValue = v
}

// SetIndentInterval sets nesting level (0 = top level)
func (l *List) SetIndentInterval(i int) {
	l.config.indentLevel = i
}

// AddText adds text to list
func (l *List) AddText(t string, s *Style) {
	l.items = append(l.items, newText(t, s.copy().withListRef(l.style.name)).withTag(_textTagP))
}

// AddParagraph adds paragraph to list
func (l *List) AddParagraph(p *Paragraph) {
	l.items = append(l.items, p)
}

// AddList adds sublist to list
func (l *List) AddList(subL *List) {
	l.items = append(l.items, subL)
}

func (l *List) getFilesInfo() []FileInfo {
	return []FileInfo{}
}

func (l *List) generateStyles() string {
	var stylesBuffer bytes.Buffer

	stylesBuffer.WriteString(l.style.generate())
	for _, item := range l.items {
		stylesBuffer.WriteString(item.generateStyles())
	}

	return stylesBuffer.String()
}

func (l *List) generate() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(`<text:list text:style-name="%s"`, l.style.name))
	if l.config.startValue > 1 {
		buf.WriteString(fmt.Sprintf(` text:start-value="%d"`, l.config.startValue))
	}
	if l.config.continueNumber {
		buf.WriteString(` text:continue-numbering="true"`)
	}
	if l.config.indentLevel > 0 {
		buf.WriteString(fmt.Sprintf(` text:level="%d"`, l.config.indentLevel+1))
	}
	buf.WriteString(">")

	for _, item := range l.items {
		buf.WriteString("<text:list-item>")
		buf.WriteString(item.generate())
		buf.WriteString("</text:list-item>")
	}

	buf.WriteString("</text:list>")
	return buf.String()
}
