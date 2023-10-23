package application

import (
	"context"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (s *Application) Home(ctx context.Context) error {
	return nil
}

func (s *Application) Task(ctx context.Context) error {
	return nil
}
