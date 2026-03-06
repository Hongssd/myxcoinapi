package myxcoinapi

type PrivateRestAssetDepositAddressAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetDepositAddressReq
}

type PrivateRestAssetDepositAddressReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Currency    *string `json:"currency"`    // false 币种，如 BTC
	ChainType   *string `json:"chainType"`   // false 链的类型
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAssetDepositAddressAPI) AccountName(accountName string) *PrivateRestAssetDepositAddressAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 币种，如 BTC
func (api *PrivateRestAssetDepositAddressAPI) Currency(currency string) *PrivateRestAssetDepositAddressAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// string false 链的类型
func (api *PrivateRestAssetDepositAddressAPI) ChainType(chainType string) *PrivateRestAssetDepositAddressAPI {
	api.req.ChainType = GetPointer(chainType)
	return api
}

type PrivateRestAssetAccountInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetAccountInfoReq
}

type PrivateRestAssetAccountInfoReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAssetAccountInfoAPI) AccountName(accountName string) *PrivateRestAssetAccountInfoAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

type PrivateRestAssetBalancesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetBalancesReq
}

type PrivateRestAssetBalancesReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAssetBalancesAPI) AccountName(accountName string) *PrivateRestAssetBalancesAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

type PrivateRestAssetTransferAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetTransferReq
}

type PrivateRestAssetTransferReq struct {
	AccountName      *string `json:"accountName"`      // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Currency         *string `json:"currency"`         // true 币种，如 BTC
	Amount           *string `json:"amount"`           // true 划转金额
	ClientTransferId *string `json:"clientTransferId"` // false 客户划转单ID
	LoanTrans        *bool   `json:"loanTrans"`        // false 是否开启借币转出
	FromAccountType  *string `json:"fromAccountType"`  // true 转出账户
	ToAccountType    *string `json:"toAccountType"`    // true 划入账户
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAssetTransferAPI) AccountName(accountName string) *PrivateRestAssetTransferAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string true 币种，如 BTC
func (api *PrivateRestAssetTransferAPI) Currency(currency string) *PrivateRestAssetTransferAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// string true 划转金额
func (api *PrivateRestAssetTransferAPI) Amount(amount string) *PrivateRestAssetTransferAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

// string false 客户划转单ID
func (api *PrivateRestAssetTransferAPI) ClientTransferId(clientTransferId string) *PrivateRestAssetTransferAPI {
	api.req.ClientTransferId = GetPointer(clientTransferId)
	return api
}

// string false 是否开启借币转出
func (api *PrivateRestAssetTransferAPI) LoanTrans(loanTrans bool) *PrivateRestAssetTransferAPI {
	api.req.LoanTrans = GetPointer(loanTrans)
	return api
}

// string true 转出账户
func (api *PrivateRestAssetTransferAPI) FromAccountType(fromAccountType string) *PrivateRestAssetTransferAPI {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

// string true 划入账户
func (api *PrivateRestAssetTransferAPI) ToAccountType(toAccountType string) *PrivateRestAssetTransferAPI {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}

type PrivateRestAssetTransferHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetTransferHistoryReq
}

type PrivateRestAssetTransferHistoryReq struct {
	AccountName *string `json:"accountName"` // false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
	Currency    *string `json:"currency"`    // false 币种，如 BTC
	BeginTime   *string `json:"beginTime"`   // false 筛选划转创建的开始时间戳 Unix时间戳的毫秒数格式，如 1732158178000
	EndTime     *string `json:"endTime"`     // false 筛选划转创建的结束时间戳 Unix时间戳的毫秒数格式，如 1732182494000
	BeginId     *string `json:"beginId"`     // false 请求此ID之后的分页内容，传的值对应接口的id字段
	EndId       *string `json:"endId"`       // false 请求此ID之前的分页内容，传的值对应接口的id字段
	Limit       *string `json:"limit"`       // false 分页返回的结果集数量 最大为100，不填默认返回100条
}

// string false 账户名称 账户级APIKEY时，此字段为空；成员级APIKEY时，此字段必填
func (api *PrivateRestAssetTransferHistoryAPI) AccountName(accountName string) *PrivateRestAssetTransferHistoryAPI {
	api.req.AccountName = GetPointer(accountName)
	return api
}

// string false 币种，如 BTC
func (api *PrivateRestAssetTransferHistoryAPI) Currency(currency string) *PrivateRestAssetTransferHistoryAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// string false 筛选划转创建的开始时间戳 Unix时间戳的毫秒数格式，如 1732158178000
func (api *PrivateRestAssetTransferHistoryAPI) BeginTime(beginTime string) *PrivateRestAssetTransferHistoryAPI {
	api.req.BeginTime = GetPointer(beginTime)
	return api
}

// string false 筛选划转创建的结束时间戳 Unix时间戳的毫秒数格式，如 1732182494000
func (api *PrivateRestAssetTransferHistoryAPI) EndTime(endTime string) *PrivateRestAssetTransferHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// string false 请求此ID之后的分页内容，传的值对应接口的id字段
func (api *PrivateRestAssetTransferHistoryAPI) BeginId(beginId string) *PrivateRestAssetTransferHistoryAPI {
	api.req.BeginId = GetPointer(beginId)
	return api
}

// string false 请求此ID之前的分页内容，传的值对应接口的id字段
func (api *PrivateRestAssetTransferHistoryAPI) EndId(endId string) *PrivateRestAssetTransferHistoryAPI {
	api.req.EndId = GetPointer(endId)
	return api
}

// string false 分页返回的结果集数量 最大为100，不填默认返回100条
func (api *PrivateRestAssetTransferHistoryAPI) Limit(limit string) *PrivateRestAssetTransferHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
