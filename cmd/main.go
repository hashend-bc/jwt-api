package main

import (
	"log"
	"net/http"
	"user-auth-api/handlers"
	"user-auth-api/middleware"
)

func main() {

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/protected", middleware.JWTMiddleware(handlers.Protected))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}