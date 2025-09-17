package worker

import (
	"log/slog"

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
	totalJobs := 0
	var firstError error

	for _, aggregator := range w.JobAggregators {
		jobRequests, err := aggregator.FetchJobs(fetchOptions)
		if err != nil {
			w.logger.Error("failed to fetch jobs from aggregator",
				"error", err)

			// Remember first error but continue with other aggregators
			if firstError == nil {
				firstError = err
			}
			continue
		}

		count := w.SaveJobs(jobRequests)
		totalJobs += count
		w.logger.Info("processed jobs from aggregator",
			"jobs_processed", count)
	}

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
