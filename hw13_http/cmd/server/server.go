package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/config/server"
	"github.com/AndreiRubinJ/otusgo/hw13_http/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	serverURL := getServerURL()
	router := setupRouter()
	startServer(serverURL, router)
}
func startServer(serverURL string, router *chi.Mux) {
	fmt.Printf("Starting server with url %s...", serverURL)
	serv := &http.Server{
		Addr:         serverURL,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := serv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getServerURL() string {
	cfg, err := server.ParseConfig()
	if err != nil {
		fmt.Printf("Error parsing configuration: %v\n", err)
		os.Exit(1)
	}
	if err := cfg.Validate(); err != nil {
		fmt.Printf("Configuration validation error: %v\n", err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s:%s", cfg.URL, cfg.Port)
}
func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/getUser", handlers.GetHandler)
	router.Post("/addUser", handlers.PostHandler)
	return router
}

func ValidateURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", fmt.Errorf("unsupported URL scheme: %s", parsedURL.Scheme)
	}
	if parsedURL.Host == "" {
		return "", fmt.Errorf("URL must have a host")
	}
	return parsedURL.String(), nil
}
