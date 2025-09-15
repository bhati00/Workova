package bootstrap

import (
	"time"

	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/docs"
	"github.com/bhati00/workova/backend/internal/job"
	. "github.com/bhati00/workova/backend/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title workova API
// @version 1.0
// @description This is the backend API for Fynelo.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@fynelo.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func InitializeApp() *gin.Engine {
	cfg := config.LoadConfig()
	db := ConnectDatabase(*cfg)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your Next.js frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jobModule := job.InitializeJobModule(db, cfg)
	jobModule.RegisterRoutes(r)
	return r
}
