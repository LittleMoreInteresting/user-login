package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"user-login/userlogin/internal/logic"
	"user-login/userlogin/internal/svc"
	"user-login/userlogin/internal/types"
)

func TagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTagsLogic(r.Context(), svcCtx)
		resp, err := l.Tags(&req)
		if err != nil {
			logx.Info(err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
