package cmd

import (
	"context"
	"database/sql"
	"log"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
	"github.com/vvvkkkggg/xp-homework-service/internal/db"
	"github.com/vvvkkkggg/xp-homework-service/internal/rest"
)

func Run() {
	ctx, _ := context.WithCancel(context.Background())
	pool, err := db.NewPool(ctx)
	if err != nil {
		log.Fatal("cant connect to db")
	}
	defer pool.Close()
	app := application.New(pool)

	router := rest.NewRouter(rest.NewHandler(app))

	if err = rest.CreateAndServe(ctx, router); err != nil {
		log.Fatal("unable to run server")
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
