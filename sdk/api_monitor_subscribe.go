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

func NewApiMonitorSubscribe(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) *ApiMonitorSubscribe {
	return &ApiMonitorSubscribe{config, logger}
}

type ApiMonitorSubscribe struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj *ApiMonitorSubscribe) GetRequest(params request.MonitorSubscribeRequest) request.MonitorSubscribeRequest {
	req := request.NewMonitorSubscribeRequest()
	req.KdniaoRequest.SetConfig(obj.config)
	copier.Copy(&req, &params)
	return req
}

func (obj ApiMonitorSubscribe) GetResponse(req request.MonitorSubscribeRequest) (response.MonitorSubscribeResponse, error) {
	url := enum.GATEWAY + enum.URI_API

	req.UpdateRequestData()
	var resp response.MonitorSubscribeResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMonitorSubscribe,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiMonitorSubscribe,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMonitorSubscribe,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiMonitorSubscribe," + resp.GetError())
	}
	return resp, nil
}
