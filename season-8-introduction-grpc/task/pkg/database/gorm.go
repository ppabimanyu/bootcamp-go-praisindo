package database

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GormConfig struct {
	DbHost   string
	DbUser   string
	DbPass   string
	DbName   string
	DbPort   string
	DbPrefix string
	DbDriver string
	Debug    bool
}

type Database struct {
	db *gorm.DB
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func NewDatabase(cfg *GormConfig) *Database {
	var db *gorm.DB
	var err error
	var dialect gorm.Dialector

	switch cfg.DbDriver {
	case "postgres", "pgsql":
		dialect = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kuala_Lumpur", cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort))
	case "mysql":
		dialect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName))
	case "sqlserver":
		dialect = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbName))
	case "oracle":
		slog.Warn("oracle driver is not supported yet")
		os.Exit(1)
	default:
		slog.Warn("unknown database driver")
		os.Exit(1)
	}

	for {
		configGorm := &gorm.Config{}
		if cfg.Debug {
			configGorm.Logger = slogGorm.New(
				slogGorm.WithHandler(slog.Default().Handler()),
				slogGorm.WithTraceAll(),
				slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug),
			)
		}
		db, err = gorm.Open(dialect, configGorm)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to connect to %s database", cfg.DbDriver), "error", err.Error())
			slog.Info(fmt.Sprintf("retrying to connect to %s database in 5 seconds...", cfg.DbDriver))
			time.Sleep(5 * time.Second)
			continue
		}
		slog.Info(fmt.Sprintf("successfully connected to %s database", cfg.DbDriver))
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to configure connection pool", "error", err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{db: db}
}

func (d *Database) MigrateDB(entities ...Entity) {
	var dst []any
	for _, e := range entities {
		dst = append(dst, e)
	}
	err := d.db.AutoMigrate(dst...)
	if err != nil {
		slog.Error("failed to migrate db", "error", err.Error())
		os.Exit(1)
	}
	for _, e := range entities {
		slog.Info(fmt.Sprintf("successfully migrated %s", e.TableName()))
	}
}

type Entity interface {
	TableName() string
}
