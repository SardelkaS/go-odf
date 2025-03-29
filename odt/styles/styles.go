package styles

type Styles struct{}

func New() Styles {
	return Styles{}
}

// Generate generates xml code
func (s Styles) Generate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
    xmlns:ooo="http://openoffice.org/2004/office"
    xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0"
    xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/"
    xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0"
    xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0"
    xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0"
    xmlns:rpt="http://openoffice.org/2005/report"
    xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0"
    xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0"
    xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0"
    xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0"
    xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0"
    xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0"
    xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc"
    xmlns:css3t="http://www.w3.org/TR/css3-text/"
    xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2"
    xmlns:tableooo="http://openoffice.org/2009/table"
    xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0"
    xmlns:drawooo="http://openoffice.org/2010/draw" xmlns:xhtml="http://www.w3.org/1999/xhtml"
    xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0"
    xmlns:grddl="http://www.w3.org/2003/g/data-view#"
    xmlns:field="urn:openoffice:names:experimental:ooo-ms-interop:xmlns:field:1.0"
    xmlns:math="http://www.w3.org/1998/Math/MathML"
    xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0"
    xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0"
    xmlns:dom="http://www.w3.org/2001/xml-events"
    xmlns:officeooo="http://openoffice.org/2009/office" office:version="1.4">
    <office:font-face-decls>
        <style:font-face style:name="Liberation Sans" svg:font-family="&apos;Liberation Sans&apos;"
            style:font-family-generic="swiss" style:font-pitch="variable" />
        <style:font-face style:name="Liberation Serif"
            svg:font-family="&apos;Liberation Serif&apos;" style:font-family-generic="roman"
            style:font-pitch="variable" />
        <style:font-face style:name="Lucida Sans" svg:font-family="&apos;Lucida Sans&apos;"
            style:font-family-generic="swiss" />
        <style:font-face style:name="Lucida Sans1" svg:font-family="&apos;Lucida Sans&apos;"
            style:font-family-generic="system" style:font-pitch="variable" />
        <style:font-face style:name="Microsoft YaHei" svg:font-family="&apos;Microsoft YaHei&apos;"
            style:font-family-generic="system" style:font-pitch="variable" />
        <style:font-face style:name="NSimSun" svg:font-family="NSimSun"
            style:font-family-generic="system" style:font-pitch="variable" />
    </office:font-face-decls>
    <office:styles>
        <style:default-style style:family="graphic">
            <style:graphic-properties svg:stroke-color="#3465a4" draw:fill-color="#729fcf"
                fo:wrap-option="no-wrap" draw:shadow-offset-x="0.3cm" draw:shadow-offset-y="0.3cm"
                draw:start-line-spacing-horizontal="0.283cm"
                draw:start-line-spacing-vertical="0.283cm"
                draw:end-line-spacing-horizontal="0.283cm" draw:end-line-spacing-vertical="0.283cm"
                style:flow-with-text="false" />
            <style:paragraph-properties style:text-autospace="ideograph-alpha"
                style:line-break="strict" loext:tab-stop-distance="0cm" style:writing-mode="lr-tb"
                style:font-independent-line-spacing="false">
                <style:tab-stops />
            </style:paragraph-properties>
            <style:text-properties style:use-window-font-color="true" loext:opacity="0%"
                style:font-name="Liberation Serif" fo:font-size="12pt" fo:language="ru"
                fo:country="RU" style:letter-kerning="true" style:font-name-asian="NSimSun"
                style:font-size-asian="10.5pt" style:language-asian="zh" style:country-asian="CN"
                style:font-name-complex="Lucida Sans1" style:font-size-complex="12pt"
                style:language-complex="hi" style:country-complex="IN" />
        </style:default-style>
        <style:default-style style:family="paragraph">
            <style:paragraph-properties fo:orphans="2" fo:widows="2"
                fo:hyphenation-ladder-count="no-limit" fo:hyphenation-keep="auto"
                loext:hyphenation-keep-type="column" style:text-autospace="ideograph-alpha"
                style:punctuation-wrap="hanging" style:line-break="strict"
                style:tab-stop-distance="1.251cm" style:writing-mode="page" />
            <style:text-properties style:use-window-font-color="true" loext:opacity="0%"
                style:font-name="Liberation Serif" fo:font-size="12pt" fo:language="ru"
                fo:country="RU" style:letter-kerning="true" style:font-name-asian="NSimSun"
                style:font-size-asian="10.5pt" style:language-asian="zh" style:country-asian="CN"
                style:font-name-complex="Lucida Sans1" style:font-size-complex="12pt"
                style:language-complex="hi" style:country-complex="IN" fo:hyphenate="false"
                fo:hyphenation-remain-char-count="2" fo:hyphenation-push-char-count="2"
                loext:hyphenation-no-caps="false" loext:hyphenation-no-last-word="false"
                loext:hyphenation-word-char-count="5" loext:hyphenation-zone="no-limit" />
        </style:default-style>
        <style:default-style style:family="table">
            <style:table-properties table:border-model="collapsing" />
        </style:default-style>
        <style:default-style style:family="table-row">
            <style:table-row-properties fo:keep-together="auto" />
        </style:default-style>
        <style:style style:name="Standard" style:family="paragraph" style:class="text" />
        <style:style style:name="Heading" style:family="paragraph"
            style:parent-style-name="Standard" style:next-style-name="Text_20_body"
            style:class="chapter">
            <style:paragraph-properties fo:margin-top="0.423cm" fo:margin-bottom="0.212cm"
                style:contextual-spacing="false" fo:keep-with-next="always" />
            <style:text-properties style:font-name="Liberation Sans"
                fo:font-family="&apos;Liberation Sans&apos;" style:font-family-generic="swiss"
                style:font-pitch="variable" fo:font-size="14pt"
                style:font-name-asian="Microsoft YaHei"
                style:font-family-asian="&apos;Microsoft YaHei&apos;"
                style:font-family-generic-asian="system" style:font-pitch-asian="variable"
                style:font-size-asian="14pt" style:font-name-complex="Lucida Sans1"
                style:font-family-complex="&apos;Lucida Sans&apos;"
                style:font-family-generic-complex="system" style:font-pitch-complex="variable"
                style:font-size-complex="14pt" />
        </style:style>
        <style:style style:name="Text_20_body" style:display-name="Text body"
            style:family="paragraph" style:parent-style-name="Standard" style:class="text">
            <style:paragraph-properties fo:margin-top="0cm" fo:margin-bottom="0.247cm"
                style:contextual-spacing="false" fo:line-height="115%" />
        </style:style>
        <style:style style:name="List" style:family="paragraph"
            style:parent-style-name="Text_20_body" style:class="list">
            <style:text-properties style:font-size-asian="12pt"
                style:font-name-complex="Lucida Sans"
                style:font-family-complex="&apos;Lucida Sans&apos;"
                style:font-family-generic-complex="swiss" />
        </style:style>
        <style:style style:name="Caption" style:family="paragraph"
            style:parent-style-name="Standard" style:class="extra">
            <style:paragraph-properties fo:margin-top="0.212cm" fo:margin-bottom="0.212cm"
                style:contextual-spacing="false" text:number-lines="false" text:line-number="0" />
            <style:text-properties fo:font-size="12pt" fo:font-style="italic"
                style:font-size-asian="12pt" style:font-style-asian="italic"
                style:font-name-complex="Lucida Sans"
                style:font-family-complex="&apos;Lucida Sans&apos;"
                style:font-family-generic-complex="swiss" style:font-size-complex="12pt"
                style:font-style-complex="italic" />
        </style:style>
        <style:style style:name="Index" style:family="paragraph" style:parent-style-name="Standard"
            style:class="index">
            <style:paragraph-properties text:number-lines="false" text:line-number="0" />
            <style:text-properties style:font-size-asian="12pt"
                style:font-name-complex="Lucida Sans"
                style:font-family-complex="&apos;Lucida Sans&apos;"
                style:font-family-generic-complex="swiss" />
        </style:style>
        <text:outline-style style:name="Outline">
            <text:outline-level-style text:level="1" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="2" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="3" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="4" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="5" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="6" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="7" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="8" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="9" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
            <text:outline-level-style text:level="10" style:num-format="">
                <style:list-level-properties
                    text:list-level-position-and-space-mode="label-alignment">
                    <style:list-level-label-alignment text:label-followed-by="listtab" />
                </style:list-level-properties>
            </text:outline-level-style>
        </text:outline-style>
        <text:notes-configuration text:note-class="footnote" style:num-format="1"
            text:start-value="0" text:footnotes-position="page" text:start-numbering-at="document" />
        <text:notes-configuration text:note-class="endnote" style:num-format="i"
            text:start-value="0" />
        <text:linenumbering-configuration text:number-lines="false" text:offset="0.499cm"
            style:num-format="1" text:number-position="left" text:increment="5" />
        <loext:theme loext:name="Office">
            <loext:theme-colors loext:name="LibreOffice">
                <loext:color loext:name="dark1" loext:color="#000000" />
                <loext:color loext:name="light1" loext:color="#ffffff" />
                <loext:color loext:name="dark2" loext:color="#000000" />
                <loext:color loext:name="light2" loext:color="#ffffff" />
                <loext:color loext:name="accent1" loext:color="#18a303" />
                <loext:color loext:name="accent2" loext:color="#0369a3" />
                <loext:color loext:name="accent3" loext:color="#a33e03" />
                <loext:color loext:name="accent4" loext:color="#8e03a3" />
                <loext:color loext:name="accent5" loext:color="#c99c00" />
                <loext:color loext:name="accent6" loext:color="#c9211e" />
                <loext:color loext:name="hyperlink" loext:color="#0000ee" />
                <loext:color loext:name="followed-hyperlink" loext:color="#551a8b" />
            </loext:theme-colors>
        </loext:theme>
    </office:styles>
    <office:automatic-styles>
        <style:page-layout style:name="Mpm1">
            <style:page-layout-properties fo:page-width="21.001cm" fo:page-height="29.7cm"
                style:num-format="1" style:print-orientation="portrait" fo:margin-top="2cm"
                fo:margin-bottom="2cm" fo:margin-left="2cm" fo:margin-right="2cm"
                style:writing-mode="lr-tb" style:footnote-max-height="0cm" loext:margin-gutter="0cm">
                <style:footnote-sep style:width="0.018cm" style:distance-before-sep="0.101cm"
                    style:distance-after-sep="0.101cm" style:line-style="solid"
                    style:adjustment="left" style:rel-width="25%" style:color="#000000" />
            </style:page-layout-properties>
            <style:header-style />
            <style:footer-style />
        </style:page-layout>
    </office:automatic-styles>
    <office:master-styles>
        <style:master-page style:name="Standard" style:page-layout-name="Mpm1" />
    </office:master-styles>
</office:document-styles>`
}
