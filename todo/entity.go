package todo

import (
	"time"
	"todoapp-refactory/user"
)

type Todo struct {
	ID          int
	UserID      int
	Title       string
	Task        string
	IsCompleted int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.User
}
