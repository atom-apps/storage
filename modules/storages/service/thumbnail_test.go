package service

import (
	"context"
	"testing"

	"github.com/atom-apps/storage/common"
	"github.com/rogeecn/fabfile"
	. "github.com/smartystreets/goconvey/convey"
)

func TestThumbnailService_Resize(t *testing.T) {
	Convey("Test_Resize", t, func() {
		sizes := []common.Size{
			common.SizeWidth100,
			common.SizeWidth200,
			common.Size100x100,
			common.Size100x200,
			common.Size200x100,
			common.Size200x200,
		}
		for _, size := range sizes {
			ins := &ThumbnailService{}
			err := ins.Resize(context.TODO(), fabfile.MustFind("tests/images/jpeg.jpeg"), "a", size)
			So(err, ShouldBeNil)
		}
	})
}
