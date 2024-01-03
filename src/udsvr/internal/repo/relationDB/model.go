package relationDB

import (
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/src/udsvr/internal/domain/scene"
)

// 示例
type UdExample struct {
	ID int64 `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"` // id编号
}

type UdSceneInfo struct {
	ID              int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`              // id编号
	TenantCode      stores.TenantCode `gorm:"column:tenant_code;type:VARCHAR(50);uniqueIndex:name;NOT NULL"` // 租户编码
	UserID          int64             `gorm:"column:user_id;type:BIGINT;NOT NULL"`                           // 用户id
	Name            string            `gorm:"column:name;type:varchar(100);uniqueIndex:name;NOT NULL"`       // 名称
	Desc            string            `gorm:"column:desc;type:varchar(200);NOT NULL"`                        // 描述
	*UdSceneTrigger `gorm:"embedded;embeddedPrefix:trigger_"`
	*UdSceneWhen    `gorm:"embedded;embeddedPrefix:when_"`
	*UdSceneThen    `gorm:"embedded;embeddedPrefix:then_"`
}

func (m *UdSceneInfo) TableName() string {
	return "ud_scene_info"
}

type UdSceneTrigger struct {
	Type   string               `gorm:"column:type;type:VARCHAR(25);NOT NULL"` //触发类型 device: 设备触发 timer: 定时触发 manual:手动触发
	Device scene.TriggerDevices `gorm:"column:device;type:json;serializer:json"`
	Timer  scene.Timer          `gorm:"column:timer;type:json;serializer:json"`
}

type UdSceneWhen struct {
	TermCondType string          `gorm:"column:type;type:VARCHAR(25);NOT NULL"` //条件类型 and: 与 or: 或
	ValidRange   scene.WhenRange `gorm:"column:validRange;type:json;serializer:json"`
	Terms        scene.Terms     `gorm:"column:terms;type:json;serializer:json"`
}

type UdSceneThen struct {
	Actions scene.Actions `gorm:"column:actions;type:json;serializer:json"`
}
