package user

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "hertz-ucenter/internal/models/dto"
  "hertz-ucenter/internal/service"
  "hertz-ucenter/pkg/core"
  "hertz-ucenter/pkg/errno"
)

func (ctrl *cUser) delete(ctx context.Context, c *app.RequestContext) {
  var req dto.IdInPathDTO
  if err := c.BindAndValidate(&req); err != nil {
    core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
    return
  }

  err := service.User().Delete(ctx, req.ID)
  if err != nil {
    core.SendResponse(c, err, false)
    return
  }

  core.SendResponse(c, nil, true)
}
