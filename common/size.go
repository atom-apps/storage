package common

import "fmt"

type Size struct {
	Width  int
	Height int
}

func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s.Width, s.Height)
}

var (
	SizeWidth100 = Size{100, 0}
	SizeWidth200 = Size{200, 0}
	Size100x100  = Size{100, 100}
	Size100x200  = Size{100, 200}
	Size200x100  = Size{200, 100}
	Size200x200  = Size{200, 200}
)

func HumanReadableSize(size uint64) string {
	const unit = uint64(1024)
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := unit, 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
