package repository

import (
	"database/sql"
	"fmt"
	"project_pertama/model"
)

type personRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *personRepository {
	return &personRepository{
		db: db,
	}
}

func (pr *personRepository) Create(newPerson model.Person) (model.Person, error) {
	query := "insert into person(name, address, status) values($1, $2, true) returning *"

	row := pr.db.QueryRow(query, newPerson.Name, newPerson.Address)
	err := row.Scan(&newPerson.Id, &newPerson.Name, &newPerson.Address, &newPerson.Status)

	return newPerson, err
}

func (pr *personRepository) GetAll() ([]model.Person, error) {
	persons := []model.Person{}

	query := "select * from person"
	rows, err := pr.db.Query(query)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		var p model.Person
		err := rows.Scan(&p.Id, &p.Name, &p.Address, &p.Status)
		// fmt.Println(err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		persons = append(persons, p)
	}

	return persons, nil
}

func (pr *personRepository) Update(id int, updatedPerson model.Person) (model.Person, error) {
	// query := "UPDATE person SET name=$1, city=$2, WHERE id=$3 returning *"
	query := "UPDATE person SET name=$1, address=$2 WHERE id=$3 returning *"

	// result, err := pr.db.Exec(query, updatedPerson.Name, updatedPerson.Address, id)
	// if err != nil {
	// 	return updatedPerson, err
	// }

	// count, err := result.RowsAffected()
	// if err != nil {
	// 	return updatedPerson, err
	// }
	// fmt.Println(count)

	row := pr.db.QueryRow(query, updatedPerson.Name, updatedPerson.Address, id)
	err := row.Scan(&updatedPerson.Id, &updatedPerson.Name, &updatedPerson.Address, &updatedPerson.Status)

	// return updatedPerson, nil
	return updatedPerson, err
}

func (pr *personRepository) Delete(id int) (model.Person, error) {
	var deletedPerson model.Person

	// query := `
	// 	DELETE FROM person
	// 	WHERE id=$1
	// 	RETURNING *
	// `

	query := `
		UPDATE PERSON
		SET status = false
		WHERE id=$1
		RETURNING *
	`

	row := pr.db.QueryRow(query, id)
	err := row.Scan(&deletedPerson.Id, &deletedPerson.Name, &deletedPerson.Address, &deletedPerson.Status)
	fmt.Println(deletedPerson.Name, deletedPerson.Address, &deletedPerson.Status)
	if err != nil {
		return deletedPerson, err
	}

	return deletedPerson, nil
}
