package migration

import (
	"usersvc/internal/entity"
	"usersvc/pkg/database"
)

func AutoMigration(db *database.Database) {
	db.MigrateDB(
		&entity.User{},
	)
}
