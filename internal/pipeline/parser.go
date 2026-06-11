package pipeline

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Step struct {
	Name    string   `yaml:"name"`
	Run     string   `yaml:"run"`
	WorkDir string   `yaml:"workdir"`
	Env     []string `yaml:"env"`
	OnFail  string   `yaml:"on_fail"`
}

type Job struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Pipeline struct {
	Name string `yaml:"name"`
	Jobs []Job  `yaml:"jobs"`
}

func Parse(path string) (*Pipeline, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var pipeline Pipeline
	if err := yaml.Unmarshal(data, &pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}