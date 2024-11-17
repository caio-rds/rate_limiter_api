package api

import (
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/middleware"
	packPkg "go_limiter_rate/internal/pack"
	"gorm.io/gorm"
)

func startPackRoutes(r *gin.Engine, s *gorm.DB) {
	p := r.Group("/pack")
	pack := packPkg.NewSQLite(s)
	{
		p.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			userID := c.GetUint("user_id")
			if userID == 0 {
				c.JSON(400, gin.H{"error": "no user_id"})
				return
			}
			pack.Read(c, userID)
		})
		p.POST("/:amount", middleware.AuthMiddleware(), func(c *gin.Context) {
			userID := c.GetUint("user_id")
			if userID == 0 {
				c.JSON(400, gin.H{"error": "no user_id"})
				return
			}
			pack.Create(c, userID)
		})
	}
}
