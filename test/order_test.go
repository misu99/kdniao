package test

import (
	"github.com/yangzhenrui/kdniao"
	"github.com/yangzhenrui/kdniao/request"
	"github.com/yangzhenrui/kdniao/sdk"
	"testing"
)

func TestOrder(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	apiSdk := sdk.NewOrder(config, logger)
	reqParams := request.OrderRequest{}
	req := apiSdk.GetRequest(reqParams)
	resp, err := apiSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessID)
	//t.Log("resp.OrderCode", resp.OrderCode)
	//t.Log("resp.ShipperCode", resp.ShipperCode)
	//t.Log("resp.LogisticCode", resp.LogisticCode)
	//t.Log("resp.State", resp.State)
	//t.Log("resp.StateEx", resp.StateEx)
	//t.Log("resp.Location", resp.Location)
	//for _, trace := range resp.Traces {
	//	t.Log(trace.AcceptTime, "|", trace.AcceptStation, "|", trace.Location, "|", trace.Action, "|", trace.Remark)
	//}
}
