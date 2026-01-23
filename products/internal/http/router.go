package https

import (
	"net/http"

	"github.com/dwaynedwards/sell-u-lar/pkg/https"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /api/v1/products", https.MakeHTTPHandlerFunc(handleProducts()))
	mux.Handle("GET /api/v1/products/{sku}", https.MakeHTTPHandlerFunc(handleProduct()))

	return mux
}
