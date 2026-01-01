package config

import (
    "database/sql"
    "log"
    
    _ "github.com/lib/pq"
)

func ConnectDB(databaseURL string) *sql.DB {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    if err = db.Ping(); err != nil {
        log.Fatal("Database not reachable:", err)
    }
    
    // Set connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    
    log.Println("Database connection established")
    return db
}