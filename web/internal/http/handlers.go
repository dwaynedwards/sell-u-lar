package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"github.com/dwaynedwards/sell-u-lar/web/internal/errors"
	"github.com/dwaynedwards/sell-u-lar/web/internal/templates/pages"
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

		return pages.Home("Products", res.Products).Render(r.Context(), w)
	}
}

func handleProduct() HandlerFunc {
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

		id := r.PathValue("id")
		var product types.Product
		for _, p := range res.Products {
			if strings.ToLower(p.Sku) == id {
				product = p
				break
			}
		}

		if product.ID == 0 {
			return errors.NotFoundError(fmt.Sprintf("Product not found with ID [%s]", id))
		}

		return pages.Product(product).Render(r.Context(), w)
	}
}

func handleNotFound() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return errors.NotFoundError("Product not found")
	}
}
