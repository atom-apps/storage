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
type DriverDao struct {
	query *query.Query
}

func (dao *DriverDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *DriverDao) Context(ctx context.Context) query.IDriverDo {
	return dao.query.Driver.WithContext(ctx)
}

func (dao *DriverDao) decorateSortQueryFilter(query query.IDriverDo, sortFilter *common.SortQueryFilter) query.IDriverDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Driver.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Driver.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *DriverDao) decorateQueryFilter(query query.IDriverDo, queryFilter *dto.DriverListQueryFilter) query.IDriverDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.Driver.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Endpoint != nil {
		query = query.Where(dao.query.Driver.Endpoint.Eq(*queryFilter.Endpoint))
	}
	if queryFilter.AccessKey != nil {
		query = query.Where(dao.query.Driver.AccessKey.Eq(*queryFilter.AccessKey))
	}
	if queryFilter.AccessSecret != nil {
		query = query.Where(dao.query.Driver.AccessSecret.Eq(*queryFilter.AccessSecret))
	}
	if queryFilter.Bucket != nil {
		query = query.Where(dao.query.Driver.Bucket.Eq(*queryFilter.Bucket))
	}

	return query
}

func (dao *DriverDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Driver.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *DriverDao) Update(ctx context.Context, model *models.Driver) error {
	_, err := dao.Context(ctx).Where(dao.query.Driver.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *DriverDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Driver.ID.Eq(id)).Delete()
	return err
}

func (dao *DriverDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Driver.ID.Eq(id)).Delete()
	return err
}

func (dao *DriverDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Driver.ID.Eq(id)).UpdateSimple(dao.query.Driver.DeletedAt.Null())
	return err
}

func (dao *DriverDao) Create(ctx context.Context, model *models.Driver) error {
	return dao.Context(ctx).Create(model)
}

func (dao *DriverDao) GetByID(ctx context.Context, id uint64) (*models.Driver, error) {
	return dao.Context(ctx).Where(dao.query.Driver.ID.Eq(id)).First()
}

func (dao *DriverDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Driver, error) {
	return dao.Context(ctx).Where(dao.query.Driver.ID.In(ids...)).Find()
}

func (dao *DriverDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DriverListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Driver, int64, error) {
	query := dao.query.Driver
	driverQuery := query.WithContext(ctx)
	driverQuery = dao.decorateQueryFilter(driverQuery, queryFilter)
	driverQuery = dao.decorateSortQueryFilter(driverQuery, sortFilter)
	return driverQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *DriverDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DriverListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Driver, error) {
	query := dao.query.Driver
	driverQuery := query.WithContext(ctx)
	driverQuery = dao.decorateQueryFilter(driverQuery, queryFilter)
	driverQuery = dao.decorateSortQueryFilter(driverQuery, sortFilter)
	return driverQuery.Find()
}

func (dao *DriverDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.DriverListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Driver, error) {
	query := dao.query.Driver
	driverQuery := query.WithContext(ctx)
	driverQuery = dao.decorateQueryFilter(driverQuery, queryFilter)
	driverQuery = dao.decorateSortQueryFilter(driverQuery, sortFilter)
	return driverQuery.First()
}
