package model

import "time"

// var handler DBHandler

// func init() {
// 	handler = newSqliteHandler()
// }

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(string) *Todo
	RemoveTodo(int) bool
	CompleteTodo(int, bool) bool
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}
