package request

import (
	"encoding/json"
	"github.com/misu99/kdniao/enum"
)

func NewOrderCencelRequest() OrderCancelRequest {
	req := OrderCancelRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_EORDER_CANCEL)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type OrderCancelRequest struct {
	KdniaoRequest
	ShipperCode  string `json:"ShipperCode"`            // 快递公司编码
	OrderCode    string `json:"OrderCode"`              // 订单编号
	ExpNo        string `json:"ExpNo"`                  // 快递单号
	CustomerName string `json:"CustomerName,omitempty"` // 电子面单客户号
	CustomerPwd  string `json:"CustomerPwd,omitempty"`  // 电子面单密码
}

func (req *OrderCancelRequest) UpdateRequestData() *OrderCancelRequest {
	req.requestData = req.ToJson()
	return req
}

func (req OrderCancelRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
