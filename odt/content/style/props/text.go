package props

import (
	"fmt"
	types2 "github.com/SardelkaS/go-odf/odt/content/style/types"
	"strconv"
)

type TextProps struct {
	fontSize        uint64
	fontSizeAsian   uint64
	fontSizeComplex uint64
	fontName        types2.FontName
	fontNameComplex types2.FontName
	language        types2.Language
	country         types2.Country
}

// newTextProps creates new TextProps with default values
func newTextProps() TextProps {
	return TextProps{
		fontSize:        14,
		fontSizeAsian:   14,
		fontSizeComplex: 14,
		fontName:        types2.FontName_TimesNewRoman,
		fontNameComplex: types2.FontName_UNDEFINED,
		language:        types2.Language_en,
		country:         types2.Country_US,
	}
}

// SetFontSize set font size value
func (t TextProps) SetFontSize(fs uint64) error {
	if fs < 1 || fs > 72 {
		return fmt.Errorf("ivalid font size: must be in 1...72")
	}

	t.fontSize = fs
	return nil
}

// SetFontName set font name
func (t TextProps) SetFontName(fn types2.FontName) {
	t.fontName = fn
}

// SetLanguage set language
func (t TextProps) SetLanguage(l types2.Language) {
	t.language = l
}

// SetCountry set country
func (t TextProps) SetCountry(c types2.Country) {
	t.country = c
}

// Generate generates xml code
func (t TextProps) Generate() string {
	if t.fontSizeAsian == 0 {
		t.fontSizeAsian = t.fontSize
	}
	if t.fontSizeComplex == 0 {
		t.fontSizeComplex = t.fontSize
	}
	if t.fontNameComplex == types2.FontName_UNDEFINED {
		t.fontNameComplex = t.fontName
	}

	return fmt.Sprintf(`<text-properties style:font-name=%s style:font-name-complex=%s fo:font-size="%spt" style:font-size-asian="%spt" style:font-size-complex="%spt" fo:language="%s" fo:country="%s" />`,
		t.fontName, t.fontNameComplex,
		strconv.FormatUint(t.fontSize, 10), strconv.FormatUint(t.fontSizeAsian, 10), strconv.FormatUint(t.fontSizeComplex, 10),
		t.language, t.country)
}
