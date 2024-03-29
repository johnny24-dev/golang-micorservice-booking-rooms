package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/nekizz/final-project/backend/booking/migration/versions"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "20221110000002",
			Migrate: versions.Version20221110000002,
		},
		{
			ID:      "20221110000001",
			Migrate: versions.Version20221110000001,
		},
	})

	return m.Migrate()
}
