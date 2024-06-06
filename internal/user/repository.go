package user

import (
	"github.com/rms-diego/bank-go-rest-api/models"
	"github.com/rms-diego/bank-go-rest-api/pkg/database"
)

type accountRepository struct{}

func newRepository() accountRepository {
	return accountRepository{}
}

func (ctx accountRepository) createUser(user models.User) (models.User, error) {
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
