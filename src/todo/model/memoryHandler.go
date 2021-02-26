package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

func (m *memoryHandler) GetTodos(sessionID string) []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(sessionID, name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{
		ID:        id,
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
	}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if _, ok := m.todoMap[id]; ok {
		m.todoMap[id].Completed = complete
		return true
	}
	return false
}

func (m *memoryHandler) Close() {}
