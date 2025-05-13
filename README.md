# students_api

This is a simple REST API built using Golang and PostgreSQL. It supports basic CRUD operations for student records.

## Features

- Create a student
- Read student(s)
- Update a student
- Delete a student

## Run in GitHub Codespaces

1. Open the repository in GitHub.
2. Click **Code > Codespaces > Create codespace**.
3. Use the built-in terminal:
   ```bash
   psql -U postgres -d students_db
   ```
   And run:
   ```sql
   CREATE TABLE students (
     id SERIAL PRIMARY KEY,
     name TEXT NOT NULL,
     age INT,
     grade TEXT
   );
   ```
4. Start the server:
   ```bash
   go run main.go
   ```

Access the API at: `http://localhost:8080/students`
