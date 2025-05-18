package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 123456")
	fetchedKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
		return
	}
	if fetchedKey != "123456" {
		t.Fatalf("expected 123456, got %s", fetchedKey)
		return
	}
	headers = http.Header{}
	fetchedKey, err = GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
		return
	}
	if fetchedKey != "" {
		t.Fatalf("expected empty string, got %s", fetchedKey)
		return
	}
	headers.Set("Authorization", "ApiKey")
	fetchedKey, err = GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
		return
	}
	if fetchedKey != "" {
		t.Fatalf("expected empty string, got %s", fetchedKey)
		return
	}
}
