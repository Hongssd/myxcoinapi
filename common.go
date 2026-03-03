package myxcoinapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

var httpTimeout = 100 * time.Second

func SetHttpTimeout(timeout time.Duration) {
	httpTimeout = timeout
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

type MyXcoin struct{}

const (
	XCOIN_API_HTTP = "api.xcoin.com"
	BASE_PATH      = "/api"

	IS_GZIP = false
)

type APIType int

const (
	REST APIType = iota
	WS_PUBLIC
	WS_PRIVATE
	WS_BUSINESS
)

type Client struct {
	APIKey    string
	APISecret string
}
type RestClient struct {
	c *Client
}
type PublicRestClient RestClient

type PrivateRestClient RestClient

func (*MyXcoin) NewRestClient(apiKey string, apiSecret string) *RestClient {
	client := &RestClient{
		&Client{
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
	}
	return client
}

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

var serverTimeDelta int64 = 0

func SetServerTimeDelta(delta int64) {
	serverTimeDelta = delta
}

// 通用接口调用
func xcoinCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*XcoinRestRes[T], error) {
	body, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 通用鉴权接口调用
func xcoinCallApiWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*XcoinRestRes[T], error) {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixMilli(), BIT_BASE_10)
	requestPath := strings.TrimPrefix(url.Path, BASE_PATH)
	query := url.RawQuery

	hmacSha256Data := timestamp + method + requestPath
	if query != "" {
		hmacSha256Data += "?" + query
	}
	if len(reqBody) != 0 {
		hmacSha256Data += string(reqBody)
	}
	sign := hex.EncodeToString(HmacSha256(client.APISecret, hmacSha256Data))

	// log.Warn(hmacSha256Data)
	// log.Warn("timestamp: ", timestamp)
	// log.Warn("method: ", method)
	// log.Warn("requestPath: ", requestPath)
	// log.Warn("query: ", query)
	// log.Warn("reqBody: ", string(reqBody))
	// log.Warn("hmacSha256Data: ", hmacSha256Data)
	// log.Warn("sign: ", sign)

	body, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"X-ACCESS-APIKEY":    client.APIKey,
			"X-ACCESS-TIMESTAMP": timestamp,
			"X-ACCESS-SIGN":      sign,
			// "X-ACCESS-RECV-WINDOW": "5000",
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// URL标准封装 带路径参数
func xcoinHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	query := xcoinHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     XCOIN_API_HTTP,
		Path:     BASE_PATH + name,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func xcoinHandlerRequestAPIWithoutPathQueryParam(apiType APIType, name string) url.URL {
	u := url.URL{
		Scheme:   "https",
		Host:     XCOIN_API_HTTP,
		Path:     BASE_PATH + name,
		RawQuery: "",
	}
	return u
}

func xcoinHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}
