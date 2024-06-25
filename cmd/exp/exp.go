package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// PostgresConfig is a struct that holds the configuration for a Postgres database.
// This can be used to connect to the database.
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// String returns the connection string for the PostgresConfig.
// This can be used to connect to the database.
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

// main is the entry point for the application.
func main() {
	// Create a new PostgresConfig struct and set the fields.
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	// Open a new database connection using the pgx driver.
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}

	// Defer the closing of the database connection.
	defer db.Close()

	// Check if the connection to the database is successful.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new table in the database.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created successfully!")

	// Insert some data into the users table.
	name := "Johnny Halliday"
	email := "johnny@hallyday.com"
	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id;
	`, name, email)
	var id int
	err = row.Scan(&id) // Scan the id of the inserted row. &id is a pointer to the id variable. It is used to store the value of the id column in the database.
	if err != nil {
		panic(err)
	}
	fmt.Println("User created. id = ", id)
}
