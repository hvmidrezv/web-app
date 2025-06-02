package main

import (
	"github.com/hvmidrezv/web-app/api"
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/data/cache"
	"github.com/hvmidrezv/web-app/data/db"
	"github.com/hvmidrezv/web-app/data/db/migrations"
	"github.com/hvmidrezv/web-app/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()

	migrations.Up_1()

	api.InitServer(cfg)

}
