package self

import (
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/result"
	"github.com/i-Things/things/src/apisvr/internal/logic/system/user/self"
	"github.com/i-Things/things/src/apisvr/internal/middleware"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}
		r = ctxs.NotLoginedInit(r)
		userCtx, err := middleware.NewCheckTokenWareMiddleware(svcCtx.Config, svcCtx.UserRpc, svcCtx.RoleRpc).UserAuth(w, r)
		if err == nil { //登录态也需要支持
			//注入 用户信息 到 ctx
			ctx2 := ctxs.SetUserCtx(r.Context(), userCtx)
			r = r.WithContext(ctx2)
		}
		l := self.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha(&req)
		result.Http(w, r, resp, err)
	}
}