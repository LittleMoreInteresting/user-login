package tags

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user-login/userlogin/internal/types"
)

var _ BlogTagModel = (*customBlogTagModel)(nil)

type (
	// BlogTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogTagModel.
	BlogTagModel interface {
		blogTagModel
		TagList(context.Context, *types.TagListRequest) (*types.TagListResponse, error)
	}

	customBlogTagModel struct {
		*defaultBlogTagModel
	}
)

// NewBlogTagModel returns a model for the database table.
func NewBlogTagModel(conn sqlx.SqlConn, c cache.CacheConf) BlogTagModel {
	return &customBlogTagModel{
		defaultBlogTagModel: newBlogTagModel(conn, c),
	}
}

func (c *customBlogTagModel) TagList(ctx context.Context, request *types.TagListRequest) (*types.TagListResponse, error) {
	result := &types.TagListResponse{}
	result.Matedata = request.Pager
	sql := fmt.Sprintf("select %s from %s where 1  limit %d,%d",
		blogTagRows, c.table, request.Pager.Offsite(), request.PageSize)
	err := c.QueryRowsNoCacheCtx(ctx, result, sql)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
