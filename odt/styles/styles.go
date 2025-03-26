package styles

type Styles struct{}

func New() Styles {
	return Styles{}
}

// Generate generates xml code
func (s Styles) Generate() string {
	return `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<document-styles xmlns:anim="urn:oasis:names:tc:opendocument:xmlns:animation:1.0" xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0" xmlns:db="urn:oasis:names:tc:opendocument:xmlns:database:1.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0" xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0" xmlns:grddl="http://www.w3.org/2003/g/data-view#" xmlns:math="http://www.w3.org/1998/Math/MathML" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:presentation="urn:oasis:names:tc:opendocument:xmlns:presentation:1.0" xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0" xmlns:smil="urn:oasis:names:tc:opendocument:xmlns:smil-compatible:1.0" xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:xhtml="http://www.w3.org/1999/xhtml" xmlns:xlink="http://www.w3.org/1999/xlink" office:version="1.3">
	<font-face-decls>
		<font-face style:name="Aptos" svg:font-family="Aptos" style:font-family-generic="swiss" style:font-pitch="variable" />
		<font-face style:name="Times New Roman" svg:font-family="Times New Roman" style:font-family-generic="roman" style:font-pitch="variable" svg:panose-1="2 2 6 3 5 4 5 2 3 4" />
		<font-face style:name="Aptos Display" svg:font-family="Aptos Display" style:font-family-generic="swiss" style:font-pitch="variable" />
		<font-face style:name="Calibri" svg:font-family="Calibri" style:font-family-generic="swiss" style:font-pitch="variable" svg:panose-1="2 15 5 2 2 2 4 3 2 4" />
	</font-face-decls>
	<styles>
		<default-style style:family="table">
			<table-properties fo:margin-left="0in" table:border-model="collapsing" style:writing-mode="lr-tb" table:align="left" />
		</default-style>
		<default-style style:family="table-column">
			<table-column-properties style:use-optimal-column-width="true" />
		</default-style>
		<default-style style:family="table-row">
			<table-row-properties style:min-row-height="0in" style:use-optimal-row-height="true" fo:keep-together="auto" />
		</default-style>
		<default-style style:family="table-cell">
			<table-cell-properties fo:background-color="transparent" style:glyph-orientation-vertical="auto" style:vertical-align="top" fo:wrap-option="wrap" />
		</default-style>
		<default-style style:family="paragraph">
			<paragraph-properties fo:keep-with-next="auto" fo:keep-together="auto" fo:widows="2" fo:orphans="2" fo:break-before="auto" text:number-lines="true" fo:border="none" fo:padding="0in" style:shadow="none" style:line-break="strict" style:punctuation-wrap="hanging" style:text-autospace="ideograph-alpha" style:snap-to-layout-grid="true" style:contextual-spacing="false" fo:text-align="start" style:writing-mode="lr-tb" style:vertical-align="auto" fo:margin-bottom="0.1111in" fo:line-height="107%" fo:background-color="transparent" style:tab-stop-distance="0.4916in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Aptos" style:font-name-complex="Times New Roman" fo:font-weight="normal" style:font-weight-asian="normal" style:font-weight-complex="normal" fo:font-style="normal" style:font-style-asian="normal" style:font-style-complex="normal" fo:text-transform="none" fo:font-variant="normal" style:text-line-through-type="none" style:text-outline="false" style:font-relief="none" style:use-window-font-color="true" fo:letter-spacing="normal" style:text-scale="100%" style:letter-kerning="true" style:text-position="0% 100%" fo:font-size="11pt" style:font-size-asian="11pt" style:font-size-complex="11pt" fo:background-color="transparent" style:text-underline-type="none" style:text-underline-color="font-color" style:text-emphasize="none" fo:language="ru" fo:country="RU" style:language-asian="en" style:country-asian="US" style:language-complex="ar" style:country-complex="SA" style:text-combine="none" fo:hyphenate="true" />
		</default-style>
		<style style:name="Заголовок1" style:display-name="Заголовок 1" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="1">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-top="0.1666in" />
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#0F4761" fo:font-size="18pt" style:font-size-asian="18pt" style:font-size-complex="16pt" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок2" style:display-name="Заголовок 2" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="2">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:text-align="center" fo:margin-top="0.0277in" />
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-size="16pt" style:font-size-asian="16pt" style:font-size-complex="13pt" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок3" style:display-name="Заголовок 3" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="3">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:text-align="center" fo:margin-top="0.0277in" />
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" style:font-size-complex="12pt" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок4" style:display-name="Заголовок 4" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="4">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-top="0.0555in" fo:margin-bottom="0.0277in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#0F4761" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок5" style:display-name="Заголовок 5" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="5">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-top="0.0555in" fo:margin-bottom="0.0277in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#0F4761" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок6" style:display-name="Заголовок 6" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="6">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-top="0.0277in" fo:margin-bottom="0in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#595959" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок7" style:display-name="Заголовок 7" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="7">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-top="0.0277in" fo:margin-bottom="0in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#595959" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок8" style:display-name="Заголовок 8" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="8">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-bottom="0in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#272727" fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок9" style:display-name="Заголовок 9" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный" style:default-outline-level="9">
			<paragraph-properties fo:keep-with-next="always" fo:keep-together="always" fo:margin-bottom="0in" />
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#272727" fo:hyphenate="false" />
		</style>
		<style style:name="Обычный" style:display-name="Обычный" style:family="paragraph">
			<text-properties style:font-name="Times New Roman" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" fo:hyphenate="false" />
		</style>
		<style style:name="Основнойшрифтабзаца" style:display-name="Основной шрифт абзаца" style:family="text" />
		<style style:name="Абзацсписка" style:display-name="Абзац списка" style:family="paragraph" style:parent-style-name="Обычный">
			<paragraph-properties style:contextual-spacing="true" fo:margin-left="0.5in">
				<tab-stops />
			</paragraph-properties>
			<text-properties fo:hyphenate="false" />
		</style>
		<style style:name="Заголовок3Знак" style:display-name="Заголовок 3 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-size="14pt" style:font-size-asian="14pt" style:font-size-complex="12pt" />
		</style>
		<style style:name="Заголовок2Знак" style:display-name="Заголовок 2 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-size="16pt" style:font-size-asian="16pt" style:font-size-complex="13pt" />
		</style>
		<style style:name="Заголовок1Знак" style:display-name="Заголовок 1 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#0F4761" fo:font-size="18pt" style:font-size-asian="18pt" style:font-size-complex="16pt" />
		</style>
		<style style:name="Заголовок4Знак" style:display-name="Заголовок 4 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#0F4761" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок5Знак" style:display-name="Заголовок 5 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#0F4761" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок6Знак" style:display-name="Заголовок 6 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#595959" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок7Знак" style:display-name="Заголовок 7 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#595959" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок8Знак" style:display-name="Заголовок 8 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#272727" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок9Знак" style:display-name="Заголовок 9 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#272727" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Заголовок" style:display-name="Заголовок" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный">
			<paragraph-properties style:contextual-spacing="true" fo:margin-bottom="0.0555in" fo:line-height="100%" />
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:letter-spacing="-0.0069in" style:letter-kerning="true" fo:font-size="28pt" style:font-size-asian="28pt" style:font-size-complex="28pt" fo:hyphenate="false" />
		</style>
		<style style:name="ЗаголовокЗнак" style:display-name="Заголовок Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Aptos Display" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:letter-spacing="-0.0069in" style:letter-kerning="true" fo:font-size="28pt" style:font-size-asian="28pt" style:font-size-complex="28pt" />
		</style>
		<style style:name="Подзаголовок" style:display-name="Подзаголовок" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный">
			<text-properties style:font-name="Aptos" style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#595959" fo:letter-spacing="0.0104in" style:font-size-complex="14pt" fo:hyphenate="false" />
		</style>
		<style style:name="ПодзаголовокЗнак" style:display-name="Подзаголовок Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name-asian="Times New Roman" style:font-name-complex="Times New Roman" fo:color="#595959" fo:letter-spacing="0.0104in" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" style:font-size-complex="14pt" />
		</style>
		<style style:name="Цитата2" style:display-name="Цитата 2" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный">
			<paragraph-properties fo:text-align="center" fo:margin-top="0.1111in" />
			<text-properties fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#404040" fo:hyphenate="false" />
		</style>
		<style style:name="Цитата2Знак" style:display-name="Цитата 2 Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#404040" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Сильноевыделение" style:display-name="Сильное выделение" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#0F4761" />
		</style>
		<style style:name="Выделеннаяцитата" style:display-name="Выделенная цитата" style:family="paragraph" style:parent-style-name="Обычный" style:next-style-name="Обычный">
			<paragraph-properties fo:border-top="0.0069in solid #0F4761" fo:border-left="none" fo:border-bottom="0.0069in solid #0F4761" fo:border-right="none" fo:padding-top="0.1388in" fo:padding-left="0in" fo:padding-bottom="0.1388in" fo:padding-right="0in" style:shadow="none" fo:text-align="center" fo:margin-top="0.25in" fo:margin-bottom="0.25in" fo:margin-left="0.6in" fo:margin-right="0.6in">
				<tab-stops />
			</paragraph-properties>
			<text-properties fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#0F4761" fo:hyphenate="false" />
		</style>
		<style style:name="ВыделеннаяцитатаЗнак" style:display-name="Выделенная цитата Знак" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties style:font-name="Times New Roman" fo:font-style="italic" style:font-style-asian="italic" style:font-style-complex="italic" fo:color="#0F4761" style:letter-kerning="false" fo:font-size="14pt" style:font-size-asian="14pt" />
		</style>
		<style style:name="Сильнаяссылка" style:display-name="Сильная ссылка" style:family="text" style:parent-style-name="Основнойшрифтабзаца">
			<text-properties fo:font-weight="bold" style:font-weight-asian="bold" style:font-weight-complex="bold" fo:font-variant="small-caps" fo:color="#0F4761" fo:letter-spacing="0.0034in" />
		</style>
		<notes-configuration text:note-class="footnote" text:start-value="0" style:num-format="1" text:start-numbering-at="document" text:footnotes-position="page" />
		<notes-configuration text:note-class="endnote" text:start-value="0" style:num-format="i" text:start-numbering-at="document" text:footnotes-position="document" />
		<outline-style style:name="WW_OutlineListStyle">
			<outline-level-style text:level="1" style:num-format="" />
			<outline-level-style text:level="2" style:num-format="" />
			<outline-level-style text:level="3" style:num-format="" />
			<outline-level-style text:level="4" style:num-format="" />
			<outline-level-style text:level="5" style:num-format="" />
			<outline-level-style text:level="6" style:num-format="" />
			<outline-level-style text:level="7" style:num-format="" />
			<outline-level-style text:level="8" style:num-format="" />
			<outline-level-style text:level="9" style:num-format="" />
		</outline-style>
		<default-page-layout>
			<page-layout-properties style:layout-grid-standard-mode="true" />
		</default-page-layout>
		<default-style style:family="graphic">
			<graphic-properties draw:fill="solid" draw:fill-color="#156082" draw:opacity="100%" draw:stroke="solid" svg:stroke-width="0.01389in" svg:stroke-color="#042433" svg:stroke-opacity="100%" draw:stroke-linejoin="miter" svg:stroke-linecap="butt" />
		</default-style>
	</styles>
	<automatic-styles>
		<page-layout style:name="PL0">
			<page-layout-properties fo:page-width="8.268in" fo:page-height="11.693in" style:print-orientation="portrait" fo:margin-top="0.7875in" fo:margin-left="1.1812in" fo:margin-bottom="0.7875in" fo:margin-right="0.5902in" style:num-format="1" style:writing-mode="lr-tb">
				<footnote-sep style:width="0.007in" style:rel-width="33%" style:color="#000000" style:line-style="solid" style:adjustment="left" />
			</page-layout-properties>
		</page-layout>
	</automatic-styles>
	<master-styles>
		<master-page style:name="MP0" style:page-layout-name="PL0" />
	</master-styles>
</document-styles>`
}
