package myxcoinapi

import "strings"

type Ticker24hr struct {
	Symbol             string `json:"symbol"`             // Trading pair name, e.g., BTC-USDT(spot), BTC-USDT-PERP(perpetual), BTC-USDT-26DEC25(futures)
	LastPrice          string `json:"lastPrice"`          // Last fill price
	PriceChange        string `json:"priceChange"`        // Price change amount in USDT
	PriceChangePercent string `json:"priceChangePercent"` // Price change percentage
	HighPrice          string `json:"highPrice"`          // Highest price
	LowPrice           string `json:"lowPrice"`           // Lowest price
	FillQty            string `json:"fillQty"`            // Filled quantity
	FillAmount         string `json:"fillAmount"`         // Filled amount in USDT
	Count              string `json:"count"`              // Number of trades
}

type WsTicker24hr struct {
	WsSubscribeArg
	Ticker24hr
	Ts int64 `json:"ts"`
}

type WsTicker24hrMiddle struct {
	WsSubscribeArg
	Data []Ticker24hr `json:"data"`
	Ts   int64        `json:"ts"`
}

func handleWsTicker24hr(data []byte) (*WsTicker24hr, error) {
	wsTicker24hrMiddle := &WsTicker24hrMiddle{}
	err := json.Unmarshal(data, wsTicker24hrMiddle)
	if err != nil {
		return nil, err
	}

	ticker := wsTicker24hrMiddle.Data[0]
	wsTicker24hr := &WsTicker24hr{
		WsSubscribeArg: wsTicker24hrMiddle.WsSubscribeArg,
		Ticker24hr:     ticker,
		Ts:             wsTicker24hrMiddle.Ts,
	}

	return wsTicker24hr, nil
}

type Klines struct {
	Symbol             string `json:"symbol"`             // 交易对
	Period             string `json:"period"`             // 周期
	OpenTime           string `json:"openTime"`           // 开盘时间
	CloseTime          string `json:"closeTime"`          // 收盘时间
	OpenPrice          string `json:"openPrice"`          // 开盘价
	ClosePrice         string `json:"closePrice"`         // 收盘价
	HighPrice          string `json:"highPrice"`          // 最高价
	LowPrice           string `json:"lowPrice"`           // 最低价
	Volume             string `json:"volume"`             // 成交量
	QuoteVolume        string `json:"quoteVolume"`        // 成交额
	Count              string `json:"count"`              // 成交笔数
	PriceChange        string `json:"priceChange"`        // 涨跌额
	PriceChangePercent string `json:"priceChangePercent"` // 涨跌幅
}

type WsKline struct {
	WsSubscribeArg
	Klines
	Ts int64 `json:"ts"`
}

type WsKlinesMiddle struct {
	WsSubscribeArg
	Data []Klines `json:"data"`
	Ts   int64    `json:"ts"`
}

func handleWsKlines(data []byte) (*[]WsKline, error) {
	wsKlinesMiddle := &WsKlinesMiddle{}
	err := json.Unmarshal(data, wsKlinesMiddle)
	if err != nil {
		return nil, err
	}

	wsKlines := []WsKline{}
	for _, kline := range wsKlinesMiddle.Data {
		wsKlines = append(wsKlines, WsKline{
			WsSubscribeArg: wsKlinesMiddle.WsSubscribeArg,
			Klines:         kline,
			Ts:             wsKlinesMiddle.Ts,
		})
	}
	return &wsKlines, nil
}

type Depth struct {
	Symbol       string     `json:"symbol"`
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	PreUpdateId  string     `json:"preUpdateId"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type WsDepth struct {
	WsSubscribeArg
	Symbol       string  `json:"symbol"`
	Asks         []Books `json:"asks"`
	Bids         []Books `json:"bids"`
	PreUpdateId  string  `json:"preUpdateId"`
	LastUpdateId string  `json:"lastUpdateId"`
	Ts           int64   `json:"ts"`
}

type WsDepthMiddle struct {
	WsSubscribeArg
	Data []Depth `json:"data"`
	Ts   int64   `json:"ts"`
}

func handleWsDepth(data []byte) (*[]WsDepth, error) {
	wsDepthMiddle := &WsDepthMiddle{}
	err := json.Unmarshal(data, wsDepthMiddle)
	if err != nil {
		return nil, err
	}

	wsDepths := []WsDepth{}
	for _, depth := range wsDepthMiddle.Data {
		asks := make([]Books, len(depth.Asks))
		for i, ask := range depth.Asks {
			asks[i] = Books{
				Price:    ask[0],
				Quantity: ask[1],
			}
		}
		bids := make([]Books, len(depth.Bids))
		for i, bid := range depth.Bids {
			bids[i] = Books{
				Price:    bid[0],
				Quantity: bid[1],
			}
		}
		wsDepths = append(wsDepths, WsDepth{
			WsSubscribeArg: wsDepthMiddle.WsSubscribeArg,
			Symbol:         depth.Symbol,
			Asks:           asks,
			Bids:           bids,
			PreUpdateId:    depth.PreUpdateId,
			LastUpdateId:   depth.LastUpdateId,
			Ts:             wsDepthMiddle.Ts,
		})
	}

	return &wsDepths, nil
}

type DepthLevels struct {
	Symbol       string     `json:"symbol"`
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateId string     `json:"lastUpdateId"`
	Group        string     `json:"group"`
}

type WsDepthLevels struct {
	WsSubscribeArg
	Symbol       string  `json:"symbol"`
	Asks         []Books `json:"asks"`
	Bids         []Books `json:"bids"`
	LastUpdateId string  `json:"lastUpdateId"`
	Group        string  `json:"group"`
	Ts           int64   `json:"ts"`
}

type WsDepthLevelsMiddle struct {
	WsSubscribeArg
	Data []DepthLevels `json:"data"`
	Ts   int64         `json:"ts"`
}

func handleWsDepthLevels(data []byte) (*[]WsDepthLevels, error) {
	wsDepthLevelsMiddle := &WsDepthLevelsMiddle{}
	err := json.Unmarshal(data, wsDepthLevelsMiddle)
	if err != nil {
		return nil, err
	}

	// 将 stream 截取到第二个 '#' 之前，例如
	// depthlevels#100ms#5#none -> depthlevels#100ms
	stream := wsDepthLevelsMiddle.WsSubscribeArg.Stream
	if stream != "" {
		if first := strings.Index(stream, "#"); first != -1 {
			if second := strings.Index(stream[first+1:], "#"); second != -1 {
				wsDepthLevelsMiddle.WsSubscribeArg.Stream = stream[:first+1+second]
			}
		}
	}

	wsDepthLevels := []WsDepthLevels{}
	for _, depthLevels := range wsDepthLevelsMiddle.Data {
		asks := make([]Books, len(depthLevels.Asks))
		for i, ask := range depthLevels.Asks {
			asks[i] = Books{
				Price:    ask[0],
				Quantity: ask[1],
			}
		}
		bids := make([]Books, len(depthLevels.Bids))
		for i, bid := range depthLevels.Bids {
			bids[i] = Books{
				Price:    bid[0],
				Quantity: bid[1],
			}
		}
		wsDepthLevels = append(wsDepthLevels, WsDepthLevels{
			WsSubscribeArg: wsDepthLevelsMiddle.WsSubscribeArg,
			Symbol:         depthLevels.Symbol,
			Asks:           asks,
			Bids:           bids,
			LastUpdateId:   depthLevels.LastUpdateId,
			Group:          depthLevels.Group,
			Ts:             wsDepthLevelsMiddle.Ts,
		})
	}
	return &wsDepthLevels, nil
}

// 当前最优挂单频道
type OrderBook struct {
	Symbol       string     `json:"symbol"`
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	PreUpdateId  string     `json:"preUpdateId"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type WsOrderbook struct {
	WsSubscribeArg
	Symbol       string  `json:"symbol"`
	Asks         []Books `json:"asks"`
	Bids         []Books `json:"bids"`
	PreUpdateId  string  `json:"preUpdateId"`
	LastUpdateId string  `json:"lastUpdateId"`
	Ts           int64   `json:"ts"`
}

type WsOrderbookMiddle struct {
	WsSubscribeArg
	Data []OrderBook `json:"data"`
	Ts   int64       `json:"ts"`
}

func handleWsOrderbook(data []byte) (*[]WsOrderbook, error) {
	wsOrderbookMiddle := &WsOrderbookMiddle{}
	err := json.Unmarshal(data, wsOrderbookMiddle)
	if err != nil {
		return nil, err
	}

	wsOrderbooks := []WsOrderbook{}
	for _, orderbook := range wsOrderbookMiddle.Data {
		asks := make([]Books, len(orderbook.Asks))
		for i, ask := range orderbook.Asks {
			asks[i] = Books{
				Price:    ask[0],
				Quantity: ask[1],
			}
		}
		bids := make([]Books, len(orderbook.Bids))
		for i, bid := range orderbook.Bids {
			bids[i] = Books{
				Price:    bid[0],
				Quantity: bid[1],
			}
		}
		wsOrderbooks = append(wsOrderbooks, WsOrderbook{
			WsSubscribeArg: wsOrderbookMiddle.WsSubscribeArg,
			Symbol:         orderbook.Symbol,
			Asks:           asks,
			Bids:           bids,
			PreUpdateId:    orderbook.PreUpdateId,
			LastUpdateId:   orderbook.LastUpdateId,
			Ts:             wsOrderbookMiddle.Ts,
		})
	}
	return &wsOrderbooks, nil
}

type Trade struct {
	Symbol string `json:"symbol"`
	Id     string `json:"id"`
	Side   string `json:"side"`
	Price  string `json:"price"`
	Qty    string `json:"qty"`
	Time   string `json:"time"`
}

type WsTrade struct {
	WsSubscribeArg
	Trade
	Ts int64 `json:"ts"`
}

type WsTradeMiddle struct {
	WsSubscribeArg
	Data []Trade `json:"data"`
	Ts   int64   `json:"ts"`
}

func handleWsTrade(data []byte) (*[]WsTrade, error) {
	wsTradeMiddle := &WsTradeMiddle{}
	err := json.Unmarshal(data, wsTradeMiddle)
	if err != nil {
		return nil, err
	}

	wsTrades := []WsTrade{}
	for _, trade := range wsTradeMiddle.Data {
		wsTrades = append(wsTrades, WsTrade{
			WsSubscribeArg: wsTradeMiddle.WsSubscribeArg,
			Trade:          trade,
			Ts:             wsTradeMiddle.Ts,
		})
	}
	return &wsTrades, nil
}

type Position struct {
	BusinessType     string `json:"businessType"`     // 业务线，linear_perpetual 表示 U 本位永续，linear_futures 表示 U 本位合约
	Symbol           string `json:"symbol"`           // 交易对名称，如 BTC-USDT-PERP（永续）、BTC-USDT-26DEC25（交割）等
	PositionQty      string `json:"positionQty"`      // 持仓数量，正值为多，负值为空；合约单位为币
	AvgPrice         string `json:"avgPrice"`         // 开仓平均价
	Upl              string `json:"upl"`              // 未实现盈亏，以标记价格计算
	Lever            string `json:"lever"`            // 杠杆倍数
	LiquidationPrice string `json:"liquidationPrice"` // 预估强平价
	MarkPrice        string `json:"markPrice"`        // 最新标记价格
	Im               string `json:"im"`               // 初始保证金
	IndexPrice       string `json:"indexPrice"`       // 最新指数价格
	Pnl              string `json:"pnl"`              // 累计平仓盈亏
	Fee              string `json:"fee"`              // 累计手续费支出
	FundingFee       string `json:"fundingFee"`       // 累计资金费用
	Pid              string `json:"pid"`              // 账户唯一 ID
	TradedType       string `json:"tradedType"`       // 交易类型：CLOSE 表示平仓，OPEN 表示开仓
	CreateTime       string `json:"createTime"`       // 创建时间，Unix 时间戳毫秒格式，如 1732158178000
	UpdateTime       string `json:"updateTime"`       // 更新时间，Unix 时间戳毫秒格式，如 1732158178000
	Delta            string `json:"delta"`            // Delta
	Gamma            string `json:"gamma"`            // Gamma
	Vega             string `json:"vega"`             // Vega
	Theta            string `json:"theta"`            // Theta
}

type WsPosition struct {
	WsSubscribeArg
	Position
}

type WsPositionMiddle struct {
	WsSubscribeArg
	Data []Position `json:"data"`
}

func handleWsPosition(data []byte) (*[]WsPosition, error) {
	wsPositionMiddle := &WsPositionMiddle{}
	err := json.Unmarshal(data, wsPositionMiddle)
	if err != nil {
		return nil, err
	}

	wsPositions := []WsPosition{}
	for _, position := range wsPositionMiddle.Data {
		arg := wsPositionMiddle.WsSubscribeArg
		arg.BusinessType = position.BusinessType
		arg.Symbol = position.Symbol
		wsPositions = append(wsPositions, WsPosition{
			WsSubscribeArg: arg,
			Position:       position,
		})
	}
	return &wsPositions, nil
}

type WsOrder struct {
	BusinessType  string `json:"businessType"`
	Symbol        string `json:"symbol"`
	OrderId       string `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	EventId       string `json:"eventId"`
	Price         string `json:"price"`
	Qty           string `json:"qty"`
	QuoteQty      string `json:"quoteQty"`
	OrderType     string `json:"orderType"`
	Side          string `json:"side"`
	TotalFillQty  string `json:"totalFillQty"`
	AvgPrice      string `json:"avgPrice"`
	Status        string `json:"status"`
	Lever         string `json:"lever"`
	BaseFee       string `json:"baseFee"`
	QuoteFee      string `json:"quoteFee"`
	Source        string `json:"source"`
	CreateType    string `json:"createType"`
	CancelSource  string `json:"cancelSource"`
	CancelUid     string `json:"cancelUid"`
	ReduceOnly    bool   `json:"reduceOnly"`
	TimeInForce   string `json:"timeInForce"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
	PosSide       string `json:"posSide"`
	RiskReducing  bool   `json:"riskReducing"`
	ParentOrderId string `json:"parentOrderId"`
	TpslOrder     struct {
		TpslClOrdId    string `json:"tpslClOrdId"`
		TpslMode       string `json:"tpslMode"`
		TakeProfitType string `json:"takeProfitType"`
		StopLossType   string `json:"stopLossType"`
		TakeProfit     string `json:"takeProfit"`
		StopLoss       string `json:"stopLoss"`
		TpOrderType    string `json:"tpOrderType"`
		SlOrderType    string `json:"slOrderType"`
		TpLimitPrice   string `json:"tpLimitPrice"`
		SlLimitPrice   string `json:"slLimitPrice"`
		MassQuoteOrder struct {
			Quote           bool   `json:"quote"`
			QuoteId         string `json:"quoteId"`
			QuoteSetId      string `json:"quoteSetId"`
			MmpGroup        string `json:"mmpGroup"`
			PriceAdjustment bool   `json:"priceAdjustment"`
			TradeList       []struct {
				OrderId       string `json:"orderId"`
				ClientOrderId string `json:"clientOrderId"`
				BusinessType  string `json:"businessType"`
				Symbol        string `json:"symbol"`
				Pnl           string `json:"pnl"`
				OrderType     string `json:"orderType"`
				Side          string `json:"side"`
				FillPrice     string `json:"fillPrice"`
				TradeId       string `json:"tradeId"`
				IndexPrice    string `json:"indexPrice"`
				MarkPrice     string `json:"markPrice"`
				EventId       string `json:"eventId"`
				QuoteId       string `json:"quoteId"`
				QuoteSetId    string `json:"quoteSetId"`
				RiskReducing  bool   `json:"riskReducing"`
				Role          string `json:"role"`
				FillQty       string `json:"fillQty"`
				FillTime      string `json:"fillTime"`
				ExecType      string `json:"execType"`
				Lever         string `json:"lever"`
				FeeCurrency   string `json:"feeCurrency"`
				Fee           string `json:"fee"`
				AccountName   string `json:"accountName"`
				Pid           string `json:"pid"`
			} `json:"tradeList"`
		} `json:"massQuoteOrder"`
	} `json:"tpslOrder"`
}

type WsOrderWithArg struct {
	WsSubscribeArg
	WsOrder
}

type WsOrderMiddle struct {
	WsSubscribeArg
	Data []WsOrder `json:"data"`
}

func handleWsOrder(data []byte) (*[]WsOrderWithArg, error) {
	wsOrderMiddle := &WsOrderMiddle{}
	err := json.Unmarshal(data, wsOrderMiddle)
	if err != nil {
		return nil, err
	}

	wsOrdersWithArg := []WsOrderWithArg{}
	for _, wsOrder := range wsOrderMiddle.Data {
		arg := wsOrderMiddle.WsSubscribeArg
		arg.BusinessType = wsOrder.BusinessType
		arg.Symbol = wsOrder.Symbol
		wsOrdersWithArg = append(wsOrdersWithArg, WsOrderWithArg{
			WsSubscribeArg: arg,
			WsOrder:        wsOrder,
		})
	}
	return &wsOrdersWithArg, nil
}

type TradingAccount struct {
	Currency               string `json:"currency"`               //币种
	Equity                 string `json:"equity"`                 //币种总权益
	Balance                string `json:"balance"`                //币种余额，包含现金宝部分
	RealLiability          string `json:"realLiability"`          //真实负债
	PotentialLiability     string `json:"potentialLiability"`     //潜在负债
	AccruedInterest        string `json:"accruedInterest"`        //累计利息
	Upl                    string `json:"upl"`                    //未实现盈亏，仅包含合约和交割合约未实现盈亏
	PositionInitialMargin  string `json:"positionInitialMargin"`  //合约仓位占用保证金
	OrderInitialMargin     string `json:"orderInitialMargin"`     //订单占用保证金
	Frozen                 string `json:"frozen"`                 //冻结资产
	Withdraw               string `json:"withdraw"`               //可转金额
	AvailableMargin        string `json:"availableMargin"`        //可用保证金
	LiabilityInitialMargin string `json:"liabilityInitialMargin"` //负债占用保证金
	InitialMargin          string `json:"initialMargin"`          //币种占用保证金，币种占用保证金=合约占用保证金+负债占用保证金
	RealLiabilityValue     string `json:"realLiabilityValue"`     //真实负债
}

type WsTradingAccount struct {
	WsSubscribeArg
	TradingAccount
}

type WsTradingAccountMiddle struct {
	WsSubscribeArg
	Data []TradingAccount `json:"data"`
}

func handleWsTradingAccount(data []byte) (*[]WsTradingAccount, error) {
	wsTradingAccountMiddle := &WsTradingAccountMiddle{}
	err := json.Unmarshal(data, wsTradingAccountMiddle)
	if err != nil {
		return nil, err
	}

	wsTradingAccounts := []WsTradingAccount{}

	for _, tradingAccount := range wsTradingAccountMiddle.Data {
		wsTradingAccount := WsTradingAccount{
			WsSubscribeArg: wsTradingAccountMiddle.WsSubscribeArg,
			TradingAccount: tradingAccount,
		}
		wsTradingAccounts = append(wsTradingAccounts, wsTradingAccount)
	}
	return &wsTradingAccounts, nil
}
