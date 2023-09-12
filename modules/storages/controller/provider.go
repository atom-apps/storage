package controller

import (
	"github.com/atom-apps/storage/modules/storages/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		driverSvc *service.DriverService,
	) (*DriverController, error) {
		obj := &DriverController{
			driverSvc: driverSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		filesystemSvc *service.FilesystemService,
	) (*FilesystemController, error) {
		obj := &FilesystemController{
			filesystemSvc: filesystemSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
