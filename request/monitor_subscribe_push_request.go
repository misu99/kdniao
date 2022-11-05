package request

/**
 * 物流轨迹推送，
 * 轨迹订阅接口
 * 1、此接口用于向快递鸟订阅物流轨迹信息
 * 2、支持服务：预约取件
 */

type MonitorSubscribePushRequest struct {
	EBusinessID           string           `json:"EBusinessID"` // 用户ID
	PushTime              string           `json:"PushTime"`    // 推送时间，示例 2021-01-01 09:00:00
	Count                 string           `json:"Count"`       // 推送轨迹的快递单号个数
	Data                  []Data           `json:"Data"`
	ReceiverCityLatAndLng string           `json:"ReceiverCityLatAndLng"` // 收件城市经纬度
	SenderCityLatAndLng   string           `json:"SenderCityLatAndLng"`   // 发件城市经纬度
	Coordinates           []Coordinates    `json:"Coordinates"`
	PreCoordinates        []PreCoordinates `json:"PreCoordinates"` //
	RouteMapUrl           string           `json:"RouteMapUrl"`    // 轨迹地图 URL
}

type Coordinates struct {
	Location  string `json:"Location"`  // 预设路径经过城市
	LatAndLng string `json:"LatAndLng"` // 当前城市经纬度
}

type PreCoordinates struct {
	Location  string `json:"Location"`  // 预设路径经过城市
	LatAndLng string `json:"LatAndLng"` // 预设路径城市经纬度
}

type Data struct {
	EBusinessID    string   `json:"EBusinessID"`    // 用户 ID
	ShipperCode    string   `json:"ShipperCode"`    // 快递公司编码
	LogisticCode   string   `json:"LogisticCode"`   // 快递单号
	Success        bool     `json:"Success"`        // 成功与否(true/false)
	Reason         string   `json:"Reason"`         // 失败原因
	OrderCode      string   `json:"OrderCode"`      // 订单编号
	State          string   `json:"State"`          // 普通物流状态： 0-暂无轨迹信息 1-已揽收 2-在途中 3-签收 4-问题件
	StateEx        string   `json:"StateEx"`        // 增值物流状态： 0-暂无轨迹信息 1-已揽收 2-在途中 201-到达派件城市 202-派件中 211-已放入快递柜或驿 站3-已签收 301-正常签收 302-派件异常后最终 签收304-代收签收 311-快递柜或驿站签收 4-问题件 401-发货无信息 402-超时未签收 403-超时未更新 404-拒收(退件) 405-派件异常 406-退货签收 407-退货未签收 412-快递柜或驿站超时 未取413-单号已拦截 10-待揽件
	Location       string   `json:"Location"`       // 所在城市
	Callback       string   `json:"Callback"`       // 用户自定义回传字段
	Traces         []Traces `json:"Traces"`         //
	Station        string   `json:"Station"`        // 派件网点的名称
	StationTel     string   `json:"StationTel"`     // 派件网点的电话
	StationAdd     string   `json:"StationAdd"`     // 派件网点的地址
	DeliveryMan    string   `json:"DeliveryMan"`    // 派件快递员
	DeliveryManTel string   `json:"DeliveryManTel"` // 派件快递员手机号
	NextCity       string   `json:"NextCity"`       // 下一站城市
}

type Traces struct {
	AcceptTime    string `json:"AcceptTime"`    // 轨迹发生时间，示例： 2021-01-01 09:00:00
	AcceptStation string `json:"AcceptStation"` // 轨迹描述
	Location      string `json:"Location"`      // 当前城市
	Action        string `json:"Action"`        // 增值物流状态：0-暂无轨迹信息 1-已揽收 2-在途中 201-到达派件城市 202-派件中 211-已放入快递柜或驿 站3-已签收 301-正常签收 302-派件异常后最终 签收304-代收签收 311-快递柜或驿站签收 4-问题件 401-发货无信息 402-超时未签收 403-超时未更新 404-拒收(退件) 405-派件异常 406-退货签收 407-退货未签收 412-快递柜或驿站超时 未取10-待揽件
	Remark        string `json:"Remark"`        // 备注
}
