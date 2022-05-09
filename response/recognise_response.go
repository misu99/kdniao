package response

type RecogniseResponse struct {
	EBusinessId  string                      `json:"EBusinessID"`  // 用户 ID
	LogisticCode string                      `json:"LogisticCode"` // 快递单号
	Success      bool                        `json:"Success"`      // 成功与否(true/false)
	Code         string                      `json:"Code"`         // 失败原因
	Shippers     []RecogniseResponseShippers `json:"Shippers"`
	Reason       string                      `json:"Reason"`
}

type RecogniseResponseShippers struct {
	ShipperCode string `json:"ShipperCode"` // 快递公司编码
	ShipperName string `json:"ShipperName"` // 快递公司名称
}

func (resp *RecogniseResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *RecogniseResponse) GetError() string {
	msg := ""
	if "" != resp.Code {
		msg = "[" + resp.Code + "]"
	}
	if "" != resp.Reason {
		msg = msg + resp.Reason
	}
	return msg
}
