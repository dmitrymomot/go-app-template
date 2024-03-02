package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/dmitrymomot/binder"
	authsvc "github.com/dmitrymomot/go-app-template/internal/auth"
	"github.com/dmitrymomot/go-app-template/pkg/validator"
	"github.com/dmitrymomot/go-app-template/web/templates/views/auth"
	"github.com/dmitrymomot/random"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// NewHTTPHandler creates a new HTTP handler for the auth service.
// It takes a pointer to a Service struct as a parameter and returns an http.Handler.
// The returned handler is responsible for handling HTTP requests related to the auth service.
func NewHTTPHandler(
	log *zap.SugaredLogger,
	gAuth *oauth2.Config,
	es *authsvc.EmailService,
	gs *authsvc.GoogleService,
	sm *scs.SessionManager,
) http.Handler {
	r := chi.NewRouter()

	r.Get("/signup", templ.Handler(auth.SignupPage()).ServeHTTP)
	r.Get("/login", templ.Handler(auth.LoginPage()).ServeHTTP)
	r.Get("/forgot-password", templ.Handler(auth.ForgotPasswordPage()).ServeHTTP)
	r.Get("/reset-password", templ.Handler(auth.ResetPasswordPage()).ServeHTTP)
	r.Get("/confirm-email", templ.Handler(auth.ResetPasswordPage()).ServeHTTP)

	r.Post("/signup", signupHandler(es, log))
	r.Post("/login", loginHandler(es, log))
	r.Post("/forgot-password", forgotPasswordHandler(es, log))
	r.Post("/reset-password", resetPasswordHandler(es, log))

	// Google OAuth2
	r.Get("/login/google", googleLoginHandler(gAuth, sm))
	r.Get("/login/google/callback", googleLoginCallbackHandler(gAuth, gs, sm, log))

	return r
}

// render renders the given view to the HTTP response writer.
// It logs any errors that occur during rendering and sends an internal server error response if necessary.
func render(w http.ResponseWriter, r *http.Request, view templ.Component, log *zap.SugaredLogger) {
	if err := view.Render(r.Context(), w); err != nil {
		log.Errorw("Failed to render view", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// signupHandler is an HTTP handler for the signup endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the signup endpoint.
func signupHandler(_ *authsvc.EmailService, log *zap.SugaredLogger) http.HandlerFunc {
	type requestPayload struct {
		Email    string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" label:"Email address"`
		Password string `form:"password" validate:"required|password" label:"Password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &requestPayload{}
		if err := binder.BindForm(r, req); err == nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			sendErrorResponse(w, r, http.StatusInternalServerError, errors.New("test error")) // TODO: Fix this
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

// loginHandler is an HTTP handler for the login endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the login endpoint.
func loginHandler(_ *authsvc.EmailService, log *zap.SugaredLogger) http.HandlerFunc {
	type requestPayload struct {
		Email    string `form:"email" validate:"required|email" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
		Password string `form:"password" validate:"required" message:"Password is required" label:"Password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &requestPayload{}
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

// forgotPasswordHandler is an HTTP handler for the forgot-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the forgot-password endpoint.
func forgotPasswordHandler(_ *authsvc.EmailService, log *zap.SugaredLogger) http.HandlerFunc {
	type requestPayload struct {
		Email string `form:"email" validate:"required|email|realEmail" filter:"sanitizeEmail" message:"Email is invalid" label:"Email address"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		// Parse the form.
		req := &requestPayload{}
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

// resetPasswordHandler is an HTTP handler for the reset-password endpoint.
// It takes a pointer to a Service struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the reset-password endpoint.
func resetPasswordHandler(_ *authsvc.EmailService, log *zap.SugaredLogger) http.HandlerFunc {
	type requestPayload struct {
		Token           string `form:"token" validate:"required" message:"Token is required" label:"Token"`
		Password        string `form:"password" validate:"required|password" label:"Password"`
		PasswordConfirm string `form:"password_confirmation" validate:"required|eqField:Password" message:"Passwords do not match" label:"Password confirmation"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		req := &requestPayload{}
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

// googleAuthStateKey is the key used to store the Google OAuth2 auth state in the session.
const googleAuthStateKey = "google_auth_state"

// googleLoginHandler is an HTTP handler for the login/google endpoint.
// It takes a pointer to a oauth2.Config struct as a parameter and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the login/google endpoint.
func googleLoginHandler(c *oauth2.Config, sm *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := random.String(32)
		// Store the auth code in the session.
		sm.Put(r.Context(), googleAuthStateKey, state)
		// Redirect to the Google OAuth2 login page.
		http.Redirect(w, r, c.AuthCodeURL(state, oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
	}
}

// googleLoginCallbackHandler is an HTTP handler for the login/google/callback endpoint.
// It takes a pointer to a oauth2.Config struct and a pointer to a Service struct as parameters and returns an http.HandlerFunc.
// The returned handler is responsible for handling HTTP requests to the login/google/callback endpoint.
func googleLoginCallbackHandler(c *oauth2.Config, _ *authsvc.GoogleService, sm *scs.SessionManager, log *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form.
		code := r.URL.Query().Get("code")
		if code == "" {
			sendErrorResponse(w, r, http.StatusBadRequest, errors.New("Missing auth code"))
			return
		}

		// Check the state.
		state := r.URL.Query().Get("state")
		if state == "" {
			sendErrorResponse(w, r, http.StatusBadRequest, errors.New("Missing auth state"))
			return
		}

		// Get the auth code from the session.
		authState := sm.GetString(r.Context(), googleAuthStateKey)
		if authState != state {
			log.Errorw("Invalid auth state", "error", "Invalid auth state", "state", state, "authState", authState)
			sendErrorResponse(w, r, http.StatusBadRequest, errors.New("Invalid auth state"))
			return
		}

		// Delete the auth state from the session.
		sm.Remove(r.Context(), googleAuthStateKey)

		// Exchange the code for a token.
		token, err := c.Exchange(r.Context(), code)
		if err != nil {
			log.Errorw("Failed to exchange code for token", "error", err)
			http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
			return
		}

		// Get the user's profile.
		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			log.Errorw("Failed to get user info", "error", err)
			http.Error(w, "Failed to get user info", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		// Read the response body.
		content, err := io.ReadAll(response.Body)
		if err != nil {
			log.Errorw("Failed to read response body", "error", err)
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Response: %s", content)
		log.Infow("Google OAuth2 callback", "response", string(content))
	}
}
