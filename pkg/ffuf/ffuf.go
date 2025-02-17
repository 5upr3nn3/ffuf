package ffuf

import (
	"context"
	"fmt"
)

// Scanner represents the main ffuf scanning engine
type Scanner struct {
	config *Config
	job    *Job
}

// NewScanner creates a new ffuf scanner instance with the given configuration
func NewScanner(config *Config) (*Scanner, error) {
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	// Set defaults if not specified
	if config.Threads <= 0 {
		config.Threads = 10
	}
	if config.Timeout <= 0 {
		config.Timeout = 10
	}
	if config.Method == "" {
		config.Method = "GET"
	}

	return &Scanner{
		config: config,
	}, nil
}

// Run executes the scan with the configured settings
func (s *Scanner) Run(ctx context.Context) error {
	var err error
	
	// Create job
	s.job, err = NewJob(s.config)
	if err != nil {
		return fmt.Errorf("could not create job: %v", err)
	}

	// Run the job
	if err := s.job.Start(ctx); err != nil {
		return fmt.Errorf("job execution failed: %v", err)
	}

	return nil
}

// Stop gracefully stops the scanning process
func (s *Scanner) Stop() {
	if s.job != nil {
		s.job.Stop()
	}
}

// Results returns the scan results
func (s *Scanner) Results() []Result {
	if s.job != nil && s.job.Output != nil {
		return s.job.Output.GetResults()
	}
	return nil
} 