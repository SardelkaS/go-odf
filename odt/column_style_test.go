package odt

import "testing"

func Test_ColumnStyle(t *testing.T) {
	t.Parallel()

	expectedRes := `<style:style style:name="TblCS1" style:family="table-column"><style:table-column-properties style:column-width="3.5cm" style:rel-column-width="3*" fo:background-color="#F8F8F8" fo:border="0.05cm solid #CCCCCC" fo:break-before="column" style:use-optimal-column-width="true"/></style:style>`

	colStyle := NewColumnStyle()
	colStyle.
		WithWidth("3.5cm").
		WithRelativeWidth("3*").
		WithBackground("#F8F8F8").
		WithBorder("0.05cm solid #CCCCCC").
		WithBreakBefore("column").
		WithOptimalWidth(true)

	expectEqual(t, expectedRes, colStyle.generate())
}
