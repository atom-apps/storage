package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Options struct {
	Timeout       *uint   `json:"timeout,omitempty" form:"timeout"`
	EnableMD5     *bool   `json:"enable_md5,omitempty" form:"enable_md5"`
	EnableCRC     *bool   `json:"enable_crc,omitempty" form:"enable_crc"`
	UseCname      *bool   `json:"use_cname,omitempty" form:"use_cname"`
	SecurityToken *string `json:"security_token,omitempty" form:"security_token"`
}

func (j Options) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *Options) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &j)
}
