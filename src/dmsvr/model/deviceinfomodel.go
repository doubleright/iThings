package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	deviceInfoFieldNames          = builderx.RawFieldNames(&DeviceInfo{})
	deviceInfoRows                = strings.Join(deviceInfoFieldNames, ",")
	deviceInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(deviceInfoFieldNames, "`create_time`", "`update_time`"), ",")
	deviceInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(deviceInfoFieldNames, "`deviceID`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDeviceInfoDeviceIDPrefix   = "cache#deviceInfo#deviceID#"
	cacheDeviceInfoDeviceNamePrefix = "cache#deviceInfo#deviceName#"
)

type (
	DeviceInfoModel interface {
		Insert(data DeviceInfo) (sql.Result, error)
		FindOne(deviceID int64) (*DeviceInfo, error)
		FindOneByDeviceName(deviceName string) (*DeviceInfo, error)
		Update(data DeviceInfo) error
		Delete(deviceID int64) error
	}

	defaultDeviceInfoModel struct {
		sqlc.CachedConn
		table string
	}

	DeviceInfo struct {
		DeviceID    int64        `db:"deviceID"`   // 设备id
		ProductID   int64        `db:"productID"`  // 产品id
		DeviceName  string       `db:"deviceName"` // 设备名称
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime sql.NullTime `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
	}
)

func NewDeviceInfoModel(conn sqlx.SqlConn, c cache.CacheConf) DeviceInfoModel {
	return &defaultDeviceInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`device_info`",
	}
}

func (m *defaultDeviceInfoModel) Insert(data DeviceInfo) (sql.Result, error) {
	deviceInfoDeviceNameKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceNamePrefix, data.DeviceName)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, deviceInfoRowsExpectAutoSet)
		return conn.Exec(query, data.DeviceID, data.ProductID, data.DeviceName, data.CreatedTime, data.UpdatedTime, data.DeletedTime)
	}, deviceInfoDeviceNameKey)
	return ret, err
}

func (m *defaultDeviceInfoModel) FindOne(deviceID int64) (*DeviceInfo, error) {
	deviceInfoDeviceIDKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceIDPrefix, deviceID)
	var resp DeviceInfo
	err := m.QueryRow(&resp, deviceInfoDeviceIDKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `deviceID` = ? limit 1", deviceInfoRows, m.table)
		return conn.QueryRow(v, query, deviceID)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceInfoModel) FindOneByDeviceName(deviceName string) (*DeviceInfo, error) {
	deviceInfoDeviceNameKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceNamePrefix, deviceName)
	var resp DeviceInfo
	err := m.QueryRowIndex(&resp, deviceInfoDeviceNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `deviceName` = ? limit 1", deviceInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, deviceName); err != nil {
			return nil, err
		}
		return resp.DeviceID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeviceInfoModel) Update(data DeviceInfo) error {
	deviceInfoDeviceIDKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceIDPrefix, data.DeviceID)
	deviceInfoDeviceNameKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceNamePrefix, data.DeviceName)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `deviceID` = ?", m.table, deviceInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.ProductID, data.DeviceName, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.DeviceID)
	}, deviceInfoDeviceIDKey, deviceInfoDeviceNameKey)
	return err
}

func (m *defaultDeviceInfoModel) Delete(deviceID int64) error {
	data, err := m.FindOne(deviceID)
	if err != nil {
		return err
	}

	deviceInfoDeviceIDKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceIDPrefix, deviceID)
	deviceInfoDeviceNameKey := fmt.Sprintf("%s%v", cacheDeviceInfoDeviceNamePrefix, data.DeviceName)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `deviceID` = ?", m.table)
		return conn.Exec(query, deviceID)
	}, deviceInfoDeviceIDKey, deviceInfoDeviceNameKey)
	return err
}

func (m *defaultDeviceInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDeviceInfoDeviceIDPrefix, primary)
}

func (m *defaultDeviceInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `deviceID` = ? limit 1", deviceInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}