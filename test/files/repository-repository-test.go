package repository

import (
	"github.com/go-codegen/go-codegen/test"
	"gorm.io/gorm"
	"time"
)

type RepositoryTestRepository struct {
	db *gorm.DB
}

type RepositoryTestRepositoryImpl interface {
	NewRepositoryTestRepository(db *gorm.DB) *RepositoryTestRepository
	Create(r1 *test.RepositoryTest) (*test.RepositoryTest, error)
	FindByID(id string) (*test.RepositoryTest, error)
	FindByUserName(Name string) ([]*test.RepositoryTest, error)
	FindByUserEmail(Email string) ([]*test.RepositoryTest, error)
	FindByUserPassword(Password string) ([]*test.RepositoryTest, error)
	FindByUserNameAndEmail(Name string, Email string) ([]*test.RepositoryTest, error)
	FindByUserNameAndPassword(Name string, Password string) ([]*test.RepositoryTest, error)
	FindByUserEmailAndPassword(Email string, Password string) ([]*test.RepositoryTest, error)
	FindByUserNameAndEmailAndPassword(Name string, Email string, Password string) ([]*test.RepositoryTest, error)
	FindByAccessToken(AccessToken string) ([]*test.RepositoryTest, error)
	FindByRefreshToken(RefreshToken string) ([]*test.RepositoryTest, error)
	FindByExpiresAt(ExpiresAt time.Time) ([]*test.RepositoryTest, error)
	FindByUserID(UserID int) ([]*test.RepositoryTest, error)
	FindByIP(IP string) ([]*test.RepositoryTest, error)
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

func (r *RepositoryTestRepository) FindByUserName(Name string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.name = ?", Name).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserEmail(Email string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.email = ?", Email).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserPassword(Password string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.password = ?", Password).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserNameAndEmail(Name string, Email string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.name = ? AND users.email = ?", Name, Email).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserNameAndPassword(Name string, Password string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.name = ? AND users.password = ?", Name, Password).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserEmailAndPassword(Email string, Password string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.email = ? AND users.password = ?", Email, Password).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserNameAndEmailAndPassword(Name string, Email string, Password string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user_id").
		Where("users.name = ? AND users.email = ? AND users.password = ?", Name, Email, Password).
		Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByAccessToken(AccessToken string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ?", AccessToken).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByRefreshToken(RefreshToken string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("refresh_token = ?", RefreshToken).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByExpiresAt(ExpiresAt time.Time) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("expires_at = ?", ExpiresAt).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByUserID(UserID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("user_id = ?", UserID).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepository) FindByIP(IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("ip = ?", IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
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
