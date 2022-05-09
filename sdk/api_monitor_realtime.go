package sdk

import (
	"encoding/json"
	"github.com/yangzhenrui/kdniao"
	"github.com/yangzhenrui/kdniao/enum"
	"github.com/yangzhenrui/kdniao/request"
	"github.com/yangzhenrui/kdniao/response"
	"github.com/yangzhenrui/kdniao/util"
	"github.com/yangzhenrui/kdniao/util/http"
	"strconv"
)

func NewApiMonitorRealtime(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiMonitorRealtime {
	return ApiMonitorRealtime{config, logger}
}

type ApiMonitorRealtime struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj ApiMonitorRealtime) GetRequest(logisticCode, shipperCode string) request.MonitorRealtimeRequest {
	req := request.NewMonitorRealtimeRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	req.SetShipperCode(shipperCode)
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
