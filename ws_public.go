package myxcoinapi

import (
	"errors"
	"fmt"
)

func getTicker24hrArg(businessType string, symbol string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       "ticker24hr",
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// 订阅24H行情频道
func (ws *PublicWsStreamClient) SubscribeTicker24hrMulti(businessType string, symbols ...string) (*Subscription[WsTicker24hr], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}
	if len(symbols) == 0 {
		return nil, errors.New("symbols is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getTicker24hrArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getTicker24hrArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeBooks Success: args:%v", doSub.Args)

	sub := &Subscription[WsTicker24hr]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsTicker24hr),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.ticker24hrSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅24H行情频道
func (ws *PublicWsStreamClient) UnsubscribeTicker24hrMulti(businessType string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}
	if len(symbols) == 0 {
		return errors.New("symbols is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getTicker24hrArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getTicker24hrArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeTicker24hrMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})

		keyData, _ := json.Marshal(arg)
		ws.ticker24hrSubMap.Delete(string(keyData))
	}

	return nil
}

func getKlineArg(businessType string, symbol string, period string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       fmt.Sprintf("kline#%s", period),
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// 订阅K线频道
func (ws *PublicWsStreamClient) SubscribeKlineMulti(businessType string, periods []string, symbols ...string) (*Subscription[WsKline], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}
	if len(periods) == 0 {
		return nil, errors.New("periods is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getKlineArg(businessType, "", ""))
	} else {
		for _, s := range symbols {
			for _, period := range periods {
				args = append(args, getKlineArg(businessType, s, period))
			}
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	log.Infof("SubscribeKlineMulti Success: args:%v", doSub.Args)

	sub := &Subscription[WsKline]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsKline),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.klineSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅K线频道
func (ws *PublicWsStreamClient) UnsubscribeKlineMulti(businessType string, periods []string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}
	if len(periods) == 0 {
		return errors.New("periods is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getKlineArg(businessType, "", ""))
	} else {
		for _, s := range symbols {
			for _, period := range periods {
				args = append(args, getKlineArg(businessType, s, period))
			}
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeKlineMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.klineSubMap.Delete(string(keyData))
	}

	return nil
}

func getDepthArg(businessType string, symbol string, interval string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       fmt.Sprintf("depth#%s", interval),
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

type WsDepthIntervalType string

// interval:数据接收频率，单位毫秒；不填默认为 100ms 频率
// 取值：100ms，500ms，1000ms
// 默认：100ms
const (
	WS_DEPTH_INTERVAL_TYPE_100  WsDepthIntervalType = "100ms"
	WS_DEPTH_INTERVAL_TYPE_500  WsDepthIntervalType = "500ms"
	WS_DEPTH_INTERVAL_TYPE_1000 WsDepthIntervalType = "1000ms"
)

func (i WsDepthIntervalType) String() string {
	return string(i)
}

// 订阅增量深度频道
func (ws *PublicWsStreamClient) SubscribeDepthMulti(businessType string, interval WsDepthIntervalType, symbols ...string) (*Subscription[WsDepth], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getDepthArg(businessType, "", ""))
	} else {
		for _, s := range symbols {
			args = append(args, getDepthArg(businessType, s, interval.String()))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	log.Infof("SubscribeDepthMulti Success: args:%v", doSub.Args)

	sub := &Subscription[WsDepth]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsDepth),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.depthSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅增量深度频道
func (ws *PublicWsStreamClient) UnsubscribeDepthMulti(businessType string, interval WsDepthIntervalType, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getDepthArg(businessType, "", ""))
	} else {
		for _, s := range symbols {
			args = append(args, getDepthArg(businessType, s, interval.String()))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeDepthMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.depthSubMap.Delete(string(keyData))
	}

	return nil
}

type WsDepthLevelsType string

// levels 表示几档买卖单信息，可选 5/10/20/30 档，默认 5 档
const (
	WS_DEPTH_LEVELS_TYPE_5  WsDepthLevelsType = "5"
	WS_DEPTH_LEVELS_TYPE_10 WsDepthLevelsType = "10"
	WS_DEPTH_LEVELS_TYPE_20 WsDepthLevelsType = "20"
	WS_DEPTH_LEVELS_TYPE_30 WsDepthLevelsType = "30"
)

func (i WsDepthLevelsType) String() string {
	return string(i)
}

type WsDepthLevelsIntervalType string

// 数据接收频率，单位毫秒；不填默认为 100ms 频率 取值：100ms，500ms，1000ms
const (
	WS_DEPTH_LEVELS_INTERVAL_TYPE_100  WsDepthLevelsIntervalType = "100ms"
	WS_DEPTH_LEVELS_INTERVAL_TYPE_500  WsDepthLevelsIntervalType = "500ms"
	WS_DEPTH_LEVELS_INTERVAL_TYPE_1000 WsDepthLevelsIntervalType = "1000ms"
)

func (i WsDepthLevelsIntervalType) String() string {
	return string(i)
}

func getDepthLevelsArg(businessType string, symbol string, interval WsDepthLevelsIntervalType, levels WsDepthLevelsType, group string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       fmt.Sprintf("depthlevels#%s", interval.String()),
		BusinessType: businessType,
		Symbol:       symbol,
		Levels:       levels.String(),
		Group:        group,
	}
}

// 订阅有限档深度快照频道
func (ws *PublicWsStreamClient) SubscribeDepthLevelsMulti(businessType string, interval WsDepthLevelsIntervalType, levels WsDepthLevelsType, group string, symbols []string) (*Subscription[WsDepthLevels], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}
	if len(symbols) == 0 {
		return nil, errors.New("symbols is required")
	}

	args := []WsSubscribeArg{}
	for _, s := range symbols {
		args = append(args, getDepthLevelsArg(businessType, s, interval, levels, group))
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	log.Infof("SubscribeDepthLevelsMulti Success: args:%v", doSub.Args)

	sub := &Subscription[WsDepthLevels]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsDepthLevels),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(WsSubscribeArg{
			Stream:       arg.Stream,
			BusinessType: arg.BusinessType,
			Symbol:       arg.Symbol,
		})
		ws.depthLevelsSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅有限档深度快照频道
func (ws *PublicWsStreamClient) UnsubscribeDepthLevelsMulti(businessType string, interval WsDepthLevelsIntervalType, levels WsDepthLevelsType, group string, symbols []string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}
	if len(symbols) == 0 {
		return errors.New("symbols is required")
	}

	args := []WsSubscribeArg{}
	for _, s := range symbols {
		args = append(args, getDepthLevelsArg(businessType, s, interval, levels, group))
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeDepthLevelsMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{{
			Stream:       arg.Stream,
			BusinessType: arg.BusinessType,
			Symbol:       arg.Symbol,
		}})
		keyData, _ := json.Marshal(WsSubscribeArg{
			Stream:       arg.Stream,
			BusinessType: arg.BusinessType,
			Symbol:       arg.Symbol,
		})
		ws.depthLevelsSubMap.Delete(string(keyData))
	}

	return nil
}

func getOrderbookArg(businessType string, symbol string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       "orderBook",
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// 订阅当前最优挂单频道
func (ws *PublicWsStreamClient) SubscribeOrderbookMulti(businessType string, symbols ...string) (*Subscription[WsOrderbook], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getOrderbookArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getOrderbookArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	log.Infof("SubscribeOrderbookMulti Success: args:%v", doSub.Args)

	sub := &Subscription[WsOrderbook]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsOrderbook),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		ws.orderbookSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅当前最优挂单频道
func (ws *PublicWsStreamClient) UnsubscribeOrderbookMulti(businessType string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getOrderbookArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getOrderbookArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeOrderbookMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.orderbookSubMap.Delete(string(keyData))
	}

	return nil
}

func getTradeArg(businessType string, symbol string) WsSubscribeArg {
	return WsSubscribeArg{
		Stream:       "trade",
		BusinessType: businessType,
		Symbol:       symbol,
	}
}

// 订阅成交频道
func (ws *PublicWsStreamClient) SubscribeTradeMulti(businessType string, symbols ...string) (*Subscription[WsTrade], error) {
	if businessType == "" {
		return nil, errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getTradeArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getTradeArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, SUBSCRIBE)
	if err != nil {
		return nil, err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}

	log.Infof("SubscribeTradeMulti Success: args:%v", doSub.Args)

	sub := &Subscription[WsTrade]{
		SubId:        doSub.SubId,
		Ws:           &ws.WsStreamClient,
		Event:        SUBSCRIBE,
		Args:         args,
		resultChan:   make(chan WsTrade),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(arg)
		log.Infof("keyData: %s", string(keyData))
		ws.tradeSubMap.Store(string(keyData), sub)
	}

	return sub, nil
}

// 取消订阅成交频道
func (ws *PublicWsStreamClient) UnsubscribeTradeMulti(businessType string, symbols ...string) error {
	if businessType == "" {
		return errors.New("businessType is required")
	}

	args := []WsSubscribeArg{}
	if len(symbols) == 0 {
		args = append(args, getTradeArg(businessType, ""))
	} else {
		for _, s := range symbols {
			args = append(args, getTradeArg(businessType, s))
		}
	}

	doSub, err := subscribe[WsSubscribeResult](&ws.WsStreamClient, args, UNSUBSCRIBE)
	if err != nil {
		return err
	}

	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}

	log.Infof("UnsubscribeTradeMulti Success: args:%v", doSub.Args)

	for _, arg := range args {
		doSub.Ws.sendUnSubscribeSuccessToCloseChan([]WsSubscribeArg{arg})
		keyData, _ := json.Marshal(arg)
		ws.tradeSubMap.Delete(string(keyData))
	}
	return nil
}
