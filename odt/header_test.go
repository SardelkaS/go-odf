package odt

import "testing"

func Test_Header(t *testing.T) {
	t.Parallel()

	expectedRes := `<text:h text:style-name="T2" text:outline-level="1"><text:bookmark-start text:name="__RefHeading__1" />test<text:bookmark-end text:name="__RefHeading__1" /></text:h>`
	s := &Style{}

	h := NewHeader().WithText("test").WithStyle(s).WithLevel(1)
	h.style.name = "T2"
	expectEqual(t, expectedRes, h.generate())
}
