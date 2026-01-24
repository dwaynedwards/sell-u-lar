package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateProductsTbl, downCreateProductsTbl)
}

func upCreateProductsTbl(ctx context.Context, tx *sql.Tx) error {
	query := `
CREATE TABLE IF NOT EXISTS products_tbl (
	sku TEXT NOT NULL,
	title TEXT NOT NULL,
	brand TEXT NOT NULL,
	description TEXT NOT NULL,
	price INTEGER NOT NULL,
	rating INTEGER NOT NULL,
	image_url TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	CONSTRAINT pk_products PRIMARY KEY (sku),
	CONSTRAINT check_sku_length CHECK (char_length(sku)<=50),
	CONSTRAINT check_title_length CHECK (char_length(title)<=50),
	CONSTRAINT check_brand_length CHECK (char_length(brand)<=50),
	CONSTRAINT check_description_length CHECK (char_length(description)<=256),
	CONSTRAINT check_image_url_length CHECK (char_length(image_url)<=256)
);

CREATE OR REPLACE FUNCTION set_timestamp()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.modified_at = NOW();
		RETURN NEW;
	END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trigger_set_timestamp
	BEFORE UPDATE ON products_tbl
	FOR EACH ROW
	EXECUTE FUNCTION set_timestamp();
`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}
	return nil
}

func downCreateProductsTbl(ctx context.Context, tx *sql.Tx) error {
	query := `
DROP TRIGGER IF EXISTS trigger_set_timestamp ON products_tbl;
DROP FUNCTION IF EXISTS set_timestamp;
DROP TABLE IF EXISTS products_tbl;
`

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}
	return nil
}
