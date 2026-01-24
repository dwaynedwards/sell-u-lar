package https

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/dwaynedwards/sell-u-lar/pkg/db"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"github.com/dwaynedwards/sell-u-lar/products/internal/services"
	"github.com/dwaynedwards/sell-u-lar/products/internal/store"
)

const ShutdownTimeout = 5 * time.Second

type ProductsService interface {
	ListProducts(ctx context.Context) (types.Products, error)
	ListProductsByBrand(ctx context.Context, brand string) (types.Products, error)
	GetProductByBrandAndSku(ctx context.Context, brand, sku string) (types.Product, error)
}

type Server struct {
	listener net.Listener
	server   *http.Server
	router   *http.ServeMux
	db       db.Database
	products ProductsService
}

func NewServer(db db.Database) *Server {
	s := &Server{
		server:   &http.Server{},
		router:   http.NewServeMux(),
		db:       db,
		products: services.NewProductsService(store.NewProductsStore(db)),
	}

	s.server.Handler = http.HandlerFunc(s.router.ServeHTTP)

	s.registerProductRoutes()

	return s
}

func (s *Server) Start() (err error) {
	slog.Info("Opening connection to DB")
	if err = s.db.Open(); err != nil {
		return
	}
	slog.Info("Connected to DB")

	if s.listener, err = net.Listen("tcp", ":3001"); err != nil {
		return err
	}
	slog.Info("Servering on:", "Host", "localhost", "Port", "3001")

	go s.server.Serve(s.listener)

	return nil
}

func (s *Server) Stop() error {
	slog.Info("Closing Server on:", "Host", "localhost", "Port", "3001")
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	slog.Info("Closing connection to DB")
	return s.db.Close()
}
