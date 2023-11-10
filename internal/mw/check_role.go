package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/internal/models/vo"
	"hertz-ucenter/pkg/core"
	"hertz-ucenter/pkg/errno"
)

func CheckRole(role int32) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		value, exists := c.Get(consts.LoginUser)
		loginUser := value.(*vo.UserVO)
		if !exists {
			core.SendResponse(c, errno.ErrUnauthorization, nil)
			c.Abort()
			return
		}

		if loginUser.UserRole != role {
			core.SendResponse(c, errno.ErrForbidden.SetDescription("权限不足"), nil)
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
