package todo

import "OAuth/user"

type GetTodoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTodoInput struct {
	Task             	string `json:"task"  binding:"required"`
	Title 				string `json:"title"  binding:"required"`
	User             	user.User
}

type DeleteTodoInput struct {
	User 	user.User
}

type UpdateTodoInput struct {
	Task             	string `json:"task"  binding:"required"`
	Title 				string `json:"title"  binding:"required"`
	IsCompleted      	int `json:"is_completed" binding:"required"`
	User             	user.User
}