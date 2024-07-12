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

	"walletsvc/pkg/logger"
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
	Logger   logger.Logger
}

type Database struct {
	db     *gorm.DB
	logger logger.Logger
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
		cfg.Logger.Warn("oracle driver is not supported yet")
		os.Exit(1)
	default:
		cfg.Logger.Warn("unknown database driver")
		os.Exit(1)
	}

	for {
		configGorm := &gorm.Config{}
		if cfg.Debug {
			if cfg.Logger.Type() == "slog" {
				configGorm.Logger = slogGorm.New(
					slogGorm.WithHandler(cfg.Logger.Default().(*slog.Logger).Handler()),
					slogGorm.WithTraceAll(),
					slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug),
				)
			}
		}
		db, err = gorm.Open(dialect, configGorm)
		if err != nil {
			cfg.Logger.Error(fmt.Sprintf("failed to connect to %s database", cfg.DbDriver), "error", err.Error())
			cfg.Logger.Info(fmt.Sprintf("retrying to connect to %s database in 5 seconds...", cfg.DbDriver))
			time.Sleep(5 * time.Second)
			continue
		}
		cfg.Logger.Info(fmt.Sprintf("successfully connected to %s database", cfg.DbDriver))
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		cfg.Logger.Error("failed to configure connection pool", "error", err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{db: db, logger: cfg.Logger}
}

func (d *Database) MigrateDB(entities ...Entity) {
	var dst []any
	for _, e := range entities {
		dst = append(dst, e)
	}
	err := d.db.AutoMigrate(dst...)
	if err != nil {
		d.logger.Error("failed to migrate db", "error", err.Error())
		os.Exit(1)
	}
	for _, e := range entities {
		d.logger.Info(fmt.Sprintf("successfully migrated %s", e.TableName()))
	}
}

type Entity interface {
	TableName() string
}
