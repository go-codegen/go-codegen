package repository

import (
	"github.com/go-codegen/go-codegen/test"
	"gorm.io/gorm"
)

type HelloRepository struct {
	db *gorm.DB
}

type HelloRepositoryImpl interface {
	Create(h *test.Hello) (*test.Hello, error)
	FindByID(id string) (*test.Hello, error)
	Update(h *test.Hello) (*test.Hello, error)
	Delete(id string) error
}

func (r *HelloRepository) Create(h *test.Hello) (*test.Hello, error) {
	if err := r.db.Create(&h).Error; err != nil {
		return nil, err
	}

	return h, nil
}
func (r *HelloRepository) FindByID(id string) (*test.Hello, error) {
	var h test.Hello

	if err := r.db.Where("id = ?", id).First(&h).Error; err != nil {
		return nil, err
	}

	return &h, nil
}
func (r *HelloRepository) Update(h *test.Hello) (*test.Hello, error) {
	if err := r.db.Save(&h).Error; err != nil {
		return nil, err
	}

	return h, nil
}
func (r *HelloRepository) Delete(id string) error {
	if err := r.db.Delete(&test.Hello{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
