package https

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/dwaynedwards/sell-u-lar/pkg/https"
	"github.com/dwaynedwards/sell-u-lar/web/internal/templates/pages"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (s *Server) registerProductRoutes() {
	s.router.Handle("GET /devices", https.MakeHTTPHandlerFunc(s.handleDevices()))
	s.router.Handle("GET /devices/{brand}", https.MakeHTTPHandlerFunc(s.handleDeviceBrands()))
	s.router.Handle("GET /devices/{brand}/{sku}", https.MakeHTTPHandlerFunc(s.handleDevice()))
	s.router.Handle("/", https.MakeHTTPHandlerFunc(s.handleDevicesRedirect()))
}

func (s *Server) handleDevices() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		slog.InfoContext(r.Context(), "List Products")

		products, err := s.products.ListProducts(r.Context())
		if err != nil {
			return err
		}

		return pages.Devices("Devices", products).Render(r.Context(), w)
	}
}

func (s *Server) handleDeviceBrands() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		brand := url.PathEscape(r.PathValue("brand"))

		slog.InfoContext(r.Context(), "List Products By Brand:", "brand", brand)

		products, err := s.products.ListProductsByBrand(r.Context(), brand)
		if err != nil {
			return err
		}

		caser := cases.Title(language.English)
		return pages.DeviceBrands("Devices", caser.String(brand), products).Render(r.Context(), w)
	}
}

func (s *Server) handleDevice() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		brand := url.PathEscape(r.PathValue("brand"))
		sku := url.PathEscape(r.PathValue("sku"))

		slog.InfoContext(r.Context(), "Get Product By Brand and Sku:", "brand", brand, "sku", sku)

		product, err := s.products.GetProductByBrandAndSku(r.Context(), brand, sku)
		if err != nil {
			return err
		}

		return pages.Device(product).Render(r.Context(), w)
	}
}

func (s *Server) handleDevicesRedirect() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		slog.InfoContext(r.Context(), "List Products Redirect")

		http.Redirect(w, r, "/devices", http.StatusTemporaryRedirect)
		return nil
	}
}
