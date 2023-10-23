package model

type User struct {
	UserID  int
	Role    int
	GroupID int
	Name    string
}

type Task struct {
	TaskID      int
	GroupID     int
	name        int
	description string
}

type Assignment struct {
	AssignmentID int
	TaskID       int
	UserID       int
	Status       int
	File         string
	Feedback     string
}

// Выводим на главной странице
type UserTask struct {
	TaskInfo Task
	Status   int
}

type HomePage struct {
	UserInfo User
	Tasks    []UserTask
}

// ----------------------------
// Выводим на странице задачи
type UserAssignment struct {
	TaskInfo       Task
	AssignmentInfo Assignment
}

// ----------------------------
