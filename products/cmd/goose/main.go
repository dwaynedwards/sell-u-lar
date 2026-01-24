package main

import (
	"context"
	"log"
	"os"

	"github.com/dwaynedwards/sell-u-lar/pkg/db"
	"github.com/dwaynedwards/sell-u-lar/products"
	_ "github.com/dwaynedwards/sell-u-lar/products/migrations"
	"github.com/pressly/goose/v3"
)

func main() {
	db := db.NewPostgres(products.Config.DatabaseURL)
	if err := db.Open(); err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()
	command := os.Args[1]
	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db.DB(), "./migrations"); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
