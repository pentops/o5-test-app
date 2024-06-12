package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pentops/o5-test-app/internal/test/runner"
)

func main() {
	envFile := flag.String("env", "", "env file for the environment")

	flag.Parse()

	if *envFile != "" {
		if err := loadEnv(*envFile); err != nil {
			fmt.Printf("LoadEnv: %s\n", err)
			os.Exit(1)
		}
	}

	cfg := &runner.APIConfig{
		APIRoot:     os.Getenv("O5_E2E_API_ADDR"),
		MetaRoot:    os.Getenv("O5_E2E_O5_ADDR"),
		BearerToken: os.Getenv("O5_TOKEN"),
	}

	ctx := context.Background()
	tags := flag.Args()
	if err := runner.Run(ctx, cfg, tags); err != nil {
		fmt.Printf("e2e.Run: %s\n", err)
		os.Exit(1)
	}
}

func loadEnv(filename string) error {
	if filename == "" {
		return nil
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
