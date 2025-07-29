package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/web-app/api/handlers"
	"github.com/hvmidrezv/web-app/config"
)

func CarType(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarTypeHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Gearbox(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewGearboxHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModel(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModelColor(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelColorHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModelYear(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelYearHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
