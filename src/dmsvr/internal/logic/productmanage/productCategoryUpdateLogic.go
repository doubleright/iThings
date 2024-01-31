package productmanagelogic

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/oss"
	"github.com/i-Things/things/shared/oss/common"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryUpdateLogic {
	return &ProductCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新产品
func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(in *dm.ProductCategory) (*dm.Response, error) {
	old, err := relationDB.NewProductCategoryRepo(l.ctx).FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Name != "" {
		old.Name = in.Name
	}
	if in.Desc != nil {
		old.Desc = utils.ToEmptyString(in.Desc)
	}
	if in.IsUpdateHeadImg && in.HeadImg != "" {
		if old.HeadImg != "" {
			err := l.svcCtx.OssClient.PrivateBucket().Delete(l.ctx, old.HeadImg, common.OptionKv{})
			if err != nil {
				l.Errorf("Delete file err path:%v,err:%v", old.HeadImg, err)
			}
		}
		nwePath := oss.GenFilePath(l.ctx, l.svcCtx.Config.Name, oss.BusinessProductManage, oss.SceneCategoryImg, fmt.Sprintf("%d/%s", in.Id, oss.GetFileNameWithPath(in.HeadImg)))
		path, err := l.svcCtx.OssClient.PrivateBucket().CopyFromTempBucket(in.HeadImg, nwePath)
		if err != nil {
			return nil, errors.System.AddDetail(err)
		}
		old.HeadImg = path
	}

	err = relationDB.NewProductCategoryRepo(l.ctx).Update(l.ctx, old)
	return &dm.Response{}, err
}