package mimedb

import "testing"

func TestMimeLookup(t *testing.T) {
	// Test mimeLookup.
	contentType := DB["txt"].ContentType
	if contentType != "text/plain" {
		t.Fatalf("Invalid content type are found expected \"application/x-msdownload\", got %s", contentType)
	}
	compressible := DB["txt"].Compressible
	if compressible {
		t.Fatalf("Invalid content type are found expected \"false\", got %t", compressible)
	}
}

func TestTypeByExtension(t *testing.T) {
	// Test TypeByExtension.
	contentType := TypeByExtension(".txt")
	if contentType != "text/plain" {
		t.Fatalf("Invalid content type are found expected \"text/plain\", got %s", contentType)
	}
	// Test non-existent type resolution
	contentType = TypeByExtension(".abc")
	if contentType != "application/octet-stream" {
		t.Fatalf("Invalid content type are found expected \"application/octet-stream\", got %s", contentType)
	}
}
