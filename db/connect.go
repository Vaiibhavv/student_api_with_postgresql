// package db

// import (
//     "database/sql"
//     "fmt"
//     "os"

//     _ "github.com/lib/pq"
// )

// func Connect() (*sql.DB, error) {
//     connStr := fmt.Sprintf(
//         "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//         os.Getenv("DB_HOST"),
//         os.Getenv("DB_PORT"),
//         os.Getenv("DB_USER"),
//         os.Getenv("DB_PASSWORD"),
//         os.Getenv("DB_NAME"),
//     )

//     return sql.Open("postgres", connStr)
// }

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect establishes a connection to the PostgreSQL database using environment variables.
func Connect() (*sql.DB, error) {
	// Construct connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Optional: Print for debug
	log.Println("Connecting to PostgreSQL with:", connStr)

	// Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Ping to verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database.")
	return db, nil
}
