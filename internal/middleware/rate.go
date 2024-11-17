package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

var ctx = context.Background()

func RateLimiterMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit_%s", ip)
		limit := 10
		duration := 10 * time.Second

		count, err := rdb.Get(ctx, key).Int()
		if errors.Is(err, redis.Nil) {
			rdb.Set(ctx, key, 1, duration)
		} else if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else if count >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		} else {
			rdb.Incr(ctx, key)
		}

		c.Next()
	}
}
