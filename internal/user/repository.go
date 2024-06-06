package user

import (
	"fmt"

	"github.com/rms-diego/bank-go-rest-api/models"
	"github.com/rms-diego/bank-go-rest-api/pkg/database"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ctx UserRepository) CreateUser(user models.User) (models.User, error) {
	var userCreated models.User

	query := `
		INSERT INTO
			users(name, last_name, email, password, tax_id, birth_date)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id, name, last_name, email,  password, tax_id, birth_date	
	`

	err := database.Db.QueryRow(
		query,
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		user.TaxId,
		user.BirthDate,
	).Scan(
		&userCreated.Id,
		&userCreated.Name,
		&userCreated.LastName,
		&userCreated.Email,
		&userCreated.Password,
		&userCreated.TaxId,
		&userCreated.BirthDate,
	)

	if err != nil {
		return models.User{}, err
	}

	return userCreated, nil
}

func (ctx UserRepository) FindByMail(email string) (models.User, error) {
	var userFound models.User

	query := `
		SELECT 
			id, name, last_name AS lastName, email, password, tax_id AS taxId, birth_date AS birthDate
		FROM 
			users
		WHERE email = $1
	`

	err := database.Db.QueryRow(query, email).
		Scan(
			&userFound.Id,
			&userFound.Name,
			&userFound.LastName,
			&userFound.Email,
			&userFound.Password,
			&userFound.TaxId,
			&userFound.BirthDate,
		)

	if err != nil {
		fmt.Println(err.Error())
		return models.User{}, fmt.Errorf("user not found")
	}

	return userFound, nil
}
