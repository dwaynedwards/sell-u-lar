package https

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"github.com/dwaynedwards/sell-u-lar/web"
	"github.com/dwaynedwards/sell-u-lar/web/internal/clients/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ShutdownTimeout = 5 * time.Second

type ProductsClient interface {
	ListProducts(ctx context.Context) (*types.Products, error)
	ListProductsByBrand(ctx context.Context, brand string) (*types.Products, error)
	GetProductByBrandAndSku(ctx context.Context, brand, sku string) (*types.Product, error)
	Close() error
}

type Server struct {
	listener net.Listener
	server   *http.Server
	router   *http.ServeMux
	products ProductsClient
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: http.NewServeMux(),
	}

	s.server.Handler = http.HandlerFunc(s.router.ServeHTTP)

	// Include the static content.
	s.router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	s.registerProductRoutes()
	return s
}

func (s *Server) Start() error {
	slog.Info("Servering on:", "Addr", web.Config.WebServerAddr)
	var err error
	if s.listener, err = net.Listen("tcp", web.Config.WebServerAddr); err != nil {
		return err
	}

	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	productsConn, err := grpc.NewClient(web.Config.ProductsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	s.products = products.NewClient(productsConn)
	return nil
}

func (s *Server) Stop() error {
	slog.Info("Closing Server on:", "Addr", web.Config.WebServerAddr)
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := s.products.Close(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return s.server.Shutdown(ctx)
}
