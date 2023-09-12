package seeders

import (
	"log"

	"github.com/atom-apps/storage/common/consts"
	"github.com/atom-apps/storage/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

type FilesystemSeeder struct{}

func NewFilesystemSeeder() contracts.Seeder {
	return &FilesystemSeeder{}
}

func (s *FilesystemSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	times := 50
	for i := 0; i < times; i++ {
		data := s.Generate(faker, i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *FilesystemSeeder) Generate(faker *gofakeit.Faker, idx int) models.Filesystem {
	return models.Filesystem{
		CreatedAt: faker.Date(),
		UpdatedAt: faker.Date(),
		TenantID:  1,
		UserID:    1,
		DriverID:  1,
		Filename:  faker.AppName(),
		Type:      consts.Filesystem(faker.RandomString([]string{"dir", "file"})),
		ParentID:  0,
		Status:    consts.FileStatus(faker.RandomString([]string{"", "uploading"})),
		Mime:      faker.FileMimeType(),
		Ext:       faker.FileExtension(),
		ShareUUID: "",
	}
}
