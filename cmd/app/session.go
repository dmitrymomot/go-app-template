package main

import (
	"net/http"

	"github.com/alexedwards/scs/goredisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/redis/go-redis/v9"
)

// Initialize a new session manager and configure the session lifetime.
func initSessionManager(redisClient *redis.Client) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = sessionTTL
	sessionManager.Cookie.Name = sessionName
	sessionManager.Cookie.Secure = appEnv == EnvProduction
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Store = goredisstore.NewWithPrefix(redisClient, sessionPrefix)
	return sessionManager
}
