package errorx

import (
	"net/http"

	"github.com/rogeecn/fen"
)

var (
	ErrForbidden         = fen.NewBusError(http.StatusForbidden, http.StatusForbidden, "拒绝操作")
	ErrInvalidRequest    = fen.NewBusError(http.StatusBadRequest, http.StatusBadGateway, "无效请求")
	ErrInvalidVerifyCode = fen.NewBusError(http.StatusBadRequest, http.StatusBadGateway, "验证码错误")
	ErrRecordExists      = fen.NewBusError(http.StatusConflict, http.StatusConflict, "记录已存在")
)
