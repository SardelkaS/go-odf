package odt

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

type position struct {
	Type       string `json:"type"`       // Reference type: "paragraph"/"page"/"char"/"frame"
	Anchor     string `json:"anchor"`     // Anchor type: matches Position.Type
	Horizontal string `json:"horizontal"` // "left"/"center"/"right"/"from-left"
	Vertical   string `json:"vertical"`   // "top"/"middle"/"bottom"/"from-top"
	XOffset    string `json:"x_offset"`   // Offset when using "from-left"
	YOffset    string `json:"y_offset"`   // Offset when using "from-top"
}

type textWrap struct {
	Type   string `json:"type"`   // "none"/"parallel"/"dynamic"
	Side   string `json:"side"`   // "left"/"right"/"biggest"/"both"
	Margin string `json:"margin"` // Space between text and image
}

type Image struct {
	name             string
	styleName        string
	captionStyleName string
	src              string
	width            string
	height           string
	caption          string
	captionStyle     *Style
	contentType      string

	// Positioning configuration
	position position

	// text wrapping behavior
	textWrap textWrap
}

const (
	_defaultWidth  = "100px"
	_defaultHeight = "100px"

	_picturesFolder = "Pictures"
)

var imgNameIter = atomic.Uint64{}

// NewImage creates new image from base64 string
//
// supported MIME-types:
//
//	image/png
//	image/jpeg
//	image/gif
//	image/svg+xml
//	image/bmp
//	image/webp
//	image/tiff
func NewImage(data string) (*Image, error) {
	iter := imgNameIter.Load()
	imgNameIter.Add(1)
	if iter == 0 {
		iter = 1
		imgNameIter.Add(1)
	}

	contentType, err := detectContentType(data)
	if err != nil {
		return nil, err
	}

	return &Image{
		name:             fmt.Sprintf("Image%s", strconv.FormatUint(iter, 10)),
		styleName:        fmt.Sprintf("Im%s", strconv.FormatUint(iter, 10)),
		captionStyleName: fmt.Sprintf("Imc%s", strconv.FormatUint(iter, 10)),
		src:              data,
		width:            _defaultWidth,
		height:           _defaultHeight,
		caption:          "",
		contentType:      contentType,
		position: position{
			Type:   PositionTypeParagraph,
			Anchor: PositionTypeParagraph,
		},
		textWrap: textWrap{
			Type: WrapNone,
		},
	}, nil
}

// SetWidth set picture width in px. Default value is 100px
func (i *Image) SetWidth(w string) {
	i.width = w
}

// SetHeight set picture height in px. Default value is 100px
func (i *Image) SetHeight(h string) {
	i.height = h
}

// SetCaption set caption text for picture
func (i *Image) SetCaption(at string) {
	i.caption = at
}

// SetCaptionStyle set caption text style
func (i *Image) SetCaptionStyle(s *Style) {
	i.captionStyle = s
}

// SetContentType set MIME-type of picture. It is determined automatically. Use this only if you are sure
func (i *Image) SetContentType(ct string) {
	i.contentType = ct
}

// SetPositionType set reference type: "paragraph"/"page"/"char"/"frame"
// default: paragraph
func (i *Image) SetPositionType(t string) {
	i.position.Type = t
	i.position.Anchor = t
}

// SetPositionHorizontal set horizontal alignment "left"/"center"/"right"/"from-left"
func (i *Image) SetPositionHorizontal(h string) {
	i.position.Horizontal = h
}

// SetPositionVertical set vertical alignment "top"/"middle"/"bottom"/"from-top"
func (i *Image) SetPositionVertical(v string) {
	i.position.Vertical = v
}

// SetPositionXOffset set offset when using "from-left"
func (i *Image) SetPositionXOffset(x string) {
	i.position.XOffset = x
}

// SetPositionYOffset set offset when using "from-top"
func (i *Image) SetPositionYOffset(y string) {
	i.position.YOffset = y
}

// SetTextWrapType set text wrap type: "none"/"parallel"/"dynamic"
// default: none
func (i *Image) SetTextWrapType(t string) {
	i.textWrap.Type = t
}

// SetTextWrapSide set text wrap side: "left"/"right"/"biggest"/"both"
func (i *Image) SetTextWrapSide(s string) {
	i.textWrap.Side = s
}

// SetTextWrapMargin set space between text and image
func (i *Image) SetTextWrapMargin(m string) {
	i.textWrap.Margin = m
}

// GetElementType returns element type
func (i *Image) getElementType() string {
	return _imageElement
}

// getFileInfo returns path for saving image in zip and it's MIME-type
func (i *Image) getFileInfo() fileInfo {
	if len(i.src) == 0 || i.name == "" || i.contentType == "" {
		return fileInfo{}
	}

	ext, ok := supportedImageTypes[i.contentType]
	if !ok {
		return fileInfo{}
	}

	data, err := base64.StdEncoding.DecodeString(cleanBase64Data(i.src))
	if err != nil {
		return fileInfo{}
	}

	return fileInfo{
		Path:        fmt.Sprintf("%s/%s%s", _picturesFolder, i.name, ext),
		ContentType: i.contentType,
		Data:        data,
	}
}

func (i *Image) getCaptionStyle() string {
	if i.captionStyle == nil {
		return ""
	}

	return i.captionStyle.generate()
}

func (i *Image) getCaptionFrameStyle() string {
	if i.caption == "" {
		return ""
	}

	return fmt.Sprintf(`<style:style style:name="%s" style:family="graphic" style:parent-style-name="Graphics">
            <style:graphic-properties fo:margin-left="0cm" fo:margin-right="0cm" fo:margin-top="0cm"
                fo:margin-bottom="0cm" style:run-through="foreground" style:wrap="none"
                style:vertical-pos="top" style:vertical-rel="paragraph-content"
                style:horizontal-pos="center" style:horizontal-rel="paragraph-content"
                fo:padding="0cm" fo:border="none" style:shadow="none" draw:shadow-opacity="100%%"
                style:mirror="none" fo:clip="rect(0cm, 0cm, 0cm, 0cm)" draw:luminance="0%%"
                draw:contrast="0%%" draw:red="0%%" draw:green="0%%" draw:blue="0%%" draw:gamma="100%%"
                draw:color-inversion="false" draw:image-opacity="100%%" draw:color-mode="standard"
                loext:rel-width-rel="paragraph" />
        </style:style>`, i.captionStyleName)
}

// generateStyles returns image style
func (i *Image) generateStyles() string {
	var builder strings.Builder
	builder.WriteString(`<style:style style:name="` + i.styleName + `" style:family="graphic" `)
	builder.WriteString(`style:parent-style-name="Graphics">`)
	builder.WriteString(`<style:graphic-properties `)

	// Positioning properties
	if i.position.Anchor != "" {
		builder.WriteString(`text:anchor-type="` + i.position.Anchor + `" `)
	}
	if i.position.Horizontal != "" {
		builder.WriteString(`style:horizontal-pos="` + i.position.Horizontal + `" `)
	}
	if i.position.Vertical != "" {
		builder.WriteString(`style:vertical-pos="` + i.position.Vertical + `" `)
	}
	if i.position.XOffset != "" && i.position.Horizontal == "from-left" {
		builder.WriteString(`svg:x="` + i.position.XOffset + `" `)
	}
	if i.position.YOffset != "" && i.position.Vertical == "from-top" {
		builder.WriteString(`svg:y="` + i.position.YOffset + `" `)
	}

	// Wrapping properties
	if i.textWrap.Type != "" {
		builder.WriteString(`style:wrap="` + i.textWrap.Type + `" `)
	}
	if i.textWrap.Side != "" {
		builder.WriteString(`fo:wrap-contour="` + i.textWrap.Side + `" `)
	}
	if i.textWrap.Margin != "" {
		builder.WriteString(`fo:margin="` + i.textWrap.Margin + `" `)
	}

	builder.WriteString(`/></style:style>`)
	builder.WriteString(" " + i.getCaptionFrameStyle())
	builder.WriteString(" " + i.getCaptionStyle())
	return builder.String()
}

// generate generates xml code
func (i *Image) generate() string {
	info := i.getFileInfo()

	if i.caption != "" {
		caption := i.caption
		if i.captionStyle != nil {
			caption = fmt.Sprintf(`<text:span text:style-name="%s">%s</text:span>`,
				i.captionStyle.getName(), i.caption)
		}

		return fmt.Sprintf(
			`<draw:frame draw:style-name="%s" draw:name="%s" text:anchor-type="paragraph" svg:width="%s" svg:height="%s" draw:z-index="0">
					<draw:text-box fo:min-height="7.999cm">
                        <text:p text:style-name="Caption">
                            <draw:frame draw:style-name="Caption"
                                draw:name="Image2" text:anchor-type="paragraph" svg:width="7.938cm"
                                style:rel-width="100%%" svg:height="7.999cm" style:rel-height="scale"
                                draw:z-index="2">
                                <draw:image xlink:href="%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/>
                            </draw:frame>
                            %s
                        </text:p>
                    </draw:text-box>
        </draw:frame>`,
			i.styleName, i.name, i.width, i.height, info.Path, caption)
	}

	return fmt.Sprintf(
		`<draw:frame draw:style-name="%s" draw:name="%s" text:anchor-type="paragraph" svg:width="%s" svg:height="%s" draw:z-index="0">
            <draw:image xlink:href="%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/>
        </draw:frame>`,
		i.styleName, i.name, i.width, i.height, info.Path)
}
