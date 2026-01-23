package https

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
	"github.com/dwaynedwards/sell-u-lar/pkg/https"
	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
	"github.com/dwaynedwards/sell-u-lar/web/internal/templates/pages"
)

func handleProducts() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		resp, err := http.Get("http://localhost:3001/api/v1/products")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.NotFoundError(resp.Status)
		}

		jsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var res responses.ProductsResponse
		err = json.Unmarshal(jsonBytes, &res)
		if err != nil {
			return err
		}

		return pages.Home("Products", res.Products).Render(r.Context(), w)
	}
}

func handleProduct() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		sku := url.PathEscape(r.PathValue("sku"))
		slog.Info("Escaped", "Sku", sku)
		resp, err := http.Get("http://localhost:3001/api/v1/products/" + sku)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.NotFoundError(resp.Status)
		}

		jsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var res responses.ProductResponse
		err = json.Unmarshal(jsonBytes, &res)
		if err != nil {
			return err
		}

		return pages.Product(res.Product).Render(r.Context(), w)
	}
}

func handleNotFound() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return errors.NotFoundError("Page not found")
	}
}
