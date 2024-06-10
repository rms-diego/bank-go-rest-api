package user

import "net/http"

func Routes(app *http.ServeMux) {
	userRepository := NewUserRepository()
	userService := newService(userRepository)
	userHandler := newUserHandler(userService)

	app.HandleFunc("/user", userHandler.createUserHandler)
	app.HandleFunc("/user/", userHandler.findById)
}
