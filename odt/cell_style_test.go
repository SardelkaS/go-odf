package odt

import (
	"testing"
)

func Test_CellStyle(t *testing.T) {
	t.Parallel()

	expectedRes := `<style:style style:name="TblCellS1" style:family="table-cell"><style:table-cell-properties fo:background-color="#E6E6E6" fo:border="0.5pt solid #CCCCCC" fo:border-top="1pt solid #000" fo:border-bottom="none" fo:border-left="1pt solid #000" fo:border-right="none" fo:padding="0.2cm" fo:padding-top="0.3cm" fo:padding-bottom="0.1cm" fo:padding-left="0.2cm" fo:padding-right="0.2cm" fo:text-align="center" style:vertical-align="middle" fo:wrap-option="wrap" style:rotation-angle="90" style:shrink-to-fit="true"/></style:style>`

	cellStyle := NewCellStyle()
	cellStyle.
		WithBackground("#E6E6E6").
		WithBorder("0.5pt solid #CCCCCC").
		WithIndividualBorders("1pt solid #000", "none", "1pt solid #000", "none").
		WithPadding("0.2cm").
		WithIndividualPadding("0.3cm", "0.1cm", "0.2cm", "0.2cm").
		WithAlignment("center", "middle").
		WithTextWrap("wrap").
		WithTextRotation(90).
		WithShrinkToFit(true)

	expectEqual(t, expectedRes, cellStyle.generate())
}
