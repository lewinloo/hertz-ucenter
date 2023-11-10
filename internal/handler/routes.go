package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"hertz-ucenter/internal/handler/user"
)

type Router interface {
	ConfigRoutes(h *route.RouterGroup)
}

var controllers = []Router{user.NewController()}

func RegisterRoutes(h *server.Hertz) {
	base := h.Group("/api")
	for _, controller := range controllers {
		controller.ConfigRoutes(base)
	}
}
