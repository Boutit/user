package store

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Boutit/user/internal/config"
	"github.com/Boutit/user/internal/models"
)

type postgresStore struct {
	db *sqlx.DB
}

type UserStore interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) error
}

func CreatePostgresStore(cfg config.Config) (UserStore, error) {
	db, err := sqlx.Connect("postgres", cfg.PostgresConfig.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return &postgresStore{
		db: db,
	}, nil
}

func (store *postgresStore) CreateUser(ctx context.Context,  user *models.User) error {
	_, err := store.db.NamedExecContext(ctx, createUserSQL, user)
	if err != nil {
		fmt.Println("farther", err)
		return err
	}
	return nil
}

func (store *postgresStore) GetUserById(ctx context.Context,  userId string) error {
	_, err := store.db.NamedExecContext(ctx, getUserByIdSQL, userId)
	if err != nil {
		return err
	}
	
	return nil
}