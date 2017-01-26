package ace

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAce(t *testing.T) {
	rootDir := "./"

	ace := Ace{
		Root:    rootDir,
		FileSys: http.Dir(rootDir),
		// Configs: []*Config{,}
	}

	req, err := http.NewRequest("GET", "/photos/test.html", nil)
	if err != nil {
		t.Fatalf("Test: Could not create HTTP request: %v", err)
	}

	rec := httptest.NewRecorder()

	ace.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Wrong status, expected: %d and got %d", http.StatusOK, rec.Code)
	}
	respBody := rec.Body.String()
	expectedBody := ``

	if respBody != expectedBody {
		t.Fatalf("Expected body: %v got: %v", expectedBody, respBody)
	}
}
