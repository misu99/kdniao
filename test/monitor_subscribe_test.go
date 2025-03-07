package test

import (
	"github.com/misu99/kdniao"
	"github.com/misu99/kdniao/request"
	"github.com/misu99/kdniao/sdk"
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
	reqParams.LogisticCode = "4303618027892"
	reqParams.ShipperCode = "YD"
	req := apiMonitorSubscribeSdk.GetRequest(reqParams)
	resp, err := apiMonitorSubscribeSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessID)
	t.Log("resp.UpdateTime", resp.UpdateTime)
}
