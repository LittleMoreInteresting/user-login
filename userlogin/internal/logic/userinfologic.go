package logic

import (
	"context"
	"encoding/json"

	"user-login/userlogin/internal/svc"
	"user-login/userlogin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		ID:    one.Id,
		Name:  one.Name,
		Email: one.Email,
	}, nil
}
