package migrations

import (
	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230912_085743CreateDriver) table() interface{} {
	type Driver struct {
		Model
		Type         consts.FilesystemDriver `gorm:"size:64;not null;comment:类型"`
		Name         string                  `gorm:"size:128;not null;uniqueIndex:idx_name;comment:名称"`
		Endpoint     string                  `gorm:"size:198;not null;comment:地址"`
		AccessKey    string                  `gorm:"size:128;not null;comment:AccessKey"`
		AccessSecret string                  `gorm:"size:128;not null;comment:AccessSecret"`
		Bucket       string                  `gorm:"size:128;not null;comment:Bucket"`
		Options      common.Options          `gorm:"comment:配置"`
	}

	return &Driver{}
}

func (m *Migration20230912_085743CreateDriver) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230912_085743CreateDriver) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230912_085743CreateDriverMigration)
}

type Migration20230912_085743CreateDriver struct {
	id string
}

func New20230912_085743CreateDriverMigration() contracts.Migration {
	return &Migration20230912_085743CreateDriver{id: "20230912_085743_create_driver"}
}

func (m *Migration20230912_085743CreateDriver) ID() string {
	return m.id
}
