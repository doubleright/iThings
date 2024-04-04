// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package schemamanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                = dm.ActionRespReq
	ActionSendReq                = dm.ActionSendReq
	ActionSendResp               = dm.ActionSendResp
	CommonSchemaCreateReq        = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq         = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp        = dm.CommonSchemaIndexResp
	CommonSchemaInfo             = dm.CommonSchemaInfo
	CommonSchemaUpdateReq        = dm.CommonSchemaUpdateReq
	CustomTopic                  = dm.CustomTopic
	DeviceCore                   = dm.DeviceCore
	DeviceCountInfo              = dm.DeviceCountInfo
	DeviceCountReq               = dm.DeviceCountReq
	DeviceCountResp              = dm.DeviceCountResp
	DeviceGatewayBindDevice      = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq        = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp       = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq  = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq  = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign            = dm.DeviceGatewaySign
	DeviceInfo                   = dm.DeviceInfo
	DeviceInfoCount              = dm.DeviceInfoCount
	DeviceInfoCountReq           = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq          = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq           = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp          = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq     = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq            = dm.DeviceInfoReadReq
	DeviceTypeCountReq           = dm.DeviceTypeCountReq
	DeviceTypeCountResp          = dm.DeviceTypeCountResp
	Empty                        = dm.Empty
	EventLogIndexReq             = dm.EventLogIndexReq
	EventLogIndexResp            = dm.EventLogIndexResp
	EventLogInfo                 = dm.EventLogInfo
	Firmware                     = dm.Firmware
	FirmwareFile                 = dm.FirmwareFile
	FirmwareInfo                 = dm.FirmwareInfo
	FirmwareInfoDeleteReq        = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp       = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq         = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp        = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq          = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp         = dm.FirmwareInfoReadResp
	FirmwareResp                 = dm.FirmwareResp
	GroupDeviceIndexReq          = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp         = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq    = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq      = dm.GroupDeviceMultiSaveReq
	GroupInfo                    = dm.GroupInfo
	GroupInfoCreateReq           = dm.GroupInfoCreateReq
	GroupInfoIndexReq            = dm.GroupInfoIndexReq
	GroupInfoIndexResp           = dm.GroupInfoIndexResp
	GroupInfoUpdateReq           = dm.GroupInfoUpdateReq
	HubLogIndexReq               = dm.HubLogIndexReq
	HubLogIndexResp              = dm.HubLogIndexResp
	HubLogInfo                   = dm.HubLogInfo
	OtaFirmwareDeviceCancelReq   = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceIndexReq    = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp   = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo        = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq    = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile              = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq      = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp     = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo          = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq           = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp          = dm.OtaFirmwareFileResp
	OtaFirmwareInfo              = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq     = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq      = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp     = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq     = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq       = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp      = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo           = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq       = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo            = dm.OtaJobDynamicInfo
	OtaJobStaticInfo             = dm.OtaJobStaticInfo
	OtaModuleInfo                = dm.OtaModuleInfo
	OtaModuleInfoIndexReq        = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp       = dm.OtaModuleInfoIndexResp
	OtaPromptIndexReq            = dm.OtaPromptIndexReq
	OtaPromptIndexResp           = dm.OtaPromptIndexResp
	PageInfo                     = dm.PageInfo
	PageInfo_OrderBy             = dm.PageInfo_OrderBy
	Point                        = dm.Point
	ProductCategory              = dm.ProductCategory
	ProductCategoryIndexReq      = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp     = dm.ProductCategoryIndexResp
	ProductCategoryReadReq       = dm.ProductCategoryReadReq
	ProductCustom                = dm.ProductCustom
	ProductCustomReadReq         = dm.ProductCustomReadReq
	ProductInfo                  = dm.ProductInfo
	ProductInfoDeleteReq         = dm.ProductInfoDeleteReq
	ProductInfoIndexReq          = dm.ProductInfoIndexReq
	ProductInfoIndexResp         = dm.ProductInfoIndexResp
	ProductInfoReadReq           = dm.ProductInfoReadReq
	ProductInitReq               = dm.ProductInitReq
	ProductRemoteConfig          = dm.ProductRemoteConfig
	ProductSchemaCreateReq       = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq       = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq        = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp       = dm.ProductSchemaIndexResp
	ProductSchemaInfo            = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq  = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq    = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq      = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp     = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq       = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq  = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg       = dm.PropertyControlSendMsg
	PropertyControlSendReq       = dm.PropertyControlSendReq
	PropertyControlSendResp      = dm.PropertyControlSendResp
	PropertyGetReportSendReq     = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp    = dm.PropertyGetReportSendResp
	PropertyLogIndexReq          = dm.PropertyLogIndexReq
	PropertyLogIndexResp         = dm.PropertyLogIndexResp
	PropertyLogInfo              = dm.PropertyLogInfo
	PropertyLogLatestIndexReq    = dm.PropertyLogLatestIndexReq
	ProtocolConfigField          = dm.ProtocolConfigField
	ProtocolConfigInfo           = dm.ProtocolConfigInfo
	ProtocolInfo                 = dm.ProtocolInfo
	ProtocolInfoIndexReq         = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp        = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq        = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq         = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp        = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq      = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp     = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq       = dm.RemoteConfigPushAllReq
	RespReadReq                  = dm.RespReadReq
	RootCheckReq                 = dm.RootCheckReq
	SdkLogIndexReq               = dm.SdkLogIndexReq
	SdkLogIndexResp              = dm.SdkLogIndexResp
	SdkLogInfo                   = dm.SdkLogInfo
	SendLogIndexReq              = dm.SendLogIndexReq
	SendLogIndexResp             = dm.SendLogIndexResp
	SendLogInfo                  = dm.SendLogInfo
	SendMsgReq                   = dm.SendMsgReq
	SendMsgResp                  = dm.SendMsgResp
	SendOption                   = dm.SendOption
	ShadowIndex                  = dm.ShadowIndex
	ShadowIndexResp              = dm.ShadowIndexResp
	StatusLogIndexReq            = dm.StatusLogIndexReq
	StatusLogIndexResp           = dm.StatusLogIndexResp
	StatusLogInfo                = dm.StatusLogInfo
	TimeRange                    = dm.TimeRange
	UserDeviceCollectSave        = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq      = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp     = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo          = dm.UserDeviceShareInfo
	UserDeviceShareReadReq       = dm.UserDeviceShareReadReq
	WithID                       = dm.WithID
	WithIDCode                   = dm.WithIDCode

	SchemaManage interface {
		// 更新产品物模型
		CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error)
		// 新增产品
		CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error)
		// 删除产品
		CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 获取产品信息列表
		CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error)
	}

	defaultSchemaManage struct {
		cli zrpc.Client
	}

	directSchemaManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.SchemaManageServer
	}
)

func NewSchemaManage(cli zrpc.Client) SchemaManage {
	return &defaultSchemaManage{
		cli: cli,
	}
}

func NewDirectSchemaManage(svcCtx *svc.ServiceContext, svr dm.SchemaManageServer) SchemaManage {
	return &directSchemaManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 更新产品物模型
func (m *defaultSchemaManage) CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewSchemaManageClient(m.cli.Conn())
	return client.CommonSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directSchemaManage) CommonSchemaUpdate(ctx context.Context, in *CommonSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.CommonSchemaUpdate(ctx, in)
}

// 新增产品
func (m *defaultSchemaManage) CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewSchemaManageClient(m.cli.Conn())
	return client.CommonSchemaCreate(ctx, in, opts...)
}

// 新增产品
func (d *directSchemaManage) CommonSchemaCreate(ctx context.Context, in *CommonSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.CommonSchemaCreate(ctx, in)
}

// 删除产品
func (m *defaultSchemaManage) CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewSchemaManageClient(m.cli.Conn())
	return client.CommonSchemaDelete(ctx, in, opts...)
}

// 删除产品
func (d *directSchemaManage) CommonSchemaDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.CommonSchemaDelete(ctx, in)
}

// 获取产品信息列表
func (m *defaultSchemaManage) CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error) {
	client := dm.NewSchemaManageClient(m.cli.Conn())
	return client.CommonSchemaIndex(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directSchemaManage) CommonSchemaIndex(ctx context.Context, in *CommonSchemaIndexReq, opts ...grpc.CallOption) (*CommonSchemaIndexResp, error) {
	return d.svr.CommonSchemaIndex(ctx, in)
}
