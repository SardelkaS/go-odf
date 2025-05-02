package ods

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"
)

type Document struct {
	Meta     *meta
	settings settings
	styles   *styles
	content  *content
	mimetype mimeType
	manifest *manifest
}

// New creates new empty Document
func New() Document {
	return Document{
		Meta:     newMeta(),
		settings: newSettings(),
		styles:   newStyles(),
		content:  newContent(),
		mimetype: newMIME(),
		manifest: newManifest(),
	}
}

// Sheet adds new Sheet
func (d Document) Sheet(p *Sheet) {
	d.content.add(p)
}

// SaveToFile save generated data to file
//
// example SaveToFile("./files/test.odt")
func (d Document) SaveToFile(filePath string) error {
	lastSlashIdx := strings.LastIndex(filePath, "/")
	if lastSlashIdx != -1 {
		path := filePath[lastSlashIdx:]
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, os.ModeDir)
			if err != nil {
				return err
			}
		}
	}

	buf, err := d.GetBytes()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, buf.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

// GetBytes return generated file in bytes.Buffer
func (d Document) GetBytes() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer func(zipWriter *zip.Writer) {
		_ = zipWriter.Close()
	}(zipWriter)

	files := map[string]string{
		_metaFileName: d.Meta.generate(),
		//_settingsFileName: d.settings.generate(),
		_stylesFileName:   d.styles.generate(),
		_contentFileName:  d.content.generate(),
		_manifestFileName: d.manifest.generate(),
	}

	mimetypeHeader := &zip.FileHeader{
		Name:   "mimetype",
		Method: zip.Store, // No compression
	}
	mimetypeWriter, err := zipWriter.CreateHeader(mimetypeHeader)
	if err != nil {
		return nil, err
	}
	_, err = mimetypeWriter.Write([]byte(d.mimetype.generate()))
	if err != nil {
		return nil, err
	}

	for file, data := range files {
		zipFile, err := zipWriter.Create(file)
		if err != nil {
			return nil, err
		}

		_, err = io.WriteString(zipFile, data)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}
