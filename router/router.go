package router

import (
	"net/http"
	"sample/handler"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// All user routes encapsulated behind this function.
	ApplyUserRoutes(mux)

	// Can also add one-off routes here
	mux.HandleFunc("/", handler.RootHandler)
	return mux
}
