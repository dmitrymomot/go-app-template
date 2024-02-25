package auth

import (
	"net/http"
	"net/url"
	"time"

	"github.com/a-h/templ"
	"github.com/dmitrymomot/binder"
	"github.com/dmitrymomot/go-app-template/pkg/validator"
	"github.com/dmitrymomot/go-app-template/web/templates/views/auth"
	"github.com/go-chi/chi/v5"
)

// NewHTTPHandler creates a new HTTP handler for the auth service.
// It takes a pointer to a Service struct as a parameter and returns an http.Handler.
// The returned handler is responsible for handling HTTP requests related to the auth service.
func NewHTTPHandler(s *Service) http.Handler {
	r := chi.NewRouter()

	r.Get("/signup", templ.Handler(auth.SignupPage()).ServeHTTP)
	r.Get("/login", templ.Handler(auth.LoginPage()).ServeHTTP)
	r.Get("/forgot-password", templ.Handler(auth.ForgotPasswordPage()).ServeHTTP)
	r.Get("/reset-password", templ.Handler(auth.ResetPasswordPage()).ServeHTTP)

	r.Post("/signup", signupHandler(s))
	r.Post("/login", loginHandler(s))
	r.Post("/forgot-password", forgotPasswordHandler(s))
	r.Post("/reset-password", resetPasswordHandler(s))

	return r
}

type signupRequest struct {
	Email    string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" label:"Email address"`
	Password string `form:"password" validate:"required|password" label:"Password"`
}

// signupHandler is an HTTP handler for the signup endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the signup endpoint.
func signupHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &signupRequest{}
		if err := binder.BindForm(r, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Validate the request.
		if verr := validator.ValidateStruct(req); len(verr) > 0 {
			// Respond to the client.
			render(w, r, auth.SignupForm(auth.SignupFormPayload{
				Form:   r.Form,
				Errors: verr,
			}))
			return
		}

		// ...

		// Respond to the client.
		render(w, r, auth.SignupForm(auth.SignupFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}))
	}
}

type loginRequest struct {
	Email    string `form:"email" validate:"required|email" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
	Password string `form:"password" validate:"required" message:"Password is required" label:"Password"`
}

// loginHandler is an HTTP handler for the login endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the login endpoint.
func loginHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &loginRequest{}
		if err := binder.BindForm(r, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Validate the request.
		if verr := validator.ValidateStruct(req); len(verr) > 0 {
			// Respond to the client.
			render(w, r, auth.LoginForm(auth.LoginFormPayload{
				Form:   r.Form,
				Errors: verr,
			}))
			return
		}

		// Respond to the client.
		render(w, r, auth.LoginForm(auth.LoginFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}))
	}
}

type forgotPasswordRequest struct {
	Email string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
}

// forgotPasswordHandler is an HTTP handler for the forgot-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the forgot-password endpoint.
func forgotPasswordHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		// Parse the form.
		req := &forgotPasswordRequest{}
		if err := binder.BindForm(r, req); err == nil {
			http.Error(w, "test error", http.StatusInternalServerError)
			return
		}

		// Validate the request.
		if verr := validator.ValidateStruct(req); len(verr) > 0 {
			// Respond to the client.
			render(w, r, auth.ForgotPasswordForm(auth.ForgotPasswordFormPayload{
				Form:   r.Form,
				Errors: verr,
			}))
			return
		}

		// Respond to the client.
		render(w, r, auth.ForgotPasswordForm(auth.ForgotPasswordFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}))
	}
}

type resetPasswordRequest struct {
	// Token           string `form:"token" validate:"required" message:"Token is required" label:"Token"`
	Password        string `form:"password" validate:"required|password" label:"Password"`
	PasswordConfirm string `form:"password_confirmation" validate:"required|eqField:Password" message:"Passwords do not match" label:"Password confirmation"`
}

// resetPasswordHandler is an HTTP handler for the reset-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the reset-password endpoint.
func resetPasswordHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &resetPasswordRequest{}
		if err := binder.BindForm(r, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Validate the request.
		if verr := validator.ValidateStruct(req); len(verr) > 0 {
			// Respond to the client.
			render(w, r, auth.ResetPasswordForm(auth.ResetPasswordFormPayload{
				Form:   r.Form,
				Errors: verr,
			}))
			return
		}

		// Respond to the client.
		render(w, r, auth.ResetPasswordForm(auth.ResetPasswordFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}))
	}
}

func render(w http.ResponseWriter, r *http.Request, view templ.Component) {
	if err := view.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
