package dao

import (
	"github.com/atom-apps/storage/database/query"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		query *query.Query,
	) (*DriverDao, error) {
		obj := &DriverDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*FilesystemDao, error) {
		obj := &FilesystemDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
