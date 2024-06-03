package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"original_url"`
	ShortURL     string `json:"short_url"`
	CreationDate string `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {

	hasher := md5.New()

	hasher.Write([]byte(OriginalURL))
	// fmt.Println("hasher : ", hasher)

	data := hasher.Sum(nil)
	// fmt.Println("data : ", data)

	hash := hex.EncodeToString(data)
	fmt.Println("hash : ", hash)

	return hash[:8]
}

func createURL(OriginalURL string) string {

	shortURL := generateShortURL(OriginalURL)

	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now().String(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {

	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

// home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "URL Shortner")
}

// short url
func shortURLHandler(w http.ResponseWriter, r *http.Request) {

	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// redirect
func redirectHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/redirect/"):]

	url, err := getURL(id)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

func main() {

	fmt.Println("URL Shortner")

	// routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectHandler)

	// start server
	fmt.Println("Server is running on port 8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("Error on starting server : ", err)
	}

}
