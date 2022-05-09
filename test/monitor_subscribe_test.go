package test

import (
	"github.com/yangzhenrui/kdniao"
	"github.com/yangzhenrui/kdniao/request"
	"github.com/yangzhenrui/kdniao/sdk"
	"testing"
)

func TestMonitorSubscribe(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	apiMonitorSubscribeSdk := sdk.NewApiMonitorSubscribe(config, logger)
	reqParams := request.MonitorSubscribeRequest{}
	req := apiMonitorSubscribeSdk.GetRequest(reqParams)
	resp, err := apiMonitorSubscribeSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.UpdateTime", resp.UpdateTime)
}
