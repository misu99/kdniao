package sdk

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/yangzhenrui/kdniao"
	"github.com/yangzhenrui/kdniao/enum"
	"github.com/yangzhenrui/kdniao/request"
	"github.com/yangzhenrui/kdniao/response"
	"github.com/yangzhenrui/kdniao/util"
	"github.com/yangzhenrui/kdniao/util/http"
	"strconv"
)

func NewOrderCancel(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiOrderCancel {
	return ApiOrderCancel{config, logger}
}

type ApiOrderCancel struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj *ApiOrderCancel) GetRequest(params request.OrderCancelRequest) request.OrderCancelRequest {
	req := request.NewOrderCencelRequest()
	req.KdniaoRequest.SetConfig(obj.config)
	copier.Copy(&req, &params)
	return req
}

func (obj ApiOrderCancel) GetResponse(req request.OrderCancelRequest) (response.OrderCancelResponse, error) {
	url := enum.GATEWAY + enum.URI_ORDER

	req.UpdateRequestData()
	var resp response.OrderCancelResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiRecognise,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiRecognise,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiRecognise,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiRecognise," + resp.GetError())
	}
	return resp, nil
}
