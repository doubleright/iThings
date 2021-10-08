// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/godLei6/things/src/dmsvr/dm"
	"gitee.com/godLei6/things/src/dmsvr/internal/logic"
	"gitee.com/godLei6/things/src/dmsvr/internal/svc"
)

type DmServer struct {
	svcCtx *svc.ServiceContext
}

func NewDmServer(svcCtx *svc.ServiceContext) *DmServer {
	return &DmServer{
		svcCtx: svcCtx,
	}
}

func (s *DmServer) LoginAuth(ctx context.Context, in *dm.LoginAuthReq) (*dm.Response, error) {
	l := logic.NewLoginAuthLogic(ctx, s.svcCtx)
	return l.LoginAuth(in)
}

func (s *DmServer) AccessAuth(ctx context.Context, in *dm.AccessAuthReq) (*dm.Response, error) {
	l := logic.NewAccessAuthLogic(ctx, s.svcCtx)
	return l.AccessAuth(in)
}

func (s *DmServer) ManageDevice(ctx context.Context, in *dm.ManageDeviceReq) (*dm.DeviceInfo, error) {
	l := logic.NewManageDeviceLogic(ctx, s.svcCtx)
	return l.ManageDevice(in)
}

func (s *DmServer) ManageProduct(ctx context.Context, in *dm.ManageProductReq) (*dm.ProductInfo, error) {
	l := logic.NewManageProductLogic(ctx, s.svcCtx)
	return l.ManageProduct(in)
}

func (s *DmServer) GetProductInfo(ctx context.Context, in *dm.GetProductInfoReq) (*dm.GetProductInfoResp, error) {
	l := logic.NewGetProductInfoLogic(ctx, s.svcCtx)
	return l.GetProductInfo(in)
}

func (s *DmServer) GetDeviceInfo(ctx context.Context, in *dm.GetDeviceInfoReq) (*dm.GetDeviceInfoResp, error) {
	l := logic.NewGetDeviceInfoLogic(ctx, s.svcCtx)
	return l.GetDeviceInfo(in)
}

func (s *DmServer) GetDeviceData(ctx context.Context, in *dm.GetDeviceDataReq) (*dm.GetDeviceDataResp, error) {
	l := logic.NewGetDeviceDataLogic(ctx, s.svcCtx)
	return l.GetDeviceData(in)
}

func (s *DmServer) GetDeviceLog(ctx context.Context, in *dm.GetDeviceLogReq) (*dm.GetDeviceLogResp, error) {
	l := logic.NewGetDeviceLogLogic(ctx, s.svcCtx)
	return l.GetDeviceLog(in)
}

func (s *DmServer) SendAction(ctx context.Context, in *dm.SendActionReq) (*dm.SendActionResp, error) {
	l := logic.NewSendActionLogic(ctx, s.svcCtx)
	return l.SendAction(in)
}

func (s *DmServer) SendProperty(ctx context.Context, in *dm.SendPropertyReq) (*dm.SendPropertyResp, error) {
	l := logic.NewSendPropertyLogic(ctx, s.svcCtx)
	return l.SendProperty(in)
}
