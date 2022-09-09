package services

import (
	"dataimpact/test/golang/models"

)

type UserService interface {
	CreateUser(*[]interface{}) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	LoginUser(string, string) (error, error)
}
