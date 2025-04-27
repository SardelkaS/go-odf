package odt

import "testing"

func Test_Image(t *testing.T) {
	t.Parallel()

	data := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEBLAEsAAD/2wBDAAICAgICAQICAg"
	s := &Style{
		name: "caption style",
	}
	expectedRes := `<draw:frame draw:style-name="Im1" draw:name="Image1" text:anchor-type="paragraph" svg:width="200px" svg:height="300px" draw:z-index="0">
        					<draw:text-box fo:min-height="7.999cm">
                                <text:p text:style-name="Caption">
                                    <draw:frame draw:style-name="Caption"
                                        draw:name="Image2" text:anchor-type="paragraph" svg:width="7.938cm"
                                        style:rel-width="100%" svg:height="7.999cm" style:rel-height="scale"
                                        draw:z-index="2">
                                        <draw:image xlink:href="" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/>
                                    </draw:frame>
                                    <text:span text:style-name="caption style">caption</text:span>
                                </text:p>
                            </draw:text-box>
                </draw:frame>`

	img, err := NewImage(data)
	expectNoError(t, err)

	img.SetWidth("200px")
	img.SetHeight("300px")
	img.SetCaption("caption")
	img.SetCaptionStyle(s)
	img.SetPositionType("char")
	img.SetPositionHorizontal("left")
	img.SetPositionVertical("top")
	img.SetTextWrapType("parallel")
	img.SetTextWrapSide("left")

	expectEqual(t, cleanString(expectedRes), cleanString(img.generate()))
}
