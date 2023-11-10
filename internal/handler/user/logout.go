package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/pkg/core"
)

func (ctrl *cUser) logout(ctx context.Context, c *app.RequestContext) {
	err := core.RemoveUserLoginState(c)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, true)
}
