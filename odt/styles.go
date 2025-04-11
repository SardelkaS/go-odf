package odt

type styles struct{}

func newStyles() styles {
	return styles{}
}

// generate generates xml code
func (s styles) generate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" 
    xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" 
    xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" 
    xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0">
    <office:styles>
        <style:style style:name="Standard" style:family="Paragraph" style:class="text">
            <style:text-properties fo:font-size="12pt" style:font-name="Liberation Sans"/>
        </style:style>
    </office:styles>
</office:document-styles>`
}
