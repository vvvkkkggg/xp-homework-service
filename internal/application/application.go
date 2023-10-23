package application

import (
	"context"
	"html/template"
	"log"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/vvvkkkggg/xp-homework-service/internal/model"
)

type Application struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Application {
	return &Application{
		db: db,
	}
}

/*
home.html
task.html
*/
func (s *Application) Home(ctx context.Context, userID int) (*template.Template, model.Info, error) {
	query := `SELECT role, group_id, name FROM users where user_id = $1`
	user := model.User{}
	row := s.db.QueryRow(ctx, query, strconv.Itoa(userID))
	user.UserID = userID
	if err := row.Scan(
		&user.Role,
		&user.GroupID,
		&user.Name,
	); err != nil {
		return nil, err
	}

	tasks := make([]model.Task, 0, 20)
	query = `SELECT task_id, name, description FROM tasks WHERE group_id = $1`
	rows, err := s.db.Query(ctx, query, strconv.Itoa(user.GroupID))
	if err != nil {
		log.Print("cannot retrieve tasks")
	}
	for rows.Next() {
		var task model.Task
		task.GroupID = user.GroupID
		if err := rows.Scan(
			&task.TaskID,
			&task.Name,
			&task.Description,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	tmpl, err := template.New("").ParseFiles("internal/templates/home.html", "internal/templates/base.html")
	if err != nil {
		log.Print("smth went wrong")
	}

	return tmpl, nil
}

func (s *Application) Task(ctx context.Context, user int, task int) (*template.Template, error) {
	return nil
}

//func (s *Application) Submit(ctx context.Context, ) error {
//	return nil
//}
