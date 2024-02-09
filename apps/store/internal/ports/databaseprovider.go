package ports

import "database/sql"

type DatabaseProvider interface {
	GetDatabase(name string) (*sql.DB, error)
}
