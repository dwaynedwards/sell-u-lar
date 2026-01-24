package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dwaynedwards/sell-u-lar/pkg/db"
	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	_ "github.com/lib/pq"
)

type ProductsStore struct {
	db.Database
}

func NewProductsStore(db db.Database) *ProductsStore {
	return &ProductsStore{
		Database: db,
	}
}

func (ps *ProductsStore) List(ctx context.Context) (types.Products, error) {
	query := `
SELECT sku, title, brand, description, price, rating, image_url
	FROM products_tbl;
`
	rows, err := ps.DB().QueryContext(ctx, query)
	if err != nil {
		return types.Products{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	products := make(types.Products, 0)
	for rows.Next() {
		var product types.Product
		values := []any{
			&product.Sku,
			&product.Title,
			&product.Brand,
			&product.Description,
			&product.Price,
			&product.Rating,
			&product.ImageUrl,
		}
		if err = rows.Scan(values...); err != nil {
			return types.Products{}, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return types.Products{}, err
	}
	return products, nil
}

func (ps *ProductsStore) ListByBrand(ctx context.Context, brand string) (types.Products, error) {
	query := `
SELECT sku, title, brand, description, price, rating, image_url
	FROM products_tbl
	WHERE LOWER(brand) = LOWER($1);
`

	rows, err := ps.DB().QueryContext(ctx, query, brand)
	if err != nil {
		return types.Products{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	products := make(types.Products, 0)
	for rows.Next() {
		var product types.Product
		values := []any{
			&product.Sku,
			&product.Title,
			&product.Brand,
			&product.Description,
			&product.Price,
			&product.Rating,
			&product.ImageUrl,
		}
		if err = rows.Scan(values...); err != nil {
			return types.Products{}, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return types.Products{}, err
	}
	return products, nil
}

func (ps *ProductsStore) GetByBrandAndSku(ctx context.Context, brand, sku string) (types.Product, error) {
	query := `
SELECT sku, title, brand, description, price, rating, image_url
	FROM products_tbl
	WHERE LOWER(brand) = LOWER($1) AND LOWER(sku )= LOWER($2);
`
	var product types.Product
	values := []any{
		&product.Sku,
		&product.Title,
		&product.Brand,
		&product.Description,
		&product.Price,
		&product.Rating,
		&product.ImageUrl,
	}
	if err := ps.DB().QueryRowContext(ctx, query, brand, sku).Scan(values...); err != nil {
		if ok := errors.Is(err, sql.ErrNoRows); ok {
			err = errors.NotFoundError(fmt.Sprintf("Product not found with brand: [%s] and sku: [%s]", brand, sku))
		}
		return types.Product{}, err
	}
	return product, nil
}
