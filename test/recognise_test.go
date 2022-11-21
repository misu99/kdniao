package test

import (
	"github.com/misu99/kdniao"
	"github.com/misu99/kdniao/sdk"
	"testing"
)

func TestRecognise(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	apiRecogniseSdk := sdk.NewApiRecognise(config, logger)
	req := apiRecogniseSdk.GetRequest("550000609031770")
	resp, err := apiRecogniseSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.LogisticCode", resp.LogisticCode)
	for _, shipper := range resp.Shippers {
		t.Log(shipper.ShipperCode, shipper.ShipperName)
	}
}
