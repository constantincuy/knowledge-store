package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Provider struct {
	username string
	password string
	host     string
	sslMode  string
}

func (p Provider) GetDatabase(name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", p.username, p.password, p.host, name, p.sslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewProvider(username, password, host, sslMode string) Provider {
	return Provider{username, password, host, sslMode}
}
