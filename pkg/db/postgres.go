package db

import (
	"database/sql"

	"github.com/dwaynedwards/sell-u-lar/pkg/errors"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db    *sql.DB
	dbUrl string
}

func NewPostgres(dbUrl string) *Postgres {
	return &Postgres{
		dbUrl: dbUrl,
	}
}

func (p *Postgres) Open() (err error) {
	if p.dbUrl == "" {
		return errors.InternalError("db url required")
	}

	if p.db, err = sql.Open("postgres", p.dbUrl); err != nil {
		return
	}

	if err = p.db.Ping(); err != nil {
		return
	}

	return
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}
