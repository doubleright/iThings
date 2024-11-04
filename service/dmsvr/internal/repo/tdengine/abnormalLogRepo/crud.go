package abnormalLogRepo

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/stores"
	sq "gitee.com/unitedrhino/squirrel"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceLog"
)

func (s *AbnormalLogRepo) fillFilter(sql sq.SelectBuilder, filter deviceLog.AbnormalFilter) sq.SelectBuilder {
	if len(filter.ProductID) != 0 {
		sql = sql.Where("`product_id`=?", filter.ProductID)
	}
	if len(filter.DeviceName) != 0 {
		sql = sql.Where("`device_name`=?", filter.DeviceName)
	}
	return sql
}

func (s *AbnormalLogRepo) GetCountLog(ctx context.Context, filter deviceLog.AbnormalFilter, page def.PageInfo2) (int64, error) {
	sqSql := sq.Select("Count(1)").From(s.GetLogStableName())
	sqSql = s.fillFilter(sqSql, filter)
	sqSql = page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := s.t.QueryRowContext(ctx, sqlStr, value...)
	if err != nil {
		return 0, err
	}
	var (
		total int64
	)

	err = row.Scan(&total)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return total, nil
}

func (s *AbnormalLogRepo) GetDeviceLog(ctx context.Context, filter deviceLog.AbnormalFilter, page def.PageInfo2) (
	[]*deviceLog.Abnormal, error) {
	sql := sq.Select("*").From(s.GetLogStableName()).OrderBy("`ts` desc")
	sql = s.fillFilter(sql, filter)
	sql = page.FmtSql(sql)
	sqlStr, value, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.t.QueryContext(ctx, sqlStr, value...)
	if err != nil {
		return nil, err
	}
	var datas []map[string]any
	stores.Scan(rows, &datas)
	retLogs := make([]*deviceLog.Abnormal, 0, len(datas))
	for _, v := range datas {
		retLogs = append(retLogs, ToDeviceLog(v))
	}
	return retLogs, nil
}

func (s *AbnormalLogRepo) Insert(ctx context.Context, data *deviceLog.Abnormal) error {
	sql := fmt.Sprintf("  %s using %s tags('%s','%s')(`ts`, `type`,`reason` ,`action`  ,`trace_id` ) values (?,?,?,?,?) ",
		s.GetLogTableName(data.ProductID, data.DeviceName), s.GetLogStableName(), data.ProductID, data.DeviceName)
	s.t.AsyncInsert(sql, data.Timestamp, data.Type, data.Reason, data.Action, data.TraceID)
	return nil
}