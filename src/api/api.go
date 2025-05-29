package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hvmidrezv/web-app/api/middlewares"
	"github.com/hvmidrezv/web-app/api/routers"
	validation "github.com/hvmidrezv/web-app/api/validations"
	"github.com/hvmidrezv/web-app/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}

	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)

	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))

}
