package response

type MonitorSubscribeResponse struct {
	EBusinessId  string `json:"eBusinessId"`  // 用户ID
	UpdateTime   string `json:"updateTime"`   // 更新时间，示例： 2021-01-01 09:00:00
	Success      bool   `json:"success"`      // 成功与否(true/false)
	ShipperCode  string `json:"shipperCode"`  // 快递公司编码
	LogisticCode string `json:"logisticCode"` // 快递单号
	Reason       string `json:"reason"`
}

func (resp *MonitorSubscribeResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *MonitorSubscribeResponse) GetError() string {
	return resp.Reason
}
