package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/csush/Go-URL-Shortener/models"
	"github.com/csush/Go-URL-Shortener/storage"
)

type ShortenRequest struct {
	URL string
}

type Handler struct {
	store storage.IStore
	cg    models.ICodeGenerator
}

func NewHandler(store storage.IStore, cg models.ICodeGenerator) *Handler {
	return &Handler{store: store, cg: cg}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortenRequestData := ShortenRequest{}
	err := json.NewDecoder(r.Body).Decode(&shortenRequestData)
	if err != nil {
		panic(err)
	}

	generatedURL, err := h.cg.GenerateUniqueID()
	if err != nil {
		panic(err)
	}

	h.store.Save(generatedURL, shortenRequestData.URL)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:5001/" + generatedURL))
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:]

	longURL, err := h.store.Get(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
