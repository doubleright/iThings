package handler

import (
	"net/http"

	"yl/src/user/api/test/internal/logic/user"
	"yl/src/user/api/test/internal/svc"
	"yl/src/user/api/test/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterCoreHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterCoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterCoreLogic(r.Context(), ctx)
		resp, err := l.RegisterCore(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
