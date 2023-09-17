package common

import "github.com/samber/lo"

var imageMiME = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
	"image/bmp",
	"image/webp",
	"image/tiff",
	"image/svg+xml",
	"image/x-icon",
	"image/vnd.microsoft.icon",
	"image/heif",
	"image/heif-sequence",
	"image/heic",
	"image/heic-sequence",
	"image/avif",
	"image/jp2",
}

type Mime struct {
	mime string
}

func NewMime(mime string) *Mime {
	return &Mime{mime: mime}
}

// IsImage
func (m *Mime) IsImage() bool {
	return lo.Contains(imageMiME, m.mime)
}
