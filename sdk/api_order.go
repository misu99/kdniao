package sdk

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/misu99/kdniao"
	"github.com/misu99/kdniao/enum"
	"github.com/misu99/kdniao/request"
	"github.com/misu99/kdniao/response"
	"github.com/misu99/kdniao/util"
	"github.com/misu99/kdniao/util/http"
	"strconv"
)

func NewOrder(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiOrder {
	return ApiOrder{config, logger}
}

type ApiOrder struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj *ApiOrder) GetRequest(params request.OrderRequest) request.OrderRequest {
	req := request.NewOrderRequest()
	req.KdniaoRequest.SetConfig(obj.config)
	copier.Copy(&req, &params)
	return req
}

func (obj ApiOrder) GetResponse(req request.OrderRequest) (response.OrderResponse, error) {
	url := enum.GATEWAY + enum.URI_ORDER

	req.UpdateRequestData()
	var resp response.OrderResponse
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

// 沙箱
func (obj ApiOrder) GetTestResponse(req request.OrderRequest) (response.OrderResponse, error) {
	url := enum.GATEWAY_SANDBOX

	req.UpdateRequestData()
	var resp response.OrderResponse
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
