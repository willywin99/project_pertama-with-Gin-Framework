package repository

import (
	"project_pertama/model"
)

type IOrderRepository interface {
	GetAll() ([]model.Order, error)
	GetAllByUserId(userId string) ([]model.Order, error)
	Create(model.Order) (model.Order, error)
	Delete(uuid string) error
}
