package server

import (
	todohandler "github.com/Chakkarin/service-todolist/modules/todo/todo-handler"
	todorepository "github.com/Chakkarin/service-todolist/modules/todo/todo-repository"
	todousecase "github.com/Chakkarin/service-todolist/modules/todo/todo-usecase"
)

func (s *ginServer) initTodosRouter() {
	router := s.app.Group("/v1/todo")

	todoRepository := todorepository.NewTodoRepository(s.db)
	todoUsecase := todousecase.NewTodosUsecase(todoRepository)
	todoshandler := todohandler.NewTodosHandler(todoUsecase)

	router.GET("/", todoshandler.SearchTodo)
	router.GET("/:id", todoshandler.GetTodoById)
	router.POST("/", todoshandler.CreateTodo)
	router.PUT("/:id", todoshandler.UpdateTodo)
	router.PATCH("/:id/complete", todoshandler.CompleteTodo)
	router.DELETE("/:id", todoshandler.DeleteTodo)
}
