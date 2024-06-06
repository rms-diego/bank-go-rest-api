package auth

import (
	"net/http"

	"github.com/rms-diego/bank-go-rest-api/internal/user"
)

func Routes(app *http.ServeMux) {
	userRepository := user.NewUserRepository()
	authService := newAuthService(userRepository)
	authHandler := newAuthHandler(authService)

	app.HandleFunc("/auth", authHandler.loginHandler)

}
