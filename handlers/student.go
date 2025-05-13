package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "students_api/db"
    "students_api/models"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    json.NewDecoder(r.Body).Decode(&student)

    conn, _ := db.Connect()
    defer conn.Close()

    err := conn.QueryRow("INSERT INTO students (name, age, grade) VALUES ($1, $2, $3) RETURNING id",
        student.Name, student.Age, student.Grade).Scan(&student.ID)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
    conn, _ := db.Connect()
    defer conn.Close()

    rows, _ := conn.Query("SELECT id, name, age, grade FROM students")
    defer rows.Close()

    var students []models.Student
    for rows.Next() {
        var s models.Student
        rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade)
        students = append(students, s)
    }

    json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    conn, _ := db.Connect()
    defer conn.Close()

    var s models.Student
    err := conn.QueryRow("SELECT id, name, age, grade FROM students WHERE id = $1", id).
        Scan(&s.ID, &s.Name, &s.Age, &s.Grade)

    if err != nil {
        http.Error(w, "Student not found", 404)
        return
    }

    json.NewEncoder(w).Encode(s)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var s models.Student
    json.NewDecoder(r.Body).Decode(&s)

    conn, _ := db.Connect()
    defer conn.Close()

    _, err := conn.Exec("UPDATE students SET name = $1, age = $2, grade = $3 WHERE id = $4",
        s.Name, s.Age, s.Grade, id)

    if err != nil {
        http.Error(w, "Update failed", 500)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    conn, _ := db.Connect()
    defer conn.Close()

    _, err := conn.Exec("DELETE FROM students WHERE id = $1", id)
    if err != nil {
        http.Error(w, "Delete failed", 500)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}