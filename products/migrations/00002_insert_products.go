package migrations

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/dwaynedwards/sell-u-lar/pkg/types"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInsertProducts, downInsertProducts)
}

func upInsertProducts(ctx context.Context, tx *sql.Tx) error {
	jsonBytes, err := os.ReadFile("./static/json/products.json")
	if err != nil {
		return err
	}
	var products types.Products
	if err = json.Unmarshal(jsonBytes, &products); err != nil {
		return err
	}

	query := `
INSERT INTO products_tbl (sku, title, brand, description, price, rating, image_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
`
	for _, product := range products {
		args := []any{
			product.Sku,
			product.Title,
			product.Brand,
			product.Description,
			product.Price,
			product.Rating,
			product.ImageUrl,
		}
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return err
		}
	}
	return nil
}

func downInsertProducts(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, "DELETE FROM products_tbl;"); err != nil {
		return err
	}
	return nil
}
