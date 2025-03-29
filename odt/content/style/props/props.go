package props

type Props struct {
	Text      *TextProps
	Paragraph ParagraphProps
}

// New creates new Props with default values
func New() Props {
	return Props{
		Text:      newTextProps(),
		Paragraph: newParagraphProps(),
	}
}

// Generate generates xml code
func (p Props) Generate() string {
	paragraphXml := p.Paragraph.Generate()
	textXml := p.Text.Generate()

	if paragraphXml == "" {
		return textXml
	}

	return paragraphXml + "\n" + textXml
}
