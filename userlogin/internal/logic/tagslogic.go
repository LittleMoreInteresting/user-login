package logic

import (
	"context"

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

func (l *TagsLogic) Tags(req *types.TagListRequest) (resp *types.TagListResponse, err error) {
	return l.svcCtx.TagsModel.TagList(l.ctx, req)
}
