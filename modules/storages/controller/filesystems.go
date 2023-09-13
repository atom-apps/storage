package controller

import (
	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/modules/storages/dto"
	"github.com/atom-apps/storage/modules/storages/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type FilesystemController struct {
	filesystemSvc *service.FilesystemService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FilesystemID"
//	@Success		200	{object}	dto.FilesystemItem
//	@Router			/v1/storages/filesystems/{id} [get]
func (c *FilesystemController) Show(ctx *fiber.Ctx, id uint64) (*dto.FilesystemItem, error) {
	item, err := c.filesystemSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.filesystemSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.FilesystemListQueryFilter	true	"FilesystemListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.FilesystemItem}
//	@Router			/v1/storages/filesystems [get]
func (c *FilesystemController) List(
	ctx *fiber.Ctx,
	claim *jwt.Claims,
	queryFilter *dto.FilesystemListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	if !claim.IsAdmin() {
		queryFilter.UserID = &claim.UserID
		queryFilter.TenantID = &claim.TenantID
	}
	queryFilter.Type = lo.ToPtr[uint32](1)

	items, total, err := c.filesystemSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.filesystemSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.FilesystemForm	true	"FilesystemForm"
//	@Success		200		{string}	FilesystemID
//	@Router			/v1/storages/filesystems [post]
func (c *FilesystemController) Create(ctx *fiber.Ctx, body *dto.FilesystemForm) error {
	return c.filesystemSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"FilesystemID"
//	@Param			body	body		dto.FilesystemForm	true	"FilesystemForm"
//	@Success		200		{string}	FilesystemID
//	@Router			/v1/storages/filesystems/{id} [put]
func (c *FilesystemController) Update(ctx *fiber.Ctx, id uint64, body *dto.FilesystemForm) error {
	return c.filesystemSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FilesystemID"
//	@Success		200	{string}	FilesystemID
//	@Router			/v1/storages/filesystems/{id} [delete]
func (c *FilesystemController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.filesystemSvc.Delete(ctx.Context(), id)
}

// Directory
//
//	@Summary		Directory
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FilesystemID"
//	@Success		200	{string}	FilesystemID
//	@Router			/v1/storages/filesystems/{id}/directory/{directory} [post]
func (c *FilesystemController) Directory(ctx *fiber.Ctx, claim *jwt.Claims, id uint64, directory string) error {
	return c.filesystemSvc.CreateSubDirectory(ctx.Context(), claim.TenantID, claim.UserID, id, directory)
}

// 目录列表

// DirectoryTree
//
//	@Summary		Directory
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	FilesystemID
//	@Router			/v1/storages/filesystems/directories/tree [get]
func (c *FilesystemController) DirectoryTree(ctx *fiber.Ctx, claim *jwt.Claims) ([]*dto.FilesystemItem, error) {
	return c.filesystemSvc.GetDirectoryTree(ctx.Context(), claim.TenantID, claim.UserID)
}

// 目录列表
