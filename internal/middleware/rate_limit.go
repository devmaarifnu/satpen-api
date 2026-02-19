package middleware

import (
	"context"
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
	visitors      = make(map[string]*visitor)
	mu            sync.RWMutex
	cleanupTicker *time.Ticker
	cleanupOnce   sync.Once
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

// StartCleanup removes old visitors periodically - should be called only ONCE at startup
func StartCleanup(ctx context.Context, interval time.Duration) {
	cleanupOnce.Do(func() {
		cleanupTicker = time.NewTicker(interval)
		go func() {
			for {
				select {
				case <-cleanupTicker.C:
					mu.Lock()
					now := time.Now()
					for ip, v := range visitors {
						if now.Sub(v.lastSeen) > 5*time.Minute {
							delete(visitors, ip)
						}
					}
					mu.Unlock()
				case <-ctx.Done():
					cleanupTicker.Stop()
					return
				}
			}
		}()
	})
}

// StopCleanup stops the cleanup ticker - call this on graceful shutdown
func StopCleanup() {
	if cleanupTicker != nil {
		cleanupTicker.Stop()
	}
}
