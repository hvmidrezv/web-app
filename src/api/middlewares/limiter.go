package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/didip/tollbooth/v7"

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.Next()
		}
	}
}
