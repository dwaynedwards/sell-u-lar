package http

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
	"github.com/dwaynedwards/sell-u-lar/web/internal/templates/pages/home"
)

func handleHome() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		jsonBytes, err := os.ReadFile("./static/json/products.json")
		if err != nil {
			return err
		}
		var res responses.ProductsResponse
		err = json.Unmarshal(jsonBytes, &res.Products)
		if err != nil {
			return err
		}

		return home.Page("Products", res.Products).Render(r.Context(), w)
	}
}
