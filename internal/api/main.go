package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go_limiter_rate/internal/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func StartApp(sqlite *gorm.DB, rdb *redis.Client) {
	r := gin.Default()
	r.Use(middleware.RateLimiterMiddleware(rdb))
	startUserRoutes(r, sqlite)
	startLoginRouter(r, sqlite)
	startKeyRoutes(r, sqlite)
	startPackRoutes(r, sqlite)
	startRequestRoutes(r, sqlite)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy"})
	})

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
