package migration

import (
	"task_1_crud_user/internal/entity"
	"task_1_crud_user/pkg/database"
)

func AutoMigration(db *database.Database) {
	db.MigrateDB(
		&entity.User{},
	)
}
