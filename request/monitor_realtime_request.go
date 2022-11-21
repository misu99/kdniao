package request

import (
	"encoding/json"
	"github.com/misu99/kdniao/enum"
)

/**
 * 在途监控-国内版
 * 即时查询接口，此接口用于向快递公司实时查询物流轨迹信息
 */

func NewMonitorRealtimeRequest() MonitorRealtimeRequest {
	req := MonitorRealtimeRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_MONITOR_REALTIME)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type MonitorRealtimeRequest struct {
	KdniaoRequest
	OrderCode    string `json:"OrderCode,omitempty"`    // 订单编号
	CustomerName string `json:"CustomerName,omitempty"` // ShipperCode 为 SF 时必填，对应寄件人 /收件人手机号后四位；ShipperCode 为其他快递时，可不填或保 留字段，不可传值
	Sort         int    `json:"Sort,omitempty"`         // 轨迹排序，0-升序，1-降序， 默认 0-升序
	ShipperCode  string `json:"ShipperCode"`            // 快递公司编码
	LogisticCode string `json:"LogisticCode"`           // 快递单号
}

func (req *MonitorRealtimeRequest) SetLogisticCode(logisticCode string) *MonitorRealtimeRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req MonitorRealtimeRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *MonitorRealtimeRequest) SetShipperCode(shipperCode string) *MonitorRealtimeRequest {
	req.ShipperCode = shipperCode
	return req
}

func (req MonitorRealtimeRequest) GetShipperCode() string {
	return req.ShipperCode
}

func (req *MonitorRealtimeRequest) SetOrderCode(orderCode string) *MonitorRealtimeRequest {
	req.OrderCode = orderCode
	return req
}

func (req MonitorRealtimeRequest) GetOrderCode() string {
	return req.OrderCode
}

func (req *MonitorRealtimeRequest) SetSort(sort int) *MonitorRealtimeRequest {
	req.Sort = sort
	return req
}

func (req MonitorRealtimeRequest) GetSort() int {
	return req.Sort
}

func (req *MonitorRealtimeRequest) SetCustomerName(customerName string) *MonitorRealtimeRequest {
	req.CustomerName = customerName
	return req
}

func (req MonitorRealtimeRequest) GetCustomerName() string {
	return req.CustomerName
}

func (req *MonitorRealtimeRequest) UpdateRequestData() *MonitorRealtimeRequest {
	req.requestData = req.ToJson()
	return req
}

func (req MonitorRealtimeRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
