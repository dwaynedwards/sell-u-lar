package https

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func MakeHTTPHandlerFunc(h HandlerFunc) http.HandlerFunc {
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

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
