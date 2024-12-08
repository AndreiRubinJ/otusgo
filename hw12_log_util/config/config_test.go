package config

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	originalEnv := os.Environ()
	originalArgs := os.Args
	defer restoreEnv(originalEnv, originalArgs)

	tests := []struct {
		name             string
		env              map[string]string
		args             []string
		expectedLogFile  string
		expectedLogLevel string
		expectedOutput   string
	}{
		{
			name: "default values",
			env: map[string]string{
				"CONFIG_ENV":          "../.env",
				"LOG_ANALYZER_FILE":   "log.txt",
				"LOG_ANALYZER_LEVEL":  "info",
				"LOG_ANALYZER_OUTPUT": "output.log",
			},
			args:             []string{"cmd"},
			expectedLogFile:  "log.txt",
			expectedLogLevel: "info",
			expectedOutput:   "output.log",
		},
		{
			name: "command line arguments override",
			env: map[string]string{
				"CONFIG_ENV":          "",
				"LOG_ANALYZER_FILE":   "log.txt",
				"LOG_ANALYZER_LEVEL":  "info",
				"LOG_ANALYZER_OUTPUT": "output.log",
			},
			args:             []string{"cmd", "--file=override_log.txt", "--level=debug", "--output=override_output.log"},
			expectedLogFile:  "override_log.txt",
			expectedLogLevel: "debug",
			expectedOutput:   "override_output.log",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setupTest(tc.env, tc.args)
			config := GetConfig()
			verifyConfig(t, &config, tc.expectedLogFile, tc.expectedLogLevel, tc.expectedOutput)
		})
	}
}

func setupTest(env map[string]string, args []string) {
	for k, v := range env {
		os.Setenv(k, v)
	}
	os.Args = args
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
}

func verifyConfig(t *testing.T, config *Config, expectedLogFile, expectedLogLevel, expectedOutput string) {
	t.Helper()
	assert.Equal(t, expectedLogFile, config.LogFile, "LogFile should match")
	assert.Equal(t, expectedLogLevel, config.LogLevel, "LogLevel should match")
	assert.Equal(t, expectedOutput, config.Output, "Output should match")
}

func restoreEnv(originalEnv []string, originalArgs []string) {
	os.Clearenv()
	os.Args = originalArgs
	for _, e := range originalEnv {
		parts := splitNLen(e, "=", 2)
		os.Setenv(parts[0], parts[1])
	}
}

func splitNLen(s, sep string, n int) (result []string) {
	return strings.SplitN(s, sep, n)
}
