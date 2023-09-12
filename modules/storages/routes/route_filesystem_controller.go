// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/storage/modules/storages/controller"
	"github.com/atom-apps/storage/modules/storages/dto"
	"github.com/atom-apps/storage/common"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeFilesystemController(engine fiber.Router, controller *controller.FilesystemController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", basePath), DataFunc1(controller.Show, Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/storages/filesystems", basePath), DataFunc3(controller.List, Query[dto.FilesystemListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems", basePath), Func1(controller.Create, Body[dto.FilesystemForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", basePath), Func2(controller.Update, Integer[uint64]("id", PathParamError), Body[dto.FilesystemForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", basePath), Func1(controller.Delete, Integer[uint64]("id", PathParamError)))
}
