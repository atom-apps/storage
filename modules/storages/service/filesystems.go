package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/database/models"
	"github.com/atom-apps/storage/modules/storages/dao"
	"github.com/atom-apps/storage/modules/storages/dto"
	"github.com/pkg/errors"
	"github.com/samber/lo"

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
		Ext:       model.Ext,
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

// CreateSubDirectory
func (svc *FilesystemService) CreateSubDirectory(ctx context.Context, tenantID, userID, parentID uint64, name string) error {
	names := strings.Split(name, "/")

	for _, name := range names {
		dirs := []string{name}
		changeParentID := true
		if strings.HasPrefix(name, "{") && strings.HasSuffix(name, "}") {
			dirs = strings.Split(strings.Trim(name, "{}"), ",")
			changeParentID = false
		}
		for _, dir := range dirs {
			model := &models.Filesystem{
				TenantID:  tenantID,
				UserID:    userID,
				DriverID:  0,
				Filename:  dir,
				Type:      0,
				ParentID:  parentID,
				Status:    0,
				Mime:      "",
				Ext:       "",
				ShareUUID: "",
			}
			if err := svc.CreateFromModel(ctx, model); err != nil {
				return err
			}

			if changeParentID {
				parentID = model.ID
			}
		}

	}
	return nil
}

// GetDirectoryTree
func (svc *FilesystemService) GetDirectoryTree(ctx context.Context, tenantID, userID uint64) ([]*dto.FilesystemItem, error) {
	queryFilter := &dto.FilesystemListQueryFilter{TenantID: &tenantID, UserID: &userID, Type: lo.ToPtr[uint32](0)}
	items, err := svc.FindByQueryFilter(ctx, queryFilter, &common.SortQueryFilter{})
	if err != nil {
		return nil, err
	}

	var result []*dto.FilesystemItem
	for _, item := range items {
		result = append(result, svc.DecorateItem(item, 0))
	}

	return svc.genTree(result, 0), nil
}

func (svc *FilesystemService) genTree(items []*dto.FilesystemItem, parentID uint64) []*dto.FilesystemItem {
	var result []*dto.FilesystemItem
	for _, item := range items {
		if item.ParentID == parentID {
			item.Children = svc.genTree(items, item.ID)
			result = append(result, item)
		}
	}
	return result
}

func (svc *FilesystemService) GetByIDWithTenantInfo(ctx context.Context, tenantID, userID, id uint64) (*models.Filesystem, error) {
	return svc.filesystemDao.GetByIDWithTenantInfo(ctx, tenantID, userID, id)
}

func (svc *FilesystemService) GetPath(ctx context.Context, tenantID, userID, id uint64) (string, error) {
	if id == 0 {
		return "", nil
	}

	paths := []string{}
	for {
		m, err := svc.GetByIDWithTenantInfo(ctx, tenantID, userID, id)
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("%d", id))
		}
		paths = append(paths, m.Filename)
		if m.ParentID == 0 {
			break
		}
		id = m.ParentID
	}

	return filepath.Join(lo.Reverse(paths)...), nil
}
