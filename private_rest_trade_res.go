package myxcoinapi

type PrivateRestTradeOrderRes struct {
	OrderId       string `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type PrivateRestTradeBatchOrderResRow struct {
	Code          string `json:"code"`
	Msg           string `json:"msg"`
	OrderId       string `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	Ts            string `json:"ts"`
}

type PrivateRestTradeBatchOrderRes []PrivateRestTradeBatchOrderResRow

type PrivateRestTradeCancelOrderRes struct {
	OrderId string `json:"orderId"`
}

type PrivateRestTradeBatchCancelOrderResRow struct {
	Code    int64  `json:"code"`
	Msg     string `json:"msg"`
	OrderId string `json:"orderId"`
	Ts      string `json:"ts"`
}

type PrivateRestTradeBatchCancelOrderRes []PrivateRestTradeBatchCancelOrderResRow

type PrivateRestTradeCancelAllOrderRes string

type Order struct {
	Id            string `json:"id"`            // 数据唯一ID
	BusinessType  string `json:"businessType"`  // 产品业务线，如spot
	Symbol        string `json:"symbol"`        // 交易对名称，如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	OrderId       string `json:"orderId"`       // 订单ID
	ClientOrderId string `json:"clientOrderId"` // 客户自定义订单ID
	EventId       string `json:"eventId"`       // 事件编号，保障订单生命周期发生事件完整性
	Price         string `json:"price"`         // 委托价格
	Qty           string `json:"qty"`           // 委托数量  对于现货杠杆，单位默认为基础货币，若下单请求参数marketUnit为quoteCoin, 则单位为USDT；对于合约，单位为币
	QuoteQty      string `json:"quoteQty"`      // 计价货币数量，单位为USDT  对于现货杠杆，单位默认为基础货币，若下单请求参数marketUnit为quoteCoin, 则单位为USDT；对于合约，单位为币
	Pnl           string `json:"pnl"`           // 收益，仅适用于合约交易有成交的平仓订单，其他情况均为0
	OrderType     string `json:"orderType"`     // 订单类型 market：市价单，limit：限价单，post_only：只挂单
	OrderFilter   string `json:"orderFilter"`   // 订单品种 order：普通订单，默认值；oco：止盈止损单
	Side          string `json:"side"`          // 订单方向，buy：买，sell：卖
	TotalFillQty  string `json:"totalFillQty"`  // 累计成交数量 对于现货杠杆，单位为基础币种 对于合约，单位为币
	AvgPrice      string `json:"avgPrice"`      // 成交均价
	Status        string `json:"status"`        // 订单状态 untriggered: 条件未触发 new：挂单成功 partially_filled：部分成交 partially_canceled：部分成交已撤单 canceled：已撤单 filled：完全成交
	Lever         string `json:"lever"`         // 杠杆倍数 仅适用于合约交易，其他场景返回空字符串
	BaseFee       string `json:"baseFee"`       // 基础币种手续费
	QuoteFee      string `json:"quoteFee"`      // 计价币种手续费
	Source        string `json:"source"`        // 订单来源，详细字段请见公共参数
	CreateType    string `json:"createType"`    // 订单类型  normal：普通订单 ADL：自动减仓 liuquidation：强制平仓 system_cancel：系统撤单 delivery：交割订单
	CancelSource  string `json:"cancelSource"`  // 撤单来源 api：来源于api web：来源于web端 ios：来源于ios设备 android：来源于android设备 systemClosing：非强平、强减、交割的系统订单 riskEngine：强平 adl：ADL订单 settlement：交割 mmpByUser：用户主动通过mmp接口撤单 mmpBySystem：因mmp冻结窗口触发撤单
	CancelUid     string `json:"cancelUid"`     // 撤单操作人，仅适用于已撤单场景
	ReduceOnly    bool   `json:"reduceOnly"`    // 是否只减仓，true或false，默认false，仅适用于合约交易
	TimeInForce   string `json:"timeInForce"`   // 订单生效类型，默认gtc gtc：长期有效直至取消 ioc：立即成交并取消剩余 fok：全部成交或立即取消
	CreateTime    string `json:"createTime"`    // 创建时间，Unix时间戳的毫秒数格式，如 1732158178000
	UpdateTime    string `json:"updateTime"`    // 更新时间，Unix时间戳的毫秒数格式，如 1732158178000
	PosSide       string `json:"posSide"`       // 持仓方向 net：单向持仓，long：正持仓，short：卖持仓
	RiskReducing  bool   `json:"riskReducing"`  // 风险降低订单标识，true代表是
	ParentOrderId string `json:"parentOrderId"` // 父订单ID，止盈止损单、计划委托生效子订单，此字段保持父订单一致
	TpslOrder     struct {
		TpslClOrdId    string `json:"tpslClOrdId"`    // 下单附带止盈止损订单时，客户定义止盈止损单触发执行订单的clientOrderId
		TpslMode       string `json:"tpslMode"`       // 止盈止损模式 all_position：全部仓位模式 partially_position：部分仓位模式
		TakeProfitType string `json:"takeProfitType"` // 止盈触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
		StopLossType   string `json:"stopLossType"`   // 止损触发价格类型，字典值：last_price：最新价，默认值；mark_price：标记价；index_price：指数价
		TakeProfit     string `json:"takeProfit"`     // 止盈触发价格
		StopLoss       string `json:"stopLoss"`       // 止损触发价格
		TpOrderType    string `json:"tpOrderType"`    // 止盈执行价格类型，字典：market：市价单，limit：限价单
		SlOrderType    string `json:"slOrderType"`    // 止损执行价格类型，字典：market：市价单，limit：限价单
		TpLimitPrice   string `json:"tpLimitPrice"`   // 止盈执行价格类型为限价时，止盈执行价
		SlLimitPrice   string `json:"slLimitPrice"`   // 止损执行价格类型为限价时，止损执行价
	} `json:"tpslOrder"` // 下单附带止盈止损信息
	MassQuoteOrder struct {
		Quote           bool   `json:"quote"`           // 若为MassQuote单，则为true；普通订单为false
		QuoteId         string `json:"quoteId"`         // 申报ID，用户自定义编号值，当前挂单中唯一
		QuoteSetId      string `json:"quoteSetId"`      // 用户自定义的ID，可基于quoteSetId进行批量撤单
		MmpGroup        string `json:"mmpGroup"`        // MMP组信息 当massQuoteOrder为true时，该字段返回
		PriceAdjustment bool   `json:"priceAdjustment"` // 因MassQuote下单“postOnly”参数为true导致价格调整，则为true，否则则为false
	} `json:"massQuoteOrder"` // 下面附带massQuote订单参数
	AccountName string `json:"accountName"` // 账户名称
	Pid         string `json:"pid"`         // 账户唯一ID
	Uid         string `json:"uid"`         // 操作员唯一ID
	Cid         string `json:"cid"`         // 操作主体唯一ID
}

type PrivateRestTradeOpenOrdersRes []Order

type PrivateRestTradeOrderInfoRes Order

type PrivateRestTradeHistoryOrdersRes []Order

type PrivateRestTradeOrderOperationsResRow struct {
	OrderId       string `json:"orderId"`       // 订单编号，用于唯一标识订单
	OperationType string `json:"operationType"` // 操作类型
	Price         string `json:"price"`         // 订单价格
	Qty           string `json:"qty"`           // 订单数量
	AccountName   string `json:"accountName"`   // 账户名称
	Pid           string `json:"pid"`           // 账户唯一ID
	Uid           string `json:"uid"`           // 操作员唯一ID
	Cid           string `json:"cid"`           // 操作主体唯一ID
	CreateTime    string `json:"createTime"`    // 操作时间，Unix时间戳的毫秒数格式，如 1732158178000
}
type PrivateRestTradeOrderOperationsRes []PrivateRestTradeOrderOperationsResRow

type PrivateRestTradeHistoryTradesResRow struct {
	Id            string `json:"id"`            // 数据唯一ID
	OrderId       string `json:"orderId"`       // 订单编号，用于唯一标识订单
	ClientOrderId string `json:"clientOrderId"` // 客户自定义订单编号，用于用户自定义标识该笔订单
	BusinessType  string `json:"businessType"`  // 业务线 spot表示现货/现货杠杆，linear_perpetual表示U本位永续，linear_futures表示交割合约
	Symbol        string `json:"symbol"`        // 交易对名称，如 BTC-USDT（现货）、BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	Pnl           string `json:"pnl"`           // 收益，仅适用于合约交易有成交的平仓订单，其他情况均为0
	OrderType     string `json:"orderType"`     // 订单类型 market：市价单，limit：限价单，post_only：只挂单
	Side          string `json:"side"`          // 订单方向
	FillPrice     string `json:"fillPrice"`     // 成交价格
	IndexPrice    string `json:"indexPrice"`    // 成交时的指数价格
	MarkPrice     string `json:"markPrice"`     // 成交时的标记价格
	TradeId       string `json:"tradeId"`       // 成交ID
	EventId       string `json:"eventId"`       // 事件编号，保障订单生命周期发生事件完整性
	QuoteId       string `json:"quoteId"`       // 申报ID，用户自定义编号值，当前挂单中唯一；此ID会在订单信息和成交信息中都体现，建议使用递增序列；
	QuoteSetId    string `json:"quoteSetId"`    // 用户自定义的ID，可基于quoteSetId进行批量撤单
	RiskReducing  bool   `json:"riskReducing"`  // 风险降低订单标识，true代表是
	Role          string `json:"role"`          // 成交角色
	FillQty       string `json:"fillQty"`       // 成交数量，单位为币
	FillTime      string `json:"fillTime"`      // 成交时间，Unix时间戳的毫秒数格式，如 1732158178000
	ExecType      string `json:"execType"`      // 成交类型 trade: 普通成交 riskEngine：强平 adl：ADL订单 settlement：交割单 systemClosing：非强平、强减、交割的系统订单；
	Lever         string `json:"lever"`         // 杠杆倍数 仅适用于合约交易，其他场景返回空字符串
	FeeCurrency   string `json:"feeCurrency"`   // 手续费币种，手续费的单位
	Fee           string `json:"fee"`           // 手续费
	AccountName   string `json:"accountName"`   // 账户名称
	Pid           string `json:"pid"`           // 账户唯一ID
	Uid           string `json:"uid"`           // 操作员唯一ID
	Cid           string `json:"cid"`           // 操作主体唯一ID
}
type PrivateRestTradeHistoryTradesRes []PrivateRestTradeHistoryTradesResRow

type PrivateRestTradePositionResRow struct {
	PositionId       string `json:"positionId"`       // 持仓ID
	BusinessType     string `json:"businessType"`     // 业务线
	Symbol           string `json:"symbol"`           // 交易对名称
	PositionQty      string `json:"positionQty"`      // 持仓数量
	AvgPrice         string `json:"avgPrice"`         // 开仓平均价
	Upl              string `json:"upl"`              // 未实现盈亏
	Lever            string `json:"lever"`            // 杠杆
	LiquidationPrice string `json:"liquidationPrice"` // 预估强平价
	MarkPrice        string `json:"markPrice"`        // 最新标记价格
	Im               string `json:"im"`               // 初始保证金
	IndexPrice       string `json:"indexPrice"`       // 最新指数价格
	TakeProfit       string `json:"takeProfit"`       // 止盈触发价格
	StopLoss         string `json:"stopLoss"`         // 止损触发价格
	AccountName      string `json:"accountName"`      // 账户名称
	Pid              string `json:"pid"`              // 账户唯一 ID
	Uid              string `json:"uid"`              // 操作员唯一 ID
	Cid              string `json:"cid"`              // 操作主体唯一 ID
	CreateTime       string `json:"createTime"`       // 创建时间
	UpdateTime       string `json:"updateTime"`       // 更新时间
	Delta            string `json:"delta"`            // Delta
	Gamma            string `json:"gamma"`            // Gamma
	Vega             string `json:"vega"`             // Vega
	Theta            string `json:"theta"`            // Theta
}
type PrivateRestTradePositionRes []PrivateRestTradePositionResRow

type PrivateRestTradeLeverCommonRes struct {
	AccountName string `json:"accountName"` // 账户名称
	Symbol      string `json:"symbol"`      // 交易对名称，如BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	Currency    string `json:"currency"`    // 币种名称，如BTC
	Lever       string `json:"lever"`       // 杠杆倍数
	Pid         string `json:"pid"`         // 账户唯一id
	Cid         string `json:"cid"`         // 注册主体唯一id
}

type PrivateRestTradeStopPositionRes struct {
	OrderId       string `json:"orderId"`       // 订单ID
	ClientOrderId string `json:"clientOrderId"` // 客户自定义订单ID
}
