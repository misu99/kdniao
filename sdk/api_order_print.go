package sdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/misu99/kdniao"
	"github.com/misu99/kdniao/enum"
	"github.com/misu99/kdniao/request"
	"github.com/misu99/kdniao/util"
)

func NewOrderPrint(config kdniao.KdniaoConfig, logger kdniao.KdniaoLoggerInterface) ApiOrderPrint {
	return ApiOrderPrint{config, logger}
}

type ApiOrderPrint struct {
	config kdniao.KdniaoConfig
	logger kdniao.KdniaoLoggerInterface
}

func (obj *ApiOrderPrint) GetRequest(params request.OrderPrintRequest) request.OrderPrintRequest {
	req := request.NewOrderPrintRequest()
	req.KdniaoRequest.SetConfig(obj.config)
	copier.Copy(&req, &params)
	return req
}

func (obj *ApiOrderPrint) Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

//func (obj *ApiOrderPrint) GetIp() string {
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return ""
//	}
//
//	for _, address := range addrs {
//		// 检查ip地址判断是否回环地址
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				return ipnet.IP.String()
//			}
//		}
//	}
//	return ""
//} //本地ip地址，示范

type Resp struct {
	RequestData string `json:"RequestData"`
	EBusinessId string `json:"EBusinessId"`
	DataSign    string `json:"DataSign"`
	IsPreview   string `json:"IsPreview"`
}

func (obj *ApiOrderPrint) BuildForm(req request.OrderPrintRequest) (resp Resp) {
	url := enum.GATEWAY + enum.URI_PRINT
	requestDataJson, _ := json.Marshal(req.RequestData)
	dataSign := obj.DataSign(req.IP, string(requestDataJson), obj.config.GetAppKey())
	form := "<form id='form1' method='POST' action='" + url + "'><input type='text' name='RequestData' value='" + string(requestDataJson) + "'/><input type='text' name='EBusinessID' value='" + obj.config.GetEBusinessId() + "'/><input type='text' name='DataSign' value='" + dataSign + "'/><input type='text' name='IsPriview' value='" + req.IsPreview + "'/></form><script>form1.submit();</script>"
	fmt.Println(form)

	resp.RequestData = string(requestDataJson)
	resp.EBusinessId = obj.config.GetEBusinessId()
	resp.DataSign = dataSign
	resp.IsPreview = req.IsPreview
	return resp
}

func (obj *ApiOrderPrint) DataSign(ip string, requestDataJson string, apiKey string) string {
	str := util.Md5(ip + requestDataJson + apiKey)
	return base64.StdEncoding.EncodeToString([]byte(str))
}
