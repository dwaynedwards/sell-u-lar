package https

import (
	"encoding/json"
	"net/http"

	"github.com/dwaynedwards/sell-u-lar/pkg/https"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Include the static content.
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.Handle("GET /", https.MakeHTTPHandlerFunc(handleProducts()))
	mux.Handle("GET /products", https.MakeHTTPHandlerFunc(handleNotFound()))
	mux.Handle("GET /products/{sku}", https.MakeHTTPHandlerFunc(handleProduct()))

	return mux
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
