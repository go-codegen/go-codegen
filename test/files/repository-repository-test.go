package repository

import (
	"github.com/go-codegen/go-codegen/test"
	"gorm.io/gorm"
)

type RepositoryTestRepository struct {
	db *gorm.DB
}

type RepositoryTestRepositoryImpl interface {
	NewRepositoryTestRepository(db *gorm.DB) *RepositoryTestRepository
	Create(r1 *test.RepositoryTest) (*test.RepositoryTest, error)
	FindByID(id string) (*test.RepositoryTest, error)
	FindByHelloID(ID int) ([]*test.RepositoryTest, error)
	FindByHelloData(Data string) ([]*test.RepositoryTest, error)
	FindByHelloIDAndData(ID int, Data string) ([]*test.RepositoryTest, error)
	FindByNameAction(NameAction string) ([]*test.RepositoryTest, error)
	FindByAge(Age int) (*test.RepositoryTest, error)
	Update(r1 *test.RepositoryTest) (*test.RepositoryTest, error)
	Delete(id string) error
}

func NewRepositoryTestRepository(db *gorm.DB) *RepositoryTestRepository {
	return &RepositoryTestRepository{
		db: db,
	}
}

func (r *RepositoryTestRepository) Create(r1 *test.RepositoryTest) (*test.RepositoryTest, error) {
	if err := r.db.Create(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByID(id string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("id = ?", id).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepository) FindByHelloID(ID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN hellos ON hellos.id = repository_tests.hello_id").
		Where("hellos.id = ?", ID).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByHelloData(Data string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN hellos ON hellos.id = repository_tests.hello_id").
		Where("hellos.data = ?", Data).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByHelloIDAndData(ID int, Data string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN hellos ON hellos.id = repository_tests.hello_id").
		Where("hellos.id = ? AND hellos.data = ?", ID, Data).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByNameAction(NameAction string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("name_action = ?", NameAction).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByAge(Age int) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("age = ?", Age).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepository) Update(r1 *test.RepositoryTest) (*test.RepositoryTest, error) {
	if err := r.db.Save(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) Delete(id string) error {
	if err := r.db.Delete(&test.RepositoryTest{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
