package controller

import (
	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/modules/storages/dto"
	"github.com/atom-apps/storage/modules/storages/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type DriverController struct {
	driverSvc *service.DriverService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DriverID"
//	@Success		200	{object}	dto.DriverItem
//	@Router			/v1/storages/drivers/{id} [get]
func (c *DriverController) Show(ctx *fiber.Ctx, id uint64) (*dto.DriverItem, error) {
	item, err := c.driverSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.driverSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.DriverListQueryFilter	true	"DriverListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.DriverItem}
//	@Router			/v1/storages/drivers [get]
func (c *DriverController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.DriverListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.driverSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.driverSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.DriverForm	true	"DriverForm"
//	@Success		200		{string}	DriverID
//	@Router			/v1/storages/drivers [post]
func (c *DriverController) Create(ctx *fiber.Ctx, body *dto.DriverForm) error {
	return c.driverSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"DriverID"
//	@Param			body	body		dto.DriverForm	true	"DriverForm"
//	@Success		200		{string}	DriverID
//	@Router			/v1/storages/drivers/{id} [put]
func (c *DriverController) Update(ctx *fiber.Ctx, id uint64, body *dto.DriverForm) error {
	return c.driverSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			Storage
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"DriverID"
//	@Success		200	{string}	DriverID
//	@Router			/v1/storages/drivers/{id} [delete]
func (c *DriverController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.driverSvc.Delete(ctx.Context(), id)
}
