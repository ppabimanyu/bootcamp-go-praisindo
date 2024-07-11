package migration

import (
	"intro-grpc-task/internal/entity"
	"intro-grpc-task/pkg/database"
)

func AutoMigration(db *database.Database) {
	db.MigrateDB(
		&entity.User{},
	)
}
