package response

type RecogniseResponse struct {
	EBusinessId  string                      `json:"eBusinessId"`  // 用户 ID
	LogisticCode string                      `json:"logisticCode"` // 快递单号
	Success      bool                        `json:"success"`      // 成功与否(true/false)
	Code         string                      `json:"code"`         // 失败原因
	Shippers     []RecogniseResponseShippers `json:"shippers"`
	Reason       string                      `json:"reason"`
}

type RecogniseResponseShippers struct {
	ShipperCode string `json:"shipperCode"` // 快递公司编码
	ShipperName string `json:"shipperName"` // 快递公司名称
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
