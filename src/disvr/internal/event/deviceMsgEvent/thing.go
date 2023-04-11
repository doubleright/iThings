package deviceMsgEvent

import (
	"context"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/domain/application"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgHubLog"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgThing"
	"github.com/i-Things/things/src/disvr/internal/repo/cache"
	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ThingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	schema *schema.Model
	dreq   msgThing.Req
	dd     msgThing.SchemaDataRepo
}

func NewThingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThingLogic {
	return &ThingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThingLogic) initMsg(msg *deviceMsg.PublishMsg) error {
	var err error
	l.schema, err = l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, msg.ProductID)
	if err != nil {
		return errors.Database.AddDetail(err)
	}
	err = utils.Unmarshal(msg.Payload, &l.dreq)
	if err != nil {
		return errors.Parameter.AddDetailf("payload unmarshal payload:%v err:%v", string(msg.Payload), err)
	}
	l.dd = l.svcCtx.SchemaMsgRepo
	return nil
}

func (l *ThingLogic) DeviceResp(msg *deviceMsg.PublishMsg, err error, data any) *deviceMsg.PublishMsg {
	resp := &deviceMsg.CommonMsg{
		Method:      deviceMsg.GetRespMethod(l.dreq.Method),
		ClientToken: l.dreq.ClientToken,
		Timestamp:   time.Now().UnixMilli(),
		Data:        data,
	}
	return &deviceMsg.PublishMsg{
		Handle:     msg.Handle,
		Type:       msg.Type,
		Payload:    resp.AddStatus(err).Bytes(),
		Timestamp:  time.Now().UnixMilli(),
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}
}

func (l *ThingLogic) HandlePropertyReport(msg *deviceMsg.PublishMsg, req msgThing.Req) (respMsg *deviceMsg.PublishMsg, err error) {
	tp, err := req.VerifyReqParam(l.schema, schema.ParamProperty)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	} else if len(tp) == 0 {
		err := errors.Parameter.AddDetail("need right param")

		return l.DeviceResp(msg, err, nil), err
	}

	params := msgThing.ToVal(tp)
	timeStamp := req.GetTimeStamp(msg.Timestamp)
	core := devices.Core{
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}
	paramValues := ToParamValues(tp)
	for identifier, param := range paramValues {
		err := l.svcCtx.PubApp.DeviceThingPropertyReport(l.ctx, application.PropertyReport{
			Device: core, Timestamp: timeStamp.UnixMilli(),
			Identifier: identifier, Param: param,
		})
		if err != nil {
			l.Errorf("%s.DeviceThingPropertyReport  identifier:%v, param:%v,err:%v", utils.FuncName(), identifier, param, err)
		}
	}

	err = l.dd.InsertPropertiesData(l.ctx, l.schema, msg.ProductID, msg.DeviceName, params, timeStamp)
	if err != nil {
		l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
		return l.DeviceResp(msg, errors.Database, nil), err
	}

	return l.DeviceResp(msg, errors.OK, nil), nil
}

func (l *ThingLogic) HandlePropertyGetStatus(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	respData := make(map[string]any, len(l.schema.Property))
	switch l.dreq.Type { //表示获取什么类型的信息（report:表示设备上报的信息 info:信息 alert:告警 fault:故障）
	case deviceMsg.Report: //表示设备属性上报
		for id := range l.schema.Property {
			data, err := l.dd.GetLatestPropertyDataByID(l.ctx, msgThing.LatestFilter{
				ProductID:  msg.ProductID,
				DeviceName: msg.DeviceName,
				DataID:     id,
			})
			if err != nil {
				l.Errorf("%s.GetPropertyDataByID.get id:%s err:%s",
					utils.FuncName(), id, err.Error())
				return nil, err
			}
			if data == nil {
				l.Infof("%s.GetPropertyDataByID not find id:%s", utils.FuncName(), id)
				continue
			}
			respData[id] = data.Param
		}
	default:
		err := errors.Parameter.AddDetailf("not support type :%s", l.dreq.Type)

		return l.DeviceResp(msg, err, nil), err
	}

	return l.DeviceResp(msg, errors.OK, respData), nil
}

//属性上报
func (l *ThingLogic) HandleProperty(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)
	switch l.dreq.Method { //操作方法
	case deviceMsg.Report, deviceMsg.ReportInfo: //设备属性上报
		return l.HandlePropertyReport(msg, l.dreq)
	case deviceMsg.GetStatus: //设备请求获取 云端记录的最新设备信息
		return l.HandlePropertyGetStatus(msg)
	case deviceMsg.ControlReply: //设备回复的 云端下发控制指令 的处理结果
		return l.HandleResp(msg)
	default:
		return nil, errors.Method
	}
}

func (l *ThingLogic) HandleEvent(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)
	dbData := msgThing.EventData{}
	dbData.Identifier = l.dreq.EventID
	dbData.Type = l.dreq.Type
	if l.dreq.Method != deviceMsg.EventPost {
		return nil, errors.Method
	}
	tp, err := l.dreq.VerifyReqParam(l.schema, schema.ParamEvent)
	if err != nil {
		return l.DeviceResp(msg, err, nil), err
	}
	dbData.Params = msgThing.ToVal(tp)
	dbData.TimeStamp = l.dreq.GetTimeStamp(msg.Timestamp)
	paramValues := ToParamValues(tp)
	err = l.svcCtx.PubApp.DeviceThingEventReport(l.ctx, application.EventReport{
		Device:     devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName},
		Timestamp:  dbData.TimeStamp.UnixMilli(),
		Identifier: dbData.Identifier,
		Params:     paramValues,
		Type:       dbData.Type,
	})
	if err != nil {
		l.Errorf("%s.DeviceThingEventReport  err:%v", utils.FuncName(), err)
	}
	err = l.dd.InsertEventData(l.ctx, msg.ProductID, msg.DeviceName, &dbData)
	if err != nil {
		l.Errorf("%s.InsertEventData err=%+v", utils.FuncName(), err)
		return l.DeviceResp(msg, errors.Database, nil), errors.Database.AddDetail(err)
	}
	return l.DeviceResp(msg, errors.OK, nil), nil
}
func (l *ThingLogic) HandleResp(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Debugf("%s req:%v", utils.FuncName(), msg)
	var resp msgThing.Resp
	err = utils.Unmarshal(msg.Payload, &resp)
	if err != nil {
		return nil, errors.Parameter.AddDetailf("payload unmarshal payload:%v err:%v", string(msg.Payload), err)
	}
	req, err := cache.GetDeviceMsg[msgThing.Req](l.ctx, l.svcCtx.Store, deviceMsg.ReqMsg, msg.Handle, msg.Type,
		devices.Core{ProductID: msg.ProductID, DeviceName: msg.DeviceName},
		resp.ClientToken)
	if req == nil || err != nil {
		return nil, err
	}
	err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Store, deviceMsg.RespMsg, msg, resp.ClientToken)
	if err != nil {
		return nil, err
	}
	if msg.Type == msgThing.TypeProperty {
		_, err = l.HandlePropertyReport(msg, *req)
		return nil, err
	}
	return nil, nil
}

func (l *ThingLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%v", utils.FuncName(), msg)
	err = l.initMsg(msg)
	if err != nil {
		return nil, err
	}
	var action = devices.Thing
	respMsg, err = func() (respMsg *deviceMsg.PublishMsg, err error) {
		action = msg.Type
		switch msg.Type { //操作类型 从topic中提取 物模型下就是   property属性 event事件 action行为
		case msgThing.TypeProperty: //属性或信息上报
			return l.HandleProperty(msg)
		case msgThing.TypeEvent: //事件上报
			return l.HandleEvent(msg)
		case msgThing.TypeAction: //行为上报
			return l.HandleResp(msg)
		default:
			action = devices.Thing
			return nil, errors.Parameter.AddDetailf("things types is err:%v", msg.Type)
		}
	}()
	if l.dreq.NoAsk() { //如果不需要回复
		respMsg = nil
	}
	_ = l.svcCtx.HubLogRepo.Insert(l.ctx, &msgHubLog.HubLog{
		ProductID:  msg.ProductID,
		Action:     action,
		Timestamp:  time.Now(), // 操作时间
		DeviceName: msg.DeviceName,
		TranceID:   utils.TraceIdFromContext(l.ctx),
		RequestID:  l.dreq.ClientToken,
		Content:    string(msg.Payload),
		Topic:      msg.Topic,
		ResultType: errors.Fmt(err).GetCode(),
	})
	return
}
