package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/web-app/api/handlers"
	"github.com/hvmidrezv/web-app/api/middlewares"
	"github.com/hvmidrezv/web-app/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}
