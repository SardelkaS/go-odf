package odt

import "testing"

func Test_RowStyle(t *testing.T) {
	t.Parallel()

	expectedRes := `<style:style style:name="TblRS1" style:family="table-row"><style:table-row-properties style:row-height="0.8cm" style:min-row-height="0.5cm" fo:background-color="#F5F5F5" fo:break-before="page" fo:keep-together="true" style:use-optimal-row-height="true"/></style:style>`

	rowStyle := NewRowStyle()
	rowStyle.
		WithHeight("0.8cm").
		WithMinHeight("0.5cm").
		WithBackground("#F5F5F5").
		WithBreakBefore("page").
		WithKeepTogether(true).
		WithOptimalHeight(true)

	expectEqual(t, expectedRes, rowStyle.generate())
}
