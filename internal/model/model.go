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
	Feedback     string
}
