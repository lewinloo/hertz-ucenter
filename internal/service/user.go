package service

import (
	"context"
	"hertz-ucenter/internal/models/dto"
	"hertz-ucenter/internal/models/entity"
	"hertz-ucenter/internal/models/vo"
)

type IUserService interface {
	Register(ctx context.Context, account, password, checkPassword, planetCode string) (int64, error)
	Login(ctx context.Context, account, password string) (out *vo.UserVO, err error)
	GetUserVO(user *entity.User) *vo.UserVO
	Search(ctx context.Context, in dto.UserSearchQuery) (list []*vo.UserVO, total int64, err error)
	Delete(ctx context.Context, id int64) error
}

var userService IUserService

func RegisterUserService(srv IUserService) {
	userService = srv
}

func User() IUserService {
	if userService == nil {
		panic("没有实现或注册该服务")
	}
	return userService
}
