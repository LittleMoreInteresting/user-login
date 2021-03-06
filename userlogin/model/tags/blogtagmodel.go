package tags

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
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
	pager := request.Pager
	countSql := fmt.Sprintf("select %s from %s where 1 ", "count(id) as total_size", c.table)
	_ = c.QueryRowNoCacheCtx(ctx, &pager.TotalSize, countSql)
	result.Matedata = pager
	where := "1 "
	var err error
	if len(request.Name) > 0 {
		where += "AND name=? "
		sql := fmt.Sprintf("select %s from %s where %s  limit %d,%d",
			blogTagRows, c.table, where, (pager.Page-1)*pager.PageSize, request.PageSize)

		err = c.QueryRowsNoCacheCtx(ctx, &result.List, sql, request.Name)
	} else {
		sql := fmt.Sprintf("select %s from %s where %s  limit %d,%d",
			blogTagRows, c.table, where, (pager.Page-1)*pager.PageSize, request.PageSize)

		err = c.QueryRowsNoCacheCtx(ctx, &result.List, sql)
	}
	logx.Info(request)
	switch err {
	case nil:
		return result, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
