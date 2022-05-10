package response

type OrderResponse struct {
	EBusinessID           string `json:"EBusinessID"`
	Order                 Order  `json:"Order"`
	Success               bool   `json:"Success"`               // 成功与否
	ResultCode            string `json:"ResultCode"`            // 返回编码
	Reason                string `json:"Reason"`                // 失败原因
	UniquerRequestNumber  string `json:"UniquerRequestNumber"`  // 唯一标识
	PrintTemplate         string `json:"PrintTemplate"`         // 面单打印模板内容(html 格式)
	SubCount              int    `json:"SubCount"`              // 子单数量
	SubOrders             string `json:"SubOrders"`             // 子单单号（数组形式）
	SubPrintTemplates     string `json:"SubPrintTemplates"`     // 子单模板内容(html 格式)
	SignBillPrintTemplate string `json:"SignBillPrintTemplate"` // 签回单模板内容(html 格式)
}

type Order struct {
	OrderCode                   string `json:"OrderCode"`                   // 订单编号
	ShipperCode                 string `json:"AcceptStation"`               // 快递公司编码
	LogisticCode                string `json:"LogisticCode"`                // 快递单号
	MarkDestination             string `json:"MarkDestination"`             // 大头笔
	SignWaybillCode             string `json:"SignWaybillCode"`             // 签回单单号
	OriginCode                  string `json:"OriginCode"`                  // 始发地区域编码
	OriginName                  string `json:"OriginName"`                  // 始发地/始发网点
	DestinatioCode              string `json:"DestinatioCode"`              // 目的地区域编码
	DestinatioName              string `json:"DestinatioName"`              // 目的地/到达网点
	SortingCode                 string `json:"SortingCode"`                 // 分拣编码
	PackageCode                 string `json:"PackageCode"`                 // 集包编码
	PackageName                 string `json:"PackageName"`                 // 集包地
	DestinationAllocationCentre string `json:"DestinationAllocationCentre"` // 目的地分拨
	TransType                   string `json:"TransType"`                   // 京东快递(ShipperCode 为 JD)的 产品类型： 1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8：生鲜专送 16：生鲜特快 17：生鲜特惠 20：函数达 21：特惠包裹 24：特惠小件 26：冷链专送 京东快运(ShipperCode 为 JDKY) 的类型： 1：零担 2：整车 3：特惠送 4：特快送 5：特惠运 6：特准运 7：航空代理 8：航空直客 9：铁路直客 10：特惠整车 11：特惠拼车 12：同城直配 13：冷链卡班 14：冷链专车 15：冷链成配 16：送货到仓 17：医药零担 18：医药整车 25：特快重货 百腾物流(ShipperCode 为 BETWL)的类型： 1：慢件 2：快件 3：航空
	TransportType               int    `json:"TransportType"`               // 运输方式(用于自行设计京东模 板)： 0：陆运 1：航空
	ShipperInfo                 string `json:"ShipperInfo"`                 // 自行设计模板用 (仅 ShipperCode 为 SF 时返回)
}

func (resp *OrderResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *OrderResponse) GetError() string {
	return resp.Reason
}
