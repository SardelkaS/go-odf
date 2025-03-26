package props

import (
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/style/types"
)

type ParagraphProps struct {
	breakBefore types.BreakBefore
}

// newParagraphProps creates new ParagraphProps with default values
func newParagraphProps() ParagraphProps {
	return ParagraphProps{
		breakBefore: types.Break_auto,
	}
}

// SetBreakBefore set fo:break-before value
func (p ParagraphProps) SetBreakBefore(bb types.BreakBefore) {
	p.breakBefore = bb
}

// Generate generates xml code
func (p ParagraphProps) Generate() string {
	if p.breakBefore == types.Break_auto {
		return ""
	}

	return fmt.Sprintf(`<paragraph-properties fo:break-before="%s" />`, p.breakBefore)
}
