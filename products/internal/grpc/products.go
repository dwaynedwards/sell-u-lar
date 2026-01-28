package grpc

import (
	"context"
	"log/slog"

	pb "github.com/dwaynedwards/sell-u-lar/pkg/proto/products"
)

func (s *Server) ListProducts(ctx context.Context, in *pb.ProductsRequest) (*pb.ProductsResponse, error) {
	slog.InfoContext(ctx, "List Products")

	resProducts, err := s.products.ListProducts(ctx)
	if err != nil {
		return &pb.ProductsResponse{}, err
	}

	products := make([]*pb.Product, 0, len(resProducts))

	for _, product := range resProducts {
		products = append(products, &pb.Product{
			Sku:         product.Sku,
			Title:       product.Title,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageUrl:    product.ImageUrl,
		})
	}

	return &pb.ProductsResponse{
		Products: products,
	}, nil
}

func (s *Server) ListProductsByBrand(ctx context.Context, in *pb.ProductsBrandRequest) (*pb.ProductsResponse, error) {
	slog.InfoContext(ctx, "List Products By Brand:", "brand", in.Brand)

	resProducts, err := s.products.ListProductsByBrand(ctx, in.Brand)
	if err != nil {
		return &pb.ProductsResponse{}, err
	}

	products := make([]*pb.Product, 0, len(resProducts))

	for _, product := range resProducts {
		products = append(products, &pb.Product{
			Sku:         product.Sku,
			Title:       product.Title,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageUrl:    product.ImageUrl,
		})
	}

	return &pb.ProductsResponse{
		Products: products,
	}, nil
}

func (s *Server) GetProductByBrandAndSku(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	slog.InfoContext(ctx, "Get Product By Brand And Sku:", "brand", in.Brand, "sku", in.Sku)

	product, err := s.products.GetProductByBrandAndSku(ctx, in.Brand, in.Sku)
	if err != nil {
		return &pb.ProductResponse{}, err
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Sku:         product.Sku,
			Title:       product.Title,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageUrl:    product.ImageUrl,
		},
	}, nil
}
