package repository

import (
	todo "github.com/oloomoses/todo/internal/model"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) Create(todo *todo.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepo) FindAll() ([]todo.Todo, error) {
	var todos []todo.Todo

	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepo) FindByID(id uint64) (todo.Todo, error) {
	var todo todo.Todo

	err := r.db.First(&todo, id).Error
	return todo, err
}

func (r *TodoRepo) Update(id uint64, updates map[string]interface{}) error {
	return r.db.Model(&todo.Todo{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TodoRepo) Delete(id uint64) error {
	return r.db.Delete(&todo.Todo{}, id).Error
}
