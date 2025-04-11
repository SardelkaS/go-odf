package constants

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
