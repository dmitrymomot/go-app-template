package handlers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/a-h/templ"
	"github.com/dmitrymomot/binder"
	authsvc "github.com/dmitrymomot/go-app-template/internal/auth"
	"github.com/dmitrymomot/go-app-template/pkg/validator"
	"github.com/dmitrymomot/go-app-template/web/templates/views/auth"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// NewHTTPHandler creates a new HTTP handler for the auth service.
// It takes a pointer to a Service struct as a parameter and returns an http.Handler.
// The returned handler is responsible for handling HTTP requests related to the auth service.
func NewHTTPHandler(s *authsvc.Service, log *zap.SugaredLogger) http.Handler {
	r := chi.NewRouter()

	r.Get("/signup", templ.Handler(auth.SignupPage()).ServeHTTP)
	r.Get("/login", templ.Handler(auth.LoginPage()).ServeHTTP)
	r.Get("/forgot-password", templ.Handler(auth.ForgotPasswordPage()).ServeHTTP)
	r.Get("/reset-password", templ.Handler(auth.ResetPasswordPage()).ServeHTTP)

	r.Post("/signup", signupHandler(s, log))
	r.Post("/login", loginHandler(s, log))
	r.Post("/forgot-password", forgotPasswordHandler(s, log))
	r.Post("/reset-password", resetPasswordHandler(s, log))

	return r
}

type signupRequest struct {
	Email    string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" label:"Email address"`
	Password string `form:"password" validate:"required|password" label:"Password"`
}

// signupHandler is an HTTP handler for the signup endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the signup endpoint.
func signupHandler(_ *authsvc.Service, log *zap.SugaredLogger) http.HandlerFunc {
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
			}), log)
			return
		}

		// ...

		// Respond to the client.
		render(w, r, auth.SignupForm(auth.SignupFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}), log)
	}
}

type loginRequest struct {
	Email    string `form:"email" validate:"required|email" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
	Password string `form:"password" validate:"required" message:"Password is required" label:"Password"`
}

// loginHandler is an HTTP handler for the login endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the login endpoint.
func loginHandler(_ *authsvc.Service, log *zap.SugaredLogger) http.HandlerFunc {
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
			}), log)
			return
		}

		// Redirect to the home page.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

type forgotPasswordRequest struct {
	Email string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
}

// forgotPasswordHandler is an HTTP handler for the forgot-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the forgot-password endpoint.
func forgotPasswordHandler(_ *authsvc.Service, log *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		// Parse the form.
		req := &forgotPasswordRequest{}
		if err := binder.BindForm(r, req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Validate the request.
		if verr := validator.ValidateStruct(req); len(verr) > 0 {
			// Respond to the client.
			render(w, r, auth.ForgotPasswordForm(auth.ForgotPasswordFormPayload{
				Form:   r.Form,
				Errors: verr,
			}), log)
			return
		}

		// Respond to the client.
		render(w, r, auth.PopupNotification(auth.PopupPayload{
			Type:        auth.PopupSuccess,
			Title:       "Success!",
			Message:     "Password reset instructions have been sent to your email address. Please check your email. If you don't receive an email, please try again.",
			ActionURL:   "/auth/forgot-password",
			ActionLabel: "Try again",
		}), log)
	}
}

type resetPasswordRequest struct {
	Token           string `form:"token" validate:"required" message:"Token is required" label:"Token"`
	Password        string `form:"password" validate:"required|password" label:"Password"`
	PasswordConfirm string `form:"password_confirmation" validate:"required|eqField:Password" message:"Passwords do not match" label:"Password confirmation"`
}

// resetPasswordHandler is an HTTP handler for the reset-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the reset-password endpoint.
func resetPasswordHandler(_ *authsvc.Service, log *zap.SugaredLogger) http.HandlerFunc {
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
			}), log)
			return
		}

		// Respond to the client.
		render(w, r, auth.ResetPasswordForm(auth.ResetPasswordFormPayload{
			Form:   r.Form,
			Errors: url.Values{},
		}), log)
	}
}

func render(w http.ResponseWriter, r *http.Request, view templ.Component, log *zap.SugaredLogger) {
	if err := view.Render(r.Context(), w); err != nil {
		log.Errorw("Failed to render view", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
