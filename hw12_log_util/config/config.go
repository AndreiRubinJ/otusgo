package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

type Config struct {
	LogFile  string
	LogLevel string
	Output   string
}

func GetConfig() Config {
	env := os.Getenv("CONFIG_ENV")
	if env == "" {
		env = "./.env_dev"
	}
	err := godotenv.Load(env)
	if err != nil {
		fmt.Printf("Error loading .env.div file: %v", err)
	}
	config := Config{
		LogLevel: "info",
	}

	pflag.StringVar(&config.LogFile, "file", "", "path to the log file")
	pflag.StringVar(&config.LogLevel, "level", "", "log level for analysis")
	pflag.StringVar(&config.Output, "output", "", "output file for statistics")

	pflag.Parse() // Use pflag's Parse method

	if config.LogFile == "" {
		config.LogFile = os.Getenv("LOG_ANALYZER_FILE")
	}
	if config.LogLevel == "" {
		config.LogLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	}
	if config.Output == "" {
		config.Output = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	return config
}
