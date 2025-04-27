package odt

import "testing"

func Test_TableStyle(t *testing.T) {
	t.Parallel()

	expectedRes := `<style:style style:name="TblS1" style:family="table"><style:table-properties style:width="100px" table:align="left" fo:margin="1px 1px 1px 1px" fo:background-color="#E6E6E6" table:border-model="collapsing" fo:border="0.002cm solid #000000"/></style:style>`

	ts := NewTableStyle()
	ts.
		WithWidth("100px").
		WithAlign("left").
		WithMargin("1px 1px 1px 1px").
		WithBackgroundColor("#E6E6E6").
		WithBorderModel(BorderModelCollapsing).
		WithBorder("0.002cm solid #000000")

	expectEqual(t, expectedRes, ts.generate())
}
