package migrations

import (
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/data/db"
	"github.com/hvmidrezv/web-app/data/models"
	"github.com/hvmidrezv/web-app/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}

func Down_1() {

}
