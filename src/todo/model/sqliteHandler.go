package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id			INTEGER PRIMARY KEY AUTOINCREMENT,
			sessionID	STRING,
			name		TEXT,
			completed 	BOOLEAN,
			createdAt 	DATETIME
		);
		CREATE INDEX IF NOT EXISTS sessionIDIndexOnTodos ON todos (
			sessionID ASC
		);
		`)
	defer statement.Close()
	statement.Exec()
	return &sqliteHandler{db: database}
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func (s *sqliteHandler) GetTodos(sessionID string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query(`SELECT id, name, completed, createdAt FROM todos WHERE sessionID=?`, sessionID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}

func (s *sqliteHandler) AddTodo(sessionID, name string) *Todo {
	stmt, err := s.db.Prepare(`
		INSERT INTO todos (sessionID, name, completed, createdAt) 
		VALUES(?, ?, ?, datetime('now'));
	`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(sessionID, name, false)
	if err != nil {
		panic(err)
	}
	id, _ := rst.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	defer stmt.Close()
	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare(`
		DELETE FROM todos 
		WHERE id=?		
	`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	count, _ := rst.RowsAffected()
	return count > 0
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare(`
		UPDATE todos 
		SET completed=? 
		WHERE id=?`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(complete, id)
	count, _ := rst.RowsAffected()
	return count > 0
}
