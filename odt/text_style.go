package odt

import (
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

var tsNameIter = atomic.Uint64{}

// Style represents text formatting options for ODT documents.
// It allows customization of font properties, weight, style, and color.
type Style struct {
	name                 string
	fontName             string
	fontSize             string // e.g. "12pt", "1.5cm"
	bold                 bool
	italic               bool
	color                string // hex color like "#000000"
	textShadow           string // Text shadow (e.g., "1pt 1pt 0.5pt #000")
	letterSpacing        string // Spacing between letters (e.g., "0.5mm")
	textTransform        string // Text case transformation
	textUnderlineStyle   string // Underline style
	textUnderlineColor   string // Underline color
	textOverlineStyle    string // Overline style
	textOverlineColor    string // Overline color
	textLineThroughStyle string // Strike-through style
	textOutline          string // Text outline (e.g., "true", "1pt black")
	textEmphasize        string // Emphasis marks
	writingMode          string // Text direction (e.g., "lr-tb", "tb-rl")
	textRotationAngle    int64  // Rotation angle (0-360 degrees)
	textRotationScale    string // How rotation affects line height
	listRef              string
	parentStyleName      string
}

// NewTextStyle creates new Style with default values
func NewTextStyle() *Style {
	iter := tsNameIter.Load()
	tsNameIter.Add(1)
	if iter == 0 {
		iter = 2
		tsNameIter.Add(2)
	}

	return &Style{
		name: fmt.Sprintf("T%s", strconv.FormatUint(iter, 10)),
	}
}

func (s *Style) copy() *Style {
	iter := tsNameIter.Load()
	tsNameIter.Add(1)
	if iter == 0 {
		iter = 2
		tsNameIter.Add(2)
	}

	newStyle := *s
	newStyle.name = fmt.Sprintf("T%s", strconv.FormatUint(iter, 10))
	return &newStyle
}

func (s *Style) withListRef(ref string) *Style {
	s.listRef = ref
	return s
}

// WithParentStyle set parent style name
func (s *Style) WithParentStyle(ps *Style) *Style {
	s.parentStyleName = ps.getName()
	return s
}

// WithFontName set FontName. FontName specifies the font family (e.g., "Arial", "Times New Roman")
func (s *Style) WithFontName(fontName string) *Style {
	s.fontName = fontName
	return s
}

// WithFontSize set FontSize. FontSize specifies the text size with units (e.g., "12pt", "1.5cm")
// Common values: "10pt", "12pt", "14pt", "16pt", "18pt"
func (s *Style) WithFontSize(fontSize string) *Style {
	s.fontSize = fontSize
	return s
}

// WithBold enables bold text formatting
func (s *Style) WithBold() *Style {
	s.bold = true
	return s
}

// WithItalic enables italic text formatting
func (s *Style) WithItalic() *Style {
	s.italic = true
	return s
}

// WithColor specifies text color in hex format (e.g., "#FF0000" for red)
// Supported formats: "#RGB", "#RRGGBB"
func (s *Style) WithColor(c string) *Style {
	s.color = c
	return s
}

// WithTextShadow sets a shadow effect for the text.
// Format should be "horizontal vertical blur color" (e.g., "1pt 1pt 0.5pt #000000").
// Set to empty string to remove shadow.
// Example:
//
//	style.WithTextShadow("1pt 1pt 0.3pt #AAAAAA")
func (s *Style) WithTextShadow(shadow string) *Style {
	s.textShadow = shadow
	return s
}

// WithLetterSpacing sets the spacing between characters.
// The spacing should include units (e.g., "0.5mm", "2pt").
// Example:
//
//	style.WithLetterSpacing("0.3mm")
func (s *Style) WithLetterSpacing(spacing string) *Style {
	s.letterSpacing = spacing
	return s
}

// WithTextTransform sets the text case transformation.
// Should use one of: TransformNone, TransformUppercase,
// TransformLowercase, or TransformCapitalize.
// Example:
//
//	style.WithTextTransform(types.TransformUppercase)
func (s *Style) WithTextTransform(transform string) *Style {
	s.textTransform = transform
	return s
}

// WithUnderline enables text underlining
func (s *Style) WithUnderline() *Style {
	s.textUnderlineStyle = UnderlineSingle
	s.textUnderlineColor = "font-color"
	return s
}

// WithStyledUnderline configures text underline with style and color.
// Style should use one of the Underline* constants.
// Color should be in hex format (e.g., "#FF0000").
// Example:
//
//	style.WithUnderline(UnderlineWave, "#0000FF")
func (s *Style) WithStyledUnderline(style, color string) *Style {
	s.textUnderlineStyle = style
	s.textUnderlineColor = color
	return s
}

// WithOverline configures text overline (above text) appearance.
// style: OverlineSingle or OverlineDouble
// color: Hex color code
// Example:
//
//	style.WithOverline(OverlineSingle, "#00FF00")
func (s *Style) WithOverline(style, color string) *Style {
	s.textOverlineStyle = style
	s.textOverlineColor = color
	return s
}

// WithLineThrough configures strikethrough appearance.
// style: LineThroughSolid or LineThroughWave
// Example:
//
//	style.WithLineThrough(LineThroughSolid)
func (s *Style) WithLineThrough(style string) *Style {
	s.textLineThroughStyle = style
	return s
}

// WithTextOutline sets text outline (border around text).
// Format: "width color" or "true" for default (e.g., "1pt #000000").
// Example:
//
//	style.WithTextOutline("0.5pt #FFFFFF")
func (s *Style) WithTextOutline(outline string) *Style {
	s.textOutline = outline
	return s
}

// WithTextEmphasis sets emphasis marks for text (used in East Asian typography).
// Format: "shape position" (e.g., "filled circle above").
// Example:
//
//	style.WithTextEmphasis("filled dot below")
func (s *Style) WithTextEmphasis(emphasis string) *Style {
	s.textEmphasize = emphasis
	return s
}

// WithWritingMode sets the text direction and layout.
// Options: "lr-tb" (left-right, top-bottom), "tb-rl" (top-bottom, right-left),
// "rl-tb" (right-left, top-bottom), etc.
// Example:
//
//	style.WithWritingMode("tb-rl") // Vertical text
func (s *Style) WithWritingMode(mode string) *Style {
	s.writingMode = mode
	return s
}

// WithRotation sets the text rotation angle and scaling behavior.
// Angle should be between 0 and 360 degrees.
// Scale should be either RotationScaleFixed or RotationScaleLineHeight.
// Example:
//
//	style.WithRotation(90, RotationScaleLineHeight)
func (s *Style) WithRotation(angle int64, scale string) *Style {
	s.textRotationAngle = angle
	s.textRotationScale = scale
	return s
}

// getName returns style name
func (s *Style) getName() string {
	return s.name
}

// generate generates xml code
func (s *Style) generate() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf(`<style:style style:name="%s"`, s.name))

	if s.listRef != "" {
		builder.WriteString(fmt.Sprintf(` style:family="paragraph" style:list-style-name="%s"`, s.listRef))
	} else {
		builder.WriteString(` style:family="text"`)
	}

	if s.parentStyleName != "" {
		builder.WriteString(fmt.Sprintf(` style:parent-style-name="%s"`, s.parentStyleName))
	}

	builder.WriteString(`>`)

	builder.WriteString("<style:text-properties")

	if s.fontName != "" {
		builder.WriteString(fmt.Sprintf(` style:font-name="%s"`, s.fontName))
	}
	if s.fontSize != "" {
		builder.WriteString(fmt.Sprintf(` fo:font-size="%s"`, s.fontSize))
	}
	if s.bold {
		builder.WriteString(` fo:font-weight="bold"`)
	}
	if s.italic {
		builder.WriteString(` fo:font-style="italic"`)
	}
	if s.textUnderlineStyle != "" && s.textUnderlineColor != "" {
		builder.WriteString(fmt.Sprintf(` style:text-underline-style="%s" style:text-underline-width="auto" style:text-underline-color="%s"`,
			s.textUnderlineStyle, s.textUnderlineColor))
	}
	if s.color != "" {
		builder.WriteString(fmt.Sprintf(` fo:color="%s"`, s.color))
	}
	if s.textShadow != "" {
		builder.WriteString(fmt.Sprintf(` fo:text-shadow="%s"`, s.textShadow))
	}
	if s.letterSpacing != "" {
		builder.WriteString(fmt.Sprintf(` fo:letter-spacing="%s"`, s.letterSpacing))
	}
	if s.textTransform != "" {
		builder.WriteString(fmt.Sprintf(` fo:text-transform="%s"`, s.textTransform))
	}
	if s.textOverlineStyle != "" && s.textOverlineColor != "" {
		builder.WriteString(fmt.Sprintf(` style:text-overline-style="%s" style:text-overline-color="%s"`,
			s.textOverlineStyle, s.textOverlineColor))
	}
	if s.textLineThroughStyle != "" {
		builder.WriteString(fmt.Sprintf(` style:text-line-through-style="%s"`, s.textLineThroughStyle))
	}
	if s.textOutline != "" {
		builder.WriteString(fmt.Sprintf(` fo:text-outline="%s"`, s.textOutline))
	}
	if s.textEmphasize != "" {
		builder.WriteString(fmt.Sprintf(` style:text-emphasize="%s"`, s.textEmphasize))
	}
	if s.writingMode != "" {
		builder.WriteString(fmt.Sprintf(` style:writing-mode="%s"`, s.writingMode))
	}
	if s.textRotationScale != "" {
		builder.WriteString(fmt.Sprintf(` style:text-rotation-scale="%s" style:text-rotation-angle="%s"`,
			s.textRotationScale, strconv.FormatInt(s.textRotationAngle, 10)))
	}

	builder.WriteString("/></style:style>")
	return builder.String()
}
