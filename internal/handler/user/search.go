package user

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "hertz-ucenter/internal/models/dto"
  "hertz-ucenter/internal/models/vo"
  "hertz-ucenter/internal/service"
  "hertz-ucenter/pkg/core"
  "hertz-ucenter/pkg/errno"
)

func (ctrl *cUser) search(ctx context.Context, c *app.RequestContext) {
  var req dto.UserSearchQuery
  if err := c.BindAndValidate(&req); err != nil {
    core.SendResponse(c, errno.ErrParameterInvalid.SetDescription(err.Error()), nil)
    return
  }

  list, total, err := service.User().Search(ctx, req)
  if err != nil {
    core.SendResponse(c, err, nil)
    return
  }

  core.SendResponse(c, nil, vo.PageResult{
    Records: list,
    Total:   total,
  })
}
