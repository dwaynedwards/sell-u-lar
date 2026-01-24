package db

import (
	"database/sql"
)

type Database interface {
	Open() error
	Close() error
	DB() *sql.DB
}
