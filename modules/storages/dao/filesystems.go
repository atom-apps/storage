package dao

import (
	"context"

	"github.com/atom-apps/storage/common"
	"github.com/atom-apps/storage/database/models"
	"github.com/atom-apps/storage/database/query"
	"github.com/atom-apps/storage/modules/storages/dto"

	"gorm.io/gen/field"
)

// @provider
type FilesystemDao struct {
	query *query.Query
}

func (dao *FilesystemDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *FilesystemDao) Context(ctx context.Context) query.IFilesystemDo {
	return dao.query.Filesystem.WithContext(ctx)
}

func (dao *FilesystemDao) decorateSortQueryFilter(query query.IFilesystemDo, sortFilter *common.SortQueryFilter) query.IFilesystemDo {
	if sortFilter == nil {
		return query
	}

	hasCreatedAt := false
	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if v == "type" {
			continue
		}

		if v == "created_at" {
			hasCreatedAt = true
		}

		if expr, ok := dao.query.Filesystem.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if v == "type" {
			continue
		}
		if v == "created_at" {
			hasCreatedAt = true
		}
		if expr, ok := dao.query.Filesystem.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}

	orderExprs = append(orderExprs, dao.query.Filesystem.Type)
	if !hasCreatedAt {
		orderExprs = append(orderExprs, dao.query.Filesystem.CreatedAt)
	}

	return query.Order(orderExprs...)
}

func (dao *FilesystemDao) decorateQueryFilter(query query.IFilesystemDo, queryFilter *dto.FilesystemListQueryFilter) query.IFilesystemDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Filesystem.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Filesystem.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.DriverID != nil {
		query = query.Where(dao.query.Filesystem.DriverID.Eq(*queryFilter.DriverID))
	}
	if queryFilter.Filename != nil {
		query = query.Where(dao.query.Filesystem.Filename.Eq(*queryFilter.Filename))
	}
	if queryFilter.Type != nil {
		query = query.Where(dao.query.Filesystem.Type.Eq(*queryFilter.Type))
	}
	if queryFilter.ParentID != nil {
		query = query.Where(dao.query.Filesystem.ParentID.Eq(*queryFilter.ParentID))
	}
	if queryFilter.Status != nil {
		query = query.Where(dao.query.Filesystem.Status.Eq(*queryFilter.Status))
	}
	if queryFilter.Mime != nil {
		query = query.Where(dao.query.Filesystem.Mime.Eq(*queryFilter.Mime))
	}
	if queryFilter.ShareUUID != nil {
		query = query.Where(dao.query.Filesystem.ShareUUID.Eq(*queryFilter.ShareUUID))
	}

	return query
}

func (dao *FilesystemDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Filesystem.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *FilesystemDao) Update(ctx context.Context, model *models.Filesystem) error {
	_, err := dao.Context(ctx).Where(dao.query.Filesystem.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *FilesystemDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Filesystem.ID.Eq(id)).Delete()
	return err
}

func (dao *FilesystemDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Filesystem.ID.Eq(id)).Delete()
	return err
}

func (dao *FilesystemDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Filesystem.ID.Eq(id)).UpdateSimple(dao.query.Filesystem.DeletedAt.Null())
	return err
}

func (dao *FilesystemDao) Create(ctx context.Context, model *models.Filesystem) error {
	return dao.Context(ctx).Create(model)
}

func (dao *FilesystemDao) CreateInBatch(ctx context.Context, models []*models.Filesystem) error {
	return dao.Context(ctx).CreateInBatches(models, 100)
}

func (dao *FilesystemDao) GetByID(ctx context.Context, id uint64) (*models.Filesystem, error) {
	return dao.Context(ctx).Where(dao.query.Filesystem.ID.Eq(id)).First()
}

func (dao *FilesystemDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Filesystem, error) {
	return dao.Context(ctx).Where(dao.query.Filesystem.ID.In(ids...)).Find()
}

func (dao *FilesystemDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.FilesystemListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Filesystem, int64, error) {
	query := dao.query.Filesystem
	filesystemQuery := query.WithContext(ctx)
	filesystemQuery = dao.decorateQueryFilter(filesystemQuery, queryFilter)
	filesystemQuery = dao.decorateSortQueryFilter(filesystemQuery, sortFilter)
	return filesystemQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *FilesystemDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.FilesystemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Filesystem, error) {
	query := dao.query.Filesystem
	filesystemQuery := query.WithContext(ctx)
	filesystemQuery = dao.decorateQueryFilter(filesystemQuery, queryFilter)
	filesystemQuery = dao.decorateSortQueryFilter(filesystemQuery, sortFilter)
	return filesystemQuery.Find()
}

func (dao *FilesystemDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.FilesystemListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Filesystem, error) {
	query := dao.query.Filesystem
	filesystemQuery := query.WithContext(ctx)
	filesystemQuery = dao.decorateQueryFilter(filesystemQuery, queryFilter)
	filesystemQuery = dao.decorateSortQueryFilter(filesystemQuery, sortFilter)
	return filesystemQuery.First()
}

func (dao *FilesystemDao) GetByIDWithTenantInfo(ctx context.Context, tenantID, userID, id uint64) (*models.Filesystem, error) {
	query, table := dao.Context(ctx), dao.query.Filesystem
	return query.Where(
		table.TenantID.Eq(tenantID),
		table.UserID.Eq(userID),
		table.ID.Eq(id),
	).First()
}

func (dao *FilesystemDao) GetByIDsWithTenantInfo(ctx context.Context, tenantID, userID uint64, id []uint64) ([]*models.Filesystem, error) {
	query, table := dao.Context(ctx), dao.query.Filesystem
	return query.Where(
		table.TenantID.Eq(tenantID),
		table.UserID.Eq(userID),
		table.ID.In(id...),
	).Find()
}

func (dao *FilesystemDao) MoveFiles(ctx context.Context, tenantID, userID, parentID uint64, files []uint64) error {
	query, table := dao.Context(ctx), dao.query.Filesystem
	_, err := query.Where(
		table.TenantID.Eq(tenantID),
		table.UserID.Eq(userID),
		table.ID.In(files...),
	).Update(table.ParentID, parentID)
	return err
}

// GetByRealNames
func (dao *FilesystemDao) GetByRealNames(ctx context.Context, realNames []string) ([]*models.Filesystem, error) {
	query, table := dao.Context(ctx), dao.query.Filesystem
	return query.Where(table.RealName.In(realNames...)).Find()
}

// GetByNameOfParent
func (dao *FilesystemDao) GetByNameOfParent(ctx context.Context, tenantID, userID, parentID uint64, name string) (*models.Filesystem, error) {
	query, table := dao.Context(ctx), dao.query.Filesystem
	return query.Where(
		table.ParentID.Eq(parentID),
		table.Filename.Eq(name),
	).First()
}
