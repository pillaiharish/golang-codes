package main

import (
	"fmt"
	"net/http"
)

// Mock session store
var sessionStore = make(map[string]map[string]interface{})

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// In a real application, authentication would be handled here
	sessionID := "abc123" // This would be generated dynamically
	userSession := map[string]interface{}{
		"username": "Harishkumar",
		"role":     "admin",
	}

	sessionStore[sessionID] = userSession
	fmt.Fprintf(w, "Login successful %v", sessionStore[sessionID])
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	fmt.Fprintf(w, "session data: %v", sessionID)
	userSession, exists := sessionStore[sessionID]
	if !exists {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s!", userSession["username"])
}

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/dashboard", handleDashboard)
	http.ListenAndServe(":8080", nil)
}
