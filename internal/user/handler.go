package user

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/pkg/httpResponse"
)

type userHandler struct{ service userService }

func newUserHandler(service userService) userHandler {
	return userHandler{service: service}
}

func (ctx userHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	userCreated, err := ctx.service.createUser(r.Body)

	if err != nil {
		httpResponse.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	httpResponse.NewJsonResponse(w, http.StatusNoContent, userCreated)
}
