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
type FilesystemService struct {
	filesystemDao *dao.FilesystemDao
}

func (svc *FilesystemService) DecorateItem(model *models.Filesystem, id int) *dto.FilesystemItem {
	return &dto.FilesystemItem{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		TenantID:  model.TenantID,
		UserID:    model.UserID,
		DriverID:  model.DriverID,
		Filename:  model.Filename,
		Type:      model.Type,
		ParentID:  model.ParentID,
		Status:    model.Status,
		Mime:      model.Mime,
		ShareUUID: model.ShareUUID,
		Metadata:  model.Metadata,
	}
}

func (svc *FilesystemService) GetByID(ctx context.Context, id uint64) (*models.Filesystem, error) {
	return svc.filesystemDao.GetByID(ctx, id)
}

func (svc *FilesystemService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.FilesystemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Filesystem, error) {
	return svc.filesystemDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *FilesystemService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.FilesystemListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Filesystem, int64, error) {
	return svc.filesystemDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *FilesystemService) CreateFromModel(ctx context.Context, model *models.Filesystem) error {
	return svc.filesystemDao.Create(ctx, model)
}

// Create
func (svc *FilesystemService) Create(ctx context.Context, body *dto.FilesystemForm) error {
	model := &models.Filesystem{}
	_ = copier.Copy(model, body)
	return svc.filesystemDao.Create(ctx, model)
}

// Update
func (svc *FilesystemService) Update(ctx context.Context, id uint64, body *dto.FilesystemForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.filesystemDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *FilesystemService) UpdateFromModel(ctx context.Context, model *models.Filesystem) error {
	return svc.filesystemDao.Update(ctx, model)
}

// Delete
func (svc *FilesystemService) Delete(ctx context.Context, id uint64) error {
	return svc.filesystemDao.Delete(ctx, id)
}
