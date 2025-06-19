package migrations

import (
	"embed"
)

//go:embed migration-sql/*.sql
var EmbeddedMigrations embed.FS
