package main

// create context type to avoid name collisions
type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
