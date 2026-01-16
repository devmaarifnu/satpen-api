package middleware

import (
	"net/http"
	"satpen-api/internal/config"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	lastSeen time.Time
	count    int
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.RWMutex
)

func RateLimit(cfg *config.Config, rule config.RateLimitRule) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !cfg.RateLimit.Enabled {
			c.Next()
			return
		}

		ip := c.ClientIP()

		mu.Lock()
		v, exists := visitors[ip]
		if !exists {
			visitors[ip] = &visitor{
				lastSeen: time.Now(),
				count:    1,
			}
			mu.Unlock()
			c.Next()
			return
		}

		// Reset count if window has passed
		if time.Since(v.lastSeen) > time.Duration(rule.Window)*time.Second {
			v.count = 1
			v.lastSeen = time.Now()
			mu.Unlock()
			c.Next()
			return
		}

		// Check if limit exceeded
		if v.count >= rule.Requests {
			mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded",
				"error":   "Too many requests, please try again later",
			})
			c.Abort()
			return
		}

		v.count++
		mu.Unlock()
		c.Next()
	}
}

// CleanupVisitors removes old visitors periodically
func CleanupVisitors(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			mu.Lock()
			for ip, v := range visitors {
				if time.Since(v.lastSeen) > 5*time.Minute {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()
}
