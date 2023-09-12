package boot

import (
	"github.com/atom-apps/storage/common"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/fen"
)

func httpMiddlewareJWT(j *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, err := common.GetJwtToken(ctx)
		if err != nil {
			return err
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals(fen.JwtCtxKey, claims)

		return ctx.Next()
	}
}
