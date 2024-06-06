package user

import "net/http"

func Routes(app *http.ServeMux) {
	app.HandleFunc("/account", createUser)
}
