package router

import (
	"log"

	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetupRoutes(r *gin.Engine) {
	r.Group("/api")
	cfg := config.LoadConfig()
	database.ConnectDatabase(*cfg)
	Migrate(cfg.DBpath)

}

func Migrate(dbPath string) {
	// Example dbPath: "./jobs.db"
	dsn := "sqlite3://" + dbPath

	m, err := migrate.New(
		"file://db/migrations", // path to migrations folder
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
