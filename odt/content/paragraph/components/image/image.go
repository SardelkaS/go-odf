package image

import (
	"encoding/base64"
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components"
	"github.com/SardelkaS/go-odf/odt/model"
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
	altText          string
	contentType      string

	// Positioning configuration
	position position

	// Text wrapping behavior
	textWrap textWrap
}

const (
	_defaultWidth  = "100px"
	_defaultHeight = "100px"

	_picturesFolder = "Pictures"
)

var nameIter = atomic.Uint64{}

// New creates new image from base64 string
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
func New(data string) (*Image, error) {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 1
		nameIter.Add(1)
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
		altText:          "",
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

// SetAltText set alternative text for picture
func (i *Image) SetAltText(at string) {
	i.altText = at
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
func (i *Image) GetElementType() string {
	return components.ImageElement
}

// GetFileInfo returns path for saving image in zip and it's MIME-type
func (i *Image) GetFileInfo() model.FileInfo {
	if len(i.src) == 0 || i.name == "" || i.contentType == "" {
		return model.FileInfo{}
	}

	ext, ok := supportedImageTypes[i.contentType]
	if !ok {
		return model.FileInfo{}
	}

	data, err := base64.StdEncoding.DecodeString(cleanBase64Data(i.src))
	if err != nil {
		return model.FileInfo{}
	}

	return model.FileInfo{
		Path:        fmt.Sprintf("%s/%s%s", _picturesFolder, i.name, ext),
		ContentType: i.contentType,
		Data:        data,
	}
}

func (i *Image) getCaptionStyle() string {
	if i.altText == "" {
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

// GetStyle returns image style
func (i *Image) GetStyle() string {
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
	builder.WriteString(" " + i.getCaptionStyle())
	return builder.String()
}

// Generate generates xml code
func (i *Image) Generate() string {
	info := i.GetFileInfo()

	if i.altText != "" {
		return fmt.Sprintf(
			`<draw:frame draw:style-name="%s" draw:name="%s" text:anchor-type="paragraph" svg:width="%s" svg:height="%s" draw:z-index="0">
					<draw:text-box fo:min-height="7.999cm">
                        <text:p text:style-name="Caption">
                            <draw:frame draw:style-name="%s"
                                draw:name="Image2" text:anchor-type="paragraph" svg:width="7.938cm"
                                style:rel-width="100%%" svg:height="7.999cm" style:rel-height="scale"
                                draw:z-index="2">
                                <draw:image xlink:href="%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/>
                            </draw:frame>
                            %s
                        </text:p>
                    </draw:text-box>
        </draw:frame>`,
			i.styleName, i.name, i.width, i.height, i.captionStyleName, info.Path, i.altText)
	}

	return fmt.Sprintf(
		`<draw:frame draw:style-name="%s" draw:name="%s" text:anchor-type="paragraph" svg:width="%s" svg:height="%s" draw:z-index="0">
            <draw:image xlink:href="%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad"/>
        </draw:frame>`,
		i.styleName, i.name, i.width, i.height, info.Path)
}
