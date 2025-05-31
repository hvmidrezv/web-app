package main

import (
	"github.com/hvmidrezv/web-app/api"
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/data/cache"
	"github.com/hvmidrezv/web-app/data/db"
	"log"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()
	api.InitServer(cfg)

}
