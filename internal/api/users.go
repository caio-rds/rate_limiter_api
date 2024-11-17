package api

import (
	"github.com/gin-gonic/gin"
	userPkg "go_limiter_rate/internal/user"
	"gorm.io/gorm"
)

func startUserRoutes(r *gin.Engine, db *gorm.DB) {
	userGroup := r.Group("/user")
	user := userPkg.NewSQLite(db)
	{
		userGroup.GET("/:id", user.Read)
		userGroup.POST("/", user.Create)
		userGroup.PUT("/:id", user.Update)
		userGroup.DELETE("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "delete user by id"})
		})
	}
}
