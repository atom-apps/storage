package controller

import (
	"github.com/atom-apps/storage/modules/storages/dto"

	"github.com/gofiber/fiber/v2"
)

// @provider
type UploadController struct{}

// Upload
//
//	@Summary		Upload
//	@Tags			Storage
//	@Produce		json
//	@Success		200	{object}	dto.UploadResponse
//	@Router			/v1/storages/uploads [post]
func (c *UploadController) Upload(ctx *fiber.Ctx, id uint64) (*dto.UploadResponse, error) {
	return nil, nil
}
