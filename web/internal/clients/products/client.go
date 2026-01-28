package products

import (
	"context"

	pb "github.com/dwaynedwards/sell-u-lar/pkg/proto/products"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.ProductsClient
	conn   *grpc.ClientConn
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: pb.NewProductsClient(conn),
		conn:   conn,
	}
}

func (c *Client) ListProducts(ctx context.Context) (*types.Products, error) {
	res, err := c.client.ListProducts(ctx, &pb.ProductsRequest{})
	if err != nil {
		return &types.Products{}, err
	}

	products := make(types.Products, 0, len(res.Products))

	for _, product := range res.Products {
		products = append(products, types.Product{
			Sku:         product.Sku,
			Title:       product.Title,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageUrl:    product.ImageUrl,
		})
	}

	return &products, nil
}

func (c *Client) ListProductsByBrand(ctx context.Context, brand string) (*types.Products, error) {
	res, err := c.client.ListProductsByBrand(ctx, &pb.ProductsBrandRequest{Brand: brand})
	if err != nil {
		return &types.Products{}, err
	}

	products := make(types.Products, 0, len(res.Products))

	for _, product := range res.Products {
		products = append(products, types.Product{
			Sku:         product.Sku,
			Title:       product.Title,
			Brand:       product.Brand,
			Description: product.Description,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageUrl:    product.ImageUrl,
		})
	}

	return &products, nil
}

func (c *Client) GetProductByBrandAndSku(ctx context.Context, brand, sku string) (*types.Product, error) {
	res, err := c.client.GetProductByBrandAndSku(ctx, &pb.ProductRequest{Brand: brand, Sku: sku})
	if err != nil {
		return &types.Product{}, err
	}

	return &types.Product{
		Sku:         res.Product.Sku,
		Title:       res.Product.Title,
		Brand:       res.Product.Brand,
		Description: res.Product.Description,
		Price:       res.Product.Price,
		Rating:      res.Product.Rating,
		ImageUrl:    res.Product.ImageUrl,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
