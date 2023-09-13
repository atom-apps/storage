package seeders

import (
	"log"

	"github.com/atom-apps/storage/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type FilesystemSeeder struct{}

func NewFilesystemSeeder() contracts.Seeder {
	return &FilesystemSeeder{}
}

func (s *FilesystemSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Filesystem{}).TableName(nil))
	times := 100
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
	parentID := 0
	if idx > 3 {
		parentID = faker.Number(1, idx-1)
	}

	typ := 0
	if idx > 10 {
		typ = 1
		parentID = faker.Number(2, 10)
	}

	return models.Filesystem{
		CreatedAt: faker.Date(),
		UpdatedAt: faker.Date(),
		TenantID:  1,
		UserID:    1,
		DriverID:  1,
		Filename:  faker.AppName(),
		Type:      uint32(typ),
		ParentID:  uint64(parentID),
		Status:    uint32(faker.RandomInt([]int{0, 1})),
		Mime:      faker.FileMimeType(),
		Ext:       faker.FileExtension(),
		Size:      uint64(faker.Number(10, 10000)),
		Md5:       faker.UUID()[0:12],
		ShareUUID: "",
	}
}
