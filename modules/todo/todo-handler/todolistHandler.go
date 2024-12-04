package todohandler

import (
	"net/http"

	"github.com/Chakkarin/service-todolist/database/entities"
	"github.com/Chakkarin/service-todolist/modules/todo"
	todousecase "github.com/Chakkarin/service-todolist/modules/todo/todo-usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	TodoHandlerService interface {
		CreateTodo(c *gin.Context)
		SearchTodo(c *gin.Context)
		GetTodoById(c *gin.Context)
		UpdateTodo(c *gin.Context)
		CompleteTodo(c *gin.Context)
		DeleteTodo(c *gin.Context)
	}

	todoHandler struct {
		todoUsecase todousecase.TodoUsecaseService
	}
)

func NewTodosHandler(todoUsecase todousecase.TodoUsecaseService) TodoHandlerService {
	return &todoHandler{
		todoUsecase: todoUsecase,
	}
}

// CreateTodo godoc
// @Summary Create Todo
// @Description Create a new Todo
// @ID CreateTodoHandler
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param   todo  body     todo.CreateTodo  true  "Todo"
// @Success 201
// @Failure 400 {object} string
// @Router /v1/todo [post]
func (h *todoHandler) CreateTodo(c *gin.Context) {

	var todo todo.CreateTodo

	if ex := c.ShouldBindJSON(&todo); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	entitiesTodo := entities.Todos{
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		Priority:    todo.Priority,
	}

	if ex := h.todoUsecase.CreateTodo(&entitiesTodo); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// SearchTodo godoc
// @Summary Search Todo
// @Description Search Todo by title, due date, priority and completed
// @ID SearchTodoHandler
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param   search  query     todo.SearchTodoReq  true  "Search Todo"
// @Success 200 {object} todo.SearchTodoRes
// @Failure 400 {object} string
// @Router /v1/todo [get]
func (h *todoHandler) SearchTodo(c *gin.Context) {

	var todoSearch todo.SearchTodoReq

	if ex := c.ShouldBindQuery(&todoSearch); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	data, ex := h.todoUsecase.GetTodos(&todoSearch)
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

// GetTodoById godoc
// @Summary Get Todo By Id
// @Description Get Todo by id
// @ID GetTodoByIdHandler
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param   id  path     string  true  "Todo id"
// @Success 200 {object} todo.Todo
// @Failure 400 {object} string
// @Router /v1/todo/{id} [get]
func (h *todoHandler) GetTodoById(c *gin.Context) {

	idTodo, ex := uuid.Parse(c.Param("id"))
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	data, ex := h.todoUsecase.GetTodoDetail(&idTodo)
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})

}

// UpdateTodo godoc
// @Summary Update Todo
// @Description Update a Todo by id
// @ID UpdateTodoHandler
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param   id  query     string  true  "Todo id"
// @Param   todo  body     todo.UpdateTodo  true  "Todo"
// @Success 200
// @Failure 400 {object} string
// @Router /v1/todo/{id} [put]
func (h *todoHandler) UpdateTodo(c *gin.Context) {

	var todo todo.UpdateTodo

	id, ex := uuid.Parse(c.Query("id"))
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	if ex := c.ShouldBindJSON(&todo); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	entitiesTodo := entities.Todos{}

	if todo.Title != "" {
		entitiesTodo.Title = todo.Title
	}

	if todo.Description != "" {
		entitiesTodo.Description = todo.Description
	}

	if !todo.DueDate.IsZero() {
		entitiesTodo.DueDate = todo.DueDate
	}

	if todo.Priority != "" {
		entitiesTodo.Priority = todo.Priority
	}

	if ex := h.todoUsecase.UpdateTodo(&id, &entitiesTodo); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	data, ex := h.todoUsecase.GetTodoDetail(&id)
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)

}

// CompleteTodo godoc
// @Summary Complete Todo
// @Description Complete a Todo by id
// @ID CompleteTodoHandler
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param   id  query     string  true  "Todo id"
// @Success 200 {object} todo.Todo
// @Failure 400 {object} string
// @Router /v1/todo/{id}/complete [patch]
func (h *todoHandler) CompleteTodo(c *gin.Context) {

	id, ex := uuid.Parse(c.Query("id"))
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	entitiesTodo := entities.Todos{
		Completed: true,
	}

	if ex := h.todoUsecase.UpdateTodo(&id, &entitiesTodo); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	data, ex := h.todoUsecase.GetTodoDetail(&id)
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)

}

// DeleteTodo godoc
// @Summary Delete Todo
// @Description Delete a Todo by id
// @ID DeleteTodoHandler
// @Tags Todo
// @Param   id  query     string  true  "Todo id"
// @Success 204
// @Failure 400 {object} string
// @Router /v1/todo/{id} [delete]
func (h *todoHandler) DeleteTodo(c *gin.Context) {

	id, ex := uuid.Parse(c.Query("id"))
	if ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	if ex := h.todoUsecase.DeleteTodo(&id); ex != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": ex.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
