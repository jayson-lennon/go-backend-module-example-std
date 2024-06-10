package user

import (
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/user")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login user")
}
