// Code generated by goctl. DO NOT EDIT.
package types

type ProjectDetail struct {
	ID      int64  `json:"id,string,optional"` //项目id 只读
	Content string `json:"content"`
}

type ProjectInfoReadReq struct {
	ID int64 `json:"id,string"` //项目id
}

type ProjectInfo struct {
	ID            int64  `json:"id,string,optional"`            //项目id 只读
	IndexImage    string `json:"indexImage,optional"`           //图片地址
	Name          string `json:"name,optional"`                 //项目名称
	Desc          string `json:"desc,optional"`                 //项目描述
	CreatedUserID int64  `json:"createdUserID,string,optional"` //创建者id
	Status        int64  `json:"status,optional"`               //项目状态 1: 已发布 2: 未发布
}

type ProjectInfoDeleteReq struct {
	ID int64 `json:"id,string"` //项目id
}

type ProjectInfoIndexReq struct {
	Page *PageInfo `json:"page,optional"` //分页信息,只获取一个则不填
}

type ProjectInfoIndexResp struct {
	List  []*ProjectInfo `json:"list"`           //项目信息
	Total int64          `json:"total,optional"` //拥有的总数
	Num   int64          `json:"num,optional"`   //返回的数量
}

type PageInfo struct {
	Page int64 `json:"page,optional" form:"page,optional"` // 页码
	Size int64 `json:"size,optional" form:"size,optional"` // 每页大小
}

type CommonResp struct {
	ID int64 `json:"id,optional"` // id
}

type WithID struct {
	ID int64 `json:"id,optional"` // id
}

type WithIDOrCode struct {
	ID   int64  `json:"id,optional"` // id
	Code string `json:"code,optional"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Point struct {
	Longitude float64 `json:"longitude,range=[0:180]"` //经度
	Latitude  float64 `json:"latitude,range=[0:90]"`   //纬度
}

type DateRange struct {
	Start string `json:"start,optional"` //开始时间 格式：yyyy-mm-dd
	End   string `json:"end,optional"`   //结束时间 格式：yyyy-mm-dd
}

type TimeRange struct {
	Start int64 `json:"start,optional"` //开始时间 unix时间戳
	End   int64 `json:"end,optional"`   //结束时间 unix时间戳
}

type SendOption struct {
	TimeoutToFail  int64 `json:"timeoutToFail,optional"`  //超时失败时间
	RequestTimeout int64 `json:"requestTimeout,optional"` //请求超时,超时后会进行重试
	RetryInterval  int64 `json:"retryInterval,optional"`  //重试间隔
}

type CodeReq struct {
	Code string `json:"code"`
}