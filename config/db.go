package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() (*gorm.DB, error) {
	DbHost := GetEnv("DB_HOST", "")
	DbPort := GetEnv("DB_PORT", "")
	DbDatabase := GetEnv("DB_DATABASE", "")
	DbUsername := GetEnv("DB_USERNAME", "")
	DbPassword := GetEnv("DB_PASSWORD", "")
	DbCluster := os.Getenv("DB_CLUSTER")
	appEnv := os.Getenv("APP_ENV")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", DbHost, DbUsername, DbPassword, DbDatabase, DbPort)
	if appEnv != "development" {
		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full&options=%s", DbUsername, DbPassword, DbHost, DbPort, DbDatabase, DbCluster)
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
}
