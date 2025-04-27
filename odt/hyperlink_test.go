package odt

import "testing"

func Test_Hyperlink(t *testing.T) {
	t.Parallel()

	s := &Style{
		name: "style",
	}
	vs := &Style{
		name: "visited style",
	}
	expectedRes := `<text:a xlink:type="simple" xlink:href="test-link.com" text:style-name="style" text:visited-style-name="visited style">test</text:a>`

	h := NewHyperLink()
	h.
		WithText("test").
		WithLink("test-link.com").
		WithStyle(s).
		WithVisitedStyle(vs)

	expectEqual(t, expectedRes, h.generate())
}
