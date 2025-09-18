package worker

import (
	"log/slog"
	"sync"

	"github.com/bhati00/workova/backend/dtos"
	"github.com/bhati00/workova/backend/internal/job"
	jobaggregator "github.com/bhati00/workova/backend/pkg/job_aggregator"
)

type Worker struct {
	JobAggregators []jobaggregator.JobAggregator
	JobService     job.JobService
	logger         *slog.Logger
}

func NewWorker(jobAggregators []jobaggregator.JobAggregator, jobService job.JobService) *Worker {
	return &Worker{
		JobAggregators: jobAggregators,
		JobService:     jobService,
		logger:         slog.Default(), // Simple logging
	}
}

func (w *Worker) AggregateJobs(fetchOptions jobaggregator.FetchOptions) (int, error) {
	var (
		mu         sync.Mutex
		wg         sync.WaitGroup
		totalJobs  int
		firstError error
	)

	// Launch a goroutine for each aggregator
	wg.Add(len(w.JobAggregators))

	for _, aggregator := range w.JobAggregators {
		go func(agg jobaggregator.JobAggregator) {
			defer wg.Done() // Signal completion when this goroutine exits

			// Fetch jobs from this aggregator
			jobRequests, err := agg.FetchJobs(fetchOptions)
			if err != nil {
				w.logger.Error("failed to fetch jobs from aggregator",
					"error", err)

				// Thread-safe error handling
				mu.Lock()
				if firstError == nil {
					firstError = err
				}
				mu.Unlock()
				return // Exit this goroutine, continue with others
			}

			// Save jobs from this aggregator
			count := w.SaveJobs(jobRequests)

			// Thread-safe update of shared variables
			mu.Lock()
			totalJobs += count
			mu.Unlock()

			w.logger.Info("processed jobs from aggregator",
				"jobs_processed", count)

		}(aggregator) // Pass aggregator to avoid closure variable capture issues
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return totalJobs, firstError
}

func (w *Worker) SaveJobs(jobRequests []dtos.JobRequest) int {
	count := 0
	for _, jobRequest := range jobRequests {
		_, err := w.JobService.CreateJob(jobRequest)
		if err != nil {
			w.logger.Error("failed to create job",
				"job_title", jobRequest.Title,
				"error", err)
			continue
		}
		count++
	}
	return count
}
