package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/johsquaree/url-shortener/int/shortener"
	"github.com/johsquaree/url-shortener/pkg/storage"
)

var store storage.Storage

func SetStorage(s storage.Storage) {
	store = s
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL := shortener.GenerateShortURL(req.URL)
	store.Save(shortURL, req.URL)

	res := ShortenResponse{ShortURL: shortURL}
	json.NewEncoder(w).Encode(res)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/redirect/"):]
	originalURL, found := store.Get(shortURL)
	if !found {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}
