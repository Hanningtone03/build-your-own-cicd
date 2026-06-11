package main

import (
	"fmt"
	"os"

	"github.com/Hanningtone03/build-your-own-cicd/internal/pipeline"
	"github.com/Hanningtone03/build-your-own-cicd/internal/runner"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <pipeline.yaml>")
		os.Exit(1)
	}

	path := args[0]
	p, err := pipeline.Parse(path)
	if err != nil {
		fmt.Printf("Error parsing pipeline: %s\n", err)
		os.Exit(1)
	}

	success := runner.RunPipeline(p)
	if !success {
		os.Exit(1)
	}
}