package user

import (
	"fmt"
	"io"
	"strings"

	"github.com/rms-diego/bank-go-rest-api/models"
	pkg "github.com/rms-diego/bank-go-rest-api/pkg/hash"
	"github.com/rms-diego/bank-go-rest-api/pkg/serialize"
	"github.com/rms-diego/bank-go-rest-api/pkg/validate"
)

type userService struct{ repo accountRepository }

func newService(repo accountRepository) userService {
	return userService{repo: repo}
}

func (ctx userService) createUser(dataReader io.ReadCloser) (models.User, error) {
	userPayload, err := serialize.BodyToJSON[models.User](dataReader)
	if err != nil {
		return models.User{}, err
	}

	err = validate.ValidatePayload(userPayload)
	if err != nil {
		return models.User{}, err
	}

	hash, err := pkg.HashPassword(userPayload.Password)
	if err != nil {
		return models.User{}, err
	}

	userPayload.Password = hash
	userCreated, err := ctx.repo.createUser(userPayload)

	switch {
	case err != nil && strings.Contains(err.Error(), "violates unique"):
		return models.User{}, fmt.Errorf("user already exists")

	case err != nil && err.Error() != "":
		return models.User{}, err

	default:
		return userCreated, nil
	}
}
