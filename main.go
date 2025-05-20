// @title Student API
// @version 1.0
// @description This is a REST API for managing students.
// @host     musical-waffle-9r74v75w54qhqgg-8080.app.github.dev
// @BasePath /
// @schemes         https
// @contact.name Vaiibhavv
// @contact.email your-email@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

package main

import (
	"log"
	"net/http"
	"os"
	"students_api/db"
	"students_api/handlers"

	_ "students_api/docs" // 👉 Required for Swaggo to find docs folder

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// ✅ Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ✅ Connect to database to ensure it works
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	log.Println("Connected to the database successfully!")

	// ✅ Define routes
	router := mux.NewRouter()

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// API routes
	router.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	router.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
	router.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

	// ✅ Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s/", port)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*.app.github.dev"}, // Allow GitHub Codespace domains
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
