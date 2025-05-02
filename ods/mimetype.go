package ods

const (
	_defaultMimeType = "application/vnd.oasis.opendocument.text"
)

type mimeType struct {
	value string
}

// newMIME creates mimeType with default values
func newMIME() mimeType {
	return mimeType{
		value: _defaultMimeType,
	}
}

func (m mimeType) generate() string {
	return m.value
}
