package https

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
	"github.com/dwaynedwards/sell-u-lar/pkg/https"
	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
)

func handleProducts() https.HandlerFunc {
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

		return https.WriteJSON(w, http.StatusOK, res)
	}
}

func handleProduct() https.HandlerFunc {
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

		sku := url.PathEscape(r.PathValue("sku"))
		var product types.Product
		for _, p := range res.Products {
			if strings.ToLower(p.Sku) == sku {
				product = p
				break
			}
		}

		if product.ID == 0 {
			return errors.NotFoundError(fmt.Sprintf("Product not found with sku [%s]", sku))
		}

		return https.WriteJSON(w, http.StatusOK, responses.ProductResponse{
			Product: product,
		})
	}
}
