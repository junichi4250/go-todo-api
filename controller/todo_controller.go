package controller

import (
	"encoding/json"
	"go-todo-api/controller/dto"
	"go-todo-api/model/entity"
	"go-todo-api/model/repository"
	"net/http"
	"path"
	"strconv"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []dto.TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(todosResponse.Todos, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	// json（body）を構造体に変換
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{Title: todoRequest.Title, Content: todoRequest.Content}

	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *todoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{Id: todoId, Title: todoRequest.Title, Content: todoRequest.Content}

	err = tc.tr.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = tc.tr.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(204)
}
