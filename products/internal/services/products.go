package services

import (
	"context"
	"net/url"

	"github.com/dwaynedwards/sell-u-lar/pkg/types"
)

type ProductsStore interface {
	List(ctx context.Context) (types.Products, error)
	ListByBrand(ctx context.Context, brand string) (types.Products, error)
	GetByBrandAndSku(ctx context.Context, brand, sku string) (types.Product, error)
}

type ProductsService struct {
	store ProductsStore
}

func NewProductsService(store ProductsStore) *ProductsService {
	return &ProductsService{
		store: store,
	}
}

func (ps *ProductsService) ListProducts(ctx context.Context) (types.Products, error) {
	return ps.store.List(ctx)
}

func (ps *ProductsService) ListProductsByBrand(ctx context.Context, brand string) (types.Products, error) {
	return ps.store.ListByBrand(ctx, url.PathEscape(brand))
}

func (ps *ProductsService) GetProductByBrandAndSku(ctx context.Context, brand, sku string) (types.Product, error) {
	return ps.store.GetByBrandAndSku(ctx, url.PathEscape(brand), url.PathEscape(sku))
}
