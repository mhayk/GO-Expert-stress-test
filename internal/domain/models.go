package domain

import "time"

type TestConfig struct {
	URL           string
	TotalRequests int
	Concurrency   int
}

type TestResult struct {
	Duration      time.Duration
	TotalRequests int
	StatusCounts  map[int]int
}
