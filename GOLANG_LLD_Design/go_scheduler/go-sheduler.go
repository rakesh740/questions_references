package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Scheduler struct {
	jobs      []Job
	intervals []time.Duration
	stop      chan struct{}
	wg        sync.WaitGroup
}

// NewScheduler creates a new Scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs:      make([]Job, 0),
		intervals: make([]time.Duration, 0),
		stop:      make(chan struct{}),
	}
}

// AddJob adds a job to the scheduler with a specified interval
func (s *Scheduler) AddJob(job Job, interval time.Duration) {
	s.jobs = append(s.jobs, job)
	s.intervals = append(s.intervals, interval)
}

// Start starts the job scheduler
func (s *Scheduler) Start() {
	for i, job := range s.jobs {
		s.wg.Add(1)
		go func(j Job, interval time.Duration) {
			defer s.wg.Done()
			ticker := time.NewTicker(interval)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					j()
				case <-s.stop:
					return
				}
			}
		}(job, s.intervals[i])
	}
}

// Stop stops the job scheduler
func (s *Scheduler) Stop() {
	close(s.stop)
	s.wg.Wait()
}

func main() {
	s := NewScheduler()

	s.AddJob(func() {
		fmt.Println("Job 1 executed")
	}, 2*time.Second)

	s.AddJob(func() {
		fmt.Println("Job 2 executed")
	}, 3*time.Second)

	s.Start()

	time.Sleep(10 * time.Second)
	s.Stop()
	fmt.Println("Scheduler stopped")
}
