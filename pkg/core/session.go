package core

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/internal/models/vo"
	"hertz-ucenter/pkg/errno"
)

func SetUserLoginState(c *app.RequestContext, data any) error {
	session := sessions.Default(c)
	sessionTTL := 86400
	session.Set(consts.UserLoginState, data)
	// 设置过期时间
	session.Options(sessions.Options{MaxAge: sessionTTL})
	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserLoginState(c *app.RequestContext) error {
	session := sessions.Default(c)
	session.Delete(consts.UserLoginState)
	session.Options(sessions.Options{MaxAge: -1})
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetUserLoginState(c *app.RequestContext) (*vo.UserVO, error) {
	session := sessions.Default(c)
	obj := session.Get(consts.UserLoginState)
	if obj == nil {
		return nil, errno.ErrUnauthorization.SetDescription("未登录")
	}
	return obj.(*vo.UserVO), nil
}
