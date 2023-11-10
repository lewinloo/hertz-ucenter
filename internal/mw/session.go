package mw

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"hertz-ucenter/internal/consts"
)

func InitSession(h *server.Hertz) {
	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New(consts.HertzSession, store))
}
