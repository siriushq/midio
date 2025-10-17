package mimedb

import (
	"strings"
)

// TypeByExtension resolves the extension to its respective content-type.
func TypeByExtension(ext string) string {
	// Set default to "application/octet-stream".
	var contentType = "application/octet-stream"
	if ext != "" {
		if content, ok := DB[strings.ToLower(strings.TrimPrefix(ext, "."))]; ok {
			contentType = content.ContentType
		}
	}
	return contentType
}
