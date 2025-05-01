package odt

// document constants
const (
	_metaFileName     = "meta.xml"
	_settingsFileName = "settings.xml"
	_stylesFileName   = "styles.xml"
	_contentFileName  = "content.xml"
	_mimeTypeFileName = "mimetype"
	_manifestFileName = "META_INF/manifest.xml"
)

const (
	_textElement  = "text"
	_imageElement = "image"
	_linkElement  = "link"
)

// image
// Anchor types define how images relate to document elements
const (
	PositionTypeParagraph = "paragraph" // Relative to text paragraph
	PositionTypePage      = "page"      // Absolute page positioning
	PositionTypeChar      = "char"      // Character-relative positioning
	PositionTypeFrame     = "frame"     // Frame-relative positioning

	// Horizontal alignment options
	HorizontalLeft     = "left"      // Left-aligned
	HorizontalCenter   = "center"    // Horizontally centered
	HorizontalRight    = "right"     // Right-aligned
	HorizontalFromLeft = "from-left" // Absolute X position

	// Vertical alignment options
	VerticalTop     = "top"      // Top-aligned
	VerticalMiddle  = "middle"   // Vertically centered
	VerticalBottom  = "bottom"   // Bottom-aligned
	VerticalFromTop = "from-top" // Absolute Y position
)

// Text wrap behaviors
const (
	WrapNone     = "none"     // No text wrapping
	WrapParallel = "parallel" // Text flows parallel to image
	WrapDynamic  = "dynamic"  // Contour-based wrapping

	// Wrap side preferences
	WrapSideLeft    = "left"    // Wrap on left side only
	WrapSideRight   = "right"   // Wrap on right side only
	WrapSideBiggest = "biggest" // Automatic side selection
	WrapSideBoth    = "both"    // Wrap on both sides
)

const (
	_defaultCaptionStyleName = "Caption"
)

// table
// Table alignment options
const (
	TableAlignLeft   = "left"
	TableAlignCenter = "center"
	TableAlignRight  = "right"
	TableAlignMargin = "margins"
)

// Border styles for tables
const (
	BorderNone   = "none"
	BorderSolid  = "solid"
	BorderDotted = "dotted"
	BorderDashed = "dashed"
	BorderDouble = "double"
	BorderGroove = "groove"
	BorderRidge  = "ridge"
	BorderInset  = "inset"
	BorderOutset = "outset"
)

// Text alignment in cells
const (
	TextAlignLeft    = "left"
	TextAlignRight   = "right"
	TextAlignCenter  = "center"
	TextAlignJustify = "justify"
)

// Border model
const (
	BorderModelCollapsing = "collapsing"
	BorderModelSeparating = "separating"
)

// text
const (
	// Base
	FontName_Arial            = "Arial"
	FontName_TimesNewRoman    = "Times NewTextStyle Roman"
	FontName_CourierNew       = "Courier NewTextStyle"
	FontName_Georgia          = "Georgia"
	FontName_Verdana          = "Verdana"
	FontName_Tahoma           = "Tahoma"
	FontName_Impact           = "Impact"
	FontName_ComicSansMS      = "Comic Sans MS"
	FontName_TrebuchetMS      = "Trebuchet MS"
	FontName_PalatinoLinotype = "Palatino Linotype"
	FontName_BookAntiqua      = "Book Antiqua"
	FontName_Symbol           = "Symbol"
	FontName_Wingdings        = "Wingdings"
	FontName_Webdings         = "Webdings"

	// Windows
	FontName_Calibri              = "Calibri"
	FontName_Cambria              = "Cambria"
	FontName_Candara              = "Candara"
	FontName_Consolas             = "Consolas"
	FontName_Constantia           = "Constantia"
	FontName_Corbel               = "Corbel"
	FontName_SegoeUI              = "Segoe UI"
	FontName_FranklinGothicMedium = "Franklin Gothic Medium"
	FontName_MicrosoftSansSerif   = "Microsoft Sans Serif"
	FontName_MSReferenceSansSerif = "MS Reference Sans Serif"
	FontName_MSReferenceSerif     = "MS Reference Serif"
	FontName_Ebrima               = "Ebrima"
	FontName_LeelawadeeUI         = "Leelawadee UI"
	FontName_MalgunGothic         = "Malgun Gothic"
	FontName_Sylfaen              = "Sylfaen"
	FontName_HoloLensMDL2Assets   = "HoloLens MDL2 Assets"
	FontName_Marlett              = "Marlett"
	FontName_SimSun               = "SimSun"
	FontName_NSimSun              = "NSimSun"
	FontName_MingLiU              = "MingLiU"
	FontName_PMingLiU             = "PMingLiU"
	FontName_MSGothic             = "MS Gothic"
	FontName_MSMincho             = "MS Mincho"
	FontName_MSUIGothic           = "MS UI Gothic"

	// macOS
	FontName_SanFrancisco  = "San Francisco"
	FontName_NewYork       = "NewTextStyle York"
	FontName_SFPro         = "SF Pro"
	FontName_SFCompact     = "SF Compact"
	FontName_SFMono        = "SF Mono"
	FontName_HelveticaNeue = "Helvetica Neue"
	FontName_Helvetica     = "Helvetica"
	FontName_LucidaGrande  = "Lucida Grande"
	FontName_Geneva        = "Geneva"
	FontName_AppleSymbols  = "Apple Symbols"
	FontName_Menlo         = "Menlo"
	FontName_Monaco        = "Monaco"
	FontName_Optima        = "Optima"
	FontName_Zapfino       = "Zapfino"
	FontName_BrushScriptMT = "Brush Script MT"
	FontName_Chalkboard    = "Chalkboard"
	FontName_Chalkduster   = "Chalkduster"
	FontName_Noteworthy    = "Noteworthy"
	FontName_Papyrus       = "Papyrus"
	FontName_Thonburi      = "Thonburi"

	// Linux
	FontName_DejaVuSans      = "DejaVu Sans"
	FontName_DejaVuSerif     = "DejaVu Serif"
	FontName_DejaVuSansMono  = "DejaVu Sans Mono"
	FontName_LiberationSans  = "Liberation Sans"
	FontName_LiberationSerif = "Liberation Serif"
	FontName_LiberationMono  = "Liberation Mono"
	FontName_Ubuntu          = "Ubuntu"
	FontName_UbuntuMono      = "Ubuntu Mono"
	FontName_NotoSans        = "Noto Sans"
	FontName_NotoSerif       = "Noto Serif"
	FontName_NotoMono        = "Noto Mono"
	FontName_FreeSans        = "FreeSans"
	FontName_FreeSerif       = "FreeSerif"
	FontName_FreeMono        = "FreeMono"
	FontName_DroidSans       = "Droid Sans"
	FontName_DroidSerif      = "Droid Serif"
	FontName_DroidSansMono   = "Droid Sans Mono"
	FontName_OpenSans        = "Open Sans"
	FontName_Roboto          = "Roboto"
	FontName_SourceSansPro   = "Source Sans Pro"
	FontName_SourceSerifPro  = "Source Serif Pro"
	FontName_SourceCodePro   = "Source Code Pro"
	FontName_FiraSans        = "Fira Sans"
	FontName_FiraMono        = "Fira Mono"
	FontName_Cantarell       = "Cantarell"
	FontName_Inconsolata     = "Inconsolata"
	FontName_NimbusSans      = "Nimbus Sans"
	FontName_NimbusRoman     = "Nimbus Roman"
	FontName_NimbusMono      = "Nimbus Mono"

	FontName_UNDEFINED = "UNDEFINED"
)

// Underline styles
const (
	UnderlineNone       = "none"         // No underline
	UnderlineSingle     = "solid"        // Single solid line
	UnderlineDouble     = "double"       // Double solid line
	UnderlineDotted     = "dotted"       // Dotted line
	UnderlineDash       = "dash"         // Dashed line
	UnderlineWave       = "wave"         // Wavy line
	UnderlineBold       = "bold"         // Thick solid line
	UnderlineDotDash    = "dot-dash"     // Alternating dots and dashes
	UnderlineDotDotDash = "dot-dot-dash" // Two dots followed by a dash
	UnderlineLongDash   = "long-dash"    // Longer dashes
)

// Overline styles (similar to underline but above text)
const (
	OverlineNone   = "none"   // No overline
	OverlineSingle = "solid"  // Single solid line
	OverlineDouble = "double" // Double solid line
	OverlineDotted = "dotted" // Dotted line
	OverlineDash   = "dash"   // Dashed line
	OverlineWave   = "wave"   // Wavy line
)

// Line-through (strikethrough) styles
const (
	LineThroughNone  = "none"  // No line through text
	LineThroughSolid = "solid" // Single solid line
	LineThroughWave  = "wave"  // Wavy line
	LineThroughSlash = "slash" // Diagonal slash
	LineThroughX     = "x"     // X-shaped
)

// Text transformation options
const (
	TransformNone       = "none"       // No transformation
	TransformUppercase  = "uppercase"  // ALL UPPERCASE
	TransformLowercase  = "lowercase"  // all lowercase
	TransformCapitalize = "capitalize" // Capitalize Each Word
	TransformSmallCaps  = "small-caps" // Small Caps
)

// Text rotation scaling options
const (
	RotationScaleFixed      = "fixed"       // Line height remains unchanged
	RotationScaleLineHeight = "line-height" // Line height adjusts to fit rotated text
)

// Writing modes (text directions)
const (
	WritingModeLR_TB = "lr-tb" // Left-to-right, top-to-bottom (Western)
	WritingModeRL_TB = "rl-tb" // Right-to-left, top-to-bottom (Hebrew/Arabic)
	WritingModeTB_RL = "tb-rl" // Top-to-bottom, right-to-left (Asian vertical)
	WritingModeTB_LR = "tb-lr" // Top-to-bottom, left-to-right (Mongolian)
	WritingModePage  = "page"  // Follows page text direction
)

// Text emphasis marks (primarily for East Asian typography)
const (
	EmphasisNone         = "none"
	EmphasisDot          = "dot"
	EmphasisCircle       = "circle"
	EmphasisDisc         = "disc"
	EmphasisAccent       = "accent"
	EmphasisFilledDot    = "filled dot"
	EmphasisFilledCircle = "filled circle"
	EmphasisFilledDisc   = "filled disc"
	EmphasisFilledAccent = "filled accent"
	EmphasisAbove        = "above"
	EmphasisBelow        = "below"
	EmphasisLeft         = "left"
	EmphasisRight        = "right"
)

// Font weights
const (
	FontWeightNormal = "normal" // Regular weight
	FontWeightBold   = "bold"   // Bold weight
	FontWeight100    = "100"    // Thin
	FontWeight200    = "200"    // Extra Light
	FontWeight300    = "300"    // Light
	FontWeight400    = "400"    // Normal (same as 'normal')
	FontWeight500    = "500"    // Medium
	FontWeight600    = "600"    // Semi Bold
	FontWeight700    = "700"    // Bold (same as 'bold')
	FontWeight800    = "800"    // Extra Bold
	FontWeight900    = "900"    // Black
)

// Font styles
const (
	FontStyleNormal  = "normal"  // Upright characters
	FontStyleItalic  = "italic"  // Italic characters
	FontStyleOblique = "oblique" // Slanted characters
)

// Text tags
const (
	_textTagSpan = "span"
	_textTagP    = "p"
)

// List
const (
	NumberStyleArabic     string = "1" // 1, 2, 3, ...
	NumberStyleUpperAlpha string = "A" // A, B, C, ...
	NumberStyleLowerAlpha string = "a" // a, b, c, ...
	NumberStyleUpperRoman string = "I" // I, II, III, ...
	NumberStyleLowerRoman string = "i" // i, ii, iii, ...
	BulletStyleDisc       string = "•" // Filled circle
	BulletStyleCircle     string = "○" // Hollow circle
	BulletStyleSquare     string = "■" // Filled square
)

// Heading
const (
	_sectStyleName    = "Sect1"
	_sectStyle        = `<style:style style:name="Sect1" style:family="section"><style:section-properties style:editable="false"><style:columns fo:column-count="1" fo:column-gap="0cm" /></style:section-properties></style:style>`
	_index20StyleName = "Index_20_Link"
	_index20Style     = `<style:style style:name="Index_20_Link" style:display-name="Index Link" style:family="text" />`
)
