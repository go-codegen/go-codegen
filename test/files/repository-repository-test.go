package repository

import (
	"github.com/go-codegen/go-codegen/test"
	"gorm.io/gorm"
	"time"
)

type RepositoryTestRepositoryImpl struct {
	db *gorm.DB
}

type RepositoryTestRepository interface {
	NewRepositoryTestRepositoryImpl(db *gorm.DB) *RepositoryTestRepositoryImpl
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

func NewRepositoryTestRepositoryImpl(db *gorm.DB) *RepositoryTestRepositoryImpl {
	return &RepositoryTestRepositoryImpl{
		db: db,
	}
}

func (r *RepositoryTestRepositoryImpl) Create(r1 *test.RepositoryTest) (*test.RepositoryTest, error) {
	if err := r.db.Create(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByID(id string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("id = ?", id).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByUserName(Name string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserEmail(Email string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserPassword(Password string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserNameAndEmail(Name string, Email string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserNameAndPassword(Name string, Password string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserEmailAndPassword(Email string, Password string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByUserNameAndEmailAndPassword(Name string, Email string, Password string) ([]*test.RepositoryTest, error) {
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

func (r *RepositoryTestRepositoryImpl) FindByAccessToken(AccessToken string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ?", AccessToken).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByRefreshToken(RefreshToken string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("refresh_token = ?", RefreshToken).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByExpiresAt(ExpiresAt time.Time) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("expires_at = ?", ExpiresAt).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByUserID(UserID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("user_id = ?", UserID).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByIP(IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("ip = ?", IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) Update(r1 *test.RepositoryTest) (*test.RepositoryTest, error) {
	if err := r.db.Save(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) Delete(id string) error {
	if err := r.db.Delete(&test.RepositoryTest{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
