package logic

import (
	"context"
	"fmt"

	"user-login/userlogin/internal/svc"
	"user-login/userlogin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TagsLogic {
	return &TagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TagsLogic) Tags() (resp *types.TagResponse, err error) {
	s := fmt.Sprintf("%s--%s", l.ctx.Value("tag"), l.ctx.Value("version"))
	return &types.TagResponse{
		Tag: s,
	}, nil
}
