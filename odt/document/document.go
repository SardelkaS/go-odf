package document

import (
	"archive/zip"
	"bytes"
	"github.com/SardelkaS/go-odf/odt/content"
	"github.com/SardelkaS/go-odf/odt/content/paragraph"
	"github.com/SardelkaS/go-odf/odt/content/table"
	"github.com/SardelkaS/go-odf/odt/manifest"
	"github.com/SardelkaS/go-odf/odt/meta"
	"github.com/SardelkaS/go-odf/odt/mimetype"
	"github.com/SardelkaS/go-odf/odt/settings"
	"github.com/SardelkaS/go-odf/odt/styles"
	"io"
	"os"
	"strings"
)

type Document struct {
	Meta     *meta.Meta
	settings settings.Settings
	styles   styles.Styles
	content  *content.Content
	mimetype mimetype.MimeType
	manifest *manifest.Manifest
}

// New creates new empty Document
func New() Document {
	return Document{
		Meta:     meta.New(),
		settings: settings.New(),
		styles:   styles.New(),
		content:  content.New(),
		mimetype: mimetype.New(),
		manifest: manifest.New(),
	}
}

// Paragraph add new paragraph
func (d Document) Paragraph(p *paragraph.Paragraph) {
	d.content.Add(p)
}

// Table add new table
func (d Document) Table(t *table.Table) {
	d.content.Add(t)
}

// SaveToFile save generated data to file
//
// example SaveToFile("./files/test.odt")
func (d Document) SaveToFile(filePath string) error {
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
	filesInfo := d.content.GetFilesInfo()
	d.manifest.AddEntries(filesInfo)

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer func(zipWriter *zip.Writer) {
		_ = zipWriter.Close()
	}(zipWriter)

	files := map[string]string{
		_metaFileName: d.Meta.Generate(),
		//_settingsFileName: d.Settings.Generate(),
		_stylesFileName:   d.styles.Generate(),
		_contentFileName:  d.content.Generate(),
		_manifestFileName: d.manifest.Generate(),
	}

	mimetypeHeader := &zip.FileHeader{
		Name:   "mimetype",
		Method: zip.Store, // No compression
	}
	mimetypeWriter, err := zipWriter.CreateHeader(mimetypeHeader)
	if err != nil {
		return nil, err
	}
	_, err = mimetypeWriter.Write([]byte(d.mimetype.Generate()))
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

	for _, file := range filesInfo {
		zipFile, err := zipWriter.Create(file.Path)
		if err != nil {
			return nil, err
		}

		_, err = zipFile.Write(file.Data)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func formatXML(s string) string {
	return strings.Replace(strings.Replace(s, "\n", "", -1), "\t", "", -1)
}
