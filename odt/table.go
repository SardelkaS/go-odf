package odt

import (
	"bytes"
	"fmt"
	"strconv"
	"sync/atomic"
)

var nameIter = atomic.Uint64{}

// Table represents an ODF table
type Table struct {
	name    string
	rows    []*row
	columns []*column
	style   *TableStyle
}

// row represents a table row
type row struct {
	cells []*cell
	style *RowStyle
	name  string
}

// column represents a table column
type column struct {
	style *ColumnStyle
	name  string
}

// cell represents a table cell
type cell struct {
	value   cellValue
	style   *CellStyle
	colSpan int
	rowSpan int
	name    string
}

// cellValue represents a table cell value
type cellValue struct {
	value string
	style *Style
}

// NewTable creates new table with r rows and c columns
func NewTable(r, c int) *Table {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 1
		nameIter.Add(1)
	}

	rows := make([]*row, r)
	cols := make([]*column, c)
	for i := range rows {
		row := &row{
			cells: make([]*cell, c),
		}
		for j := range row.cells {
			row.cells[j] = &cell{}
		}

		rows[i] = row
	}
	for i := range cols {
		cols[i] = &column{}
	}

	return &Table{
		name:    fmt.Sprintf("Table%s", strconv.FormatUint(iter, 10)),
		rows:    rows,
		columns: cols,
	}
}

// SetValue sets cell value
func (t *Table) SetValue(row, col int, value string) {
	if row >= 0 && row < len(t.rows) && col >= 0 && col < len(t.rows[row].cells) {
		t.rows[row].cells[col].value.value = value
	}
}

// SetSpan sets cell span
func (t *Table) SetSpan(row, col, colSpan, rowSpan int) {
	if row >= 0 && row < len(t.rows) && col >= 0 && col < len(t.rows[row].cells) {
		t.rows[row].cells[col].colSpan = colSpan
		t.rows[row].cells[col].rowSpan = rowSpan
	}
}

// SetStyle sets table style
func (t *Table) SetStyle(s *TableStyle) {
	t.style = s
}

// SetRowStyle sets row style
func (t *Table) SetRowStyle(row int, s *RowStyle) {
	if row >= 0 && row < len(t.rows) {
		t.rows[row].style = s
	}
}

// SetColumnStyle sets column style
func (t *Table) SetColumnStyle(col int, s *ColumnStyle) {
	if col >= 0 && col < len(t.columns) {
		t.columns[col].style = s
	}
}

// SetCellStyle sets cell style
func (t *Table) SetCellStyle(row int, col int, s *CellStyle) {
	if row >= 0 && row < len(t.rows) && col >= 0 && col < len(t.rows[row].cells) {
		t.rows[row].cells[col].style = s
	}
}

// getFilesInfo returns files info
func (t *Table) getFilesInfo() []fileInfo {
	return []fileInfo{}
}

// generateStyles generates XML representation of the table styles
func (t *Table) generateStyles() string {
	var stylesBuffer bytes.Buffer

	// table style
	if t.style != nil {
		stylesBuffer.WriteString(t.style.generate())
	}

	for _, r := range t.rows {
		// rows styles
		if r.style != nil {
			stylesBuffer.WriteString(r.style.generate())
		}

		for _, c := range r.cells {
			// cells styles
			if c.style != nil {
				stylesBuffer.WriteString(c.style.generate())
			}

			// cells text styles
			if c.value.style != nil {
				stylesBuffer.WriteString(c.value.style.generate())
			}
		}
	}

	for _, c := range t.columns {
		// column styles
		if c.style != nil {
			stylesBuffer.WriteString(c.style.generate())
		}
	}

	return stylesBuffer.String()
}

// generate produces the XML representation
func (t *Table) generate() string {
	var buf bytes.Buffer

	buf.WriteString(`<table:table table:name="` + t.name + `"`)
	if t.style != nil {
		buf.WriteString(` table:style-name="` + t.style.getName() + `"`)
	}
	buf.WriteString(">")

	for _, col := range t.columns {
		buf.WriteString(`<table:table-column`)
		if col.style != nil {
			buf.WriteString(` table:style-name="` + col.style.getName() + `"`)
		}
		buf.WriteString("/>")
	}

	for _, row := range t.rows {
		buf.WriteString(`<table:table-row`)
		if row.style != nil {
			buf.WriteString(` table:style-name="` + row.style.getName() + `"`)
		}
		buf.WriteString(">\n")

		for _, cell := range row.cells {
			buf.WriteString(`<table:table-cell`)
			if cell.style != nil {
				buf.WriteString(` table:style-name="` + cell.style.getName() + `"`)
			}
			if cell.colSpan > 1 {
				buf.WriteString(` table:number-columns-spanned="` + fmt.Sprintf("%d", cell.colSpan) + `"`)
			}
			if cell.rowSpan > 1 {
				buf.WriteString(` table:number-rows-spanned="` + fmt.Sprintf("%d", cell.rowSpan) + `"`)
			}
			buf.WriteString(">")

			if cell.value.value != "" {
				cellStyle := ""
				if cell.value.style != nil {
					cellStyle = fmt.Sprintf(` text:style-name="%s"`, cell.value.style.getName())
				}
				buf.WriteString(fmt.Sprintf(`<text:p%s>`, cellStyle))
				buf.WriteString(cell.value.value)
				buf.WriteString(`</text:p>`)
			}

			buf.WriteString(`</table:table-cell>`)
			buf.WriteString("\n")
		}

		buf.WriteString(`</table:table-row>`)
		buf.WriteString("\n")
	}

	buf.WriteString(`</table:table>`)

	return buf.String()
}
