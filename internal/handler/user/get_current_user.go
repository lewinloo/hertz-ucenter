package user

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/internal/dal"
	"hertz-ucenter/internal/models/vo"
	"hertz-ucenter/internal/service"
	"hertz-ucenter/pkg/core"
	"hertz-ucenter/pkg/errno"
)

func (ctrl *cUser) getCurrentUser(ctx context.Context, c *app.RequestContext) {
	value, exists := c.Get(consts.LoginUser)
	loginUser := value.(*vo.UserVO)
	if !exists {
		core.SendResponse(c, errno.ErrUnauthorization, nil)
		return
	}
	user, err := dal.User.WithContext(ctx).Where(dal.User.ID.Eq(loginUser.ID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			core.SendResponse(c, errno.ErrEntityNull, nil)
			return
		}
		core.SendResponse(c, errno.ErrDBFailed, nil)
		return
	}
	userVO := service.User().GetUserVO(user)
	core.SendResponse(c, nil, userVO)
}
