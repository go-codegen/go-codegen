package repository

import (
	"github.com/go-codegen/go-codegen/test"
	"gorm.io/gorm"
)

type RepositoryTestRepositoryImpl struct {
	db *gorm.DB
}

type RepositoryTestRepository interface {
	Create(r1 *test.RepositoryTest) (*test.RepositoryTest, error)
	FindByID(id string) (*test.RepositoryTest, error)
	FindByAccessToken(AccessToken string) (*test.RepositoryTest, error)
	DeleteByAccessToken(AccessToken string) (*test.RepositoryTest, error)
	FindByRefreshToken(RefreshToken string) (*test.RepositoryTest, error)
	DeleteByRefreshToken(RefreshToken string) (*test.RepositoryTest, error)
	FindByUserID(UserID int) ([]*test.RepositoryTest, error)
	FindByIP(IP string) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndRefreshToken(AccessToken string, RefreshToken string) (*test.RepositoryTest, error)
	DeleteByAccessTokenAndRefreshToken(AccessToken string, RefreshToken string) (*test.RepositoryTest, error)
	FindByAccessTokenAndUserID(AccessToken string, UserID int) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndIP(AccessToken string, IP string) ([]*test.RepositoryTest, error)
	FindByRefreshTokenAndUserID(RefreshToken string, UserID int) ([]*test.RepositoryTest, error)
	FindByRefreshTokenAndIP(RefreshToken string, IP string) ([]*test.RepositoryTest, error)
	FindByUserIDAndIP(UserID int, IP string) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndRefreshTokenAndUserID(AccessToken string, RefreshToken string, UserID int) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndRefreshTokenAndIP(AccessToken string, RefreshToken string, IP string) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndUserIDAndIP(AccessToken string, UserID int, IP string) ([]*test.RepositoryTest, error)
	FindByRefreshTokenAndUserIDAndIP(RefreshToken string, UserID int, IP string) ([]*test.RepositoryTest, error)
	FindByAccessTokenAndRefreshTokenAndUserIDAndIP(AccessToken string, RefreshToken string, UserID int, IP string) ([]*test.RepositoryTest, error)
	FindByUserName(Name string) ([]*test.RepositoryTest, error)
	FindByUserEmail(Email string) ([]*test.RepositoryTest, error)
	FindByUserPassword(Password string) ([]*test.RepositoryTest, error)
	FindByUserNameAndEmail(Name string, Email string) ([]*test.RepositoryTest, error)
	FindByUserNameAndPassword(Name string, Password string) ([]*test.RepositoryTest, error)
	FindByUserEmailAndPassword(Email string, Password string) ([]*test.RepositoryTest, error)
	FindByUserNameAndEmailAndPassword(Name string, Email string, Password string) ([]*test.RepositoryTest, error)
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

func (r *RepositoryTestRepositoryImpl) FindByAccessToken(AccessToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("access_token = ?", AccessToken).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) DeleteByAccessToken(AccessToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("access_token = ?", AccessToken).Delete(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByRefreshToken(RefreshToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("refresh_token = ?", RefreshToken).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) DeleteByRefreshToken(RefreshToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("refresh_token = ?", RefreshToken).Delete(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
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

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndRefreshToken(AccessToken string, RefreshToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("access_token = ? AND refresh_token = ?", AccessToken, RefreshToken).First(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) DeleteByAccessTokenAndRefreshToken(AccessToken string, RefreshToken string) (*test.RepositoryTest, error) {
	var r1 test.RepositoryTest

	if err := r.db.Where("access_token = ? AND refresh_token = ?", AccessToken, RefreshToken).Delete(&r1).Error; err != nil {
		return nil, err
	}

	return &r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndUserID(AccessToken string, UserID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND user_id = ?", AccessToken, UserID).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndIP(AccessToken string, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND ip = ?", AccessToken, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByRefreshTokenAndUserID(RefreshToken string, UserID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("refresh_token = ? AND user_id = ?", RefreshToken, UserID).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByRefreshTokenAndIP(RefreshToken string, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("refresh_token = ? AND ip = ?", RefreshToken, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByUserIDAndIP(UserID int, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("user_id = ? AND ip = ?", UserID, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndRefreshTokenAndUserID(AccessToken string, RefreshToken string, UserID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND refresh_token = ? AND user_id = ?", AccessToken, RefreshToken, UserID).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndRefreshTokenAndIP(AccessToken string, RefreshToken string, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND refresh_token = ? AND ip = ?", AccessToken, RefreshToken, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndUserIDAndIP(AccessToken string, UserID int, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND user_id = ? AND ip = ?", AccessToken, UserID, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByRefreshTokenAndUserIDAndIP(RefreshToken string, UserID int, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("refresh_token = ? AND user_id = ? AND ip = ?", RefreshToken, UserID, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByAccessTokenAndRefreshTokenAndUserIDAndIP(AccessToken string, RefreshToken string, UserID int, IP string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("access_token = ? AND refresh_token = ? AND user_id = ? AND ip = ?", AccessToken, RefreshToken, UserID, IP).Find(&r1).Error; err != nil {
		return nil, err
	}

	return r1, nil
}

func (r *RepositoryTestRepositoryImpl) FindByUserName(Name string) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Table("repository_tests").
		Select("repository_tests.*").
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
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
		Joins("JOIN users ON users.id = repository_tests.user").
		Where("users.name = ? AND users.email = ? AND users.password = ?", Name, Email, Password).
		Find(&r1).Error; err != nil {
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
	if err := r.db.Delete(&test.RepositoryTest{}, id).Error; err != nil {
		return err
	}

	return nil
}
