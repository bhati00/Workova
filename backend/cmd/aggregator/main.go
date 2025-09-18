package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/internal/job"
	"github.com/bhati00/workova/backend/internal/job/repository"
	"github.com/bhati00/workova/backend/internal/worker"
	"github.com/bhati00/workova/backend/pkg/database"
	job_aggregator "github.com/bhati00/workova/backend/pkg/job_aggregator"
	"github.com/bhati00/workova/backend/pkg/job_aggregator/rapid"
)

func main() {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	logger.Info("Starting Job Aggregator Service")

	// Initialize dependencies
	cfg := config.LoadConfig()
	db := database.ConnectDatabase(*cfg)

	jobRepo := repository.NewJobRepository(db)
	locationRepo := repository.NewLocationRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	skillRepo := repository.NewSkillRepository(db)

	jobService := job.NewJobService(jobRepo, skillRepo, categoryRepo, locationRepo)

	// Initialize aggregators
	aggregatorList := initializeAggregators(logger)
	worker := worker.NewWorker(aggregatorList, jobService)

	// Setup fetch options
	fetchOptions := job_aggregator.FetchOptions{
		Pages:      3,
		MaxJobs:    300,
		DatePosted: getOneMonthAgo(),
	}

	// Run initial aggregation
	logger.Info("Running initial job aggregation")
	go runAggregation(worker, fetchOptions, logger)

	// Setup midnight scheduler
	scheduler := setupMidnightScheduler()
	defer scheduler.Stop()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Scheduler started - running daily at midnight")

	// Main loop
	for {
		select {
		case <-scheduler.C:
			logger.Info("Daily aggregation triggered")
			go runAggregation(worker, fetchOptions, logger)

		case <-quit:
			logger.Info("Shutting down aggregator service")
			return
		}
	}
}

func initializeAggregators(logger *slog.Logger) []job_aggregator.JobAggregator {
	var aggregators []job_aggregator.JobAggregator

	// Initialize YCombinator aggregator
	ycombAggregator := rapid.NewYCombinatorAggregator()
	aggregators = append(aggregators, ycombAggregator)
	logger.Info("YCombinator aggregator initialized")

	// Add more aggregators here as you implement them

	logger.Info("Aggregators initialized", "count", len(aggregators))
	return aggregators
}

func setupMidnightScheduler() *time.Ticker {
	now := time.Now()

	// Calculate time until next midnight
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	timeUntilMidnight := nextMidnight.Sub(now)

	// Sleep until midnight
	time.Sleep(timeUntilMidnight)

	// Then start daily ticker
	return time.NewTicker(24 * time.Hour)
}

func runAggregation(worker *worker.Worker, options job_aggregator.FetchOptions, logger *slog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	logger.Info("Starting job aggregation")
	startTime := time.Now()

	// Create a channel to receive the result
	done := make(chan struct {
		count int
		err   error
	}, 1)

	// Run aggregation in goroutine
	go func() {
		count, err := worker.AggregateJobs(options)
		done <- struct {
			count int
			err   error
		}{count, err}
	}()

	// Wait for completion or timeout
	select {
	case result := <-done:
		duration := time.Since(startTime)
		if result.err != nil {
			logger.Error("Aggregation completed with errors",
				"jobs_processed", result.count,
				"duration", duration,
				"error", result.err)
		} else {
			logger.Info("Aggregation completed successfully",
				"jobs_processed", result.count,
				"duration", duration)
		}
	case <-ctx.Done():
		logger.Error("Aggregation timed out after 30 minutes")
	}
}

func getOneMonthAgo() *time.Time {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return &oneMonthAgo
}
