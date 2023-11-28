package ancientt

import "sync"

type Run struct {
	OutputFile string   `json:"output_file"`
	Args       []string `json:"args"`
}

type Runner struct {
	mutex sync.Mutex
}

func New() *Runner {
	return &Runner{
		mutex: sync.Mutex{},
	}
}

// TODO add logic to run, check if ancientt is running, and get ancientt test results
