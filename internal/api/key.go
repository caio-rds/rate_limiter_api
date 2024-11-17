package api

import (
	"github.com/gin-gonic/gin"
	keyPkg "go_limiter_rate/internal/key"
	"go_limiter_rate/internal/middleware"
	"gorm.io/gorm"
	"strconv"
)

func startKeyRoutes(router *gin.Engine, s *gorm.DB) {
	keyRouter := router.Group("/keys")
	key := keyPkg.NewDb(s)
	{
		keyRouter.POST("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			ID := c.GetUint("user_id")
			key.Create(c, ID)
		})
		keyRouter.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			ID := c.GetUint("user_id")
			key.Read(c, ID)
		})
		keyRouter.DELETE("/:kid", middleware.AuthMiddleware(), func(c *gin.Context) {
			ID := c.GetUint("user_id")
			if kid, err := c.Params.Get("kid"); err {
				if keyId, err := strconv.ParseUint(kid, 10, 64); err == nil {
					key.Delete(c, ID, uint(keyId))
				}
			} else {
				c.JSON(400, gin.H{"error": "Invalid key id"})
			}
		})
	}
}
