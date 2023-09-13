package local

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/database/models"
)

var _ (common.Storage) = (*LocalDriver)(nil)

type LocalDriver struct {
	model *models.Driver
}

func New(params *models.Driver) common.Storage {
	return &LocalDriver{model: params}
}

func (driver *LocalDriver) ID() uint64 {
	return driver.model.ID
}

// Append implements common.Storage.
func (*LocalDriver) Append(string, io.Reader, int64) (string, error) {
	panic("unimplemented")
}

// Copy implements common.Storage.
func (*LocalDriver) Copy(string, string) error {
	panic("unimplemented")
}

// Delete implements common.Storage.
func (*LocalDriver) Delete([]string) error {
	panic("unimplemented")
}

// Exists implements common.Storage.
func (*LocalDriver) Exists(path string) (bool, error) {
	panic("unimplemented")
}

// Get implements common.Storage.
func (*LocalDriver) Get(string) (io.ReadCloser, error) {
	panic("unimplemented")
}

// GetURL implements common.Storage.
func (*LocalDriver) GetURL(string) (string, error) {
	panic("unimplemented")
}

// IsDir implements common.Storage.
func (*LocalDriver) IsDir(path string) (bool, error) {
	panic("unimplemented")
}

// IsFile implements common.Storage.
func (*LocalDriver) IsFile(path string) (bool, error) {
	panic("unimplemented")
}

// List implements common.Storage.
func (*LocalDriver) List(dir string, marker string, limit uint) ([]string, string, error) {
	panic("unimplemented")
}

// Move implements common.Storage.
func (*LocalDriver) Move(string, string) error {
	panic("unimplemented")
}

// Put implements common.Storage.
func (driver *LocalDriver) Put(dst string, in io.Reader) (string, error) {
	dst = filepath.Join(driver.model.Bucket, dst)
	if !strings.HasPrefix(dst, driver.model.Bucket) {
		return "", errors.New("invalid save path")
	}

	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return "", errors.New("create dir failed, err:" + err.Error())
	}

	fd, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	if _, err := io.Copy(fd, in); err != nil {
		return "", err
	}

	return dst[len(driver.model.Bucket):], nil
}

// PutFile implements common.Storage.
func (*LocalDriver) PutFile(string, string) (string, error) {
	panic("unimplemented")
}
