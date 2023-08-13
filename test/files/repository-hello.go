package repository

import (
	"gorm.io/gorm"
	"github.com/go-codegen/go-codegen/test"
)

type HelloRepository struct {
	db *gorm.DB
}

type HelloRepositoryImpl interface {
	NewHelloRepository(db *gorm.DB) (*HelloRepository)
	Create(h *test.Hello) (*test.Hello, error)
	FindByID(id string) (*test.Hello, error)
	Update(h *test.Hello) (*test.Hello, error)
	Delete(id string) (error)
}

func NewHelloRepository(db *gorm.DB) *HelloRepository {
	return &HelloRepository{
		db: db,	
	}
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
	if err := r.db.Delete(&test.Hello{},"id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

