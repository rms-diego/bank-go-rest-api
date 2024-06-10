package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rms-diego/bank-go-rest-api/internal/utils/httpResponse"
)

type userHandler struct{ service userService }

func newUserHandler(service userService) userHandler {
	return userHandler{service: service}
}

func (u userHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user")

	if r.Method != http.MethodPost && id == "" {
		return
	}

	userCreated, err := u.service.createUser(r.Body)
	if err != nil {
		httpResponse.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	httpResponse.NewJsonResponse(w, http.StatusNoContent, userCreated)
}

func (u userHandler) findById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		httpResponse.NewErrorResponse(w, http.StatusNotImplemented, fmt.Errorf("not implemented"))
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/user/")
	userFound, err := u.service.findById(id)
	if err != nil {
		httpResponse.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	httpResponse.NewJsonResponse(w, 200, userFound)
}
