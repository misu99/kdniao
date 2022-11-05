package request

func NewOrderPrintRequest() OrderPrintRequest {
	req := OrderPrintRequest{}
	return req
}

type OrderPrintRequest struct {
	KdniaoRequest
	RequestData []RequestData `json:"RequestData"` // 请求体，涵盖：下单订单号 和 打印机名称
	IsPreview   string        `json:"IsPreview"`   // 是否预览，1是，0否
	IP          string        `json:"IP"`
}

type RequestData struct {
	OrderCode string `json:"OrderCode"` // 订单编号
	PortName  string `json:"PortName"`  // 打印机名称
}
