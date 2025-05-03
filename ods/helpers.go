package ods

import (
	"bytes"
	"errors"
	"regexp"
	"strings"
)

func convertFormula(inputFormula string) string {
	formula := strings.TrimPrefix(inputFormula, "=")
	re := regexp.MustCompile(`([A-Za-z]+)(\d+)`)

	odfFormula := re.ReplaceAllStringFunc(formula, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		col := strings.ToUpper(parts[1])
		row := parts[2]
		return "[." + col + row + "]"
	})

	var buf bytes.Buffer
	buf.WriteString("of:=")
	buf.WriteString(odfFormula)

	return buf.String()
}

func cellToIndices(cell string) (row, col int, err error) {
	re := regexp.MustCompile(`^([A-Za-z]+)(\d+)$`)
	matches := re.FindStringSubmatch(cell)
	if len(matches) != 3 {
		return 0, 0, errors.New("invalid cell format, expected like 'A1' or 'BC23'")
	}

	colStr := strings.ToUpper(matches[1])
	col = 0
	for _, ch := range colStr {
		col = col*26 + int(ch-'A') + 1
	}
	col--

	rowStr := matches[2]
	row = 0
	for _, digit := range rowStr {
		row = row*10 + int(digit-'0')
	}
	row--

	return row, col, nil
}
