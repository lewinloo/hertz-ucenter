package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/internal/models/dto"
	"hertz-ucenter/internal/service"
	"hertz-ucenter/pkg/core"
	"hertz-ucenter/pkg/errno"
)

func (ctrl *cUser) register(ctx context.Context, c *app.RequestContext) {
	var req dto.UserRegisterDTO
	if err := c.BindAndValidate(&req); err != nil {
		core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
		return
	}
	id, err := service.User().Register(ctx, req.Account, req.Password, req.CheckPassword, req.PlanetCode)
	if err != nil {
		core.SendResponse(c, err, nil)
		return
	}
	core.SendResponse(c, nil, id)
}
