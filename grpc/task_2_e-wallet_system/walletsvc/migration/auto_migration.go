package migration

import (
	"walletsvc/internal/entity"
	"walletsvc/pkg/database"
)

func AutoMigration(db *database.Database) {
	db.MigrateDB(
		&entity.Wallet{},
		&entity.Transaction{},
	)
}
