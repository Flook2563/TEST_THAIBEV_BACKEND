package database

import (
	"fmt"
	"log"
	"os"
	"thaibev_backend/appconfig"
	"thaibev_backend/internal/repositories"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenPostgresqlDatabase(cfg appconfig.Database) (*gorm.DB, error) {
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=%s",
			cfg.Host,
			cfg.User,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.SSLMode,
			cfg.Timezone,
			cfg.SearchPath,
		)),
		&gorm.Config{
			Logger: sqlLogger,
			DryRun: false,
		})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	// Set connection pool options
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(
		&repositories.TbTUser{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate tables: %v", err)
	}

	fmt.Println("Database tables migrated successfully!")

	return db, nil
}
