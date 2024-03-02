package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dmitrymomot/go-app-template/web/templates/views"
)

// NotFoundHandler is a handler for 404 Not Found
func NotFoundHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := errors.New("Page not found")
		if isJsonRequest(r) {
			err = errors.New("Endpoint not found")
		}
		sendErrorResponse(
			w, r,
			http.StatusNotFound,
			err,
		)
	}
}

// MethodNotAllowedHandler is a handler for 405 Method Not Allowed
func MethodNotAllowedHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sendErrorResponse(
			w, r,
			http.StatusMethodNotAllowed,
			errors.New(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}
}

// Predefined http encoder content type
const (
	contentTypeHeader  = "Content-Type"
	contextTypeCharset = "charset=utf-8"
	contentTypeJSON    = "application/json"
	contentTypeHTML    = "text/html"
	contentTypeJSONUTF = contentTypeJSON + "; " + contextTypeCharset
	contentTypeHTMLUTF = contentTypeHTML + "; " + contextTypeCharset
)

// Helper function to check if an error code is valid
func isValidErrorCode(errCode int) bool {
	return errCode >= 400 && errCode < 600
}

// Is request a json request?
func isJsonRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get(contentTypeHeader)), contentTypeJSON)
}

// Helper function to send an error response
func sendErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	if !isValidErrorCode(statusCode) {
		statusCode = http.StatusInternalServerError
	}

	if isJsonRequest(r) {
		w.Header().Set(contentTypeHeader, contentTypeJSONUTF)
		w.WriteHeader(statusCode)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set(contentTypeHeader, contentTypeHTMLUTF)
	w.WriteHeader(statusCode)
	if err := views.ErrorPage(statusCode, err.Error()).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
