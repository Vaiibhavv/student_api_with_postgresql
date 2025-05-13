package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "students_api/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
    router.HandleFunc("/students", handlers.GetStudents).Methods("GET")
    router.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
    router.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
    router.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

    port := os.Getenv("SERVER_PORT")
    log.Printf("Server running at http://localhost:%s/", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}