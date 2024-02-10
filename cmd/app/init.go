package main

import (
	"encoding/gob"
	"net/url"

	_ "github.com/joho/godotenv/autoload"                // Load .env file automatically
	_ "github.com/tursodatabase/libsql-client-go/libsql" // init libsql driver (SQLite fork)
)

func init() {
	// Register the types for gob
	gob.Register(url.Values{})
	gob.Register(map[string]string{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[string][]interface{}{})
	gob.Register(map[interface{}]interface{}{})

	// ... (other init code)
}
