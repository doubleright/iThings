package productmanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSchemaTslReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductSchemaTslReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSchemaTslReadLogic {
	return &ProductSchemaTslReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取产品信息列表
func (l *ProductSchemaTslReadLogic) ProductSchemaTslRead(in *dm.ProductSchemaTslReadReq) (*dm.ProductSchemaTslReadResp, error) {
	l.Infof("%s req=%v", utils.FuncName(), utils.Fmt(in))
	model, err := l.svcCtx.ProductSchemaRepo.GetData(l.ctx, in.ProductID)
	if err != nil {
		return nil, err
	}
	return &dm.ProductSchemaTslReadResp{Tsl: model.String()}, nil
}
