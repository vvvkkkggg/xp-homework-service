package application

import (
	"context"
	"html/template"
	"net/http"
	"path"

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

type Task struct {
	Title  string
	Author string
}

/*
home.html
task.html
*/
func (s *Application) Home(ctx context.Context, user int) (*template.Template, error) {
	task := Task{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "home.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (s *Application) Task(ctx context.Context, user int, task int) (*template.Template, error) {
	return nil
}

//func (s *Application) Submit(ctx context.Context, ) error {
//	return nil
//}
