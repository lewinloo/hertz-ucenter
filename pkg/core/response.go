package core

import (
	"github.com/cloudwego/hertz/pkg/app"
	"hertz-ucenter/pkg/errno"
	"hertz-ucenter/pkg/resputil"
)

func SendResponse(c *app.RequestContext, err error, data any) {
	if err != nil {
		code, msg, desc := errno.Decode(err)
		resputil.Fail(c, code, msg, desc)
		return
	}
	resputil.Ok(c, data)
}
