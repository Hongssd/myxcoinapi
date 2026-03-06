package myxcoinapi

// GET 获取充值地址 仅支持有充值权限用户进行调用
func (client *PrivateRestClient) NewPrivateRestAssetDepositAddress() *PrivateRestAssetDepositAddressAPI {
	return &PrivateRestAssetDepositAddressAPI{
		client: client,
		req:    &PrivateRestAssetDepositAddressReq{},
	}
}

func (api *PrivateRestAssetDepositAddressAPI) Do() (*XcoinRestRes[PrivateRestAssetDepositAddressRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAssetDepositAddress])
	return xcoinCallApiWithSecret[PrivateRestAssetDepositAddressRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取账户配置信息 该接口可以查询账户保证金模式，账户状态等信息
func (client *PrivateRestClient) NewPrivateRestAssetAccountInfo() *PrivateRestAssetAccountInfoAPI {
	return &PrivateRestAssetAccountInfoAPI{
		client: client,
		req:    &PrivateRestAssetAccountInfoReq{},
	}
}

func (api *PrivateRestAssetAccountInfoAPI) Do() (*XcoinRestRes[PrivateRestAssetAccountInfoRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAssetAccountInfo])
	return xcoinCallApiWithSecret[PrivateRestAssetAccountInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET 获取资金账户余额 该接口可以查询资金账户余额信息
func (client *PrivateRestClient) NewPrivateRestAssetBalances() *PrivateRestAssetBalancesAPI {
	return &PrivateRestAssetBalancesAPI{
		client: client,
		req:    &PrivateRestAssetBalancesReq{},
	}
}

func (api *PrivateRestAssetBalancesAPI) Do() (*XcoinRestRes[PrivateRestAssetBalancesRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAssetBalances])
	return xcoinCallApiWithSecret[PrivateRestAssetBalancesRes](api.client.c, url, NIL_REQBODY, GET)
}

// POST 账户内划转申请 账户内账户类型间划转，不支持跨账户划转，额度受权限配置控制
func (client *PrivateRestClient) NewPrivateRestAssetTransfer() *PrivateRestAssetTransferAPI {
	return &PrivateRestAssetTransferAPI{
		client: client,
		req:    &PrivateRestAssetTransferReq{},
	}
}

func (api *PrivateRestAssetTransferAPI) Do() (*XcoinRestRes[PrivateRestAssetTransferRes], error) {
	url := xcoinHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAssetTransfer])
	body, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return xcoinCallApiWithSecret[PrivateRestAssetTransferRes](api.client.c, url, body, POST)
}

// GET 获取账户内划转记录 获取最近三个月资金划转记录，返回数据按照划转创建时间倒序排序
func (client *PrivateRestClient) NewPrivateRestAssetTransferHistory() *PrivateRestAssetTransferHistoryAPI {
	return &PrivateRestAssetTransferHistoryAPI{
		client: client,
		req:    &PrivateRestAssetTransferHistoryReq{},
	}
}

func (api *PrivateRestAssetTransferHistoryAPI) Do() (*XcoinRestRes[PrivateRestAssetTransferHistoryRes], error) {
	url := xcoinHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAssetTransferHistory])
	return xcoinCallApiWithSecret[PrivateRestAssetTransferHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}
