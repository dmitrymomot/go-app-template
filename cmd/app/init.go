package main

import (
	"encoding/gob"
	"net/url"
)

// init is a special function in Go that is automatically called before the main function.
// It is commonly used for initialization tasks such as registering types or setting up global variables.
// In this case, the init function is registering various types for gob serialization.
func init() {
	// Register the types for gob serialization
	gob.Register(url.Values{})
	gob.Register(map[string]string{})
	gob.Register(map[string][]string{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[string][]interface{}{})
	gob.Register(map[interface{}]interface{}{})

	// ... (other init code)
}
