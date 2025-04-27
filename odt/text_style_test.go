package odt

import "testing"

func Test_TextStyle(t *testing.T) {
	t.Parallel()

	expectedRes := `<style:style style:name="T2" style:family="text"><style:text-properties style:font-name="Arial" fo:font-size="10pt" fo:font-weight="bold" fo:font-style="italic" style:text-underline-style="solid" style:text-underline-width="auto" style:text-underline-color="font-color" fo:color="#FF0000" fo:text-shadow="1pt 1pt 0.3pt #AAAAAA" fo:letter-spacing="0.3mm" fo:text-transform="uppercase" style:text-overline-style="solid" style:text-overline-color="#00FF00" style:text-line-through-style="solid" fo:text-outline="0.5pt #FFFFFF" style:text-emphasize="filled dot below" style:writing-mode="tb-rl" style:text-rotation-scale="line-height" style:text-rotation-angle="90"/></style:style>`

	ts := NewTextStyle()
	ts.
		WithFontName("Arial").
		WithFontSize("10pt").
		WithBold().
		WithItalic().
		WithColor("#FF0000").
		WithTextShadow("1pt 1pt 0.3pt #AAAAAA").
		WithLetterSpacing("0.3mm").
		WithTextTransform(TransformUppercase).
		WithUnderline().
		WithOverline(OverlineSingle, "#00FF00").
		WithLineThrough(LineThroughSolid).
		WithTextOutline("0.5pt #FFFFFF").
		WithTextEmphasis("filled dot below").
		WithWritingMode("tb-rl").
		WithRotation(90, RotationScaleLineHeight)
	ts.name = "T2"

	expectEqual(t, expectedRes, ts.generate())
}
