package grpc

import (
	"context"
	"log"
	"log/slog"
	"net"

	"github.com/dwaynedwards/sell-u-lar/pkg/db"
	pb "github.com/dwaynedwards/sell-u-lar/pkg/proto/products"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"github.com/dwaynedwards/sell-u-lar/products"
	"github.com/dwaynedwards/sell-u-lar/products/internal/services"
	"github.com/dwaynedwards/sell-u-lar/products/internal/store"
	"google.golang.org/grpc"
)

type ProductsService interface {
	ListProducts(ctx context.Context) (types.Products, error)
	ListProductsByBrand(ctx context.Context, brand string) (types.Products, error)
	GetProductByBrandAndSku(ctx context.Context, brand, sku string) (types.Product, error)
}

type Server struct {
	listener net.Listener
	server   *grpc.Server
	db       db.Database
	products ProductsService
	pb.UnimplementedProductsServer
}

func NewServer(db db.Database) *Server {
	s := &Server{
		db:       db,
		products: services.NewProductsService(store.NewProductsStore(db)),
	}

	return s
}

func (s *Server) Start() (err error) {
	slog.Info("Opening connection to DB")
	if err = s.db.Open(); err != nil {
		return
	}
	slog.Info("Connected to DB")

	if s.listener, err = net.Listen("tcp", products.Config.ProductsServerAddr); err != nil {
		return err
	}
	slog.Info("Servering on:", "Addr", products.Config.ProductsServerAddr)

	s.server = grpc.NewServer()
	pb.RegisterProductsServer(s.server, s)

	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	slog.Info("Closing Server on:", "Addr", products.Config.ProductsServerAddr)

	s.server.Stop()

	slog.Info("Closing connection to DB")
	return s.db.Close()
}
