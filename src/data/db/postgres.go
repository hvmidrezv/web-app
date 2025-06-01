package db

import (
	"fmt"
	"github.com/hvmidrezv/web-app/config"
	"github.com/hvmidrezv/web-app/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)
	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, _ := dbClient.DB()

	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w (Ping)", err)
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.Startup, "Connected to database", nil)
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}
