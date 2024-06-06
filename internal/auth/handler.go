package auth

import (
	"fmt"
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/pkg/httpResponse"
)

type authHandler struct {
	authService authService
}

func newAuthHandler(service authService) authHandler {
	return authHandler{authService: service}
}

func (ctx authHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpResponse.NewErrorResponse(w, http.StatusNotImplemented, fmt.Errorf("route not found"))
		return
	}

	token, err := ctx.authService.loginService(r.Body)

	if err != nil {
		httpResponse.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	type res struct {
		Token string `json:"token"`
	}

	httpResponse.NewJsonResponse(w, http.StatusOK, res{Token: token})
}
