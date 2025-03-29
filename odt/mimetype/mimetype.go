package mimetype

const (
	_defaultMimeType = "application/vnd.oasis.opendocument.text"
)

type MimeType struct {
	value string
}

// New creates MimeType with default values
func New() MimeType {
	return MimeType{
		value: _defaultMimeType,
	}
}

func (m MimeType) Generate() string {
	return m.value
}
