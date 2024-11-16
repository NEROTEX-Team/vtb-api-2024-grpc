package main

import (
	"context"
	"log"
	"os"
	"time"

	"/internal/antivirus"
	"/internal/app"
)

func main() {
	ctx := context.Background()

	useAntivirus := os.Getenv("USE_ANTIVIRUS") == "true"

	scanner := antivirus.NewScanner("localhost:3310", "tcp", 10*time.Second, useAntivirus)

	a, err := app.NewApp(ctx, scanner)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
