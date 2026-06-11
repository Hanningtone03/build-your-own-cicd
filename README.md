# Build Your Own CI/CD

A CI/CD pipeline runner built in Go; runs pipelines defined in YAML with steps, logging and failure handling.

## How it works

CI/CD tools like GitHub Actions and Jenkins read pipeline definitions and execute them step by step. This project implements that from scratch:

- Parses YAML pipeline definitions into structured job and step objects
- Executes each step as a shell command in sequence
- Streams output from each step in real time
- Tracks duration and status of every step
- Supports failure handling — stop or continue on failure
- Prints a formatted summary at the end

## Project structure

```
main.go
internal/
├── pipeline/
│   └── parser.go
├── runner/
│   └── executor.go
└── logger/
    └── logger.go
```

## Running locally

```bash
go run main.go pipeline.yaml
```

## Pipeline format

```yaml
name: build-and-test
jobs:
  - name: build
    steps:
      - name: Check Go files
        run: go vet ./...
      - name: Build binary
        run: go build -o bin/cicd main.go
  - name: test
    steps:
      - name: Run tests
        run: echo "All tests passed"
        on_fail: continue
```

## Tech

- Go
- `gopkg.in/yaml.v3`
- `os/exec` module
