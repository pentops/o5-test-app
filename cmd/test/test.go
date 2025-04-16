package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pentops/o5-test-app/internal/test"
	"github.com/pentops/o5-test-app/internal/test/universe"
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

	apiRoot := os.Getenv("O5_API")
	if apiRoot == "" {
		fmt.Println("O5_API is not set")
		os.Exit(1)
	}

	replayQueueURL := os.Getenv("REPLAY_QUEUE_URL")
	if replayQueueURL == "" {
		fmt.Println("REPLAY_QUEUE_URL is not set")
		os.Exit(1)
	}

	cfg := &universe.APIConfig{
		APIRoot:        apiRoot,
		MetaRoot:       apiRoot,
		BearerToken:    os.Getenv("O5_BEARER"),
		ReplayQueueURL: replayQueueURL,
	}

	ctx := context.Background()
	tags := flag.Args()
	if err := test.Run(ctx, cfg, tags); err != nil {
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

	lines := strings.SplitSeq(string(fileData), "\n")
	for line := range lines {
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
