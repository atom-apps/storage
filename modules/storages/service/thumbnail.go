package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atom-apps/storage/common"
	"github.com/disintegration/imaging"
)

// @provider
type ThumbnailService struct{}

func (svc *ThumbnailService) preparePath(ctx context.Context, path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

func (svc *ThumbnailService) Resize(ctx context.Context, src string, size common.Size) error {
	srcImage, err := imaging.Open(src, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}
	img := imaging.Resize(srcImage, size.Width, size.Height, imaging.Lanczos)
	dstPath := fmt.Sprintf("%s/thumb/%dx%d/%s", filepath.Dir(src), size.Width, size.Height, filepath.Base(src))
	if err := svc.preparePath(ctx, dstPath); err != nil {
		return err
	}
	return imaging.Save(img, dstPath)
}
