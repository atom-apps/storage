package service

import (
	"context"

	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/database/models"
	"github.com/atom-apps/storage/modules/storages/dao"
	"github.com/atom-apps/storage/modules/storages/dto"

	"github.com/jinzhu/copier"
)

// @provider
type DriverService struct {
	driverDao *dao.DriverDao
}

func (svc *DriverService) DecorateItem(model *models.Driver, id int) *dto.DriverItem {
	var dtoItem *dto.DriverItem
	_ = copier.Copy(dtoItem, model)

	return dtoItem
}

func (svc *DriverService) GetByID(ctx context.Context, id uint64) (*models.Driver, error) {
	return svc.driverDao.GetByID(ctx, id)
}

func (svc *DriverService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DriverListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Driver, error) {
	return svc.driverDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *DriverService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DriverListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Driver, int64, error) {
	return svc.driverDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *DriverService) CreateFromModel(ctx context.Context, model *models.Driver) error {
	return svc.driverDao.Create(ctx, model)
}

// Create
func (svc *DriverService) Create(ctx context.Context, body *dto.DriverForm) error {
	model := &models.Driver{}
	_ = copier.Copy(model, body)
	return svc.driverDao.Create(ctx, model)
}

// Update
func (svc *DriverService) Update(ctx context.Context, id uint64, body *dto.DriverForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.driverDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *DriverService) UpdateFromModel(ctx context.Context, model *models.Driver) error {
	return svc.driverDao.Update(ctx, model)
}

// Delete
func (svc *DriverService) Delete(ctx context.Context, id uint64) error {
	return svc.driverDao.Delete(ctx, id)
}
