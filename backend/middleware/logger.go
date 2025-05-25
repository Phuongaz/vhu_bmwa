package middleware

import (
	"time"
	"vhu_bmwa/logging"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		startTime := time.Now()

		c.Next()

		userID, _ := c.Get("user_id")
		var userIDUint uint
		if id, ok := userID.(uint); ok {
			userIDUint = id
		}

		queryParams := make(map[string]string)
		for k, v := range c.Request.URL.Query() {
			if len(v) > 0 {
				queryParams[k] = v[0]
			}
		}

		logging.LogAccess(logging.AccessLogEntry{
			Method:      c.Request.Method,
			Path:        c.Request.URL.Path,
			StatusCode:  c.Writer.Status(),
			Duration:    time.Since(startTime),
			IP:          c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
			UserID:      userIDUint,
			RequestID:   requestID,
			QueryParams: queryParams,
		})

		if c.Writer.Status() == 401 || c.Writer.Status() == 403 {
			logging.LogSecurity(logging.SecurityLogEntry{
				EventType:   "AUTH_FAILURE",
				UserID:      userIDUint,
				IP:          c.ClientIP(),
				Description: "Authentication or authorization failure",
				Metadata: map[string]interface{}{
					"path":        c.Request.URL.Path,
					"method":      c.Request.Method,
					"status_code": c.Writer.Status(),
					"request_id":  requestID,
				},
			})
		}
	}
}
