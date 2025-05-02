package ods

import (
	"bytes"
	"fmt"
	"strconv"
)

type column struct {
	style                *ColumnStyle
	defaultCellStyle     *CellStyle
	defaultCellStyleName string
	repeatedCount        int
}

type cell struct {
	style     *CellStyle
	value     string
	valueType string
	currency  string
}

type row struct {
	style         *RowStyle
	cells         []*cell
	repeatedCount int
}

// Sheet represents table sheet in ODS document
type Sheet struct {
	name    string
	style   *SheetStyle
	columns []*column
	rows    []*row
	lastCol *column
	lastRow *row
}

// NewSheet creates new empty Sheet with provided name
func NewSheet(name string) *Sheet {
	return &Sheet{
		name:    name,
		style:   NewSheetStyle().WithGridDisplay(true),
		columns: make([]*column, 0),
		rows:    make([]*row, 0),
		lastCol: &column{
			style:                NewColumnStyle().WithBreakBefore("auto").WithWidth(3),
			defaultCellStyleName: _defaultCellStyleName,
			repeatedCount:        16384,
		},
		lastRow: &row{
			style:         NewRowStyle().WithBreakBefore("auto").WithUseOptimal(true),
			cells:         make([]*cell, 0),
			repeatedCount: 16384,
		},
	}
}

// SetColumnStyle sets column style
func (s *Sheet) SetColumnStyle(num int, cs *ColumnStyle) {
	s.expandColumns(num)
	s.columns[num].style = cs.copy()
}

// SetColumnDefaultCellStyle sets default cell style for column
func (s *Sheet) SetColumnDefaultCellStyle(num int, cs *CellStyle) {
	s.expandColumns(num)
	newCS := cs.copy()
	s.columns[num].defaultCellStyle = newCS
	s.columns[num].defaultCellStyleName = newCS.name
}

// SetRowStyle sets row style
func (s *Sheet) SetRowStyle(num int, rs *RowStyle) {
	s.expandRows(num)
	s.rows[num].style = rs.copy()
}

// SetCellStyle sets cell style for cell in row r
func (s *Sheet) SetCellStyle(r, num int, cs *CellStyle) {
	s.expandCells(r, num)
	s.rows[r].cells[num].style = cs.copy()
}

// SetCellValue sets cell value
//
// Possible value types:
//   - "float"      : Numeric values (e.g., 42 or 3.14).
//   - "percentage" : Percentage values (stored as 0.95 for 95%).
//   - "currency"   : Monetary values (requires currency parameter).
//   - "date"       : Dates in YYYY-MM-DD format.
//   - "time"       : Durations in ISO 8601 format (e.g., "PT13H45M" for 13:45).
//   - "boolean"    : True/false values.
//   - "string"     : Text content.
//   - "void"       : Empty cell (may still carry styles).
//
// For value type "currency" you need to pass the additional parameter "currency" (example: "USD", "EUR")
func (s *Sheet) SetCellValue(r, num int, value string, valueType string, currency ...string) {
	s.expandCells(r, num)
	switch valueType {
	case Float, Percentage, Date, Time, Boolean, String, Void:
		s.rows[r].cells[num].valueType = valueType
		s.rows[r].cells[num].value = value
	case Currency:
		if len(currency) == 0 {
			return
		}
		s.rows[r].cells[num].valueType = valueType
		s.rows[r].cells[num].value = value
		s.rows[r].cells[num].currency = currency[0]
	}
}

// SetSheetStyle sets sheet style
func (s *Sheet) SetSheetStyle(ss *SheetStyle) {
	s.style = ss.copy()
}

func (s *Sheet) expandColumns(num int) {
	for {
		if len(s.columns) > num {
			break
		}

		s.columns = append(s.columns, &column{
			style:                NewColumnStyle().WithBreakBefore("auto").WithWidth(3),
			defaultCellStyleName: _defaultCellStyleName,
			repeatedCount:        1,
		})
	}
}

func (s *Sheet) expandRows(num int) {
	for {
		if len(s.rows) > num {
			break
		}

		s.rows = append(s.rows, &row{
			style:         NewRowStyle().WithBreakBefore("auto").WithUseOptimal(true),
			cells:         make([]*cell, 0),
			repeatedCount: 1,
		})
	}
}

func (s *Sheet) expandCells(r int, num int) {
	s.expandRows(r)

	for {
		if len(s.rows[r].cells) > num {
			break
		}

		s.rows[r].cells = append(s.rows[r].cells, &cell{
			valueType: "string",
			value:     "",
		})
	}
}

func (s *Sheet) generateStyles() string {
	var stylesBuffer bytes.Buffer
	if s.style != nil {
		stylesBuffer.WriteString(s.style.generate())
	}
	for _, c := range s.columns {
		if c.style != nil {
			stylesBuffer.WriteString(c.style.generate())
		}
		if c.defaultCellStyle != nil {
			stylesBuffer.WriteString(c.defaultCellStyle.generate())
		}
	}

	for _, r := range s.rows {
		if r.style != nil {
			stylesBuffer.WriteString(r.style.generate())
		}

		for _, c := range r.cells {
			if c.style != nil {
				stylesBuffer.WriteString(c.style.generate())
			}
		}
	}

	return stylesBuffer.String()
}

func (s *Sheet) generate() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(` <table:table table:name="%s" table:style-name="%s">`, s.name, s.style.name))

	for _, c := range s.columns {
		buf.WriteString(fmt.Sprintf(`<table:table-column table:style-name="%s" table:number-columns-repeated="%s" table:default-cell-style-name="%s" />`,
			c.style.name, strconv.FormatInt(int64(c.repeatedCount), 10), c.defaultCellStyleName))
	}

	if s.lastCol != nil {
		buf.WriteString(fmt.Sprintf(`<table:table-column table:style-name="%s" table:number-columns-repeated="%s" table:default-cell-style-name="%s" />`,
			s.lastCol.style.name, strconv.FormatInt(int64(s.lastCol.repeatedCount), 10), s.lastCol.defaultCellStyleName))
	}

	for _, r := range s.rows {
		buf.WriteString(fmt.Sprintf(`<table:table-row table:style-name="%s">`, r.style.name))
		for _, c := range r.cells {
			cellStyle := ""
			if c.style != nil {
				cellStyle = fmt.Sprintf(`table:style-name="%s" `, c.style.name)
			}

			switch c.valueType {
			case Float, Percentage:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="%s" office:value="%s"/>`, cellStyle, c.valueType, c.value))
			case Currency:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="currency" office:value="%s" office:currency="%s"/>`, cellStyle, c.value, c.currency))
			case Date:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="date" office:date-value="%s"/>`, cellStyle, c.value))
			case Time:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="time" office:time-value="%s"/>`, cellStyle, c.value))
			case Boolean:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="boolean" office:boolean-value="%s"/>`, cellStyle, c.value))
			case String:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="string"><text:p>%s</text:p></table:table-cell>`, cellStyle, c.value))
			case Void:
				buf.WriteString(fmt.Sprintf(`<table:table-cell %soffice:value-type="void"/>`, cellStyle))
			default:
			}
		}
		buf.WriteString(fmt.Sprintf(`<table:table-cell table:number-columns-repeated="%s" />`, strconv.FormatInt(int64(16384)-int64(len(r.cells)), 10)))
		buf.WriteString(`</table:table-row>`)
	}

	if s.lastRow != nil {
		buf.WriteString(fmt.Sprintf(`<table:table-row table:style-name="%s" table:number-rows-repeated="1048571"><table:table-cell table:number-columns-repeated="%s" /></table:table-row>`,
			s.lastRow.style.name, strconv.FormatInt(int64(s.lastRow.repeatedCount), 10)))
	}

	buf.WriteString(`</table:table>`)

	return buf.String()
}
