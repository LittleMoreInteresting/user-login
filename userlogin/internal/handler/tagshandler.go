package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user-login/userlogin/internal/logic"
	"user-login/userlogin/internal/svc"
)

func TagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTagsLogic(r.Context(), svcCtx)
		resp, err := l.Tags()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
