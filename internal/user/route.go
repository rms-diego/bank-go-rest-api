package user

import "net/http"

func Routes(app *http.ServeMux) {
	userRepository := newUserRepository()
	userService := newService(userRepository)
	userHandler := newUserHandler(userService)

	app.HandleFunc("/user", userHandler.createUserHandler)
}
