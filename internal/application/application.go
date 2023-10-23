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
func (s *Application) Home(ctx context.Context, userID int) (*template.Template, *model.HomePage, error) {
	query := `SELECT role, group_id, name FROM users where user_id = $1`
	user := model.User{}
	row := s.db.QueryRow(ctx, query, strconv.Itoa(userID))
	user.UserID = userID
	if err := row.Scan(
		&user.Role,
		&user.GroupID,
		&user.Name,
	); err != nil {
		return nil, nil, err
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
			log.Print("unable to retrieve data from query")
			return nil, nil, err
		}

		tasks = append(tasks, task)
	}

	// retrieving statuses
	taskToStatus := make(map[int]int, 20)
	query = `SELECT task_id, status FROM assignments WHERE user_id = $1`
	rows, err = s.db.Query(ctx, query, strconv.Itoa(user.UserID))
	if err != nil {
		log.Print("cannot retrieve tasks")
	}
	for rows.Next() {
		var taskID int
		var status int
		if err := rows.Scan(
			&taskID,
			&status,
		); err != nil {
			return nil, nil, err
		}

		taskToStatus[taskID] = status
	}

	// filling user tasks
	userTasks := make([]model.UserTask, 0, 20)
	for _, task := range tasks {
		userTask := model.UserTask{
			TaskInfo: task,
			Status:   taskToStatus[task.TaskID],
			UserID:   userID,
		}
		userTasks = append(userTasks, userTask)
	}

	homeDir := "/Users/r.gostilo/Projects/hse/se/xp-homework-service/internal/templates/home.html"
	//baseDir := "/Users/r.gostilo/Projects/hse/se/xp-homework-service/internal/templates/base.html"
	tmpl := template.Must(template.ParseFiles(homeDir))

	return tmpl, &model.HomePage{
		UserInfo: user,
		Tasks:    userTasks,
	}, nil
}

/*
create table if not exists tasks
(
    task_id     SERIAL PRIMARY KEY,
    group_id    int,
    name        varchar(50),
    description varchar(500)
);
*/

func (s *Application) Task(ctx context.Context, userID int, taskID int) (*template.Template, *model.AssignmentPage, error) {
	var user model.User
	query := `SELECT user_id, role, group_id, name FROM users WHERE user_id = $1`
	row := s.db.QueryRow(ctx, query, strconv.Itoa(userID))
	if err := row.Scan(
		&user.UserID,
		&user.Role,
		&user.GroupID,
		&user.Name,
	); err != nil {
		return nil, nil, err
	}

	var task model.Task
	query = `SELECT task_id, group_id, name, description FROM tasks WHERE task_id = $1`
	row = s.db.QueryRow(ctx, query, strconv.Itoa(taskID))
	if err := row.Scan(
		&task.TaskID,
		&task.GroupID,
		&task.Name,
		&task.Description,
	); err != nil {
		return nil, nil, err
	}

	// retrieving statuses
	var assignment model.Assignment
	query = `SELECT task_id, user_id, status, feedback, file FROM assignments WHERE user_id = $1 AND task_id = $2`
	row = s.db.QueryRow(ctx, query, strconv.Itoa(userID), strconv.Itoa(taskID))
	if err := row.Scan(
		&assignment.TaskID,
		&assignment.UserID,
		&assignment.Status,
		&assignment.Feedback,
		&assignment.File,
	); err != nil {
		return nil, nil, err
	}

	taskDir := "internal/templates/task.html"
	tmpl := template.Must(template.ParseFiles(taskDir))

	return tmpl, &model.AssignmentPage{
		UserInfo:       user,
		TaskInfo:       task,
		AssignmentInfo: assignment,
	}, nil
}

//func (s *Application) Submit(ctx context.Context, ) error {
//	return nil
//}
