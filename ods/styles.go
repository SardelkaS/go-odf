package ods

import "fmt"

type styles struct{}

func newStyles() *styles {
	return &styles{}
}

// generate generates xml code
func (s *styles) generate() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
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
    xmlns:presentation="urn:oasis:names:tc:opendocument:xmlns:presentation:1.0" office:version="1.4">
    <office:styles>
        <style:style style:name="Standard" style:family="Paragraph" style:class="text">
            <style:text-properties fo:font-size="12pt" style:font-name="Liberation Sans"/>
        </style:style>
    </office:styles>
	<office:automatic-styles>
		<number:number-style style:name="N2">
            <number:number number:decimal-places="2" number:min-decimal-places="2"
                number:min-integer-digits="1" />
        </number:number-style>
        <style:page-layout style:name="Mpm1">
            <style:page-layout-properties style:writing-mode="lr-tb" />
            <style:header-style>
                <style:header-footer-properties fo:min-height="0.75cm" fo:margin-left="0cm"
                    fo:margin-right="0cm" fo:margin-bottom="0.25cm" />
            </style:header-style>
            <style:footer-style>
                <style:header-footer-properties fo:min-height="0.75cm" fo:margin-left="0cm"
                    fo:margin-right="0cm" fo:margin-top="0.25cm" />
            </style:footer-style>
        </style:page-layout>
        <style:page-layout style:name="Mpm2">
            <style:page-layout-properties style:writing-mode="lr-tb" />
            <style:header-style>
                <style:header-footer-properties fo:min-height="0.75cm" fo:margin-left="0cm"
                    fo:margin-right="0cm" fo:margin-bottom="0.25cm" fo:border="1.5pt solid #000000"
                    fo:padding="0.018cm" fo:background-color="#c0c0c0">
                    <style:background-image />
                </style:header-footer-properties>
            </style:header-style>
            <style:footer-style>
                <style:header-footer-properties fo:min-height="0.75cm" fo:margin-left="0cm"
                    fo:margin-right="0cm" fo:margin-top="0.25cm" fo:border="1.5pt solid #000000"
                    fo:padding="0.018cm" fo:background-color="#c0c0c0">
                    <style:background-image />
                </style:header-footer-properties>
            </style:footer-style>
        </style:page-layout>
	</office:automatic-styles>
	<office:master-styles>
        <style:master-page style:name="Default" style:page-layout-name="Mpm1">
            <style:header>
                <text:p>
                    <text:sheet-name>???</text:sheet-name>
                </text:p>
            </style:header>
            <style:header-left style:display="false" />
            <style:header-first style:display="false" />
            <style:footer>
                <text:p>Страница <text:page-number>1</text:page-number></text:p>
            </style:footer>
            <style:footer-left style:display="false" />
            <style:footer-first style:display="false" />
        </style:master-page>
        <style:master-page style:name="Report" style:page-layout-name="Mpm2">
            <style:header>
                <style:region-left>
                    <text:p><text:sheet-name>???</text:sheet-name><text:s />(<text:title>???</text:title>
                        )</text:p>
                </style:region-left>
                <style:region-right>
                    <text:p><text:date style:data-style-name="N2" text:date-value="2025-05-02">
                        00.00.0000</text:date>, <text:time>00:00:00</text:time></text:p>
                </style:region-right>
            </style:header>
            <style:header-left style:display="false" />
            <style:header-first style:display="false" />
            <style:footer>
                <text:p>Страница <text:page-number>1</text:page-number><text:s />/ <text:page-count>
                    99</text:page-count></text:p>
            </style:footer>
            <style:footer-left style:display="false" />
            <style:footer-first style:display="false" />
        </style:master-page>
    </office:master-styles>
</office:document-styles>`)
}
