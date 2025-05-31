package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/web-app/api/helper"
	"net/http"
)
import "github.com/didip/tollbooth/v7"

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, -100, err))
			return
		} else {
			c.Next()
		}
	}
}
