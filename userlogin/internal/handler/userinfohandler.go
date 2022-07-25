package handler

import (
	"net/http"

	"user-login/userlogin/common/response"
	"user-login/userlogin/internal/logic"
	"user-login/userlogin/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.HttpResult(r, w, resp, err)
	}
}
