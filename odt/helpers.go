package odt

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
)

// supportedImageTypes содержит маппинг MIME-типов и расширений
var supportedImageTypes = map[string]string{
	"image/png":     ".png",
	"image/jpeg":    ".jpg",
	"image/gif":     ".gif",
	"image/svg+xml": ".svg",
	"image/bmp":     ".bmp",
	"image/webp":    ".webp",
	"image/tiff":    ".tiff",
}

// detectContentType returns MIME-type
func detectContentType(base64Data string) (string, error) {
	if len(base64Data) < 12 {
		return "", fmt.Errorf("invalid base64 data: too short")
	}

	// check data URI scheme
	if strings.HasPrefix(base64Data, "data:") {
		parts := strings.SplitN(base64Data, ";", 2)
		if len(parts) < 2 {
			return "", fmt.Errorf("invalid data URI format")
		}

		mimeType := strings.TrimPrefix(parts[0], "data:")
		if _, supported := supportedImageTypes[mimeType]; supported {
			return mimeType, nil
		}
		return "", fmt.Errorf("unsupported image MIME type: %s", mimeType)
	}

	data, err := base64.StdEncoding.DecodeString(cleanBase64Data(base64Data))
	if err != nil {
		return "", fmt.Errorf("invalid base64 data: %v", err)
	}

	if len(data) < 12 {
		return "", fmt.Errorf("image data too short")
	}

	// check by magic numbers
	switch {
	case bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47}):
		return "image/png", nil
	case bytes.HasPrefix(data, []byte{0xFF, 0xD8}):
		return "image/jpeg", nil
	case bytes.HasPrefix(data, []byte("GIF87a")), bytes.HasPrefix(data, []byte("GIF89a")):
		return "image/gif", nil
	case bytes.HasPrefix(data, []byte{0x42, 0x4D}):
		return "image/bmp", nil
	case bytes.HasPrefix(data, []byte("WEBP")):
		return "image/webp", nil
	case bytes.HasPrefix(data, []byte{0x49, 0x49, 0x2A, 0x00}),
		bytes.HasPrefix(data, []byte{0x4D, 0x4D, 0x00, 0x2A}):
		return "image/tiff", nil
	case bytes.Contains(data[:12], []byte("<svg")):
		return "image/svg+xml", nil
	default:
		return "", fmt.Errorf("unrecognized image format")
	}
}
func cleanBase64Data(data string) string {
	if idx := strings.Index(data, ","); idx != -1 {
		return data[idx+1:]
	}
	return data
}

// escapeXML escapes symbols for xml
func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}
