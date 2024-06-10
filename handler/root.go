package handler

import (
	"fmt"
	"net/http"
)

// / Handler for /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
