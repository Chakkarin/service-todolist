package todorepository

import (
	"errors"
	"fmt"

	"github.com/Chakkarin/service-todolist/database/entities"
	"github.com/Chakkarin/service-todolist/modules/todo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	TodoRepositoryService interface {
		Create(todo *entities.Todos) (*entities.Todos, error)
		Update(id *uuid.UUID, todo *entities.Todos) error
		GetAll(search *todo.SearchTodoReq) (*todo.SearchTodoRes, error)
		FindById(id *uuid.UUID) (*entities.Todos, error)
		Delete(id *uuid.UUID) error
	}

	todoRepository struct {
		db *gorm.DB
	}
)

func NewTodoRepository(db *gorm.DB) TodoRepositoryService {
	return &todoRepository{db}
}

func (r *todoRepository) Create(todo *entities.Todos) (*entities.Todos, error) {
	insertTodo := new(entities.Todos)

	if ex := r.db.Create(todo).Scan(insertTodo).Error; ex != nil {
		return nil, errors.New("Creating todo failed ::" + ex.Error())
	}

	return insertTodo, nil
}

func (r *todoRepository) Update(id *uuid.UUID, todo *entities.Todos) error {
	if ex := r.db.Model(&entities.Todos{}).Where("id = ?", id).Updates(todo).Error; ex != nil {
		return errors.New("Failed to update todo: " + ex.Error())
	}
	return nil
}

func (r *todoRepository) GetAll(search *todo.SearchTodoReq) (*todo.SearchTodoRes, error) {
	var todos []todo.CreateTodo // ใช้ Slice ของ Todo สำหรับผลลัพธ์
	var result todo.SearchTodoRes

	query := r.db.Model(&entities.Todos{})

	// กรองตาม Title
	if search.Title != "" {
		query = query.Where("title ILIKE ?", "%"+search.Title+"%")
	}

	// กรองตาม DueDate
	if !search.DueDate.IsZero() {
		query = query.Where("due_date = ?", search.DueDate)
	}

	// กรองตาม Priority
	if search.Priority != "" {
		query = query.Where("priority = ?", search.Priority)
	}

	// กรองตาม Completed
	if search.Completed {
		query = query.Where("completed = ?", search.Completed)
	}

	// นับจำนวนทั้งหมด
	if ex := query.Count(&result.Total).Error; ex != nil {
		return nil, ex
	}

	// Pagination
	offset, limit := parsePagination(search.PaginateReq)
	if ex := query.Offset(offset).Limit(limit).Find(&todos).Error; ex != nil {
		return nil, ex
	}

	// แมปข้อมูล
	result.Items = todos
	result.Start = search.Start
	result.Limit = search.Limit

	return &result, nil
}

func (r *todoRepository) FindById(id *uuid.UUID) (*entities.Todos, error) {
	todo := new(entities.Todos)

	if ex := r.db.Where("id = ?", id).First(todo).Error; ex != nil {
		if errors.Is(ex, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Todo not found with id: %v", id)
		}
		return nil, fmt.Errorf("Find todo failed: %v", ex)
	}

	return todo, nil
}

func (r *todoRepository) Delete(id *uuid.UUID) error {

	if ex := r.db.Where("id = ?", id).Delete(&entities.Todos{}).Error; ex != nil {
		return fmt.Errorf("Failed to delete todo with id %s: %v", id, ex)
	}

	return nil
}

func parsePagination(paginate todo.PaginateReq) (offset int, limit int) {
	limit = paginate.Limit
	if paginate.Start != 0 {
		offset = paginate.Start
	}
	return
}
