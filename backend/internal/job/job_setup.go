package job

import (
	"log"

	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/internal/job/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"gorm.io/gorm"
)

// JobModule represents the complete job module with all dependencies
type JobModule struct {
	Service JobService
	Handler *JobHandler
}

// InitializeJobModule initializes the complete job module
func InitializeJobModule(db *gorm.DB, config *config.Config) *JobModule {
	// Run migrations
	Migrate(config.DBpath)
	var jobRepo repository.JobRepository
	var locationRepo repository.LocationRepository
	var categoryRepo repository.CategoryRepository
	var skillRepo repository.SkillRepository
	// Initialize repository with database connection
	jobRepo = repository.NewJobRepository(db)
	locationRepo = repository.NewLocationRepository(db)
	categoryRepo = repository.NewCategoryRepository(db)
	skillRepo = repository.NewSkillRepository(db)

	// Initialize service with repository dependency
	jobService := NewJobService(jobRepo, skillRepo, categoryRepo, locationRepo)

	// Initialize handler with service dependency
	jobHandler := NewJobHandler(jobService)

	return &JobModule{
		Handler: jobHandler,
	}
}

// RegisterRoutes registers all job routes to the router
func (jm *JobModule) RegisterRoutes(router *gin.Engine) {
	// Create API version group
	v1 := router.Group("/api/v1")

	// Register job routes
	jm.Handler.RegisterJobRoutes(v1)
}

func Migrate(dbPath string) {
	// Example dbPath: "./jobs.db"
	dsn := "sqlite3://" + dbPath

	m, err := migrate.New(
		"file://./migrations", // path to migrations folder
		dsn,
	)
	if err != nil {
		log.Fatalf("Could not initialize migration: %v", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
