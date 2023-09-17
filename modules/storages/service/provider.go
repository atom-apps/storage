package service

import (
	"github.com/atom-apps/storage/modules/storages/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		driverDao *dao.DriverDao,
	) (*DriverService, error) {
		obj := &DriverService{
			driverDao: driverDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		filesystemDao *dao.FilesystemDao,
	) (*FilesystemService, error) {
		obj := &FilesystemService{
			filesystemDao: filesystemDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func() (*ThumbnailService, error) {
		obj := &ThumbnailService{}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
