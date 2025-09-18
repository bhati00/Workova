package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/internal/job"
	"github.com/bhati00/workova/backend/internal/job/repository"
	"github.com/bhati00/workova/backend/internal/worker"
	. "github.com/bhati00/workova/backend/pkg/database"
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

	// Initialize job service (replace with your actual implementation)
	var jobRepo repository.JobRepository
	var locationRepo repository.LocationRepository
	var categoryRepo repository.CategoryRepository
	var skillRepo repository.SkillRepository
	// Initialize repository with database connection
	cfg := config.LoadConfig()
	db := ConnectDatabase(*cfg)
	jobRepo = repository.NewJobRepository(db)
	locationRepo = repository.NewLocationRepository(db)
	categoryRepo = repository.NewCategoryRepository(db)
	skillRepo = repository.NewSkillRepository(db)

	// Initialize service with repository dependency
	jobService := job.NewJobService(jobRepo, skillRepo, categoryRepo, locationRepo)

	// Initialize aggregators
	aggregatorList := initializeAggregators(logger)

	// Create worker
	worker := worker.NewWorker(aggregatorList, jobService)

	// Setup fetch options
	fetchOptions := job_aggregator.FetchOptions{
		Pages:      3,
		MaxJobs:    300,
		DatePosted: getOneMonthAgo(),
	}

	// Run immediately on startup
	logger.Info("Running initial job aggregation")
	if count, err := worker.AggregateJobs(fetchOptions); err != nil {
		logger.Error("Initial aggregation failed", "error", err, "jobs_processed", count)
	} else {
		logger.Info("Initial aggregation completed", "jobs_processed", count)
	}

	// Setup daily scheduler
	scheduler := time.NewTicker(24 * time.Hour)
	defer scheduler.Stop()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Scheduler started - running daily at midnight")

	for {
		select {
		case <-scheduler.C:
			// Run at midnight
			runAggregation(worker, fetchOptions, logger)

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

func runAggregation(worker *worker.Worker, options job_aggregator.FetchOptions, logger *slog.Logger) {
	logger.Info("Starting daily job aggregation")
	startTime := time.Now()

	count, err := worker.AggregateJobs(options)
	duration := time.Since(startTime)

	if err != nil {
		logger.Error("Daily aggregation completed with errors",
			"jobs_processed", count,
			"duration", duration,
			"error", err)
	} else {
		logger.Info("Daily aggregation completed successfully",
			"jobs_processed", count,
			"duration", duration)
	}
}
func getOneMonthAgo() *time.Time {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return &oneMonthAgo
}
