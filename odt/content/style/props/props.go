package props

type Props struct {
	text      TextProps
	paragraph ParagraphProps
}

// New creates new Props with default values
func New() Props {
	return Props{
		text:      newTextProps(),
		paragraph: newParagraphProps(),
	}
}

// Generate generates xml code
func (p Props) Generate() string {
	paragraphXml := p.paragraph.Generate()
	textXml := p.text.Generate()

	if paragraphXml == "" {
		return textXml
	}

	return paragraphXml + "\n" + textXml
}
