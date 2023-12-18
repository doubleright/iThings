package app

import (
	"context"
	app "github.com/i-Things/things/src/apisvr/internal/logic/system/app/info"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.RoleAppIndexReq) (resp *types.RoleAppIndexResp, err error) {
	ret, err := l.svcCtx.RoleRpc.RoleAppIndex(l.ctx, &sys.RoleAppIndexReq{Id: req.ID})
	if err != nil {
		return nil, err
	}
	return &types.RoleAppIndexResp{
		List:  app.ToAppInfosTypes(ret.List),
		Total: ret.Total,
	}, err
}