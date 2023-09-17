package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type FileInfoImage struct {
	Width  int
	Height int
}

type FilesystemMetadata struct {
	FileInfo      any     `json:"file_info,omitempty"`
	SharePassword *string `json:"share_password,omitempty" form:"share_password"`
	Thumbnail     *string `json:"thumbnail,omitempty"`
}

func (j FilesystemMetadata) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *FilesystemMetadata) Scan(value interface{}) error {
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
