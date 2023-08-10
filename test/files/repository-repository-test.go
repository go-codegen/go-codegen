package repository

import "gorm.io/gorm"

type RepositoryTestRepository struct {
	db *gorm.DB
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
func (r *RepositoryTestRepository) FindByHelloID(HelloID int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("hello_id = ?", HelloID).Find(&r1).Error; err != nil {
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
func (r *RepositoryTestRepository) FindByAge(Age int) ([]*test.RepositoryTest, error) {
	var r1 []*test.RepositoryTest

	if err := r.db.Where("age = ?", Age).Find(&r1).Error; err != nil {
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
	if err := r.db.Delete(&test.RepositoryTest{},"id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
