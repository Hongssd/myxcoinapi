package myxcoinapi

// GET Get Server Time.
func (client *PublicRestClient) NewPublicRestMarketTime() *PublicRestMarketTimeAPI {
	return &PublicRestMarketTimeAPI{
		client: client,
		req:    &PublicRestMarketTimeReq{},
	}
}

func (api *PublicRestMarketTimeAPI) Do() (*XcoinRestRes[PublicRestMarketTimeRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTime])
	return xcoinCallAPI[PublicRestMarketTimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get Order Book (Depth).
func (client *PublicRestClient) NewPublicRestMarketDepth() *PublicRestMarketDepthAPI {
	return &PublicRestMarketDepthAPI{
		client: client,
		req:    &PublicRestMarketDepthReq{},
	}
}

func (api *PublicRestMarketDepthAPI) Do() (*XcoinRestRes[PublicRestMarketDepthRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketDepth])
	middle, err := xcoinCallAPI[PublicRestMarketDepthResMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return &XcoinRestRes[PublicRestMarketDepthRes]{
		XcoinErrRes:  middle.XcoinErrRes,
		XcoinTimeRes: middle.XcoinTimeRes,
		Data:         *middle.Data.ConvertToRes(),
	}, nil
}

// GET Get Latest Ticker Information.
func (client *PublicRestClient) NewPublicRestMarketTickerMini() *PublicRestMarketTickerMiniAPI {
	return &PublicRestMarketTickerMiniAPI{
		client: client,
		req:    &PublicRestMarketTickerMiniReq{},
	}
}

func (api *PublicRestMarketTickerMiniAPI) Do() (*XcoinRestRes[PublicRestMarketTickerMiniRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTickerMini])
	return xcoinCallAPI[PublicRestMarketTickerMiniRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get Recent Trades.
func (client *PublicRestClient) NewPublicRestMarketTrade() *PublicRestMarketTradeAPI {
	return &PublicRestMarketTradeAPI{
		client: client,
		req:    &PublicRestMarketTradeReq{},
	}
}

func (api *PublicRestMarketTradeAPI) Do() (*XcoinRestRes[PublicRestMarketTradeRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTrade])
	return xcoinCallAPI[PublicRestMarketTradeRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get 24-Hour Ticker Data.
func (client *PublicRestClient) NewPublicRestMarketTicker24hr() *PublicRestMarketTicker24hrAPI {
	return &PublicRestMarketTicker24hrAPI{
		client: client,
		req:    &PublicRestMarketTicker24hrReq{},
	}
}

func (api *PublicRestMarketTicker24hrAPI) Do() (*XcoinRestRes[PublicRestMarketTicker24hrRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTicker24hr])
	return xcoinCallAPI[PublicRestMarketTicker24hrRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get Kline Data.
func (client *PublicRestClient) NewPublicRestMarketKline() *PublicRestMarketKlineAPI {
	return &PublicRestMarketKlineAPI{
		client: client,
		req:    &PublicRestMarketKlineReq{},
	}
}

func (api *PublicRestMarketKlineAPI) Do() (*XcoinRestRes[PublicRestMarketKlineRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketKline])
	middle, err := xcoinCallAPI[PublicRestMarketKlineResMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return &XcoinRestRes[PublicRestMarketKlineRes]{
		XcoinErrRes:  middle.XcoinErrRes,
		XcoinTimeRes: middle.XcoinTimeRes,
		Data:         *middle.Data.ConvertToRes(),
	}, nil
}

// GET 	Get market price kline.
func (client *PublicRestClient) NewPublicRestMarketMarkPriceKline() *PublicRestMarketMarkPriceKlineAPI {
	return &PublicRestMarketMarkPriceKlineAPI{
		client: client,
		req:    &PublicRestMarketMarkPriceKlineReq{},
	}
}

func (api *PublicRestMarketMarkPriceKlineAPI) Do() (*XcoinRestRes[PublicRestMarketMarkPriceKlineRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketMarkPriceKline])
	middle, err := xcoinCallAPI[PublicRestMarketMarkPriceKlineResMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return &XcoinRestRes[PublicRestMarketMarkPriceKlineRes]{
		XcoinErrRes:  middle.XcoinErrRes,
		XcoinTimeRes: middle.XcoinTimeRes,
		Data:         *middle.Data.ConvertToRes(),
	}, nil
}

// GET 	Get index price kline.
// TODO: Test needed, response data return empty for now. Check the response data later.
func (client *PublicRestClient) NewPublicRestMarketIndexPriceKline() *PublicRestMarketIndexPriceKlineAPI {
	return &PublicRestMarketIndexPriceKlineAPI{
		client: client,
		req:    &PublicRestMarketIndexPriceKlineReq{},
	}
}

func (api *PublicRestMarketIndexPriceKlineAPI) Do() (*XcoinRestRes[PublicRestMarketIndexPriceKlineRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketIndexPriceKline])
	middle, err := xcoinCallAPI[PublicRestMarketIndexPriceKlineResMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	return &XcoinRestRes[PublicRestMarketIndexPriceKlineRes]{
		XcoinErrRes:  middle.XcoinErrRes,
		XcoinTimeRes: middle.XcoinTimeRes,
		Data:         *middle.Data.ConvertToRes(),
	}, nil
}

// GET 	Get Delivery History.
func (client *PublicRestClient) NewPublicRestMarketDeliveryExcerciseHistory() *PublicRestMarketDeliveryExcerciseHistoryAPI {
	return &PublicRestMarketDeliveryExcerciseHistoryAPI{
		client: client,
		req:    &PublicRestMarketDeliveryExcerciseHistoryReq{},
	}
}

func (api *PublicRestMarketDeliveryExcerciseHistoryAPI) Do() (*XcoinRestRes[PublicRestMarketDeliveryExcerciseHistoryRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketDeliveryExcerciseHistory])
	return xcoinCallAPI[PublicRestMarketDeliveryExcerciseHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET		Get Current Funding Rate.
func (client *PublicRestClient) NewPublicRestMarketFundingRate() *PublicRestMarketFundingRateAPI {
	return &PublicRestMarketFundingRateAPI{
		client: client,
		req:    &PublicRestMarketFundingRateReq{},
	}
}

func (api *PublicRestMarketFundingRateAPI) Do() (*XcoinRestRes[PublicRestMarketFundingRateRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketFundingRate])
	return xcoinCallAPI[PublicRestMarketFundingRateRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET		Get Funding Rate History.
func (client *PublicRestClient) NewPublicRestMarketFundingRateHistory() *PublicRestMarketFundingRateHistoryAPI {
	return &PublicRestMarketFundingRateHistoryAPI{
		client: client,
		req:    &PublicRestMarketFundingRateHistoryReq{},
	}
}

func (api *PublicRestMarketFundingRateHistoryAPI) Do() (*XcoinRestRes[PublicRestMarketFundingRateHistoryRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketFundingRateHistory])
	return xcoinCallAPI[PublicRestMarketFundingRateHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}
