package main

import (
	"github.com/hvmidrezv/web-app/api"
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)

}
