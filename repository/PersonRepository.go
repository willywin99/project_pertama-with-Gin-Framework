package repository

import (
	"project_pertama/model"

	"gorm.io/gorm"
)

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository {
	return &personRepository{
		db: db,
	}
}

func (pr *personRepository) Create(newPerson model.Person) (model.Person, error) {
	tx := pr.db.Create(&newPerson)

	return newPerson, tx.Error
}

func (pr *personRepository) GetAll() ([]model.Person, error) {
	persons := []model.Person{}

	// tx := pr.db.Unscoped().Find(&persons)
	// tx := pr.db.Find(&persons)
	// tx := pr.db.Preload("Cards").Find(&persons)
	tx := pr.db.Find(&persons)

	return persons, tx.Error
}

// func (pr *personRepository) Update(id int, updatedPerson model.Person) (model.Person, error) {
// 	query := "UPDATE person SET name=$1, address=$2 WHERE id=$3 returning *"

// 	// result, err := pr.db.Exec(query, updatedPerson.Name, updatedPerson.Address, id)
// 	// if err != nil {
// 	// 	return updatedPerson, err
// 	// }

// 	// count, err := result.RowsAffected()
// 	// if err != nil {
// 	// 	return updatedPerson, err
// 	// }
// 	// fmt.Println(count)

// 	row := pr.db.QueryRow(query, updatedPerson.Name, updatedPerson.Address, id)
// 	err := row.Scan(&updatedPerson.Id, &updatedPerson.Name, &updatedPerson.Address, &updatedPerson.Status)

// 	// return updatedPerson, nil
// 	return updatedPerson, err
// }

// func (pr *personRepository) Delete(id int) (model.Person, error) {
func (pr *personRepository) Delete(uuid string) error {
	// var deletedPerson model.Person

	// // query := `
	// // 	DELETE FROM person
	// // 	WHERE id=$1
	// // 	RETURNING *
	// // `

	// query := `
	// 	UPDATE PERSON
	// 	SET status = false
	// 	WHERE id=$1
	// 	RETURNING *
	// `

	// row := pr.db.QueryRow(query, id)
	// err := row.Scan(&deletedPerson.Id, &deletedPerson.Name, &deletedPerson.Address, &deletedPerson.Status)
	// fmt.Println(deletedPerson.Name, deletedPerson.Address, &deletedPerson.Status)
	// if err != nil {
	// 	return deletedPerson, err
	// }

	// return deletedPerson, nil

	// tx := pr.db.Unscoped().Delete(&model.Person{}, "uuid = ?", uuid)
	tx := pr.db.Delete(&model.Person{}, "uuid = ?", uuid)
	return tx.Error
}
