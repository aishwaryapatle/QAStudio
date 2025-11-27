package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/aishwaryapatle/qastudio/internal/auth/model"
	"github.com/aishwaryapatle/qastudio/internal/config"
	"github.com/aishwaryapatle/qastudio/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		cfg := config.Load()

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
			cfg.DBHost,
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBName,
			cfg.DBPort,
			cfg.DBSSL,
		)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Log.Errorf("❌ failed to connect DB: %v", err)
			return
		}

		err = DB.AutoMigrate(&model.User{})
		sqlDB, err2 := DB.DB()
		if err2 != nil {
			err = err2
			return
		}

		// Connection pooling (async safe)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)

		logger.Log.Infof("✅ database connected successfully")
	})

	return DB, err
}

// Close closes database (used on graceful shutdown)
func Close() error {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	logger.Log.Infof("🔌 closing DB connection")
	return sqlDB.Close()
}