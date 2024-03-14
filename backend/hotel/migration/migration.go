package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/nekizz/final-project/backend/hotel/migration/versions"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "20221110000003",
			Migrate: versions.Version20221110000003,
		},
		{
			ID:      "20230101000000",
			Migrate: versions.Version20230101000000,
		},
		{
			ID:      "20230225000001",
			Migrate: versions.Version20230225000001,
		},
		{
			ID:      "20221110000006",
			Migrate: versions.Version20221110000006,
		},
	})

	return m.Migrate()
}
