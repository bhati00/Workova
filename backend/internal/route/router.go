package router

import (
	"github.com/bhati00/workova/backend/config"
	"github.com/bhati00/workova/backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Group("/api")
	cfg := config.LoadConfig()
	database.ConnectDatabase(*cfg)

}
