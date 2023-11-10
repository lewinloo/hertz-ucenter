package user

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hertz-ucenter/internal/consts"
	"hertz-ucenter/internal/dal"
	"hertz-ucenter/internal/models/dto"
	"hertz-ucenter/internal/models/entity"
	"hertz-ucenter/internal/models/vo"
	"hertz-ucenter/internal/service"
	"hertz-ucenter/pkg/errno"
	"hertz-ucenter/pkg/utils"
)

func init() {
	service.RegisterUserService(New())
}

func New() *userLogic {
	return &userLogic{}
}

type userLogic struct {
}

func (u *userLogic) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errno.ErrParameterInvalid
	}
	result, err := dal.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).Delete()
	if err != nil {
		hlog.CtxErrorf(ctx, "Failed to delete user. err: %s", err.Error())
		return errno.ErrDBFailed.SetDescription("删除失败")
	}
	if result.RowsAffected <= 0 {
		return errno.ErrEntityNull.SetDescription("用户已被删除")
	}
	return nil
}

func (u *userLogic) Search(ctx context.Context, in dto.UserSearchQuery) (list []*vo.UserVO, total int64, err error) {
	username := in.Username
	userDal := dal.User.WithContext(ctx)
	if strutil.IsNotBlank(username) {
		userDal = userDal.Where(dal.User.Username.Like("%" + username + "%"))
	}
	users, total, err := userDal.FindByPage((in.Current-1)*in.Size, in.Size)
	if err != nil {
		hlog.CtxErrorf(ctx, "Failed to search user. err: %s", err.Error())
		return nil, 0, errno.ErrDBFailed.SetDescription("查询失败")
	}
	list = slice.Map(users, func(index int, item *entity.User) *vo.UserVO {
		return u.GetUserVO(item)
	})

	return
}

func (u *userLogic) Login(ctx context.Context, account, password string) (out *vo.UserVO, err error) {
	// 参数不能为空
	if utils.IsAnyStringBlank(account, password) {
		return out, errno.ErrParameterInvalid
	}

	// 账号长度不小于 4
	if len(account) < 4 {
		return out, errno.ErrParameterInvalid.SetDescription("账号长度过短")
	}

	// 密码长度不小于 8
	if len(password) < 8 {
		return out, errno.ErrParameterInvalid.SetDescription("密码长度过短")
	}
	// 账号不能包含特殊字符
	if utils.HasSpecialText(account) {
		return out, errno.ErrParameterInvalid.SetDescription("账号包含特殊字符")
	}

	// 密码加密
	encryptString := cryptor.Md5String(consts.PwdSalt + password)
	// 账号和密码进行匹配
	user, err := dal.User.WithContext(ctx).Where(dal.User.UserAccount.Eq(account), dal.User.UserPassword.Eq(encryptString)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return out, errno.ErrParameterInvalid.SetDescription("账号和密码不匹配")
		}
		return out, errno.ErrDBFailed
	}
	// 用户是否禁用
	if user.UserStatus == consts.UserStatusDisabled {
		return out, errno.ErrUnauthorization.SetDescription("用户已被冻结")
	}

	// 用户信息脱敏
	userVO := u.GetUserVO(user)

	return userVO, nil
}

func (u *userLogic) Register(ctx context.Context, account, password, checkPassword, planetCode string) (out int64, err error) {
	// 参数不能为空
	if utils.IsAnyStringBlank(account, password, checkPassword, planetCode) {
		return out, errno.ErrParameterInvalid
	}

	// 账号长度不小于 4
	if len(account) < 4 {
		return out, errno.ErrParameterInvalid.SetDescription("账号长度过短")
	}

	// 密码长度不小于 8
	if len(password) < 8 || len(checkPassword) < 8 {
		return out, errno.ErrParameterInvalid.SetDescription("密码长度过短")
	}
	// 账号不能包含特殊字符
	if utils.HasSpecialText(account) {
		return out, errno.ErrParameterInvalid.SetDescription("账号包含特殊字符")
	}

	// 密码和确认密码相等
	if !compare.Equal(password, checkPassword) {
		return out, errno.ErrParameterInvalid.SetDescription("密码和校验密码不一致")
	}

	// 星球编号长度不大于 5
	if len(planetCode) > 5 {
		return out, errno.ErrParameterInvalid.SetDescription("星球编号长度过长")
	}

	// 判断用户是否已注册
	userDao := dal.User.WithContext(ctx)
	count, err := userDao.Where(dal.User.UserAccount.Eq(account)).Count()
	if err != nil {
		hlog.CtxErrorf(ctx, "用户注册失败：%s", err.Error())
		return out, errno.ErrDBFailed
	}
	if count > 0 {
		return out, errno.ErrParameterInvalid.SetDescription("用户已注册")
	}

	// 查询数据库是否存在星球编号
	count, err = userDao.Where(dal.User.PlanetCode.Eq(planetCode)).Count()
	if err != nil {
		hlog.CtxErrorf(ctx, "用户注册失败：%s", err.Error())
		return out, errno.ErrDBFailed
	}
	if count > 0 {
		return out, errno.ErrParameterInvalid.SetDescription("星球编号已存在")
	}

	// 密码加密
	encryptString := cryptor.Md5String(consts.PwdSalt + password)
	user := &entity.User{
		UserAccount:  account,
		UserPassword: encryptString,
		PlanetCode:   planetCode,
		Username:     strutil.UpperKebabCase(account),
		AvatarURL:    consts.DefaultAvatar,
	}
	// 插入数据
	if err := userDao.Create(user); err != nil {
		hlog.CtxErrorf(ctx, "注册用户失败：%s", err.Error())
		return out, errno.ErrDBFailed.SetDescription("注册失败")
	}

	return user.ID, nil
}

func (u *userLogic) GetUserVO(user *entity.User) *vo.UserVO {
	if user == nil {
		return nil
	}
	var safetyUser vo.UserVO
	_ = copier.Copy(&safetyUser, &user)
	return &safetyUser
}
