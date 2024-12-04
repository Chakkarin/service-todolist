package todousecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Chakkarin/service-todolist/database/entities"
	"github.com/Chakkarin/service-todolist/modules/todo"
	todorepository "github.com/Chakkarin/service-todolist/modules/todo/todo-repository"
	"github.com/google/uuid"
)

type (
	TodoUsecaseService interface {
		CreateTodo(todo *entities.Todos) error
		GetTodos(search *todo.SearchTodoReq) (*todo.SearchTodoRes, error)
		GetTodoDetail(id *uuid.UUID) (*todo.Todo, error)
		UpdateTodo(id *uuid.UUID, todo *entities.Todos) error
		DeleteTodo(id *uuid.UUID) error
	}

	todoUsecase struct {
		todoRepository todorepository.TodoRepositoryService
	}
)

func NewTodosUsecase(todosRepository todorepository.TodoRepositoryService) TodoUsecaseService {
	return &todoUsecase{
		todoRepository: todosRepository,
	}
}

func (u *todoUsecase) CreateTodo(todo *entities.Todos) error {

	if todo.Priority == "" {
		return errors.New("Priority is required")
	}

	if todo.DueDate.IsZero() {
		return errors.New("Due date is required")
	}

	if todo.Title == "" {
		return errors.New("Title is required")
	}

	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()

	if _, ex := u.todoRepository.Create(todo); ex != nil {
		return ex
	}

	return nil
}

func (u *todoUsecase) GetTodos(search *todo.SearchTodoReq) (*todo.SearchTodoRes, error) {

	// if search.Description == "" && search.Title == "" && search.DueDate.IsZero() && search.Priority == "" {
	// 	return nil, errors.New("At least one search criteria is required")
	// }

	res, ex := u.todoRepository.GetAll(search)
	if ex != nil {
		return nil, ex
	}

	if res.Items == nil {
		res.Items = make([]todo.CreateTodo, 0)
	}

	return res, nil

}

func (u *todoUsecase) GetTodoDetail(id *uuid.UUID) (*todo.Todo, error) {

	res, ex := u.getTodoById(id)
	if ex != nil {
		return nil, ex
	}

	return &todo.Todo{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		DueDate:     res.DueDate,
		Priority:    res.Priority,
		Completed:   res.Completed,
	}, nil
}

func (u *todoUsecase) UpdateTodo(id *uuid.UUID, todo *entities.Todos) error {
	_, ex := u.getTodoById(id)
	if ex != nil {
		return ex
	}

	todo.UpdatedAt = time.Now()

	if ex := u.todoRepository.Update(id, todo); ex != nil {
		return ex
	}

	return nil
}

func (u *todoUsecase) DeleteTodo(id *uuid.UUID) error {

	_, ex := u.getTodoById(id)
	if ex != nil {
		return ex
	}

	if ex := u.todoRepository.Delete(id); ex != nil {
		return ex
	}

	return nil
}

func (u *todoUsecase) getTodoById(id *uuid.UUID) (*entities.Todos, error) {

	if id == nil || *id == uuid.Nil {
		return nil, fmt.Errorf("Failed id is required")
	}

	row, ex := u.todoRepository.FindById(id)
	if ex != nil {
		return nil, ex
	}

	if row == nil {
		return nil, fmt.Errorf("Todo not found with id: %v", id)
	}

	return row, nil
}
