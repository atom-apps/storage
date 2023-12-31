package dto

import (
	"time"

	"github.com/atom-apps/storage/common"
)

type FilesystemForm struct {
	TenantID  uint64                    `form:"tenant_id" json:"tenant_id,omitempty"`   // 租户ID
	UserID    uint64                    `form:"user_id" json:"user_id,omitempty"`       // 用户ID
	Driver    string                    `form:"driver" json:"driver,omitempty"`         // 驱动
	Filename  string                    `form:"filename" json:"filename,omitempty"`     // 文件名
	Type      uint32                    `form:"type" json:"type,omitempty"`             // 类型
	ParentID  uint64                    `form:"parent_id" json:"parent_id,omitempty"`   // 父级ID
	Status    uint32                    `form:"status" json:"status,omitempty"`         // 状态
	Mime      string                    `form:"mime" json:"mime,omitempty"`             // MIME
	ShareUUID string                    `form:"share_uuid" json:"share_uuid,omitempty"` // 共享ID
	Metadata  common.FilesystemMetadata `form:"metadata" json:"metadata,omitempty"`     // 元数据
}

type FilesystemListQueryFilter struct {
	TenantID  *uint64                    `query:"tenant_id" json:"tenant_id,omitempty"`   // 租户ID
	UserID    *uint64                    `query:"user_id" json:"user_id,omitempty"`       // 用户ID
	DriverID  *uint64                    `query:"driver_id" json:"driver_id,omitempty"`   // 驱动
	Filename  *string                    `query:"filename" json:"filename,omitempty"`     // 文件名
	Type      *uint32                    `query:"type" json:"type,omitempty"`             // 类型
	ParentID  *uint64                    `query:"parent_id" json:"parent_id,omitempty"`   // 父级ID
	Status    *uint32                    `query:"status" json:"status,omitempty"`         // 状态
	Mime      *string                    `query:"mime" json:"mime,omitempty"`             // MIME
	ShareUUID *string                    `query:"share_uuid" json:"share_uuid,omitempty"` // 共享ID
	Metadata  *common.FilesystemMetadata `query:"metadata" json:"metadata,omitempty"`     // 元数据
}

type FilesystemItem struct {
	ID        uint64                    `json:"id,omitempty"`         // ID
	CreatedAt time.Time                 `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time                 `json:"updated_at,omitempty"` // 更新时间
	TenantID  uint64                    `json:"tenant_id,omitempty"`  // 租户ID
	UserID    uint64                    `json:"user_id,omitempty"`    // 用户ID
	DriverID  uint64                    `json:"driver_id,omitempty"`  // 驱动
	Filename  string                    `json:"filename,omitempty"`   // 文件名
	Type      uint32                    `json:"type"`                 // 类型
	ParentID  uint64                    `json:"parent_id,omitempty"`  // 父级ID
	Status    uint32                    `json:"status"`               // 状态
	Ext       string                    `json:"ext,omitempty"`        // 状态
	Mime      string                    `json:"mime,omitempty"`       // MIME
	ShareUUID string                    `json:"share_uuid,omitempty"` // 共享ID
	Metadata  common.FilesystemMetadata `json:"metadata,omitempty"`   // 元数据
	Children  []*FilesystemItem         `json:"children,omitempty"`   // 子级
}

type CreateSubDirectoryForm struct {
	Name string `json:"name" form:"name"`
}

type RealNamesForm struct {
	Names []string `json:"names" form:"names"`
}
