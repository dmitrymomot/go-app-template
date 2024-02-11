package main

import (
	"encoding/gob"
	"net/url"
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
