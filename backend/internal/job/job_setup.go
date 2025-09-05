package job

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JobModule represents the complete job module with all dependencies
type JobModule struct {
	Repository JobRepository
	Service    JobService
	Handler    *JobHandler
}

// InitializeJobModule initializes the complete job module
func InitializeJobModule(db *gorm.DB) *JobModule {
	// Initialize repository with database connection
	jobRepo := NewJobRepository(db)

	// Initialize service with repository dependency
	jobService := NewJobService(jobRepo)

	// Initialize handler with service dependency
	jobHandler := NewJobHandler(jobService)

	return &JobModule{
		Repository: jobRepo,
		Service:    jobService,
		Handler:    jobHandler,
	}
}

// RegisterRoutes registers all job routes to the router
func (jm *JobModule) RegisterRoutes(router *gin.Engine) {
	// Create API version group
	v1 := router.Group("/api/v1")

	// Register job routes
	jm.Handler.RegisterJobRoutes(v1)
}
