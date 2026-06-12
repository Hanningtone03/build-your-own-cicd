![CI](https://github.com/Hanningtone03/build-your-own-cicd/actions/workflows/ci.yml/badge.svg)

# Build Your Own CI/CD

A CI/CD pipeline runner in Go; YAML-defined pipelines, step execution, logging, failure handling.

## How it works

A YAML file defines jobs and steps. The runner parses it, executes each step as a shell command, streams output, tracks duration and status, and prints a summary. Steps can be configured to continue or stop on failure.

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
