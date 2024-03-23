package repository

import "project_pertama/model"

type IPersonRepository interface {
	Create(newPerson model.Person) (model.Person, error)
	GetAll() ([]model.Person, error)
	Update(id int, updatedPerson model.Person) (model.Person, error)
	Delete(id int) (model.Person, error)
}
