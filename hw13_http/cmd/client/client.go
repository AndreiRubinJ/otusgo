package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/config/client"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}
	fullURL := cfg.URL + cfg.Path
	switch cfg.Method {
	case "GET":
		err = sendGetRequest(fullURL)
		if err != nil {
			fmt.Printf("Error performing the HTTP request: %v\n", err)
			os.Exit(1)
		}
	case "POST":
		err = sendPostRequest(fullURL, cfg.Data)
		if err != nil {
			fmt.Printf("error performing the HTTP request: %v\n", err)
			os.Exit(1)
		}
	}
}

func getConfig() (*client.Config, error) {
	cfg, err := client.ParseConfig()
	if err != nil {
		return nil, fmt.Errorf("error parsing configuration: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation error: %w", err)
	}
	return cfg, nil
}
func sendPostRequest(serverURL string, data string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return fmt.Errorf("URL is not allowed: %s", serverURL)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing the HTTP request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	readResponseBody(resp)
	return nil
}

func sendGetRequest(serverURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	if err != nil {
		return fmt.Errorf("URL is not allowed: %s", serverURL)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing the HTTP request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received unexpected status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	readResponseBody(resp)
	return nil
}

func readResponseBody(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the server response: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)
	fmt.Println("Server Response:")
	fmt.Println(string(body))
}
