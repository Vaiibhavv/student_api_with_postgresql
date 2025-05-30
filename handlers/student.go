package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"students_api/db"
	"students_api/models"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Add a new student record to the database
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student Data"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {string} string "Internal Server Error"
// @Router /students [post]
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	// validate the json response
	validate := validator.New()
	err = validate.Struct(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, _ := db.Connect()
	defer conn.Close()

	err = conn.QueryRow("INSERT INTO students (name, age, grade) VALUES ($1, $2, $3) RETURNING id",
		student.Name, student.Age, student.Grade).Scan(&student.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := map[string]interface{}{
		"message": "Student created successfully",
		"student": student,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetStudents godoc
// @Summary Get all students
// @Description Retrieve all student records. You can filter by name, age, and grade.
// @Tags students
// @Produce json
// @Param name query string false "Filter by student name (optional)"
// @Param age query int false "Filter by student age (optional)"
// @Param grade query string false "Filter by student grade (optional)"
// @Success 200 {array} models.Student
// @Failure 500 {object} map[string]string
// @Router /students [get]
func GetStudents(w http.ResponseWriter, r *http.Request) {
	conn, _ := db.Connect()
	defer conn.Close()
	//---- Get query params ----
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")
	grade := queryParams.Get("grade")

	// --- base query
	query := "select id, name, age, grade from students where 1=1"
	args := []interface{}{}
	argsIndex := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argsIndex)
		args = append(args, "%"+name+"%")
		argsIndex++
	}
	if age != "" {
		query += fmt.Sprintf(" AND age=$%d", argsIndex)
		args = append(args, age)
		argsIndex++
	}
	if grade != "" {
		query += fmt.Sprintf(" AND grade=$%d", argsIndex)
		args = append(args, grade)
		argsIndex++
	}

	rows, err := conn.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade)
		students = append(students, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// GetStudent godoc
// @Summary Get a student by ID
// @Description Get details of a specific student
// @Tags students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 404 {string} string "Student not found"
// @Router /students/{id} [get]
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

// UpdateStudent godoc
// @Summary Update a student
// @Description Update student information by ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Updated Student"
// @Success 200 {object} map[string]string
// @Failure 500 {string} string "Update failed"
// @Router /students/{id} [put]
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "invalid response", http.StatusBadRequest)
	}
	//validate the json
	validate := validator.New()
	err = validate.Struct(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, _ := db.Connect()
	defer conn.Close()

	_, err = conn.Exec("UPDATE students SET name = $1, age = $2, grade = $3 WHERE id = $4",
		s.Name, s.Age, s.Grade, id)

	if err != nil {
		http.Error(w, "Update failed", 500)
		return
	}

	response := map[string]interface{}{
		"message": "Student Updated",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	w.WriteHeader(http.StatusOK)
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete a student by ID
// @Tags students
// @Param id path int true "Student ID"
// @Success 204 "No Content"
// @Failure 500 {string} string "Delete failed"
// @Router /students/{id} [delete]
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
