package auth

import (
	"errors"
	"net/http"
	"rema/kredit/model"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateAccount(data RegisterRequest) (model.User, int, error)
	Login(data LoginRequest) (model.User, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
func (s *service) Login(data LoginRequest) (model.User, int, error) {
	username := data.Username
	User, err := s.repo.FindUser(username)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "not found" {
			status = http.StatusUnauthorized
			err =  errors.New("username or password is wrong")
		}
		return model.User{}, status, err
	}
	

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errors.New("Password is wrong")
	}
	
	User.Password = ""
	return User, http.StatusOK, nil
}
func (s *service) CreateAccount(data RegisterRequest) (model.User, int, error) {
	found, err := s.repo.FindUser(data.Username)
	if err == nil && found.Name != "" {
		return model.User{}, http.StatusBadRequest, errors.New("duplicate data")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User := model.User{
		Username: data.Username,
		Password: string(passwordHash),
		Name:     data.Name,
	}
	res, err := s.repo.AddUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
func (s *service) UpdatePassword(oldPass string, newPass string, username string)(int, error) {
	User, err := s.repo.FindUser(username)
	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(oldPass))
	if err != nil {
		return 401, errors.New("password salah")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	update := map[string]interface{}{
		"password":string(passwordHash),
	}
	where := map[string]interface{}{
		"username":username,
	}
	err = s.repo.UpdateUser(update, where)
	if err != nil {
		return 500, err
	}
	return 200, nil

}