package user

import (
	"fmt"
	"io"
	"strings"

	"github.com/rms-diego/bank-go-rest-api/internal/utils/bcrypt"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/serialize"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/validate"
	"github.com/rms-diego/bank-go-rest-api/models"
)

type userService struct{ repo UserRepository }

func newService(repo UserRepository) userService {
	return userService{repo: repo}
}

func (u userService) createUser(dataReader io.ReadCloser) (models.User, error) {
	userPayload, err := serialize.BodyToJSON[models.User](dataReader)
	if err != nil {
		return models.User{}, err
	}

	err = validate.ValidatePayload(userPayload)
	if err != nil {
		return models.User{}, err
	}

	userPayload.Password, err = bcrypt.HashPassword(userPayload.Password)
	if err != nil {
		return models.User{}, err
	}

	userCreated, err := u.repo.CreateUser(userPayload)

	switch {
	case err != nil && strings.Contains(err.Error(), "violates unique"):
		return models.User{}, fmt.Errorf("user already exists")

	case err != nil && err.Error() != "":
		return models.User{}, err

	default:
		return userCreated, nil
	}
}

func (u userService) findById(id string) (models.User, error) {

	if id == "" {
		return models.User{}, fmt.Errorf("missing user id")
	}
	userFound, err := u.repo.findById(id)

	if err != nil {
		return models.User{}, err
	}

	return userFound, nil
}
