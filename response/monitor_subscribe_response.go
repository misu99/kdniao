package response

type MonitorSubscribeResponse struct {
	EBusinessId  string `json:"EBusinessID"`  // 用户ID
	UpdateTime   string `json:"UpdateTime"`   // 更新时间，示例： 2021-01-01 09:00:00
	Success      bool   `json:"Success"`      // 成功与否(true/false)
	ShipperCode  string `json:"ShipperCode"`  // 快递公司编码
	LogisticCode string `json:"LogisticCode"` // 快递单号
	Reason       string `json:"Reason"`
}

func (resp *MonitorSubscribeResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *MonitorSubscribeResponse) GetError() string {
	return resp.Reason
}
