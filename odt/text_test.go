package odt

import "testing"

func Test_Text(t *testing.T) {
	t.Parallel()

	s := &Style{
		name: "text1",
	}
	expectedRes := `<text:span text:style-name="text1">text</text:span>`

	txt := newText("text", s)

	expectEqual(t, expectedRes, txt.generate())
}
