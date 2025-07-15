package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SpyCodeGenerator struct{}

func (spyCg *SpyCodeGenerator) GenerateUniqueID() (string, error) {
	return "abcd", nil
}

type SpyStore struct{}

func (s *SpyStore) Save(shortURL, longURL string) {}

func (s *SpyStore) Get(shortURL string) (string, error) {
	return "https://google.com", nil
}

func TestShorten(t *testing.T) {
	validData := []byte(`{"url": "https://google.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(validData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	store := &SpyStore{}
	cg := &SpyCodeGenerator{}
	handler := NewHandler(store, cg)

	handler.ShortenURL(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got status code %q want %q", w.Code, http.StatusCreated)
	}
	response := w.Body.String()
	if response != "http://localhost:5001/abcd" {
		t.Errorf("got %q want %q", response, "http://localhost:5001/abcd")
	}
}

func TestRedirect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/abcd", strings.NewReader(""))
	w := httptest.NewRecorder()

	store := &SpyStore{}
	cg := &SpyCodeGenerator{}
	handler := NewHandler(store, cg)

	handler.RedirectURL(w, req)
	want := "https://google.com"
	got := w.Result().Header.Get("Location")
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
