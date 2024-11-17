package api

import (
	"github.com/gin-gonic/gin"
	loginPkg "go_limiter_rate/internal/login"
	"go_limiter_rate/internal/middleware"
	"gorm.io/gorm"
	"net/http"
)

func startLoginRouter(r *gin.Engine, db *gorm.DB) {
	loginGroup := r.Group("/login")
	login := loginPkg.NewDb(db)
	{
		loginGroup.POST("/", login.TryLogin)
		loginGroup.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			userId := c.GetUint("user_id")
			c.JSON(http.StatusOK, gin.H{"username": username, "user_id": userId})
		})
		loginGroup.POST("/refresh", loginPkg.RefreshToken)
	}
}
