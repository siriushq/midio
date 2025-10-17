package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Test cross domain xml handler.
func TestCrossXMLHandler(t *testing.T) {
	// Server initialization.
	router := mux.NewRouter().SkipClean(true)
	handler := setCrossDomainPolicy(router)
	srv := httptest.NewServer(handler)

	resp, err := http.Get(srv.URL + crossDomainXMLEntity)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Unexpected http status received", resp.Status)
	}
}
