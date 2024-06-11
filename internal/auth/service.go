package auth

import (
	"fmt"
	"io"

	"github.com/rms-diego/bank-go-rest-api/internal/user"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/bcrypt"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/jwt"
	"github.com/rms-diego/bank-go-rest-api/internal/utils/serialize"
)

type authPayload struct {
	Email    string
	Password string
}

type authService struct{ repo user.UserRepository }

func newAuthService(repo user.UserRepository) authService {
	return authService{repo: repo}
}

func (u authService) loginService(dataReader io.ReadCloser) (string, error) {

	authPayload, err := serialize.BodyToJSON[authPayload](dataReader)
	if err != nil {
		return "", err
	}

	userFound, nil := u.repo.FindByMail(authPayload.Email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CheckPasswordHash(authPayload.Password, userFound.Password); err != nil {
		return "", fmt.Errorf("wrong credentials")
	}

	token, err := jwt.CreateToken(userFound)

	if err != nil {
		return "", err
	}

	return token, nil
}
