package odt

import (
	"bytes"
	"fmt"
)

type manifest struct {
	additionalEntries []FileInfo
}

func newManifest() *manifest {
	return &manifest{
		additionalEntries: []FileInfo{},
	}
}

func (m *manifest) addEntries(entries []FileInfo) {
	m.additionalEntries = append(m.additionalEntries, entries[:]...)
}

// generate generates xml code
func (m *manifest) generate() string {
	var additionalEntries bytes.Buffer
	for _, e := range m.additionalEntries {
		additionalEntries.WriteString("\n\t" + fmt.Sprintf(`<manifest:file-entry manifest:full-path="%s" manifest:media-type="%s"/>`, e.Path, e.ContentType))
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<manifest:manifest xmlns:manifest="urn:oasis:names:tc:opendocument:xmlns:manifest:1.0">
    <manifest:file-entry manifest:full-path="/" manifest:media-type="application/vnd.oasis.opendocument.text"/>
    <manifest:file-entry manifest:full-path="content.xml" manifest:media-type="text/xml"/>
    <manifest:file-entry manifest:full-path="styles.xml" manifest:media-type="text/xml"/>
	<manifest:file-entry manifest:full-path="meta.xml" manifest:media-type="text/xml"/>%s
</manifest:manifest>`, additionalEntries.String())
}
