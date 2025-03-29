package settings

type Settings struct{}

func New() Settings {
	return Settings{}
}

// Generate generates xml code
func (s Settings) Generate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<office:document-settings xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0"
    xmlns:ooo="http://openoffice.org/2004/office" xmlns:xlink="http://www.w3.org/1999/xlink"
    xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0" office:version="1.4">
    <office:settings>
        <config:config-item-set config:name="ooo:view-settings">
            <config:config-item config:name="ViewAreaTop" config:type="long">0</config:config-item>
            <config:config-item config:name="ViewAreaLeft" config:type="long">0</config:config-item>
            <config:config-item config:name="ViewAreaWidth" config:type="long">66439</config:config-item>
            <config:config-item config:name="ViewAreaHeight" config:type="long">32466</config:config-item>
            <config:config-item config:name="ShowRedlineChanges" config:type="boolean">true</config:config-item>
            <config:config-item config:name="InBrowseMode" config:type="boolean">false</config:config-item>
            <config:config-item-map-indexed config:name="Views">
                <config:config-item-map-entry>
                    <config:config-item config:name="ViewId" config:type="string">view2</config:config-item>
                    <config:config-item config:name="ViewLeft" config:type="long">25307</config:config-item>
                    <config:config-item config:name="ViewTop" config:type="long">2501</config:config-item>
                    <config:config-item config:name="VisibleLeft" config:type="long">0</config:config-item>
                    <config:config-item config:name="VisibleTop" config:type="long">0</config:config-item>
                    <config:config-item config:name="VisibleRight" config:type="long">66437</config:config-item>
                    <config:config-item config:name="VisibleBottom" config:type="long">32464</config:config-item>
                    <config:config-item config:name="ZoomType" config:type="short">0</config:config-item>
                    <config:config-item config:name="ViewLayoutColumns" config:type="short">1</config:config-item>
                    <config:config-item config:name="ViewLayoutBookMode" config:type="boolean">false</config:config-item>
                    <config:config-item config:name="ZoomFactor" config:type="short">100</config:config-item>
                    <config:config-item config:name="IsSelectedFrame" config:type="boolean">false</config:config-item>
                    <config:config-item config:name="KeepRatio" config:type="boolean">false</config:config-item>
                    <config:config-item config:name="AnchoredTextOverflowLegacy"
                        config:type="boolean">false</config:config-item>
                    <config:config-item config:name="LegacySingleLineFontwork" config:type="boolean">
                        false</config:config-item>
                    <config:config-item config:name="ConnectorUseSnapRect" config:type="boolean">
                        false</config:config-item>
                    <config:config-item config:name="IgnoreBreakAfterMultilineField"
                        config:type="boolean">false</config:config-item>
                </config:config-item-map-entry>
            </config:config-item-map-indexed>
        </config:config-item-set>
        <config:config-item-set config:name="ooo:configuration-settings">
            <config:config-item config:name="ProtectForm" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrinterName" config:type="string" />
            <config:config-item config:name="EmbeddedDatabaseName" config:type="string" />
            <config:config-item config:name="CurrentDatabaseDataSource" config:type="string" />
            <config:config-item config:name="LinkUpdateMode" config:type="short">1</config:config-item>
            <config:config-item config:name="AddParaTableSpacingAtStart" config:type="boolean">true</config:config-item>
            <config:config-item config:name="UnbreakableNumberings" config:type="boolean">false</config:config-item>
            <config:config-item config:name="FieldAutoUpdate" config:type="boolean">true</config:config-item>
            <config:config-item config:name="AddVerticalFrameOffsets" config:type="boolean">false</config:config-item>
            <config:config-item config:name="AddParaTableSpacing" config:type="boolean">true</config:config-item>
            <config:config-item config:name="ChartAutoUpdate" config:type="boolean">true</config:config-item>
            <config:config-item config:name="CurrentDatabaseCommand" config:type="string" />
            <config:config-item config:name="PrinterSetup" config:type="base64Binary" />
            <config:config-item config:name="AlignTabStopPosition" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrinterPaperFromSetup" config:type="boolean">false</config:config-item>
            <config:config-item config:name="IsKernAsianPunctuation" config:type="boolean">false</config:config-item>
            <config:config-item config:name="CharacterCompressionType" config:type="short">0</config:config-item>
            <config:config-item config:name="ApplyUserData" config:type="boolean">true</config:config-item>
            <config:config-item config:name="DoNotJustifyLinesWithManualBreak" config:type="boolean">
                false</config:config-item>
            <config:config-item config:name="SaveThumbnail" config:type="boolean">true</config:config-item>
            <config:config-item config:name="SaveGlobalDocumentLinks" config:type="boolean">false</config:config-item>
            <config:config-item config:name="SmallCapsPercentage66" config:type="boolean">false</config:config-item>
            <config:config-item config:name="CurrentDatabaseCommandType" config:type="int">0</config:config-item>
            <config:config-item config:name="SaveVersionOnClose" config:type="boolean">false</config:config-item>
            <config:config-item config:name="UpdateFromTemplate" config:type="boolean">true</config:config-item>
            <config:config-item config:name="DoNotCaptureDrawObjsOnPage" config:type="boolean">false</config:config-item>
            <config:config-item config:name="UseFormerObjectPositioning" config:type="boolean">false</config:config-item>
            <config:config-item config:name="EmbedSystemFonts" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrinterIndependentLayout" config:type="string">
                high-resolution</config:config-item>
            <config:config-item config:name="IsLabelDocument" config:type="boolean">false</config:config-item>
            <config:config-item config:name="AddFrameOffsets" config:type="boolean">false</config:config-item>
            <config:config-item config:name="AddExternalLeading" config:type="boolean">true</config:config-item>
            <config:config-item config:name="MsWordCompMinLineHeightByFly" config:type="boolean">
                false</config:config-item>
            <config:config-item config:name="UseOldNumbering" config:type="boolean">false</config:config-item>
            <config:config-item config:name="OutlineLevelYieldsNumbering" config:type="boolean">
                false</config:config-item>
            <config:config-item config:name="DoNotResetParaAttrsForNumFont" config:type="boolean">
                false</config:config-item>
            <config:config-item config:name="IgnoreFirstLineIndentInNumbering" config:type="boolean">
                false</config:config-item>
            <config:config-item config:name="AllowPrintJobCancel" config:type="boolean">true</config:config-item>
            <config:config-item config:name="UseFormerLineSpacing" config:type="boolean">false</config:config-item>
            <config:config-item config:name="AddParaSpacingToTableCells" config:type="boolean">true</config:config-item>
            <config:config-item config:name="AddParaLineSpacingToTableCells" config:type="boolean">
                true</config:config-item>
            <config:config-item config:name="UseFormerTextWrapping" config:type="boolean">false</config:config-item>
            <config:config-item config:name="RedlineProtectionKey" config:type="base64Binary" />
            <config:config-item config:name="ConsiderTextWrapOnObjPos" config:type="boolean">false</config:config-item>
            <config:config-item config:name="NoGapAfterNoteNumber" config:type="boolean">false</config:config-item>
            <config:config-item config:name="TableRowKeep" config:type="boolean">false</config:config-item>
            <config:config-item config:name="TabsRelativeToIndent" config:type="boolean">true</config:config-item>
            <config:config-item config:name="IgnoreTabsAndBlanksForLineCalculation"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="IgnoreHiddenCharsForLineCalculation"
                config:type="boolean">true</config:config-item>
            <config:config-item config:name="TabAtLeftIndentForParagraphsInList"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="Rsid" config:type="int">1414294</config:config-item>
            <config:config-item config:name="RsidRoot" config:type="int">1336284</config:config-item>
            <config:config-item config:name="LoadReadonly" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ClipAsCharacterAnchoredWriterFlyFrames"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="UnxForceZeroExtLeading" config:type="boolean">false</config:config-item>
            <config:config-item config:name="UseOldPrinterMetrics" config:type="boolean">false</config:config-item>
            <config:config-item config:name="MsWordCompTrailingBlanks" config:type="boolean">false</config:config-item>
            <config:config-item config:name="MathBaselineAlignment" config:type="boolean">true</config:config-item>
            <config:config-item config:name="InvertBorderSpacing" config:type="boolean">false</config:config-item>
            <config:config-item config:name="CollapseEmptyCellPara" config:type="boolean">true</config:config-item>
            <config:config-item config:name="TabOverflow" config:type="boolean">true</config:config-item>
            <config:config-item config:name="StylesNoDefault" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ClippedPictures" config:type="boolean">false</config:config-item>
            <config:config-item config:name="BackgroundParaOverDrawings" config:type="boolean">false</config:config-item>
            <config:config-item config:name="EmbedFonts" config:type="boolean">false</config:config-item>
            <config:config-item config:name="EmbedOnlyUsedFonts" config:type="boolean">false</config:config-item>
            <config:config-item config:name="EmbedLatinScriptFonts" config:type="boolean">true</config:config-item>
            <config:config-item config:name="EmbedAsianScriptFonts" config:type="boolean">true</config:config-item>
            <config:config-item config:name="EmptyDbFieldHidesPara" config:type="boolean">true</config:config-item>
            <config:config-item config:name="EmbedComplexScriptFonts" config:type="boolean">true</config:config-item>
            <config:config-item config:name="TabOverMargin" config:type="boolean">false</config:config-item>
            <config:config-item config:name="TabOverSpacing" config:type="boolean">false</config:config-item>
            <config:config-item config:name="TreatSingleColumnBreakAsPageBreak"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="SurroundTextWrapSmall" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ApplyParagraphMarkFormatToNumbering"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="PropLineSpacingShrinksFirstLine" config:type="boolean">
                true</config:config-item>
            <config:config-item config:name="SubtractFlysAnchoredAtFlys" config:type="boolean">false</config:config-item>
            <config:config-item config:name="DisableOffPagePositioning" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ContinuousEndnotes" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ProtectBookmarks" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ProtectFields" config:type="boolean">false</config:config-item>
            <config:config-item config:name="HyphenateURLs" config:type="boolean">false</config:config-item>
            <config:config-item config:name="HeaderSpacingBelowLastPara" config:type="boolean">false</config:config-item>
            <config:config-item config:name="FrameAutowidthWithMorePara" config:type="boolean">false</config:config-item>
            <config:config-item config:name="GutterAtTop" config:type="boolean">false</config:config-item>
            <config:config-item config:name="FootnoteInColumnToPageEnd" config:type="boolean">true</config:config-item>
            <config:config-item config:name="ImagePreferredDPI" config:type="int">0</config:config-item>
            <config:config-item config:name="AutoFirstLineIndentDisregardLineSpace"
                config:type="boolean">true</config:config-item>
            <config:config-item config:name="JustifyLinesWithShrinking" config:type="boolean">false</config:config-item>
            <config:config-item config:name="NoNumberingShowFollowBy" config:type="boolean">false</config:config-item>
            <config:config-item config:name="DropCapPunctuation" config:type="boolean">true</config:config-item>
            <config:config-item config:name="UseVariableWidthNBSP" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintBlackFonts" config:type="boolean">false</config:config-item>
            <config:config-item config:name="ApplyTextAttrToEmptyLineAtEndOfParagraph"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="ApplyParagraphMarkFormatToEmptyLineAtEndOfParagraph"
                config:type="boolean">false</config:config-item>
            <config:config-item config:name="PaintHellOverHeaderFooter" config:type="boolean">false</config:config-item>
            <config:config-item config:name="MinRowHeightInclBorder" config:type="boolean">false</config:config-item>
            <config:config-item config:name="MsWordCompGridMetrics" config:type="boolean">false</config:config-item>
            <config:config-item config:name="NoClippingWithWrapPolygon" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintAnnotationMode" config:type="short">0</config:config-item>
            <config:config-item config:name="PrintGraphics" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintLeftPages" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintControls" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintPageBackground" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintTextPlaceholder" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintDrawings" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintHiddenText" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintProspect" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintTables" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintProspectRTL" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintReversed" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintRightPages" config:type="boolean">true</config:config-item>
            <config:config-item config:name="PrintFaxName" config:type="string" />
            <config:config-item config:name="PrintPaperFromSetup" config:type="boolean">false</config:config-item>
            <config:config-item config:name="PrintEmptyPages" config:type="boolean">true</config:config-item>
        </config:config-item-set>
    </office:settings>
</office:document-settings>`
}
