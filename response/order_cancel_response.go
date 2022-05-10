package response

type OrderCancelResponse struct {
	EBusinessID string `json:"EBusinessID"` // 用户ID
	Success     bool   `json:"Success"`     // 成功与否
	Reason      string `json:"Reason"`      // 失败原因
	ResultCode  string `json:"ResultCode"`  // 返回编码
}

func (resp *OrderCancelResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *OrderCancelResponse) GetError() string {
	return resp.Reason
}
