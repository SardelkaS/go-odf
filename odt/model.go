package odt

type FileInfo struct {
	Path        string
	ContentType string
	Data        []byte
}

func (f FileInfo) Valid() bool {
	return f.Path != "" && f.ContentType != "" && len(f.Data) > 0
}
