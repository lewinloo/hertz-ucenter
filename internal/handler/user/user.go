package user

import (
	"github.com/cloudwego/hertz/pkg/route"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/internal/mw"
)

func NewController() *cUser {
	return &cUser{}
}

type cUser struct {
}

func (ctrl *cUser) ConfigRoutes(h *route.RouterGroup) {
	group := h.Group("/user")
	{
		group.POST("/register", ctrl.register)
		group.POST("/login", ctrl.login)
		group.Use(mw.CheckAuth)
		group.GET("/current", ctrl.getCurrentUser)
		group.POST("/logout", ctrl.logout)
		group.Use(mw.CheckRole(consts.AdminRole))
		group.GET("/search", ctrl.search)
		group.DELETE("/:id", ctrl.delete)
	}

}
