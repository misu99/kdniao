package request

import (
	"encoding/json"
	"github.com/yangzhenrui/kdniao/enum"
)

func NewOrderRequest() OrderRequest {
	req := OrderRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_EORDER)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type OrderRequest struct {
	KdniaoRequest
	CallBack              string                  `json:"CallBack,omitempty"`     // 用户自定义回传字段
	MemberID              string                  `json:"MemberID,omitempty"`     // ERP 系统、电商平台等系统或 平台类型用户的会员 ID 或店 铺账号等唯一性标识，用于区 分其用户
	CustomerName          string                  `json:"CustomerName,omitempty"` // 电子面单客户号
	CustomerPwd           string                  `json:"CustomerPwd,omitempty"`  // 电子面单密码
	SendSite              string                  `json:"SendSite,omitempty"`     // 收件网点标识(名称)
	SendStaff             string                  `json:"SendStaff,omitempty"`
	MonthCode             string                  `json:"MonthCode,omitempty"`                          // 月结编码
	CustomArea            string                  `json:"CustomArea,omitempty"`                         // 商家自定义区域
	WareHouseID           string                  `json:"WareHouseID,omitempty"`                        // 发货仓编码(ShipperCode 为 JD 或 JDKY 时必填)
	TransType             int                     `json:"TransType,omitempty"`                          // 京东快递(ShipperCode 为 JD)的产品类型： 1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8：生鲜专送 16：生鲜特快 17：生鲜特惠 20：函数达 21：特惠包裹 24：特惠小件 26：冷链专送 京东快运(ShipperCode 为 JDKY)的类型： 1：零担 2：整车 3：特惠送 4：特快送 5：特惠运 6：特准运 7：航空代理 8：航空直客 9：铁路直客 10：特惠整车 11：特惠拼车 12：同城直配 13：冷链卡班 14：冷链专车 15：冷链成配 16：送货到仓 17：医药零担 18：医药整车 25：特快重货 百腾物流(ShipperCode 为 BETWL)的类型： 1：慢件 2：快件 3：航空
	ShipperCode           string                  `json:"ShipperCode"`                                  // 快递公司编码
	LogisticCode          string                  `json:"LogisticCode,omitempty"`                       // 快递单号
	ThrOrderCode          string                  `json:"ThrOrderCode,omitempty"`                       // 京东商城的订单号 (ShipperCode 为 JD 且 ExpType 为 1 时必填)
	OrderCode             string                  `json:"OrderCode"`                                    // 订单编号
	PayType               int                     `json:"PayType"`                                      // 邮费支付方式: 1-现付，2-到付，3-月结，4-第三方支付)
	ExpType               string                  `json:"ExpType"`                                      // 快递类型：1-标准快件
	IsReturnSignBill      int                     `json:"IsReturnSignBill,omitempty"`                   // 是否要求签回单 0-不要求，1-要求
	OperateRequire        string                  `json:"OperateRequire,omitempty"`                     // 签回单操作要求(如：签名、盖 章、身份证复印件等)
	Cost                  float64                 `json:"Cost,omitempty"`                               // 快递运费
	OtherCost             float64                 `json:"OtherCost,omitempty"`                          // 快递运费
	Receiver              OrderRequestReceiver    `json:"Receiver"`                                     // 收件信息
	Sender                OrderRequestSender      `json:"Sender"`                                       // 寄件信息
	IsNotice              int                     `json:"IsNotice,omitempty"`                           // 是否通知快递员上门揽件：0-通知；1-不通知；不填则
	StartDate             string                  `json:"StartDate,omitempty"`                          // 上门揽件时间段，格式：YYYY-MM-DD HH24:MM:SS
	EndDate               string                  `json:"EndDate,omitempty"`                            // 上门揽件结束时间，示例： 2021-01-01 17:00:00
	Weight                float64                 `json:"Weight,omitempty"`                             // 包裹总重量kg
	Quantity              int                     `json:"Quantity"`                                     // 包裹数，一个包裹对应一个运单号，如果是大于1个包裹，返回则按照子母件的方式返回母运单号和子运单号
	Volume                float64                 `json:"Volume,omitempty"`                             // 包裹总体积m3
	Remark                string                  `json:"Remark,omitempty"`                             // 备注
	AddService            OrderRequestAddService  `json:"AddService"`                                   // 增值服务
	Commodity             []OrderRequestCommodity `json:"Commodity,omitempty"`                          // 商品信息
	IsReturnPrintTemplate string                  `json:"isReturnPrintTemplate,omitempty"`              // 是否返回电子面单模板： 0-不需要，1-需要
	IsSendMessage         int                     `json:"IsSendMessage,omitempty"`                      // 是否订阅短信：0-不需要；1-需要
	IsSubscribe           string                  `json:"IsSubscribe,omitempty"`                        // 是否订阅轨迹推送 0-不订阅，1-订阅，不填默认 为 1
	TemplateSize          string                  `json:"TemplateSize,omitempty"`                       // 模板规格 (默认的模板无需传值，非默认 模板传对应模板尺寸)
	PackingType           int                     `json:"PackingType,omitempty"`                        // 包装类型(快运字段)； 0-纸，1-纤，2-木，3-托膜， 4-木托，99-其他
	DeliveryMethod        int                     `json:"DeliveryMethod,omitempty"`                     // 送货方式/派送类型/配送方式 (快运字段)： 0-自提 1-送货上门（不含上楼） 2-送货上楼 当 ShipperCode 为 JTSD 时必 填，支持以下传值： 3-派送上门 4-站点自提 5-快递柜自提 6-代收点自提 当 ShipperCode 为 DBL 或 DBLKY 时必填，支持以下传值： 1-自提 2-送货进仓 3-送货（不含上楼） 4-送货上楼 5-大件上楼 当 ShipperCode 为 ZYE 时必 填，支持以下传值： 1-送货上门 2-自提
	CurrencyCode          string                  `json:"CurrencyCode,omitempty"`                       // 货物单价的币种： CNY: 人民币 HKD: 港币 NTD: 新台币 MOP: 澳门元 (ShipperCode 为 SF 且收件地址为港澳台地区，必填)
	Dutiable              Dutiable                `json:"Dutiable,omitempty" json:"Dutiable,omitempty"` // 申报信息
}

type OrderRequestReceiver struct {
	Company      string `json:"Company,omitempty"`  // 收件人公司
	Name         string `json:"Name"`               // 收件人
	Tel          string `json:"Tel,omitempty"`      // 电话，电话与手机，必填一个
	Mobile       string `json:"Mobile,omitempty"`   // 手机，电话与手机，必填一个
	PostCode     string `json:"PostCode,omitempty"` // 收件地邮编 (ShipperCode为 EMS、YZPY 时必填)
	ProvinceName string `json:"ProvinceName"`       // 收件省 (如广东省，不要缺少“省”；如是直辖市，请直接传北京、上海等；如是自治区，请直接传广西壮族自治区等)
	CityName     string `json:"CityName"`           // 收件市 (如深圳市，不要缺少“市”)
	ExpAreaName  string `json:"ExpAreaName"`        // 收件区/县(如福田区，不要缺少“区”或“县”)
	Address      string `json:"Address"`            // 收件人详细地址
}

type OrderRequestSender struct {
	Company      string `json:"Company,omitempty"`  // 发件人公司
	Name         string `json:"Name"`               // 发件人
	Tel          string `json:"Tel,omitempty"`      // 电话，电话与手机，必填一个
	Mobile       string `json:"Mobile,omitempty"`   // 手机，电话与手机，必填一个
	PostCode     string `json:"PostCode,omitempty"` // 发件地邮编 (ShipperCode为 EMS、YZPY 时必填)
	ProvinceName string `json:"ProvinceName"`       // 发件省 (如广东省，不要缺少“省”；如是直辖市，请直接传北京、上海等；如是自治区，请直接传广西壮族自治区等)
	CityName     string `json:"CityName"`           // 发件市 (如深圳市，不要缺少“市”)
	ExpAreaName  string `json:"ExpAreaName"`        // 发件区/县(如福田区，不要缺少“区”或“县”)
	Address      string `json:"Address"`            // 发件详细地址

}

type OrderRequestAddService struct {
	Name       string `json:"Name,omitempty"`       // 增值服务名称
	Value      string `json:"Value,omitempty"`      // 增值服务值
	CustomerId string `json:"CustomerID,omitempty"` // 客户标识
}

type OrderRequestCommodity struct {
	GoodsName     string  `json:"GoodsName"`               // 商品名称
	GoodsCode     string  `json:"GoodsCode,omitempty"`     // 商品编码
	Goodsquantity int     `json:"Goodsquantity,omitempty"` // 商品件数
	GoodsPrice    float64 `json:"GoodsPrice,omitempty"`    // 商品价格
	GoodsWeight   float64 `json:"GoodsWeight,omitempty"`   // 商品重量kg
	GoodsDesc     string  `json:"GoodsDesc,omitempty"`     // 商品描述
	GoodsVol      float64 `json:"GoodsVol,omitempty"`      // 商品体积m3
}

type Dutiable struct {
	DeclaredValue float64 `json:"DeclaredValue"` // 申报价值：订单货物总声明价 值，包含子母件，精确到小数 点后 3 位 (ShipperCode 为 SF 且收件 地址为港澳台地区，必填)
}

func (req *OrderRequest) UpdateRequestData() *OrderRequest {
	req.requestData = req.ToJson()
	return req
}

func (req OrderRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
