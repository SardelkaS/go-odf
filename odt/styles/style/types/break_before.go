package types

type BreakBefore string

const (
	Break_auto   = "auto"      // Default automatic pagination (renderer decides)
	Break_page   = "page"      // Force page break before element
	Break_column = "column"    // Start new column before element
	Break_even   = "even-page" // Next even-numbered page
	Break_odd    = "odd-page"  // Next odd-numbered page

	// Special pagination controls
	Break_avoid  = "avoid"  // Attempt to avoid break (not guaranteed)
	Break_always = "always" // Legacy value (equivalent to 'page')
)
