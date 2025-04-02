package styles

type Styles struct{}

func New() Styles {
	return Styles{}
}

// Generate generates xml code
func (s Styles) Generate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" 
    xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" 
    xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" 
    xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0">
    <office:styles>
        <style:style style:name="Standard" style:family="paragraph" style:class="text">
            <style:text-properties fo:font-size="12pt" style:font-name="Liberation Sans"/>
        </style:style>
    </office:styles>
</office:document-styles>`
}
