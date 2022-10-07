package store

import (
	"context"
	_ "database/sql"
	"fmt"

	"github.com/Boutit/user/api"
	"github.com/Boutit/user/internal/config"
	"github.com/jmoiron/sqlx"
)

type postgresStore struct {
	conn *sqlx.DB
}

type UserStore interface {
	CreateUser(ctx context.Context, user *api.User) error
}

func CreatePostgresStore(cfg config.Config) (UserStore, error) {
	conn, err := sqlx.Connect("postgres", cfg.PostgresConfig.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return &postgresStore{
		conn: conn,
	}, nil
}

func (store *postgresStore) CreateUser(ctx context.Context,  user *api.User) error {
	_, err := store.conn.NamedExecContext(ctx, createAccountSQL, user)
	if err != nil {
		fmt.Println("farther", err)
		return err
	}
	return nil
}