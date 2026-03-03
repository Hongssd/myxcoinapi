package myxcoinapi

import (
	"fmt"
)

type XcoinErrRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type XcoinTimeRes struct {
	Ts string `json:"ts"`
}

type XcoinRestRes[T any] struct {
	XcoinErrRes    // 错误信息
	XcoinTimeRes   // 时间戳
	Data         T `json:"data"` // 返回结果
}

func handlerCommonRest[T any](data []byte) (*XcoinRestRes[T], error) {
	res := &XcoinRestRes[T]{}
	if err := json.Unmarshal(data, res); err != nil {
		log.Error(string(data))
		log.Error("rest返回值获取失败: ", err)
		return nil, err
	}
	return nil, res.handlerError()
}

func (err *XcoinErrRes) handlerError() error {
	if err.Code != "0" {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	}
	return nil
}
