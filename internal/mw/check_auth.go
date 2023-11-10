package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/pkg/core"
)

func CheckAuth(ctx context.Context, c *app.RequestContext) {
	loginUser, err := core.GetUserLoginState(c)
	if err != nil {
		core.SendResponse(c, err, nil)
		c.Abort()
		return
	}

	c.Set(consts.LoginUser, loginUser)

	c.Next(ctx)
}
