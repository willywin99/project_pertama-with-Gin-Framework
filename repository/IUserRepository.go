package repository

import "project_pertama/model"

type IUserRepository interface {
	Create(model.User) (model.User, error)
	GetByUsername(string) (model.User, error)
}
