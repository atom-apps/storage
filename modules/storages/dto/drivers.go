package dto

import (
	"time"

	"github.com/atom-apps/storage/common"
)

type DriverForm struct {
	Name         string         `form:"name" json:"name,omitempty"`                   // 名称
	Endpoint     string         `form:"endpoint" json:"endpoint,omitempty"`           // 地址
	AccessKey    string         `form:"access_key" json:"access_key,omitempty"`       // AccessKey
	AccessSecret string         `form:"access_secret" json:"access_secret,omitempty"` // AccessSecret
	Bucket       string         `form:"bucket" json:"bucket,omitempty"`               // Bucket
	Options      common.Options `form:"options" json:"options,omitempty"`             // 配置
}

type DriverListQueryFilter struct {
	Name         *string `query:"name" json:"name,omitempty"`                   // 名称
	Endpoint     *string `query:"endpoint" json:"endpoint,omitempty"`           // 地址
	AccessKey    *string `query:"access_key" json:"access_key,omitempty"`       // AccessKey
	AccessSecret *string `query:"access_secret" json:"access_secret,omitempty"` // AccessSecret
	Bucket       *string `query:"bucket" json:"bucket,omitempty"`               // Bucket
}

type DriverItem struct {
	ID           uint64         `json:"id,omitempty"`            // ID
	CreatedAt    time.Time      `json:"created_at,omitempty"`    // 创建时间
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`    // 更新时间
	Name         string         `json:"name,omitempty"`          // 名称
	Endpoint     string         `json:"endpoint,omitempty"`      // 地址
	AccessKey    string         `json:"access_key,omitempty"`    // AccessKey
	AccessSecret string         `json:"access_secret,omitempty"` // AccessSecret
	Bucket       string         `json:"bucket,omitempty"`        // Bucket
	Options      common.Options `json:"options,omitempty"`       // 配置
}
