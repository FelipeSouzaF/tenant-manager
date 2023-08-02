package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection initializes and returns the database connection.
func GetConnection() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:kodejifr@tcp(172.17.0.1:3306)/multiplier_central")
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    fmt.Println("Connected to MySQL!")
    return db, nil
}
