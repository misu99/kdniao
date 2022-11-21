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

func NewApiMonitorRealtime(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) *ApiMonitorRealtime {
	return &ApiMonitorRealtime{config, logger}
}

type ApiMonitorRealtime struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj *ApiMonitorRealtime) GetRequest(reqParams request.MonitorRealtimeRequest) request.MonitorRealtimeRequest {
	req := request.NewMonitorRealtimeRequest()
	req.KdniaoRequest.SetConfig(obj.config)
	copier.Copy(&req, &reqParams)
	return req
}

func (obj ApiMonitorRealtime) GetResponse(req request.MonitorRealtimeRequest) (response.MonitorRealtimeResponse, error) {
	url := enum.GATEWAY + enum.URI_BUSINESS

	req.UpdateRequestData()
	var resp response.MonitorRealtimeResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMonitorRealtime,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiMonitorRealtime,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMonitorRealtime,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiMonitorRealtime," + resp.GetError())
	}
	return resp, nil
}
