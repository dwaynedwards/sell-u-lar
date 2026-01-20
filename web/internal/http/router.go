package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dwaynedwards/sell-u-lar/web/internal/errors"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP error", "err", err.Error(), "path", r.URL.Path)
			var e errors.Error
			if errors.As(err, &e) {
				err = WriteJSON(w, e.StatusCode, e)
			} else {
				errRes := errors.InternalServerError("internal server error")
				err = WriteJSON(w, errRes.StatusCode, errRes)
			}
			if err == nil {
				return
			}
			slog.Error("HTTP error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Include the static content.
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.Handle("GET /", makeHTTPHandlerFunc(handleHome()))
	mux.Handle("GET /products", makeHTTPHandlerFunc(handleNotFound()))
	mux.Handle("GET /products/{id}", makeHTTPHandlerFunc(handleProduct()))

	return mux
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
