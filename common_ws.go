package myxcoinapi

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
)

const (
	XCOIN_API_WS_HOST = "stream.xcoin.com"

	XCOIN_API_WS_API_PUBLIC   = "/ws/public/v1/market"
	XCOIN_API_WS_API_PRIVATE  = "/ws/private/v2/notification"
	XCOIN_API_WS_API_BUSINESS = "/ws/private/trade/v2/trade"

	WS_PRIVATE_BASE_PATH = "/ws/private"
)

const (
	SUBSCRIBE   = "subscribe"   //订阅
	UNSUBSCRIBE = "unsubscribe" //取消订阅
)

var (
	WebsocketTimeout        = time.Second * 10
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

type WsAuthData struct {
	Type            string `json:"type"`                  // 认证类型，Token
	AccessKey       string `json:"accessKey"`             // 用户accessKey当检测到用户accessKey过期后服务端会主动发送给客户端用户token过期的消息，并断开和客户端的连接。客户端收到accessKey过期消息后需要重新建立连接并订阅主题。
	AccessTimestamp string `json:"accessTimestamp"`       // 当前时间戳
	AccountName     string `json:"accountName,omitempty"` // 账户名称：账户级 APIKEY 时，此字段为空；成员级 APIKEY 时，此字段必填
}

type WsAuthReq struct {
	Event      string     `json:"event"`      // 事件类型，authorization
	AccessSign string     `json:"accessSign"` // 验签，根据 data 对象进行验签，示例参考下方 Java demo
	Data       WsAuthData `json:"data"`       // 数据
}

type WsAuthResult struct {
	Event   string `json:"event"`   // Event type : authorization
	Code    int    `json:"code"`    // Result code: 0 indicates success; non-0 indicates an error. See msg for error details.
	Message string `json:"message"` // Error message field; it indicates success when code = 0
}

type WsStreamClient struct {
	client          *RestClient
	lastAuth        *WsAuthReq
	apiType         APIType
	conn            *websocket.Conn
	connId          string
	commonSubMap    MySyncMap[string, *Subscription[WsSubscribeResult]] //订阅的返回结果
	waitSubResult   *Subscription[WsSubscribeResult]                    //订阅结果
	waitSubResultMu *sync.Mutex                                         //订阅结果锁

	authResultChan chan *WsAuthResult // 鉴权结果通道

	// Public
	ticker24hrSubMap  MySyncMap[string, *Subscription[WsTicker24hr]]  //ticker24hr订阅结果
	klineSubMap       MySyncMap[string, *Subscription[WsKline]]       //kline订阅结果
	depthSubMap       MySyncMap[string, *Subscription[WsDepth]]       //depth订阅结果
	depthLevelsSubMap MySyncMap[string, *Subscription[WsDepthLevels]] //depthLevels订阅结果
	orderbookSubMap   MySyncMap[string, *Subscription[WsOrderbook]]   //orderbook订阅结果(当前最优挂单频道 BBO)
	tradeSubMap       MySyncMap[string, *Subscription[WsTrade]]       //trade订阅结果

	// Private
	positionSubMap       MySyncMap[string, *Subscription[WsPosition]]       //position订阅结果
	orderSubMap          MySyncMap[string, *Subscription[WsOrder]]          //order订阅结果
	tradingAccountSubMap MySyncMap[string, *Subscription[WsTradingAccount]] //tradingAccount订阅结果

	isAuth     bool // 是否已鉴权
	resultChan chan []byte
	errChan    chan error
	isClose    bool

	reSubscribeMu      *sync.Mutex
	AutoReConnectTimes int //自动重连次数

	writeMu *sync.Mutex
}

type PublicWsStreamClient struct {
	WsStreamClient
}

func (*MyXcoin) NewPublicWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:         WS_PUBLIC,
			commonSubMap:    NewMySyncMap[string, *Subscription[WsSubscribeResult]](),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},

			reSubscribeMu:     &sync.Mutex{},
			ticker24hrSubMap:  NewMySyncMap[string, *Subscription[WsTicker24hr]](),
			klineSubMap:       NewMySyncMap[string, *Subscription[WsKline]](),
			depthSubMap:       NewMySyncMap[string, *Subscription[WsDepth]](),
			depthLevelsSubMap: NewMySyncMap[string, *Subscription[WsDepthLevels]](),
			orderbookSubMap:   NewMySyncMap[string, *Subscription[WsOrderbook]](),
			tradeSubMap:       NewMySyncMap[string, *Subscription[WsTrade]](),
		},
	}
}

type PrivateWsStreamClient struct {
	WsStreamClient
}

func (*MyXcoin) NewPrivateWsStreamClient(client *RestClient) *PrivateWsStreamClient {
	return &PrivateWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:         WS_PRIVATE,
			client:          client,
			commonSubMap:    NewMySyncMap[string, *Subscription[WsSubscribeResult]](),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},

			authResultChan: make(chan *WsAuthResult, 1),

			positionSubMap:       NewMySyncMap[string, *Subscription[WsPosition]](),
			orderSubMap:          NewMySyncMap[string, *Subscription[WsOrder]](),
			tradingAccountSubMap: NewMySyncMap[string, *Subscription[WsTradingAccount]](),

			reSubscribeMu: &sync.Mutex{},
			writeMu:       &sync.Mutex{},
		},
	}
}

type BusinessWsStreamClient struct {
	WsStreamClient
}

func (*MyXcoin) NewBusinessWsStreamClient() *BusinessWsStreamClient {
	return &BusinessWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:         WS_BUSINESS,
			commonSubMap:    NewMySyncMap[string, *Subscription[WsSubscribeResult]](),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},

			authResultChan: make(chan *WsAuthResult, 1),

			reSubscribeMu: &sync.Mutex{},
			writeMu:       &sync.Mutex{},
		},
	}
}

type Subscription[T any] struct {
	SubId        int64            //订阅ID
	Ws           *WsStreamClient  //订阅的连接
	Event        string           //订阅方法  subscribe/unsubscribe
	Args         []WsSubscribeArg //订阅参数
	resultChan   chan T           //接收订阅结果的通道
	errChan      chan error       //接收订阅错误的通道
	closeChan    chan struct{}    //接收订阅关闭的通道
	subResultMap map[string]bool  //订阅结果
}

type WsSubscribeReq struct {
	Event string           `json:"event"`
	Data  []WsSubscribeArg `json:"data"`
}

type WsSubscribeArg struct {
	Stream       string `json:"stream"`
	BusinessType string `json:"businessType"`
	Symbol       string `json:"symbol"`

	// 有限档深度快照频道
	Levels string `json:"levels,omitempty"`
	Group  string `json:"group,omitempty"`
}

// 登陆及订阅返回结果
type WsSubscribeResult struct {
	Event   string `json:"event"` // Event type, e.g. subscribe, unsubscribe
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Code    int    `json:"code"` // API 返回数字，如 0 表示成功
		Message string `json:"message"`
		WsSubscribeArg
	} `json:"data"` // Data
	Ts int64 `json:"ts"` // Server response timestamp in Unix milliseconds, e.g., 1732158178000
}

// 获取数据流请求Path
func getWsApi(apiType APIType) string {
	switch apiType {
	case WS_PUBLIC:
		return XCOIN_API_WS_API_PUBLIC
	case WS_PRIVATE:
		return XCOIN_API_WS_API_PRIVATE
	case WS_BUSINESS:
		return XCOIN_API_WS_API_BUSINESS
	default:
		log.Error("apiType Error is ", apiType)
		return ""
	}
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType) string {
	query := ""
	u := url.URL{
		Scheme:   "wss",
		Host:     XCOIN_API_WS_HOST,
		Path:     getWsApi(apiType),
		RawQuery: query,
	}
	return u.String()
}

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > 3*timeout {
				err := c.Close()
				if err != nil {
					log.Error(err)
					return
				}
				return
			}
		}
	}()
}

// 标准订阅方法
func wsStreamServe(api string, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	if WsUseProxy {
		proxy, _ := getBestProxyAndWeight()
		if proxy == nil {
			return nil, errors.New("no proxy available")
		}
		proxyUrl, err := url.Parse(proxy.ProxyUrl)
		if err != nil {
			return nil, err
		}
		dialer = &websocket.Dialer{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
	c, _, err := dialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- message
		}
	}()
	return c, err
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.apiType)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.isClose = false
		ws.connId = ""
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.connId = ""
		log.Info("Auto ReOpenConn success to ", apiUrl)
		return err
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	args := []WsSubscribeArg{}
	ws.commonSubMap.Range(func(key string, _ *Subscription[WsSubscribeResult]) bool {
		arg := WsSubscribeArg{}
		err := json.Unmarshal([]byte(key), &arg)
		if err != nil {
			return false
		}
		args = append(args, arg)
		return true
	})
	ws.sendUnSubscribeSuccessToCloseChan(args)
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true
	ws.connId = ""

	if ws.conn != nil {
		err := ws.conn.Close()
		if err != nil {
			return err
		}
	}
	//手动关闭成功，给所有订阅发送关闭信号
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.lastAuth = nil
	ws.commonSubMap = NewMySyncMap[string, *Subscription[WsSubscribeResult]]()
	ws.ticker24hrSubMap = NewMySyncMap[string, *Subscription[WsTicker24hr]]()
	ws.klineSubMap = NewMySyncMap[string, *Subscription[WsKline]]()
	ws.depthSubMap = NewMySyncMap[string, *Subscription[WsDepth]]()
	ws.depthLevelsSubMap = NewMySyncMap[string, *Subscription[WsDepthLevels]]()
	ws.orderbookSubMap = NewMySyncMap[string, *Subscription[WsOrderbook]]()
	ws.tradeSubMap = NewMySyncMap[string, *Subscription[WsTrade]]()
	ws.positionSubMap = NewMySyncMap[string, *Subscription[WsPosition]]()
	ws.orderSubMap = NewMySyncMap[string, *Subscription[WsOrder]]()
	ws.tradingAccountSubMap = NewMySyncMap[string, *Subscription[WsTradingAccount]]()

	if ws.waitSubResult != nil {
		//给当前等待订阅结果的请求返回错误
		ws.waitSubResultMu.Lock()
		ws.waitSubResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
	}

	return nil
}

// sendAuthMessage 发送鉴权消息
func (ws *WsStreamClient) sendAuthMessage() (*WsAuthReq, error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

	if ws.client == nil || ws.client.c.APIKey == "" || ws.client.c.APISecret == "" {
		return nil, fmt.Errorf("client credentials not set")
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixMilli(), BIT_BASE_10)
	wsAuthData := WsAuthData{
		Type:            "Token",
		AccessKey:       ws.client.c.APIKey,
		AccessTimestamp: timestamp,
	}

	// calc accessSign
	authBody, err := json.Marshal(wsAuthData)
	if err != nil {
		return nil, err
	}

	var path string
	switch ws.apiType {
	case WS_PRIVATE:
		path = XCOIN_API_WS_API_PRIVATE
	case WS_BUSINESS:
		path = XCOIN_API_WS_API_BUSINESS
	default:
		return nil, fmt.Errorf("apiType not supported")
	}

	preHash := timestamp + "POST" + strings.TrimPrefix(path, WS_PRIVATE_BASE_PATH) + "" + string(authBody)
	accessSign := hex.EncodeToString(HmacSha256(ws.client.c.APISecret, preHash))

	wsAuthReq := WsAuthReq{
		Event:      "authorization",
		AccessSign: accessSign,
		Data:       wsAuthData,
	}
	data, err := json.Marshal(wsAuthReq)
	if err != nil {
		return nil, err
	}
	log.Debugf("Sending auth message: %s", string(data))

	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()
	return &wsAuthReq, ws.conn.WriteMessage(websocket.TextMessage, data)
}

func (ws *WsStreamClient) Auth() error {
	if ws == nil || ws.conn == nil || ws.isClose {
		return fmt.Errorf("websocket is close")
	}
	if ws.client == nil || ws.client.c.APIKey == "" || ws.client.c.APISecret == "" {
		return fmt.Errorf("client credentials not set")
	}

	// 初始化鉴权结果通道
	ws.authResultChan = make(chan *WsAuthResult, 1)

	// 发送鉴权消息
	authReq, err := ws.sendAuthMessage()
	if err != nil {
		return err
	}

	// 同步等待鉴权结果
	select {
	case authRes := <-ws.authResultChan:
		if authRes == nil {
			return fmt.Errorf("auth failed: nil response")
		}
		if authRes.Code != 0 {
			ws.isAuth = false
			return fmt.Errorf("auth failed: [%v] %s", authRes.Code, authRes.Message)
		}
		ws.isAuth = true
		ws.lastAuth = authReq
		log.Info("Auth success")
		return nil
	case <-time.After(10 * time.Second):
		return fmt.Errorf("auth timeout")
	}

}

// 订阅
func subscribe[T any](ws *WsStreamClient, args []WsSubscribeArg, event string) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}

	ws.waitSubResultMu.Lock()

	subscribeReq := WsSubscribeReq{
		Event: event,
		Data:  args,
	}

	data, err := json.Marshal(subscribeReq)
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}
	log.Debugf("send msg: %s", string(data))

	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}

	node, err := snowflake.NewNode(2)
	if err != nil {
		ws.waitSubResultMu.Unlock()
		return nil, err
	}

	id := node.Generate().Int64()

	sub := &Subscription[T]{
		SubId:        id,
		Ws:           ws,
		Event:        event,
		Args:         args,
		resultChan:   make(chan T),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		subResultMap: make(map[string]bool),
	}

	return sub, nil
}

func (ws *WsStreamClient) reSubscribeForReconnect() error {
	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()
	isDoReSubscribe := map[int64]bool{}
	var wErr error
	ws.commonSubMap.Range(func(_ string, sub *Subscription[WsSubscribeResult]) bool {
		if _, ok := isDoReSubscribe[sub.SubId]; ok {
			return true
		}

		reSub, err := subscribe[WsSubscribeResult](ws, sub.Args, sub.Event)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		err = ws.catchSubscribeResult(reSub)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		log.Infof("reSubscribe Success: args:%v", reSub.Args)

		sub.SubId = reSub.SubId
		isDoReSubscribe[sub.SubId] = true
		time.Sleep(500 * time.Millisecond)
		return true
	})
	return wErr
}

func (ws *WsStreamClient) DeferSub() {
	if ws.waitSubResult != nil {
		if len(ws.waitSubResult.subResultMap) == len(ws.waitSubResult.Args) {
			for _, arg := range ws.waitSubResult.Args {
				keyData, _ := json.Marshal(&arg)
				ws.commonSubMap.Store(string(keyData), ws.waitSubResult)
			}
		}
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
	}
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result *WsSubscribeResult) {
	if ws.waitSubResult != nil {
		if result.Event == "error" || result.Code != 0 {
			ws.waitSubResult.errChan <- fmt.Errorf("errHandler: %+v", result)
			return
		}
		flag := false
		for _, data := range result.Data {
			if data.Code != 0 {
				ws.waitSubResult.errChan <- fmt.Errorf("errHandler: %+v", data)
				flag = true
			}
		}
		if !flag {
			ws.waitSubResult.resultChan <- *result
		}
		return
	}
}

// 捕获订阅结果
func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsSubscribeResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()
	isBreak := false
	for {
		select {
		case err := <-sub.ErrChan():
			log.Error(err)
			return fmt.Errorf("SubAction Error: %v", err)
		case subResult := <-sub.ResultChan():
			hasErr := false
			for _, data := range subResult.Data {
				if data.Code != 0 {
					log.Error(data.Code, ":", data.Message)
					hasErr = true
					break
				}
			}
			if hasErr {
				return fmt.Errorf("errHandler: subscribe result code != 0")
			}
			// 服务端可能一条消息返回多个 Data 项，需按项写入 map 才能正确计数
			for _, data := range subResult.Data {
				keyData, _ := json.Marshal(data)
				sub.subResultMap[string(keyData)] = true
			}
			if len(sub.subResultMap) == len(sub.Args) {
				isBreak = true
			}
		}
		if isBreak {
			break
		}
	}
	log.Debug("catchResults: ", sub.subResultMap)
	return nil
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(args []WsSubscribeArg) {
	for _, arg := range args {
		data, _ := json.Marshal(arg)
		key := string(data)
		if sub, ok := ws.ticker24hrSubMap.Load(key); ok {
			ws.ticker24hrSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
	}
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					//重连
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(1500 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1
					go func() {
						//重新登陆
						if ws.lastAuth != nil && ws.client != nil {
							err = ws.Auth()
							for err != nil {
								time.Sleep(1500 * time.Millisecond)
								err = ws.Auth()
							}
						}

						//重新订阅
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()
				}
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}

				log.Debugf("receive msg: %s", string(data))

				// 处理鉴权结果
				if strings.Contains(string(data), "authorization") {
					result := &WsAuthResult{}
					err := json.Unmarshal(data, result)
					if err != nil {
						log.Error(err)
						continue
					}
					log.Debugf("receive auth result: %+v", result)
					ws.authResultChan <- result
					continue
				}

				// 处理订阅或查询订阅列表请求返回结果
				// 这里必须严格匹配字段名 "event"，防止被 "eventId" 等字段误触发
				if strings.Contains(string(data), "\"event\"") {
					result := &WsSubscribeResult{}
					err := json.Unmarshal(data, result)
					if err != nil {
						log.Error(err)
						continue
					}
					log.Debugf("receive subscribe result: %+v", result)
					ws.sendSubscribeResultToChan(result)
					continue
				}

				// ticker24hr订阅结果处理
				if strings.Contains(string(data), "ticker24hr") {
					t, err := handleWsTicker24hr(data)
					arg := t.WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					// log.Warnf("keyData: %s", string(keyData))
					if sub, ok := ws.ticker24hrSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *t
					}
					continue
				}

				// kline订阅结果处理
				if strings.Contains(string(data), "kline#") {
					k, err := handleWsKlines(data)
					if len(*k) == 0 {
						log.Warnf("kline is empty, skip")
						continue
					}
					arg := (*k)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					// log.Warnf("keyData: %s", string(keyData))
					if sub, ok := ws.klineSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, kline := range *k {
							sub.resultChan <- kline
						}
					}
					continue
				}

				// depth订阅结果处理
				if strings.Contains(string(data), "depth#") {
					d, err := handleWsDepth(data)
					if len(*d) == 0 {
						log.Warnf("depth is empty, skip")
						continue
					}
					arg := (*d)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					if sub, ok := ws.depthSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, depth := range *d {
							sub.resultChan <- depth
						}
					}
					continue
				}

				// depthLevels订阅结果处理
				if strings.Contains(string(data), "depthlevels#") {
					d, err := handleWsDepthLevels(data)
					if len(*d) == 0 {
						log.Warnf("depthLevels is empty, skip")
						continue
					}
					arg := (*d)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					if sub, ok := ws.depthLevelsSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, d := range *d {
							sub.resultChan <- d
						}
					}
					continue
				}

				// orderbook订阅结果处理
				if strings.Contains(string(data), "orderBook") {
					o, err := handleWsOrderbook(data)
					if len(*o) == 0 {
						log.Warnf("orderBook is empty, skip")
						continue
					}
					arg := (*o)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					if sub, ok := ws.orderbookSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, o := range *o {
							sub.resultChan <- o
						}
					}
					continue
				}

				// trade订阅结果处理
				if strings.Contains(string(data), "trade") {
					t, err := handleWsTrade(data)
					if len(*t) == 0 {
						log.Warnf("trade is empty, skip")
						continue
					}

					arg := (*t)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					// log.Warnf("keyData: %s", string(keyData))
					if sub, ok := ws.tradeSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, trade := range *t {
							sub.resultChan <- trade
						}
					}
					continue
				}

				// position订阅结果处理
				if strings.Contains(string(data), "position") {
					p, err := handleWsPosition(data)
					if len(*p) == 0 {
						log.Warnf("position is empty, skip")
						continue
					}
					arg := (*p)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					if sub, ok := ws.positionSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, p := range *p {
							sub.resultChan <- p
						}
					}
				}

				// order订阅结果处理
				if strings.Contains(string(data), "order") {
					o, err := handleWsOrder(data)
					if len(*o) == 0 {
						log.Warnf("order is empty, skip")
						continue
					}
					arg := (*o)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream:       arg.Stream,
						BusinessType: arg.BusinessType,
						Symbol:       arg.Symbol,
					})
					if sub, ok := ws.orderSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, order := range *o {
							sub.resultChan <- order.WsOrder
						}
					}
				}

				//tradingAccount订阅结果处理
				if strings.Contains(string(data), "trading_account") {
					t, err := handleWsTradingAccount(data)
					if len(*t) == 0 {
						log.Warnf("tradingAccount is empty, skip")
						continue
					}
					arg := (*t)[0].WsSubscribeArg
					keyData, _ := json.Marshal(WsSubscribeArg{
						Stream: arg.Stream,
					})
					if sub, ok := ws.tradingAccountSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, tradingAccount := range *t {
							sub.resultChan <- tradingAccount
						}
					}
				}
			}
		}
	}()
}
