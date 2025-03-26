package types

type Family string

const (
	// Family_Paragraph Paragraph styles - for paragraph formatting (main text, headings, indents etc.)
	Family_Paragraph = "Paragraph"

	// Family_Text Text styles - for inline text formatting (font, color, emphasis)
	Family_Text = "Text"

	// Family_Section Section styles - for document sections with special formatting
	Family_Section = "Section"

	// Family_Table Table styles - controls table properties (borders, positioning)
	Family_Table = "Table"

	// Family_TableColumn Table column styles - controls column width and spacing
	Family_TableColumn = "Table Column"

	// Family_TableRow Table row styles - controls row height and dividers
	Family_TableRow = "Table Row"

	// Family_TableCell Table cell styles - controls cell borders, background, alignment
	Family_TableCell = "Table Cell"

	// Family_Graphic Graphic styles - for images and shapes formatting
	Family_Graphic = "Graphic"

	// Family_Presentation Presentation styles - specific to presentation elements (Impress)
	Family_Presentation = "Presentation"

	// Family_DrawingPage Drawing page styles - controls canvas properties for drawings
	Family_DrawingPage = "Drawing Page"

	// Family_Chart Chart styles - for charts and graphs formatting
	Family_Chart = "Chart"

	// Family_Default Default style - base document settings
	Family_Default = "Default"

	// Family_Ruby Ruby styles - for East Asian typographic annotations
	Family_Ruby = "Ruby"

	// Family_List List styles - controls bullets, numbering and indents
	Family_List = "List"

	// Family_Footnote Footnote styles - formatting for footnotes at page bottom
	Family_Footnote = "Footnote"

	// Family_Endnote Endnote styles - formatting for endnotes at document end
	Family_Endnote = "Endnote"

	// Family_Index Index styles - formatting for indexes and references
	Family_Index = "Index"

	// Family_TableOfContent Table of content styles - controls TOC structure and indents
	Family_TableOfContent = "Table of Content"

	// Family_IllustrationIndex Illustration index styles - for figures and illustrations list
	Family_IllustrationIndex = "Illustration Index"

	// Family_ObjectIndex Object index styles - for objects and elements list
	Family_ObjectIndex = "Object Index"

	// Family_UserIndex User index styles - for custom indexes and references
	Family_UserIndex = "User Index"

	// Family_Bibliography Bibliography styles - formatting for citations and references
	Family_Bibliography = "Bibliography"
)
