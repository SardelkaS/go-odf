package table

import (
	"bytes"
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components/text/style"
	table_style "github.com/SardelkaS/go-odf/odt/content/table/style"
	"github.com/SardelkaS/go-odf/odt/model"
	"strconv"
	"sync/atomic"
)

var nameIter = atomic.Uint64{}

// Table represents an ODF table
type Table struct {
	Name    string
	Rows    []*Row
	Columns []*Column
	Style   *table_style.Style
}

// Row represents a table row
type Row struct {
	Cells []*Cell
	Style *table_style.RowStyle
	name  string
}

// Column represents a table column
type Column struct {
	Style *table_style.ColumnStyle
	name  string
}

// Cell represents a table cell
type Cell struct {
	Value   CellValue
	Style   *table_style.CellStyle
	ColSpan int
	RowSpan int
	name    string
}

// CellValue represents a table cell value
type CellValue struct {
	Value string
	Style *style.Style
}

func New(r, c int) *Table {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 1
		nameIter.Add(1)
	}

	rows := make([]*Row, r)
	cols := make([]*Column, c)
	for i := range rows {
		row := &Row{
			Cells: make([]*Cell, c),
		}
		for j := range row.Cells {
			row.Cells[j] = &Cell{}
		}

		rows[i] = row
	}
	for i := range cols {
		cols[i] = &Column{}
	}

	return &Table{
		Name:    fmt.Sprintf("Table%s", strconv.FormatUint(iter, 10)),
		Rows:    rows,
		Columns: cols,
	}
}

// SetValue sets cell value
func (t *Table) SetValue(row, col int, value string) {
	if row >= 0 && row < len(t.Rows) && col >= 0 && col < len(t.Rows[row].Cells) {
		t.Rows[row].Cells[col].Value.Value = value
	}
}

// SetSpan sets cell span
func (t *Table) SetSpan(row, col, colSpan, rowSpan int) {
	if row >= 0 && row < len(t.Rows) && col >= 0 && col < len(t.Rows[row].Cells) {
		t.Rows[row].Cells[col].ColSpan = colSpan
		t.Rows[row].Cells[col].RowSpan = rowSpan
	}
}

// SetStyle sets table style
func (t *Table) SetStyle(s *table_style.Style) {
	t.Style = s
}

// SetRowStyle sets row style
func (t *Table) SetRowStyle(row int, s *table_style.RowStyle) {
	if row >= 0 && row < len(t.Rows) {
		t.Rows[row].Style = s
	}
}

// SetColumnStyle sets column style
func (t *Table) SetColumnStyle(col int, s *table_style.ColumnStyle) {
	if col >= 0 && col < len(t.Columns) {
		t.Columns[col].Style = s
	}
}

// SetCellStyle sets cell style
func (t *Table) SetCellStyle(row int, col int, s *table_style.CellStyle) {
	if row >= 0 && row < len(t.Rows) && col >= 0 && col < len(t.Rows[row].Cells) {
		t.Rows[row].Cells[col].Style = s
	}
}

// GetFilesInfo returns files info
func (t *Table) GetFilesInfo() []model.FileInfo {
	return []model.FileInfo{}
}

// GenerateStyles generates XML representation of the table styles
func (t *Table) GenerateStyles() string {
	var stylesBuffer bytes.Buffer

	// table style
	if t.Style != nil {
		stylesBuffer.WriteString(t.Style.Generate())
	}

	for _, r := range t.Rows {
		// rows styles
		if r.Style != nil {
			stylesBuffer.WriteString(r.Style.Generate())
		}

		for _, c := range r.Cells {
			// cells styles
			if c.Style != nil {
				stylesBuffer.WriteString(c.Style.Generate())
			}

			// cells text styles
			if c.Value.Style != nil {
				stylesBuffer.WriteString(c.Value.Style.Generate())
			}
		}
	}

	for _, c := range t.Columns {
		// column styles
		if c.Style != nil {
			stylesBuffer.WriteString(c.Style.Generate())
		}
	}

	return stylesBuffer.String()
}

// Generate produces the XML representation
func (t *Table) Generate() string {
	var buf bytes.Buffer

	buf.WriteString(`<table:table table:name="` + t.Name + `"`)
	if t.Style != nil {
		buf.WriteString(` table:style-name="` + t.Style.GetName() + `"`)
	}
	buf.WriteString(">")

	for _, col := range t.Columns {
		buf.WriteString(`<table:table-column`)
		if col.Style != nil {
			buf.WriteString(` table:style-name="` + col.Style.GetName() + `"`)
		}
		buf.WriteString("/>")
	}

	for _, row := range t.Rows {
		buf.WriteString(`<table:table-row`)
		if row.Style != nil {
			buf.WriteString(` table:style-name="` + row.Style.GetName() + `"`)
		}
		buf.WriteString(">\n")

		for _, cell := range row.Cells {
			buf.WriteString(`<table:table-cell`)
			if cell.Style != nil {
				buf.WriteString(` table:style-name="` + cell.Style.GetName() + `"`)
			}
			if cell.ColSpan > 1 {
				buf.WriteString(` table:number-columns-spanned="` + fmt.Sprintf("%d", cell.ColSpan) + `"`)
			}
			if cell.RowSpan > 1 {
				buf.WriteString(` table:number-rows-spanned="` + fmt.Sprintf("%d", cell.RowSpan) + `"`)
			}
			buf.WriteString(">")

			if cell.Value.Value != "" {
				cellStyle := ""
				if cell.Value.Style != nil {
					cellStyle = fmt.Sprintf(` text:style-name="%s"`, cell.Value.Style.GetName())
				}
				buf.WriteString(fmt.Sprintf(`<text:p%s>`, cellStyle))
				buf.WriteString(cell.Value.Value)
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
