package api

import (
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/middleware"
	userPkg "go_limiter_rate/internal/user"
	"gorm.io/gorm"
)

func startUserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/user")
	user := userPkg.NewSQLite(db)
	{
		userGroup.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			UserID := c.GetUint("user_id")
			if UserID == 0 {
				c.JSON(400, gin.H{"error": "User ID is required"})
				return
			}
			user.Read(c, UserID)
		})
		userGroup.POST("/", user.Create)
		userGroup.PUT("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			UserID := c.GetUint("user_id")
			if UserID == 0 {
				c.JSON(400, gin.H{"error": "User ID is required"})
				return
			}
			user.Update(c, UserID)
		})
		userGroup.DELETE("/", func(c *gin.Context) {
			UserID := c.GetUint("user_id")
			if UserID == 0 {
				c.JSON(400, gin.H{"error": "User ID is required"})
				return
			}
			user.Delete(c, UserID)
		})
	}
}
