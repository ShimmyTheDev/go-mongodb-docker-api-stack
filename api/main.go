package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// Logger middleware for logging requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"remote": r.RemoteAddr,
		}).Info("received request")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	// Use default middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Custom logger middleware
	r.Use(Logger)

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to API!"))
	})

	r.Get("/test-logging", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Test logging endpoint hit")
		w.Write([]byte("Logging test successful!"))
	})

	// Setup logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	// Start the server
	logrus.Info("Starting server on port 8000")
	http.ListenAndServe(":8000", r)
}
