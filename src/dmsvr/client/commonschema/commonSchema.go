// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package commonschema

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonSchemaCreateReq       = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq        = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp       = dm.CommonSchemaIndexResp
	CommonSchemaInfo            = dm.CommonSchemaInfo
	CommonSchemaUpdateReq       = dm.CommonSchemaUpdateReq
	CustomTopic                 = dm.CustomTopic
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq    = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	EventIndex                  = dm.EventIndex
	EventIndexResp              = dm.EventIndexResp
	EventLogIndexReq            = dm.EventLogIndexReq
	Firmware                    = dm.Firmware
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GetPropertyReplyReq         = dm.GetPropertyReplyReq
	GetPropertyReplyResp        = dm.GetPropertyReplyResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq     = dm.GroupDeviceMultiSaveReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	HubLogIndex                 = dm.HubLogIndex
	HubLogIndexReq              = dm.HubLogIndexReq
	HubLogIndexResp             = dm.HubLogIndexResp
	MultiSendPropertyReq        = dm.MultiSendPropertyReq
	MultiSendPropertyResp       = dm.MultiSendPropertyResp
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaPageInfo                 = dm.OtaPageInfo
	OtaPromptIndexReq           = dm.OtaPromptIndexReq
	OtaPromptIndexResp          = dm.OtaPromptIndexResp
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCategory             = dm.ProductCategory
	ProductCategoryIndexReq     = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp    = dm.ProductCategoryIndexResp
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	PropertyIndex               = dm.PropertyIndex
	PropertyIndexResp           = dm.PropertyIndexResp
	PropertyLatestIndexReq      = dm.PropertyLatestIndexReq
	PropertyLogIndexReq         = dm.PropertyLogIndexReq
	ProtocolInfo                = dm.ProtocolInfo
	ProtocolInfoIndexReq        = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp       = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	RespActionReq               = dm.RespActionReq
	RespReadReq                 = dm.RespReadReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq
	SdkLogIndex                 = dm.SdkLogIndex
	SdkLogIndexReq              = dm.SdkLogIndexReq
	SdkLogIndexResp             = dm.SdkLogIndexResp
	SendActionReq               = dm.SendActionReq
	SendActionResp              = dm.SendActionResp
	SendMsgReq                  = dm.SendMsgReq
	SendMsgResp                 = dm.SendMsgResp
	SendOption                  = dm.SendOption
	SendPropertyMsg             = dm.SendPropertyMsg
	SendPropertyReq             = dm.SendPropertyReq
	SendPropertyResp            = dm.SendPropertyResp
	ShadowIndex                 = dm.ShadowIndex
	ShadowIndexResp             = dm.ShadowIndexResp
	WithID                      = dm.WithID

	CommonSchema interface {
		// 更新产品物模型
		CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
		// 新增产品
		CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 删除产品
		CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
		// 获取产品信息列表
		CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error)
	}

	defaultCommonSchema struct {
		cli zrpc.Client
	}

	directCommonSchema struct {
		svcCtx *svc.ServiceContext
		svr    dm.CommonSchemaServer
	}
)

func NewCommonSchema(cli zrpc.Client) CommonSchema {
	return &defaultCommonSchema{
		cli: cli,
	}
}

func NewDirectCommonSchema(svcCtx *svc.ServiceContext, svr dm.CommonSchemaServer) CommonSchema {
	return &directCommonSchema{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 更新产品物模型
func (m *defaultCommonSchema) CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewCommonSchemaClient(m.cli.Conn())
	return client.CommonSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directCommonSchema) CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.CommonSchemaUpdate(ctx, in)
}

// 新增产品
func (m *defaultCommonSchema) CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewCommonSchemaClient(m.cli.Conn())
	return client.CommonSchemaCreate(ctx, in, opts...)
}

// 新增产品
func (d *directCommonSchema) CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.CommonSchemaCreate(ctx, in)
}

// 删除产品
func (m *defaultCommonSchema) CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewCommonSchemaClient(m.cli.Conn())
	return client.CommonSchemaDelete(ctx, in, opts...)
}

// 删除产品
func (d *directCommonSchema) CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.CommonSchemaDelete(ctx, in)
}

// 获取产品信息列表
func (m *defaultCommonSchema) CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error) {
	client := dm.NewCommonSchemaClient(m.cli.Conn())
	return client.CommonSchemaIndex(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directCommonSchema) CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error) {
	return d.svr.CommonSchemaIndex(ctx, in)
}