package ods

type fileInfo struct {
	Path        string
	ContentType string
	Data        []byte
}

func (f fileInfo) Valid() bool {
	return f.Path != "" && f.ContentType != "" && len(f.Data) > 0
}
