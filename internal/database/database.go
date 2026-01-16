package database

import (
	"fmt"
	"log"
	"time"

	"satpen-api/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes database connection
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	// Configure GORM logger
	logLevel := logger.Silent
	if cfg.IsDevelopment() {
		logLevel = logger.Info
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get generic database object to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	log.Println("Database connected successfully")
	return db, nil
}

// AutoMigrate runs auto migration for models
// DISABLED: Menggunakan database yang sudah ada (testing_lpmaarif1)
func AutoMigrate(db *gorm.DB) error {
	log.Println("Skipping auto migration (using existing database schema)")
	// Database schema already exists, no need to migrate
	// Models are mapped to existing tables:
	// - models.Satpen -> satpen
	// - models.Provinsi -> provinsi
	// - models.Kabupaten -> kabupaten
	// - models.JenjangPendidikan -> jenjang_pendidikan
	// - models.KategoriSatpen -> kategori_satpen
	// - models.PengurusCabang -> pengurus_cabang
	// - models.PDPTK -> pdptk
	return nil
}

// Close closes database connection
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
