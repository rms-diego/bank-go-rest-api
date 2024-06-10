package user

import (
	"fmt"
	"io"
	"strings"

	"github.com/rms-diego/bank-go-rest-api/internal/utils/jwt"
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

	hash, err := jwt.CreateToken(userPayload)
	if err != nil {
		return models.User{}, err
	}

	userPayload.Password = hash
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
