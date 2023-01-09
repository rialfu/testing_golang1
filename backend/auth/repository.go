package auth

import (
	"errors"
	"rema/kredit/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindUser(username string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	UpdateUser(update map[string]interface{}, where map[string]interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindUser(username string) (model.User, error) {
	var User model.User
	// loc, _ := time.LoadLocation("GMT")
	if err := r.db.Where("username = ?", username).First(&User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("not found")
		}
		return model.User{}, err
	}
	return User, nil
}
func (r *repository) UpdateUser(update map[string]interface{}, where map[string]interface{}) error{
	if err := r.db.Model(&model.User{}).Where(where).Updates(update).Error; err != nil {
		return err
	}
	return nil
}
// 	User.LastLogin = User.LastLogin.In(loc)
// 	return User, nil
// }
func (r *repository) AddUser(user model.User) (model.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}
	return user, nil
}

