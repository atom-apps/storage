package migrations

import (
	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230912_104027CreateFilesystem) table() interface{} {
	type Filesystem struct {
		Model
		ModelWithUser

		Driver    consts.FilesystemDriver   `gorm:"size:36;not null;comment:驱动"`
		Filename  string                    `gorm:"size:128;not null;index:idx_filename;comment:文件名"`
		Type      consts.Filesystem         `gorm:"size:12;not null;comment:类型"`
		ParentID  uint                      `gorm:"comment:父级ID"`
		Status    consts.FileStatus         `gorm:"comment:状态"`
		Mime      string                    `gorm:"size:128;index:idx_mime;comment:MIME"`
		ShareUUID string                    `gorm:"size:64;comment:共享ID"`
		Metadata  common.FilesystemMetadata `gorm:"comment:元数据"`
	}

	return &Filesystem{}
}

func (m *Migration20230912_104027CreateFilesystem) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230912_104027CreateFilesystem) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230912_104027CreateFilesystemMigration)
}

type Migration20230912_104027CreateFilesystem struct {
	id string
}

func New20230912_104027CreateFilesystemMigration() contracts.Migration {
	return &Migration20230912_104027CreateFilesystem{id: "20230912_104027_create_filesystem"}
}

func (m *Migration20230912_104027CreateFilesystem) ID() string {
	return m.id
}
