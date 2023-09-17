package controller

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/atom-apps/storage/database/models"
	"github.com/atom-apps/storage/modules/storages/service"
	"github.com/atom-providers/jwt"
	"github.com/atom-providers/log"
	"github.com/atom-providers/uuid"

	"github.com/gofiber/fiber/v2"
)

// @provider
type UploadController struct {
	driver *service.DriverService
	fs     *service.FilesystemService
	uuid   *uuid.Generator
}

// Upload
//
//	@Summary		Upload
//	@Tags			Storage
//	@Produce		json
//	@Success		200	{object}	dto.UploadResponse
//	@Router			/v1/storages/uploads/{id} [post]
func (c *UploadController) Upload(ctx *fiber.Ctx, claim *jwt.Claims, id uint64) (*models.Filesystem, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	driver, err := c.driver.GetDefault(ctx.Context())
	if err != nil {
		return nil, err
	}

	dst, err := c.fs.GetPath(ctx.Context(), claim.TenantID, claim.UserID, id)
	if err != nil {
		return nil, err
	}

	uuid := c.uuid.MustGenerate()
	realName := fmt.Sprintf("%s%s", uuid, filepath.Ext(file.Filename))

	dst = filepath.Join(dst, realName)
	log.Infof("save file(%s) to path %s", file.Filename, dst)

	fd, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fd.Close()

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
		ParentID: id,
		Status:   0,
		Mime:     file.Header.Get("Content-Type"),
		Ext:      strings.Trim(filepath.Ext(file.Filename), "."),
		Size:     uint64(file.Size),
		Md5:      "",
	}
	if err := c.fs.CreateFromModel(ctx.Context(), model); err != nil {
		return nil, err
	}

	return model, err
}
