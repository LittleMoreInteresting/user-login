package logic

import (
	"context"
	"time"
	"user-login/userlogin/common/errorx"

	"user-login/userlogin/common/cryptx"
	"user-login/userlogin/common/jwtx"
	"user-login/userlogin/internal/svc"
	"user-login/userlogin/internal/types"
	"user-login/userlogin/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		if err == user.ErrNotFound {
			return nil, errorx.NewDefaultError("用户不存在")
		}
		return nil, errorx.NewCodeError(errorx.DBErrorCode, err.Error())
	}
	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if password != res.Password {
		return nil, errorx.NewCodeError(100, "密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	return &types.LoginResponse{
		Token:  accessToken,
		Expire: now + accessExpire,
	}, nil
}
