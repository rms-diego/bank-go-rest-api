package auth

import (
	"fmt"
	"io"

	"github.com/rms-diego/bank-go-rest-api/internal/user"
	pkg "github.com/rms-diego/bank-go-rest-api/pkg/hash"
	"github.com/rms-diego/bank-go-rest-api/pkg/jwt"
	"github.com/rms-diego/bank-go-rest-api/pkg/serialize"
)

type authPayload struct {
	Email    string
	Password string
}

type authService struct{ repo user.UserRepository }

func newAuthService(repo user.UserRepository) authService {
	return authService{repo: repo}
}

func (ctx authService) loginService(dataReader io.ReadCloser) (string, error) {

	authPayload, err := serialize.BodyToJSON[authPayload](dataReader)
	if err != nil {
		return "", err
	}

	userFound, nil := ctx.repo.FindByMail(authPayload.Email)

	if err != nil {
		return "", err
	}

	if err := pkg.CheckPasswordHash(authPayload.Password, userFound.Password); err != nil {
		return "", fmt.Errorf("wrong credentials")
	}

	token, err := jwt.CreateToken(userFound)

	if err != nil {
		return "", err
	}

	return token, nil
}
