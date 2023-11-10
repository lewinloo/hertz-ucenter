package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/internal/models/dto"
	"hertz-ucenter/internal/service"
	"hertz-ucenter/pkg/core"
	"hertz-ucenter/pkg/errno"
)

func (ctrl *cUser) login(ctx context.Context, c *app.RequestContext) {
	var req dto.UserLoginDTO
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}
	userVO, err := service.User().Login(ctx, req.Account, req.Password)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	err = core.SetUserLoginState(c, &userVO)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, userVO)
}
