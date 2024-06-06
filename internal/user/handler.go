package user

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/pkg/httpResponse"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	repo := newRepository()
	service := newService(repo)

	userCreated, err := service.createUser(r.Body)

	if err != nil {
		httpResponse.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	httpResponse.NewJsonResponse(w, http.StatusOK, userCreated)
}
