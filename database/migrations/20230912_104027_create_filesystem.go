package migrations

import (
	"github.com/atom-apps/storage/common"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230912_104027CreateFilesystem) table() interface{} {
	type Filesystem struct {
		Model
		ModelWithUser
		DriverID  uint                      `gorm:"comment:驱动"`
		Filename  string                    `gorm:"size:128;not null;index:idx_filename;comment:文件名"`
		RealName  string                    `gorm:"size:128;not null;comment:真实文件名"`
		Type      uint                      `gorm:"size:1;not null;comment:类型"`
		ParentID  uint                      `gorm:"comment:父级ID"`
		Status    uint                      `gorm:"size:1;comment:状态"`
		Mime      string                    `gorm:"size:256;index:idx_mime;comment:MIME"`
		Ext       string                    `gorm:"size:32;index:idx_ext;comment:后缀名"`
		Size      uint                      `gorm:"comment:文件大小"`
		Md5       string                    `gorm:"size:32;comment:MD5"`
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
