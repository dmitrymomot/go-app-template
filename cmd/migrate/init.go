package main

import (
	_ "github.com/joho/godotenv/autoload"                // Load .env file automatically
	_ "github.com/tursodatabase/libsql-client-go/libsql" // init libsql driver (SQLite fork)
)

func init() {
	// ... (other init code)
}
