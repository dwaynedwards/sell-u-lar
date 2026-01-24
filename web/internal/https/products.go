package https

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
	"github.com/dwaynedwards/sell-u-lar/pkg/https"
	"github.com/dwaynedwards/sell-u-lar/pkg/responses"
	"github.com/dwaynedwards/sell-u-lar/web"
	"github.com/dwaynedwards/sell-u-lar/web/internal/templates/pages"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (s *Server) registerProductRoutes() {
	s.router.Handle("GET /devices", https.MakeHTTPHandlerFunc(s.handleProducts()))
	s.router.Handle("GET /devices/{brand}", https.MakeHTTPHandlerFunc(s.handleDeviceBrands()))
	s.router.Handle("GET /devices/{brand}/{sku}", https.MakeHTTPHandlerFunc(s.handleDevice()))
	s.router.Handle("GET /", https.MakeHTTPHandlerFunc(s.handleDevicesRedirect()))
}

func (s *Server) handleProducts() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		resp, err := http.Get(web.Config.ProductsServiceBaseUrl + "/api/v1/products")
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

		return pages.Devices("Devices", res.Products).Render(r.Context(), w)
	}
}

func (s *Server) handleDeviceBrands() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		brand := url.PathEscape(r.PathValue("brand"))

		resp, err := http.Get(fmt.Sprintf("%s/api/v1/products/%s", web.Config.ProductsServiceBaseUrl, brand))
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

		caser := cases.Title(language.English)
		return pages.DeviceBrands("Devices", caser.String(brand), res.Products).Render(r.Context(), w)
	}
}

func (s *Server) handleDevice() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		brand := url.PathEscape(r.PathValue("brand"))
		sku := url.PathEscape(r.PathValue("sku"))
		resp, err := http.Get(fmt.Sprintf("%s/api/v1/products/%s/%s", web.Config.ProductsServiceBaseUrl, brand, sku))
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

		return pages.Device(res.Product).Render(r.Context(), w)
	}
}

func (s *Server) handleDevicesRedirect() https.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		http.Redirect(w, r, "/devices", http.StatusTemporaryRedirect)
		return nil
	}
}
