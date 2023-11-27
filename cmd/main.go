package cmd

import (
	"context"
	"log"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
	"github.com/vvvkkkggg/xp-homework-service/internal/db"
	"github.com/vvvkkkggg/xp-homework-service/internal/rest"
)

func Run() {
	ctx, _ := context.WithCancel(context.Background())
	pool, err := db.NewPool(ctx)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("cant connect to db")
	}
	defer pool.Close()
	app := application.New(pool)

	router := rest.NewRouter(rest.NewHandler(app))

	if err = rest.CreateAndServe(ctx, router); err != nil {
		log.Fatal("unable to run server")
	}
}
