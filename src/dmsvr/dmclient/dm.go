// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

//go:generate mockgen -destination ./dm_mock.go -package dmclient -source $GOFILE

package dmclient

import (
	"context"

	"gitee.com/godLei6/things/src/dmsvr/dm"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	DeviceInfo        = dm.DeviceInfo
	ManageDeviceReq   = dm.ManageDeviceReq
	GetDeviceInfoReq  = dm.GetDeviceInfoReq
	LoginAuthReq      = dm.LoginAuthReq
	AccessAuthReq     = dm.AccessAuthReq
	ProductInfo       = dm.ProductInfo
	ManageProductReq  = dm.ManageProductReq
	Response          = dm.Response
	ManResp           = dm.ManResp
	GetProductInfoReq = dm.GetProductInfoReq

	Dm interface {
		LoginAuth(ctx context.Context, in *LoginAuthReq) (*Response, error)
		AccessAuth(ctx context.Context, in *AccessAuthReq) (*Response, error)
		ManageDevice(ctx context.Context, in *ManageDeviceReq) (*DeviceInfo, error)
		ManageProduct(ctx context.Context, in *ManageProductReq) (*ProductInfo, error)
		GetProductInfo(ctx context.Context, in *GetProductInfoReq) (*ProductInfo, error)
		GetDeviceInfo(ctx context.Context, in *GetDeviceInfoReq) (*DeviceInfo, error)
	}

	defaultDm struct {
		cli zrpc.Client
	}
)

func NewDm(cli zrpc.Client) Dm {
	return &defaultDm{
		cli: cli,
	}
}

func (m *defaultDm) LoginAuth(ctx context.Context, in *LoginAuthReq) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.LoginAuth(ctx, in)
}

func (m *defaultDm) AccessAuth(ctx context.Context, in *AccessAuthReq) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.AccessAuth(ctx, in)
}

func (m *defaultDm) ManageDevice(ctx context.Context, in *ManageDeviceReq) (*DeviceInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageDevice(ctx, in)
}

func (m *defaultDm) ManageProduct(ctx context.Context, in *ManageProductReq) (*ProductInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageProduct(ctx, in)
}

func (m *defaultDm) GetProductInfo(ctx context.Context, in *GetProductInfoReq) (*ProductInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetProductInfo(ctx, in)
}

func (m *defaultDm) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoReq) (*DeviceInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetDeviceInfo(ctx, in)
}