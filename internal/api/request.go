package api

import (
	"github.com/gin-gonic/gin"
	reqPkg "go_limiter_rate/internal/request"
	"gorm.io/gorm"
)

func startRequestRoutes(r *gin.Engine, s *gorm.DB) {
	req := r.Group("/request")
	request := reqPkg.NewSQLite(s)
	{
		req.POST("/", request.Create)
		req.GET("/", request.Read)
	}

}
