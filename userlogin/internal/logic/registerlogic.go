package logic

import (
	"context"

	"google.golang.org/grpc/status"
	"user-login/userlogin/common/cryptx"
	"user-login/userlogin/internal/svc"
	"user-login/userlogin/internal/types"
	"user-login/userlogin/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {

	_, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}
	if err != user.ErrNotFound {
		return nil, status.Error(100, err.Error())
	}
	newUser := user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
	}

	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newUser.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.RegisterResponse{
		ID:    int(newUser.Id),
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}
