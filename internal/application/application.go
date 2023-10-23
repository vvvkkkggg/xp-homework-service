package application

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Application {
	return &Application{
		db: db,
	}
}

func (s *Application) Home(ctx context.Context) error {
	return nil
}

func (s *Application) Task(ctx context.Context) error {
	return nil
}
