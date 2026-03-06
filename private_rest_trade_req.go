package myxcoinapi

type PrivateRestTradeOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderReq
}
type PrivateRestTradeOrderReq struct {
	AccountName   *string    `json:"accountName"`   // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol        *string    `json:"symbol"`        // true 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	Side          *string    `json:"side"`          // true 订单方向 buy：买，sell：卖
	OrderType     *string    `json:"orderType"`     // true 订单类型 market：市价单，limit：限价单，post_only：只挂单
	Qty           *string    `json:"qty"`           // true 委托数量
	MarketUnit    *string    `json:"marketUnit"`    // false qty的单位 现货交易默认单位为基础币种，市价单可以选择基础币种或者计价币种, baseCoin：基础币种， 如买BTC-USDT，则qty的单位是BTC, quoteCoin：计价币种， 如卖BTC-USDT，则 qty的单位是USDT,对于合约交易，默认单位为币
	Price         *string    `json:"price"`         // true 委托价格 limit、post_only类型的订单，指委托价格；market：为空
	TimeInForce   *string    `json:"timeInForce"`   // true 订单生效类型, gtc：长期有效直至取消, ioc：立即成交并取消剩余, fok：全部成交或立即取消
	ClientOrderId *string    `json:"clientOrderId"` // false 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
	ReduceOnly    *bool      `json:"reduceOnly"`    // false 是否只减仓 true或false，默认false, 仅适用于合约交易
	TpslOrder     *TpslOrder `json:"tpslOrder"`     // false 下单附带止盈止损信息
}

type TpslOrder struct {
	TpslClOrdId    *string `json:"tpslClOrdId"`    // false 下单附带止盈止损订单时，客户定义止盈止损单触发执行订单的clientOrderId
	TakeProfitType *string `json:"takeProfitType"` // false 止盈触发价格类型 字典：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
	StopLossType   *string `json:"stopLossType"`   // false 止损触发价格类型 字典：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
	TakeProfit     *string `json:"takeProfit"`     // false 止盈触发价格
	StopLoss       *string `json:"stopLoss"`       // false 止损触发价格
	TpOrderType    *string `json:"tpOrderType"`    // false 止盈执行价格类型 字典：market：市价单，默认值；limit：限价单
	SlOrderType    *string `json:"slOrderType"`    // false 止损执行价格类型 字典：market：市价单，默认值；limit：限价单
	TpLimitPrice   *string `json:"tpLimitPrice"`   // false 止盈执行价格类型为limit时，止盈执行价; 止盈执行价格类型为market时, 必须输入null值
	SlLimitPrice   *string `json:"slLimitPrice"`   // false 止损执行价格类型为limit时，止损执行价; 止损执行价格类型为market时, 必须输入null值
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeOrderAPI) AccountName(accountName string) *PrivateRestTradeOrderAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string true 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeOrderAPI) Symbol(symbol string) *PrivateRestTradeOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string true 订单方向 buy：买，sell：卖
func (api *PrivateRestTradeOrderAPI) Side(side string) *PrivateRestTradeOrderAPI {
	api.req.Side = GetPointer(side)
	return api
}

// string true 订单类型 market：市价单，limit：限价单，post_only：只挂单
func (api *PrivateRestTradeOrderAPI) OrderType(orderType string) *PrivateRestTradeOrderAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

// string true 委托数量
func (api *PrivateRestTradeOrderAPI) Qty(qty string) *PrivateRestTradeOrderAPI {
	api.req.Qty = GetPointer(qty)
	return api
}

// string false qty的单位 现货交易默认单位为基础币种，市价单可以选择基础币种或者计价币种, baseCoin：基础币种， 如买BTC-USDT，则qty的单位是BTC, quoteCoin：计价币种， 如卖BTC-USDT，则 qty的单位是USDT,对于合约交易，默认单位为币
func (api *PrivateRestTradeOrderAPI) MarketUnit(marketUnit string) *PrivateRestTradeOrderAPI {
	api.req.MarketUnit = GetPointer(marketUnit)
	return api
}

// string true 委托价格 limit、post_only类型的订单，指委托价格；market：为空
func (api *PrivateRestTradeOrderAPI) Price(price string) *PrivateRestTradeOrderAPI {
	api.req.Price = GetPointer(price)
	return api
}

// string true 订单生效类型, gtc：长期有效直至取消, ioc：立即成交并取消剩余, fok：全部成交或立即取消
func (api *PrivateRestTradeOrderAPI) TimeInForce(timeInForce string) *PrivateRestTradeOrderAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// string false 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
func (api *PrivateRestTradeOrderAPI) ClientOrderId(clientOrderId string) *PrivateRestTradeOrderAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

// bool false 是否只减仓 true或false，默认false, 仅适用于合约交易
func (api *PrivateRestTradeOrderAPI) ReduceOnly(reduceOnly bool) *PrivateRestTradeOrderAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// string false 下单附带止盈止损订单时，客户定义止盈止损单触发执行订单的clientOrderId
func (api *PrivateRestTradeOrderAPI) TpslClOrdId(tpslClOrdId string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.TpslClOrdId = GetPointer(tpslClOrdId)
	return api
}

// string false 止盈触发价格类型 字典：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
func (api *PrivateRestTradeOrderAPI) TakeProfitType(takeProfitType string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.TakeProfitType = GetPointer(takeProfitType)
	return api
}

// string false 止损触发价格类型 字典：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
func (api *PrivateRestTradeOrderAPI) StopLossType(stopLossType string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.StopLossType = GetPointer(stopLossType)
	return api
}

// string false 止盈触发价格
func (api *PrivateRestTradeOrderAPI) TakeProfit(takeProfit string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.TakeProfit = GetPointer(takeProfit)
	return api
}

// string false 止损触发价格
func (api *PrivateRestTradeOrderAPI) StopLoss(stopLoss string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.StopLoss = GetPointer(stopLoss)
	return api
}

// string false 止盈执行价格类型 字典：market：市价单，默认值；limit：限价单
func (api *PrivateRestTradeOrderAPI) TpOrderType(tpOrderType string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.TpOrderType = GetPointer(tpOrderType)
	return api
}

// string false 止损执行价格类型 字典：market：市价单，默认值；limit：限价单
func (api *PrivateRestTradeOrderAPI) SlOrderType(slOrderType string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.SlOrderType = GetPointer(slOrderType)
	return api
}

// string false 止盈执行价格类型为limit时，止盈执行价; 止盈执行价格类型为market时, 必须输入null值
func (api *PrivateRestTradeOrderAPI) TpLimitPrice(tpLimitPrice string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.TpLimitPrice = GetPointer(tpLimitPrice)
	return api
}

// string false 止损执行价格类型为limit时，止损执行价; 止损执行价格类型为market时, 必须输入null值
func (api *PrivateRestTradeOrderAPI) SlLimitPrice(slLimitPrice string) *PrivateRestTradeOrderAPI {
	if api.req.TpslOrder == nil {
		api.req.TpslOrder = &TpslOrder{}
	}
	api.req.TpslOrder.SlLimitPrice = GetPointer(slLimitPrice)
	return api
}

type PrivateRestTradeBatchOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeBatchOrderReq
}

type PrivateRestTradeBatchOrderReq struct {
	OrderReqList []*PrivateRestTradeOrderReq `json:"orderReqList"`
}

func (api *PrivateRestTradeBatchOrderAPI) AddOrderReq(orderReq *PrivateRestTradeOrderReq) *PrivateRestTradeBatchOrderAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeBatchOrderReq{}
	}
	api.req.OrderReqList = append(api.req.OrderReqList, orderReq)
	return api
}

type PrivateRestTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelOrderReq
}

type PrivateRestTradeCancelOrderReq struct {
	AccountName   *string `json:"accountName"`   // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol        *string `json:"symbol"`        // true 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	OrderFilter   *string `json:"orderFilter"`   // false 订单品种 order：普通订单，默认值; oco：止盈止损单
	OrderId       *string `json:"orderId"`       // false 订单ID
	ClientOrderId *string `json:"clientOrderId"` // false 客户端订单ID，和订单ID必须一个有值
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeCancelOrderAPI) AccountName(accountName string) *PrivateRestTradeCancelOrderAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string true 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeCancelOrderAPI) Symbol(symbol string) *PrivateRestTradeCancelOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 订单品种 order：普通订单，默认值; oco：止盈止损单
func (api *PrivateRestTradeCancelOrderAPI) OrderFilter(orderFilter string) *PrivateRestTradeCancelOrderAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// string false 订单ID
func (api *PrivateRestTradeCancelOrderAPI) OrderId(orderId string) *PrivateRestTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// string false 客户端订单ID，和订单ID必须一个有值
func (api *PrivateRestTradeCancelOrderAPI) ClientOrderId(clientOrderId string) *PrivateRestTradeCancelOrderAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

type PrivateRestTradeBatchCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeBatchCancelOrderReq
}

type PrivateRestTradeBatchCancelOrderReq struct {
	OrderReqList []*PrivateRestTradeCancelOrderReq `json:"orderCancelReqList"`
}

func (api *PrivateRestTradeBatchCancelOrderAPI) AddOrderReq(orderReq *PrivateRestTradeCancelOrderReq) *PrivateRestTradeBatchCancelOrderAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeBatchCancelOrderReq{}
	}
	api.req.OrderReqList = append(api.req.OrderReqList, orderReq)
	return api
}

type PrivateRestTradeCancelAllOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelAllOrderReq
}

type PrivateRestTradeCancelAllOrderReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）、BTC-USDT-26DEC25-9100-C（期权）等
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	OrderFilter  *string `json:"orderFilter"`  // false 订单品种 order：普通订单，默认值；oco：止盈止损单
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeCancelAllOrderAPI) AccountName(accountName string) *PrivateRestTradeCancelAllOrderAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）、BTC-USDT-26DEC25-9100-C（期权）等
func (api *PrivateRestTradeCancelAllOrderAPI) Symbol(symbol string) *PrivateRestTradeCancelAllOrderAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeCancelAllOrderAPI) BusinessType(businessType string) *PrivateRestTradeCancelAllOrderAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false 订单品种 order：普通订单，默认值；oco：止盈止损单
func (api *PrivateRestTradeCancelAllOrderAPI) OrderFilter(orderFilter string) *PrivateRestTradeCancelAllOrderAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

type PrivateRestTradeOpenOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOpenOrdersReq
}

type PrivateRestTradeOpenOrdersReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	OrderFilter  *string `json:"orderFilter"`  // false 订单品种 order：普通订单，默认值；oco：止盈止损单
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeOpenOrdersAPI) AccountName(accountName string) *PrivateRestTradeOpenOrdersAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 订单品种 order：普通订单，默认值；oco：止盈止损单
func (api *PrivateRestTradeOpenOrdersAPI) OrderFilter(orderFilter string) *PrivateRestTradeOpenOrdersAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeOpenOrdersAPI) Symbol(symbol string) *PrivateRestTradeOpenOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeOpenOrdersAPI) BusinessType(businessType string) *PrivateRestTradeOpenOrdersAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PrivateRestTradeOrderInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderInfoReq
}

type PrivateRestTradeOrderInfoReq struct {
	AccountName   *string `json:"accountName"`   // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	OrderFilter   *string `json:"orderFilter"`   // false 订单品种 order：普通订单，默认值；oco：止盈止损单
	OrderId       *string `json:"orderId"`       // false 订单ID，orderId和clientOrderId必须填一个，若填两个，以orderId为主
	ClientOrderId *string `json:"clientOrderId"` // false 客户自定义订单ID 如果clientOrderId关联了多个订单，只会返回最近的那笔订单
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeOrderInfoAPI) AccountName(accountName string) *PrivateRestTradeOrderInfoAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 订单品种 order：普通订单，默认值；oco：止盈止损单
func (api *PrivateRestTradeOrderInfoAPI) OrderFilter(orderFilter string) *PrivateRestTradeOrderInfoAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// string false 订单ID，orderId和clientOrderId必须填一个，若填两个，以orderId为主
func (api *PrivateRestTradeOrderInfoAPI) OrderId(orderId string) *PrivateRestTradeOrderInfoAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// string false 客户自定义订单ID 如果clientOrderId关联了多个订单，只会返回最近的那笔订单
func (api *PrivateRestTradeOrderInfoAPI) ClientOrderId(clientOrderId string) *PrivateRestTradeOrderInfoAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

type PrivateRestTradeHistoryOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeHistoryOrdersReq
}

type PrivateRestTradeHistoryOrdersReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	OrderType    *string `json:"orderType"`    // false 订单类型 market：市价单，limit：限价单，post_only：只挂单
	OrderFilter  *string `json:"orderFilter"`  // false 订单品种 order：普通订单，默认值；oco：止盈止损单
	BeginTime    *string `json:"beginTime"`    // false 筛选订单创建的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
	EndTime      *string `json:"endTime"`      // false 筛选订单创建的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
	BeginId      *string `json:"beginId"`      // false 请求此ID之后的分页内容，传的值对应接口的id字段
	EndId        *string `json:"endId"`        // false 请求此ID之前的分页内容，传的值对应接口的id字段
	Limit        *string `json:"limit"`        // false 分页返回的结果集数量，最大为100，不填默认返回100条
	OrderId      *string `json:"orderId"`      // false 订单ID
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeHistoryOrdersAPI) AccountName(accountName string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeHistoryOrdersAPI) BusinessType(businessType string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeHistoryOrdersAPI) Symbol(symbol string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 订单类型 market：市价单，limit：限价单，post_only：只挂单
func (api *PrivateRestTradeHistoryOrdersAPI) OrderType(orderType string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

// string false 订单品种 order：普通订单，默认值；oco：止盈止损单
func (api *PrivateRestTradeHistoryOrdersAPI) OrderFilter(orderFilter string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.OrderFilter = GetPointer(orderFilter)
	return api
}

// string false 筛选订单创建的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
func (api *PrivateRestTradeHistoryOrdersAPI) BeginTime(beginTime string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false 筛选订单创建的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
func (api *PrivateRestTradeHistoryOrdersAPI) EndTime(endTime string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false 请求此ID之后的分页内容，传的值对应接口的id字段
func (api *PrivateRestTradeHistoryOrdersAPI) BeginId(beginId string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.BeginId = GetPointer(beginId)
	return api
}

// string false 请求此ID之前的分页内容，传的值对应接口的id字段
func (api *PrivateRestTradeHistoryOrdersAPI) EndId(endId string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.EndId = GetPointer(endId)
	return api
}

// string false 分页返回的结果集数量，最大为100，不填默认返回100条
func (api *PrivateRestTradeHistoryOrdersAPI) Limit(limit string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// string false 订单ID
func (api *PrivateRestTradeHistoryOrdersAPI) OrderId(orderId string) *PrivateRestTradeHistoryOrdersAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestTradeOrderOperationsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderOperationsReq
}

type PrivateRestTradeOrderOperationsReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	OrderId     *string `json:"orderId"`     // true 订单编号，用于唯一标识订单
	BeginTime   *string `json:"beginTime"`   // false 筛选操作时间的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
	EndTime     *string `json:"endTime"`     // false 筛选操作时间的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeOrderOperationsAPI) AccountName(accountName string) *PrivateRestTradeOrderOperationsAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string true 订单编号，用于唯一标识订单
func (api *PrivateRestTradeOrderOperationsAPI) OrderId(orderId string) *PrivateRestTradeOrderOperationsAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// string false 筛选操作时间的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
func (api *PrivateRestTradeOrderOperationsAPI) BeginTime(beginTime string) *PrivateRestTradeOrderOperationsAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false 筛选操作时间的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
func (api *PrivateRestTradeOrderOperationsAPI) EndTime(endTime string) *PrivateRestTradeOrderOperationsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

type PrivateRestTradeHistoryTradesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeHistoryTradesReq
}

type PrivateRestTradeHistoryTradesReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	OrderType    *string `json:"orderType"`    // false 订单类型 market：市价单，limit：限价单，post_only：只挂单
	BeginTime    *string `json:"beginTime"`    // false 筛选成交时间的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
	EndTime      *string `json:"endTime"`      // false 筛选成交时间的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
	BeginId      *string `json:"beginId"`      // false 请求此ID之后的分页内容，传的值对应接口的ID字段
	EndId        *string `json:"endId"`        // false 请求此ID之前的分页内容，传的值对应接口的ID字段
	Limit        *string `json:"limit"`        // false 分页返回的结果集数量，最大为100，不填默认返回100条
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeHistoryTradesAPI) AccountName(accountName string) *PrivateRestTradeHistoryTradesAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeHistoryTradesAPI) BusinessType(businessType string) *PrivateRestTradeHistoryTradesAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeHistoryTradesAPI) Symbol(symbol string) *PrivateRestTradeHistoryTradesAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 订单类型 market：市价单，limit：限价单，post_only：只挂单
func (api *PrivateRestTradeHistoryTradesAPI) OrderType(orderType string) *PrivateRestTradeHistoryTradesAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

// string false 筛选成交时间的开始时间戳，Unix时间戳的毫秒数格式，如 1732158178000
func (api *PrivateRestTradeHistoryTradesAPI) BeginTime(beginTime string) *PrivateRestTradeHistoryTradesAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false 筛选成交时间的结束时间戳，Unix时间戳的毫秒数格式，如 1732182494000
func (api *PrivateRestTradeHistoryTradesAPI) EndTime(endTime string) *PrivateRestTradeHistoryTradesAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false 请求此ID之后的分页内容，传的值对应接口的ID字段
func (api *PrivateRestTradeHistoryTradesAPI) BeginId(beginId string) *PrivateRestTradeHistoryTradesAPI {
	api.req.BeginId = GetPointer(beginId)
	return api
}

// string false 请求此ID之前的分页内容，传的值对应接口的ID字段
func (api *PrivateRestTradeHistoryTradesAPI) EndId(endId string) *PrivateRestTradeHistoryTradesAPI {
	api.req.EndId = GetPointer(endId)
	return api
}

// string false 分页返回的结果集数量，最大为100，不填默认返回100条
func (api *PrivateRestTradeHistoryTradesAPI) Limit(limit string) *PrivateRestTradeHistoryTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestTradePositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionReq
}

type PrivateRestTradePositionReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradePositionAPI) AccountName(accountName string) *PrivateRestTradePositionAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradePositionAPI) Symbol(symbol string) *PrivateRestTradePositionAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradePositionAPI) BusinessType(businessType string) *PrivateRestTradePositionAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PrivateRestTradeLeverPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeLeverPostReq
}

type PrivateRestTradeLeverPostReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	Lever        *string `json:"lever"`        // true 杠杆倍数
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeLeverPostAPI) AccountName(accountName string) *PrivateRestTradeLeverPostAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeLeverPostAPI) Symbol(symbol string) *PrivateRestTradeLeverPostAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeLeverPostAPI) BusinessType(businessType string) *PrivateRestTradeLeverPostAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string true 杠杆倍数
func (api *PrivateRestTradeLeverPostAPI) Lever(lever string) *PrivateRestTradeLeverPostAPI {
	api.req.Lever = GetPointer(lever)
	return api
}

type PrivateRestTradeLeverGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeLeverGetReq
}

type PrivateRestTradeLeverGetReq struct {
	AccountName  *string `json:"accountName"`  // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol       *string `json:"symbol"`       // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	BusinessType *string `json:"businessType"` // false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	Currency     *string `json:"currency"`     // false 币种名称，如BTC
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeLeverGetAPI) AccountName(accountName string) *PrivateRestTradeLeverGetAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeLeverGetAPI) Symbol(symbol string) *PrivateRestTradeLeverGetAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// string false 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
func (api *PrivateRestTradeLeverGetAPI) BusinessType(businessType string) *PrivateRestTradeLeverGetAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

// string false 币种名称，如BTC
func (api *PrivateRestTradeLeverGetAPI) Currency(currency string) *PrivateRestTradeLeverGetAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestTradeStopPositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeStopPositionReq
}

type PrivateRestTradeStopPositionReq struct {
	AccountName    *string `json:"accountName"`    // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Symbol         *string `json:"symbol"`         // false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	PositionIdx    *string `json:"positionIdx"`    // true 仓位标识，字典项：net: 单持仓模式，long：双向持仓买持仓，short：双向持仓卖持仓
	TpslClOrdId    *string `json:"tpslClOrdId"`    // false 下单附带止盈止损订单时，客户定义止盈止损单触发执行订单的clientOrderId
	TpslMode       *string `json:"tpslMode"`       // false 止盈止损模式：all_position：全部仓位模式，partially_position ：部分仓位模式
	TpslQty        *string `json:"tpslQty"`        // false 止盈止损数量，仅部分止盈止损时有效
	TakeProfitType *string `json:"takeProfitType"` // false 止盈触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价，index_price：指数价
	StopLossType   *string `json:"stopLossType"`   // false 止损触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价，index_price：指数价
	TakeProfit     *string `json:"takeProfit"`     // false 止盈触发价格
	StopLoss       *string `json:"stopLoss"`       // false 止损触发价格
	TpOrderType    *string `json:"tpOrderType"`    // false 止盈执行价格类型，字典：market：市价单，默认值；limit：限价单
	SlOrderType    *string `json:"slOrderType"`    // false 止损执行价格类型，字典：market：市价单，默认值；limit：限价单
	TpLimitPrice   *string `json:"tpLimitPrice"`   // false 止盈执行价格类型为limit时，止盈执行价; 止盈执行价格类型为market时, 必须输入null值
	SlLimitPrice   *string `json:"slLimitPrice"`   // false 止损执行价格类型为limit时，止损执行价; 止损执行价格类型为market时, 必须输入null值
}

// false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestTradeStopPositionAPI) AccountName(accountName string) *PrivateRestTradeStopPositionAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// false 交易对名称 如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
func (api *PrivateRestTradeStopPositionAPI) Symbol(symbol string) *PrivateRestTradeStopPositionAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// true 仓位标识，字典项：net: 单持仓模式，long：双向持仓买持仓，short：双向持仓卖持仓
func (api *PrivateRestTradeStopPositionAPI) PositionIdx(positionIdx string) *PrivateRestTradeStopPositionAPI {
	api.req.PositionIdx = GetPointer(positionIdx)
	return api
}

// false 下单附带止盈止损订单时，客户定义止盈止损单触发执行订单的clientOrderId
func (api *PrivateRestTradeStopPositionAPI) TpslClOrdId(tpslClOrdId string) *PrivateRestTradeStopPositionAPI {
	api.req.TpslClOrdId = GetPointer(tpslClOrdId)
	return api
}

// false 止盈止损模式：all_position：全部仓位模式，partially_position ：部分仓位模式
func (api *PrivateRestTradeStopPositionAPI) TpslMode(tpslMode string) *PrivateRestTradeStopPositionAPI {
	api.req.TpslMode = GetPointer(tpslMode)
	return api
}

// false 止盈止损数量，仅部分止盈止损时有效
func (api *PrivateRestTradeStopPositionAPI) TpslQty(tpslQty string) *PrivateRestTradeStopPositionAPI {
	api.req.TpslQty = GetPointer(tpslQty)
	return api
}

// false 止盈触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价，index_price：指数价
func (api *PrivateRestTradeStopPositionAPI) TakeProfitType(takeProfitType string) *PrivateRestTradeStopPositionAPI {
	api.req.TakeProfitType = GetPointer(takeProfitType)
	return api
}

// false 止损触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价，index_price：指数价
func (api *PrivateRestTradeStopPositionAPI) StopLossType(stopLossType string) *PrivateRestTradeStopPositionAPI {
	api.req.StopLossType = GetPointer(stopLossType)
	return api
}

// false 止盈触发价格
func (api *PrivateRestTradeStopPositionAPI) TakeProfit(takeProfit string) *PrivateRestTradeStopPositionAPI {
	api.req.TakeProfit = GetPointer(takeProfit)
	return api
}

// false 止损触发价格
func (api *PrivateRestTradeStopPositionAPI) StopLoss(stopLoss string) *PrivateRestTradeStopPositionAPI {
	api.req.StopLoss = GetPointer(stopLoss)
	return api
}

// false 止盈执行价格类型，字典：market：市价单，默认值；limit：限价单
func (api *PrivateRestTradeStopPositionAPI) TpOrderType(tpOrderType string) *PrivateRestTradeStopPositionAPI {
	api.req.TpOrderType = GetPointer(tpOrderType)
	return api
}

// false 止损执行价格类型，字典：market：市价单，默认值；limit：限价单
func (api *PrivateRestTradeStopPositionAPI) SlOrderType(slOrderType string) *PrivateRestTradeStopPositionAPI {
	api.req.SlOrderType = GetPointer(slOrderType)
	return api
}

// false 止盈执行价格类型为limit时，止盈执行价; 止盈执行价格类型为market时, 必须输入null值
func (api *PrivateRestTradeStopPositionAPI) TpLimitPrice(tpLimitPrice string) *PrivateRestTradeStopPositionAPI {
	api.req.TpLimitPrice = GetPointer(tpLimitPrice)
	return api
}

// false 止损执行价格类型为limit时，止损执行价; 止损执行价格类型为market时, 必须输入null值
func (api *PrivateRestTradeStopPositionAPI) SlLimitPrice(slLimitPrice string) *PrivateRestTradeStopPositionAPI {
	api.req.SlLimitPrice = GetPointer(slLimitPrice)
	return api
}
