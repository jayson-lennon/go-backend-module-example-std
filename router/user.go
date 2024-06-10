package router

import (
	"net/http"
	"sample/handler/user"
)

// All routes related to user go here
func ApplyUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user", user.Root)
	mux.HandleFunc("/user/login", user.Login)
}
