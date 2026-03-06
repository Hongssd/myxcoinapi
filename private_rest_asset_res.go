package myxcoinapi

type PrivateRestAssetDepositAddressResRow struct {
	Pid            string `json:"pid"`            // 账户对应唯一ID
	AccountName    string `json:"accountName"`    // 账户名称
	Uid            string `json:"uid"`            // 操作员唯一ID
	Cid            string `json:"cid"`            // 操作主体唯一ID
	Currency       string `json:"currency"`       // 币种，如 BTC
	ChainType      string `json:"chainType"`      // 链的类型
	AddressDeposit string `json:"addressDeposit"` // 充值地址
	Memo           string `json:"memo"`           // 地址备注，例如XRP等币种，此字段必填
}

type PrivateRestAssetDepositAddressRes []PrivateRestAssetDepositAddressResRow

type PrivateRestAssetAccountInfoRes struct {
	Pid           string `json:"pid"`           // 账户对应唯一ID
	Uid           string `json:"uid"`           // 操作员唯一ID
	Cid           string `json:"cid"`           // 操作主体唯一ID
	AccountStatus string `json:"accountStatus"` // 账户状态
	AccountType   string `json:"accountType"`   // 账户类型
	AutoSubscribe bool   `json:"autoSubscribe"` // 现金宝开关
	AccountName   string `json:"accountName"`   // 账户名称
	MarginMode    string `json:"marginMode"`    // 账户保证金模式
	CanWithdraw   bool   `json:"canWithdraw"`   // 可以提现
	CanDeposit    bool   `json:"canDeposit"`    // 可以充值
	CreateTime    string `json:"createTime"`    // 创建时间，Unix 时间戳的毫秒数格式，如 1732158178000
}

type PrivateRestAssetBalancesResRow struct {
	Currency     string `json:"currency"`     // 币种，如 BTC
	AccountType  string `json:"accountType"`  // 账户类型，如 funding
	Balance      string `json:"balance"`      // 可用资产
	Freeze       string `json:"freeze"`       // 冻结资产
	Equity       string `json:"equity"`       // 总资产
	WithdrawAble string `json:"withdrawAble"` // 可划转资产
	AccountName  string `json:"accountName"`  // 账户名称
	Pid          string `json:"pid"`          // 账户唯一 ID
	Uid          string `json:"uid"`          // 操作员唯一 ID
	Cid          string `json:"cid"`          // 操作主体唯一 ID
}
type PrivateRestAssetBalancesRes []PrivateRestAssetBalancesResRow

type PrivateRestAssetTransferRes bool

type PrivateRestAssetTransferHistoryResRow struct {
	Id               string `json:"id"`               // 数据唯一ID
	ClientTransferId string `json:"clientTransferId"` // 客户划转单ID
	FromAccountType  string `json:"fromAccountType"`  // 转出账户，如 funding
	ToAccountType    string `json:"toAccountType"`    // 转入账户，如 trading
	Amount           string `json:"amount"`           // 划转金额
	Currency         string `json:"currency"`         // 币种，如 BTC
	TransferType     string `json:"transferType"`     // 转换类型，如 account
	Status           string `json:"status"`           // 划转状态
	AccountName      string `json:"accountName"`      // 账户名称
	Pid              string `json:"pid"`              // 账户唯一ID
	Uid              string `json:"uid"`              // 操作员唯一ID
	Cid              string `json:"cid"`              // 操作主体唯一ID
	CreateTime       string `json:"createTime"`       // 划转时间 Unix时间戳的毫秒数格式，如 1732182494000
}
type PrivateRestAssetTransferHistoryRes []PrivateRestAssetTransferHistoryResRow