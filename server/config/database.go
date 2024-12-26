package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create videos table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS videos (
            id VARCHAR(255) PRIMARY KEY,
            title TEXT NOT NULL,
            description TEXT,
            published_at TIMESTAMP NOT NULL,
            thumbnail_url TEXT,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
        CREATE INDEX IF NOT EXISTS idx_published_at ON videos(published_at DESC);
    `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
