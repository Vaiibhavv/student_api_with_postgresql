//

package main

import (
	"log"
	"net/http"
	"os"
	"students_api/db"       // ✅ Import your DB connection package
	"students_api/handlers" // ✅ Adjust the module path if needed

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	log.Fatal(http.ListenAndServe(":"+port, router))
}
