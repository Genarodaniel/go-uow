package uow

import (
	"context"
	"database/sql"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow UowInterface) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}
