package runner

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Hanningtone03/build-your-own-cicd/internal/logger"
	"github.com/Hanningtone03/build-your-own-cicd/internal/pipeline"
)

type Result struct {
	Step     string
	Success  bool
	Duration time.Duration
	Output   string
}

func runCommand(command string, workdir string, env []string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	if workdir != "" {
		cmd.Dir = workdir
	}
	cmd.Env = append(os.Environ(), env...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func RunPipeline(p *pipeline.Pipeline) bool {
	fmt.Printf("\n========================================\n")
	fmt.Printf("  Pipeline: %s\n", p.Name)
	fmt.Printf("========================================\n")

	allPassed := true
	results := []Result{}
	start := time.Now()

	for _, job := range p.Jobs {
		logger.Info(fmt.Sprintf("Starting job: %s", job.Name))

		for _, step := range job.Steps {
			logger.Step(step.Name)
			stepStart := time.Now()

			output, err := runCommand(step.Run, step.WorkDir, step.Env)
			duration := time.Since(stepStart)

			if output != "" {
				fmt.Println(strings.TrimSpace(output))
			}

			result := Result{
				Step:     step.Name,
				Duration: duration,
				Output:   output,
			}

			if err != nil {
				result.Success = false
				logger.Error(fmt.Sprintf("Step failed: %s (%s)", step.Name, duration))
				results = append(results, result)
				if step.OnFail == "continue" {
					logger.Warn("Continuing despite failure...")
					continue
				}
				allPassed = false
				break
			}

			result.Success = true
			logger.Success(fmt.Sprintf("Step passed: %s (%s)", step.Name, duration))
			results = append(results, result)
		}
	}

	total := time.Since(start)
	fmt.Printf("\n========================================\n")
	fmt.Printf("  Summary\n")
	fmt.Printf("========================================\n")

	for _, r := range results {
		status := "PASS"
		if !r.Success {
			status = "FAIL"
		}
		fmt.Printf("  [%s] %s (%s)\n", status, r.Step, r.Duration)
	}

	fmt.Printf("\n  Total time: %s\n", total)
	if allPassed {
		logger.Success("Pipeline completed successfully")
	} else {
		logger.Error("Pipeline failed")
	}
	fmt.Printf("========================================\n\n")

	return allPassed
}