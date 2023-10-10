package seeders

import (
	"github.com/atom-apps/storage/common/consts"
	"github.com/atom-apps/storage/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type DriverSeeder struct{}

func NewDriverSeeder() contracts.Seeder {
	return &DriverSeeder{}
}

func (s *DriverSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Driver{}).TableName(nil))
	ms := []*models.Driver{
		{
			Type:         consts.FilesystemDriverLocal,
			Name:         "LocalDriver",
			Endpoint:     "https://localhost:9899/statics/",
			AccessKey:    "--",
			AccessSecret: "---",
			Bucket:       "/home/rogee/assets",
			IsDefault:    true,
		},
		{
			Type:         consts.FilesystemDriverAliyunOss,
			Name:         "AliOSS",
			Endpoint:     "https://localhost:9899/statics/",
			AccessKey:    "--",
			AccessSecret: "---",
			Bucket:       "/home/rogee/assets",
		},
	}
	db.CreateInBatches(ms, 10)
}
