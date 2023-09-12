package storages

import (
	"github.com/atom-apps/storage/modules/storages/controller"
	"github.com/atom-apps/storage/modules/storages/dao"
	"github.com/atom-apps/storage/modules/storages/routes"
	"github.com/atom-apps/storage/modules/storages/service"

	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: dao.Provide},
		{Provider: service.Provide},
		{Provider: controller.Provide},
		{Provider: routes.Provide},
	}
}
