package repository

import (
	"github.com/Yasinqurni/be-project/src/app/user/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) Create(user *model.User) error {
	err := r.db.Create(user).Error

	return err
}

func (r *userRepositoryImpl) Update(id uint, user *model.User) error {
	err := r.db.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) GetUser(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetAllUser(ids []string) (*[]model.User, error) {

	var user []model.User

	if len(ids) != 0 {
		err := r.db.Where("id IN ?", ids).Find(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Delete(id uint) error {
	var user model.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) ListUser() (*[]model.User, error) {
	var user []model.User
	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
