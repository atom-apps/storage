package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/common/consts"
	"github.com/atom-apps/storage/common/storages/local"
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
	return &dto.DriverItem{
		ID:           model.ID,
		Name:         model.Name,
		Endpoint:     model.Endpoint,
		AccessKey:    model.AccessKey,
		AccessSecret: svc.HideSecret(model.AccessSecret),
		Bucket:       model.Bucket,
		Options:      model.Options,
	}
}

func (svc *DriverService) HideSecret(secret string) string {
	if len(secret) < 3 {
		return "***"
	}
	return secret[0:len(secret)-3] + "***"
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

func (svc *DriverService) GetDefault(ctx context.Context) (common.Storage, error) {
	driver, err := svc.driverDao.GetDefaultDriver(ctx)
	if err != nil {
		return nil, err
	}

	switch driver.Type {
	case consts.FilesystemDriverLocal:
		return local.New(driver), nil
	case consts.FilesystemDriverAliyunOss:
		return nil, nil
	}

	return nil, errors.New("no valid storage driver")
}

func (svc *DriverService) GetHostFromDriver(ctx context.Context, driver *models.Driver) string {
	host := ""
	switch driver.Type {
	case consts.FilesystemDriverLocal:
		if strings.HasPrefix(driver.Endpoint, "http") {
			host = driver.Endpoint
		} else {
			host = fmt.Sprintf("//%s", driver.Endpoint)
		}
	default:
		if strings.HasPrefix(driver.Endpoint, "http") {
			host = fmt.Sprintf("//%s/%s", driver.Endpoint, driver.Bucket)
		} else {
			host = fmt.Sprintf("%s/%s", driver.Endpoint, driver.Bucket)
		}
	}

	return strings.TrimRight(host, "/")
}
