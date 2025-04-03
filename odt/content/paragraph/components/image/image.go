package image

import (
	"encoding/base64"
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/paragraph/components"
	"github.com/SardelkaS/go-odf/odt/model"
	"strconv"
	"sync/atomic"
)

type Image struct {
	name        string
	src         string
	width       string
	height      string
	altText     string
	contentType string
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
		iter = 2
		nameIter.Add(2)
	}

	contentType, err := detectContentType(data)
	if err != nil {
		return nil, err
	}

	return &Image{
		name:        fmt.Sprintf("Image%s", strconv.FormatUint(iter, 10)),
		src:         data,
		width:       _defaultWidth,
		height:      _defaultHeight,
		altText:     "",
		contentType: contentType,
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

// Generate generates xml code
func (i *Image) Generate() string {
	info := i.GetFileInfo()
	return fmt.Sprintf(
		`<draw:frame draw:name="%s" text:anchor-type="paragraph" svg:width="%s" svg:height="%s" draw:z-index="0">
            <draw:image xlink:href="%s" xlink:type="simple" xlink:show="embed" xlink:actuate="onLoad">
                <text:p>%s</text:p>
            </draw:image>
        </draw:frame>`,
		i.name, i.width, i.height, info.Path, i.altText)
}
