package client

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Method string
	URL    string
	Path   string
	Data   string
}

func ParseConfig() (*Config, error) {
	env := os.Getenv("CONFIG_ENV")
	if env == "" {
		env = "./.env_dev"
	}
	err := godotenv.Load(env)
	if err != nil {
		fmt.Printf("Error loading .env.div file: %v", err)
	}

	method := flag.String("method", "", "HTTP method: GET or POST")
	url := flag.String("url", "http://localhost:8080", "Base server URL (e.g., http://localhost:8080)")
	path := flag.String("path", "", "Resource path (e.g., /getUser)")
	data := flag.String("data", "", "Request body for POST method")
	flag.Parse()

	if *method == "" {
		*method = os.Getenv("REQUEST_METHOD")
	}
	if *path == "" {
		*path = os.Getenv("REQUEST_PATH")
	}
	if *data == "" {
		*data = os.Getenv("REQUEST_DATA")
	}

	if *url == "" {
		return nil, fmt.Errorf("missing required flag: -url")
	}

	return &Config{
		Method: *method,
		URL:    *url,
		Path:   *path,
		Data:   *data,
	}, nil
}

// Validate ensures the provided configuration is valid.
func (c *Config) Validate() error {
	if c.Method != "GET" && c.Method != "POST" {
		return fmt.Errorf("unsupported HTTP method: %s", c.Method)
	}
	return nil
}
