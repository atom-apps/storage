package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atom-apps/storage/common"
	"github.com/disintegration/imaging"
)

// @provider
type ThumbnailService struct{}

func (svc *ThumbnailService) preparePath(ctx context.Context, path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

func (svc *ThumbnailService) Resize(ctx context.Context, src, dstFilename string, size common.Size) error {
	dstPath := fmt.Sprintf(
		"%s/thumb/%dx%d/%s.%s",
		filepath.Dir(src),
		size.Width,
		size.Height,
		dstFilename,
		strings.Trim(filepath.Ext(src), "."),
	)
	// file exists
	if _, err := os.Stat(dstPath); err == nil {
		return nil
	}

	if err := svc.preparePath(ctx, dstPath); err != nil {
		return err
	}

	srcImage, err := imaging.Open(src, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}
	img := imaging.Resize(srcImage, size.Width, size.Height, imaging.Lanczos)
	return imaging.Save(img, dstPath)
}
