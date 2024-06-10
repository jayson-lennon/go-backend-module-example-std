package server

import (
	"fmt"
	"net/http"
)

func StartServer(handler http.Handler) {
	fmt.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
