package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// serveFile is a helper function to serve a single file and handle caching.
func serveFile(w http.ResponseWriter, r *http.Request, file http.File, info os.FileInfo, cacheTTL time.Duration) {
	// Generate ETag using file info
	etag := fmt.Sprintf(`"%x-%x"`, info.ModTime().Unix(), info.Size())
	lastModified := info.ModTime().UTC().Format(http.TimeFormat)

	// Set headers for caching
	w.Header().Set("ETag", etag)
	w.Header().Set("Last-Modified", lastModified)
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int(cacheTTL.Seconds())))
	w.Header().Set("Expires", time.Now().Add(cacheTTL).UTC().Format(http.TimeFormat))
	w.Header().Set("Pragma", "cache")

	// Check if file hasn't been modified since the last request
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// Check if file has been modified since the last request based on Last-Modified header
	ifModifiedSince := r.Header.Get("If-Modified-Since")
	if ifModifiedSince != "" {
		if t, err := time.Parse(http.TimeFormat, ifModifiedSince); err == nil && info.ModTime().Before(t.Add(1*time.Second)) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// Serve the file
	http.ServeContent(w, r, info.Name(), info.ModTime(), file)
}

// fileServer sets up a http.FileServer handler to serve static files from a no directory listing file system.
func fileServer(r chi.Router, publicPath string, root http.FileSystem, cacheTTL time.Duration) {
	if strings.ContainsAny(publicPath, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	publicPath = strings.TrimRight(publicPath, "/")
	r.HandleFunc(publicPath+"/*", func(w http.ResponseWriter, r *http.Request) {
		fsPath := strings.TrimPrefix(r.URL.Path, publicPath)
		file, err := root.Open(fsPath)
		if err != nil {
			// File not found
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			// Error getting file info
			http.NotFound(w, r)
			return
		}

		if info.IsDir() {
			// Path is a directory, return 404
			http.NotFound(w, r)
			return
		}

		// Serve file with caching
		serveFile(w, r, file, info, cacheTTL)
	})
}
