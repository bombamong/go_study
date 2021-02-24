package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo
var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func addTestTodo() {
	todoMap[1] = &Todo{
		ID:        1,
		Name:      "do something",
		Completed: false,
		CreatedAt: time.Now(),
	}
	todoMap[2] = &Todo{
		ID:        2,
		Name:      "do something else",
		Completed: true,
		CreatedAt: time.Now(),
	}
	todoMap[3] = &Todo{
		ID:        3,
		Name:      "do something else yey",
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTodoListHandler(w http.ResponseWriter, r *http.Request) {
	todoName := r.FormValue("name")
	todo := &Todo{ID: len(todoMap) + 1, Name: todoName, CreatedAt: time.Now(), Completed: false}
	todoMap[todo.ID] = todo
	rd.JSON(w, http.StatusOK, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("completed") == "true"
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeNewHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestTodo()
	rd = render.New()
	r := mux.NewRouter()
	r.HandleFunc("/todos", addTodoListHandler).Methods("POST")
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)
	return r
}
