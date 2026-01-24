package https

import (
	"net/http"

	"github.com/dwaynedwards/sell-u-lar/pkg/https"
	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
)

func (s *Server) registerProductRoutes() {
	s.router.Handle("GET /api/v1/products", https.MakeHTTPHandlerFunc(s.handleProducts()))
	s.router.Handle("GET /api/v1/products/{brand}", https.MakeHTTPHandlerFunc(s.handleBrandProducts()))
	s.router.Handle("GET /api/v1/products/{brand}/{sku}", https.MakeHTTPHandlerFunc(s.handleProduct()))
}

func (s *Server) handleProducts() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		products, err := s.products.ListProducts(r.Context())
		if err != nil {
			return err
		}

		return https.WriteJSON(w, http.StatusOK, responses.ProductsResponse{
			Products: products,
		})
	}
}

func (s *Server) handleBrandProducts() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		products, err := s.products.ListProductsByBrand(r.Context(), r.PathValue("brand"))
		if err != nil {
			return err
		}

		return https.WriteJSON(w, http.StatusOK, responses.ProductsResponse{
			Products: products,
		})
	}
}

func (s *Server) handleProduct() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		product, err := s.products.GetProductByBrandAndSku(r.Context(), r.PathValue("brand"), r.PathValue("sku"))
		if err != nil {
			return err
		}
		return https.WriteJSON(w, http.StatusOK, responses.ProductResponse{
			Product: product,
		})
	}
}
