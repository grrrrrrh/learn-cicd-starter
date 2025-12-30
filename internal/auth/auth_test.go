package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey abc123")

	got, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "abc123" {
		t.Fatalf("expected %q, got %q", "THIS_WILL_FAIL", got)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	h := http.Header{}

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetAPIKey_WrongScheme(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer abc123")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
