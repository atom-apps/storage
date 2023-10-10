package controller

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/atom-apps/storage/database/models"
	"github.com/atom-apps/storage/modules/storages/dto"
	"github.com/atom-apps/storage/modules/storages/service"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/atom-providers/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
)

// @provider
type UploadController struct {
	driver *service.DriverService
	fs     *service.FilesystemService
	uuid   *uuid.Generator
}

// Dir files to dir
//
//	@Summary		FileToDir
//	@Tags			Storage
//	@Produce		json
//	@Success		200	{object}	dto.FilesystemItem
//	@Router			/v1/storages/uploads/dir/{id} [post]
func (c *UploadController) Dir(ctx *fiber.Ctx, claim *jwt.Claims, id uint64) (*dto.FilesystemItem, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	dst, err := c.fs.GetPathByID(ctx.Context(), claim.TenantID, claim.UserID, id)
	if err != nil {
		return nil, err
	}
	return c.save(ctx.Context(), claim, id, file, dst)
}

// Posts Upload files to posts dir
//
//	@Summary		FileToPosts
//	@Tags			Storage
//	@Produce		json
//	@Success		200	{object}	dto.FilesystemItem
//	@Router			/v1/storages/uploads/posts [post]
func (c *UploadController) Posts(ctx *fiber.Ctx, claim *jwt.Claims) (*dto.FilesystemItem, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	now := carbon.Now()

	dir := fmt.Sprintf("posts/%d/%d/%d", now.Year(), now.Month(), now.Day())

	fs, err := c.fs.CreateDir(ctx.Context(), claim.TenantID, claim.UserID, dir)
	if err != nil {
		return nil, err
	}

	return c.save(ctx.Context(), claim, fs.ID, file, dir)
}

func (c *UploadController) save(ctx context.Context, claim *jwt.Claims, parentID uint64, file *multipart.FileHeader, dst string) (*dto.FilesystemItem, error) {
	uuid := c.uuid.MustGenerate()
	realName := fmt.Sprintf("%s%s", uuid, filepath.Ext(file.Filename))

	dst = filepath.Join(dst, realName)
	log.Infof("save file(%s) to path %s", file.Filename, dst)

	fd, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	driver, err := c.driver.GetDefault(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := driver.Put(dst, fd); err != nil {
		return nil, err
	}

	model := &models.Filesystem{
		TenantID: claim.TenantID,
		UserID:   claim.UserID,
		DriverID: driver.ID(),
		Filename: file.Filename[0 : len(file.Filename)-len(filepath.Ext(file.Filename))],
		RealName: uuid,
		Type:     1,
		ParentID: parentID,
		Status:   0,
		Mime:     file.Header.Get("Content-Type"),
		Ext:      strings.Trim(filepath.Ext(file.Filename), "."),
		Size:     uint64(file.Size),
		Md5:      "",
	}
	if err := c.fs.CreateFromModel(ctx, model); err != nil {
		return nil, err
	}

	return c.fs.DecorateItem(model, 0), err
}
