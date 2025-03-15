package config

import (
	"fmt"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Config.DBHost, Config.DBUser, Config.DBPassword, Config.DBName, Config.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.New(Log, gormLogger.Config{
			SlowThreshold:             0, // bisa atur durasi log lambat
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		Log.WithField(enum.LevelError, err).Fatal("failed to connect to database")
	}
	Log.Info("Connected to PostgreSQL successfully")
}
