package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	URL  string
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

	url := flag.String("url", "", "Base server URL (e.g., localhost:8080)")
	port := flag.String("port", "", "Resource port (e.g., 8080)")

	flag.Parse()

	if *url == "" {
		*url = os.Getenv("REQUEST_SERVER_URL")
	}
	if *port == "" {
		*port = os.Getenv("SERVER_PORT")
	}

	if *url == "" {
		return nil, fmt.Errorf("missing required flag: -url")
	}

	return &Config{
		URL:  *url,
		Port: *port,
	}, nil
}

func (c *Config) Validate() error {
	if c.URL == "" && c.Port == "" {
		return fmt.Errorf("unsupported empty port %s or url: %s", c.Port, c.URL)
	}
	return nil
}
