package myxcoinapi

type PrivateRestAPI int

const (
	// Trade
	PrivateRestTradeOrder            PrivateRestAPI = iota // POST 		下单 											支持普通订单和止盈止损订单
	PrivateRestTradeBatchOrder                             // POST		批量下单 									支持普通订单和止盈止损订单
	PrivateRestTradeCancelOrder                            // POST		取消订单 									单个订单撤单操作
	PrivateRestTradeBatchCancelOrder                       // POST		批量撤单 									撤单请求参数为List对象，每个对象定义如下，List对象最大个数为20
	PrivateRestTradeCancelAllOrder                         // POST		撤销所有订单 							批量撤所有订单
	PrivateRestTradeOpenOrders                             // GET			获取当前挂单 							该接口仅查询未成交订单列表，已成交或已撤单订单通过历史订单接口查询
	PrivateRestTradeOrderInfo                              // GET			获取订单信息 							基于订单id获取完整订单信息
	PrivateRestTradeHistoryOrders                          // GET			获取历史订单							该接口仅查询完结态订单，未完结态通过当前订单列表来查询，仅获取最近90天已完成的订单数据，返回数据按照订单创建时间倒序排序。
	PrivateRestTradeOrderOperations                        // GET			获取订单操作明细记录			 获取最近90天具体订单操作的记录，返回数据按照订单操作时间倒序排序。
	PrivateRestTradeHistoryTrades                          // GET			账户成交历史 							获取最近90天已成交订单明细数据，返回数据按照成交时间倒序排序。
	PrivateRestTradePosition                               // GET			获取交易账户仓位 					 仅适用于合约业务线
	PrivateRestTradeLeverPost                              // POST		设置杠杆倍数							设置仓位和订单的币种维度交易杠杆
	PrivateRestTradeLeverGet                               // GET			获取杠杆倍数							获取仓位和订单的币种维度交易杠杆
	PrivateRestTradeStopPosition                           // POST		设置仓位止盈止损 					该接口可以设置仓位止盈止损参数，支持全部仓位设置止盈止损和部分仓位止盈止损

	// Account
	PrivateRestAccountBalance          // GET		获取账户余额 									获取账户余额
	PrivateRestAccountTransferBalance  // GET		获取交易账户最大可转余额 			 查询交易账户指定币种的可划转余额
	PrivateRestAccountAvailableBalance // GET		获取交易账户可用余额 		 			查询交易账户指定币种的可用余额

	// Asset
	PrivateRestAssetAccountInfo     // GET 	获取账户配置信息 		该接口可以查询账户保证金模式，账户状态等信息
	PrivateRestAssetBalances        // GET 	获取资金账户余额		该接口可以查询资金账户余额信息
	PrivateRestAssetDepositAddress  // GET 	获取充值地址				仅支持有充值权限用户进行调用
	PrivateRestAssetTransfer        // POST 账户内划转申请			账户内账户类型间划转，不支持跨账户划转，额度受权限配置控制
	PrivateRestAssetTransferHistory // GET 	获取账户内划转记录	获取最近三个月资金划转记录，返回数据按照划转创建时间倒序排序
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	// Trade
	PrivateRestTradeOrder:            "/v2/trade/order",              // POST	下单 										支持普通订单和止盈止损订单
	PrivateRestTradeBatchOrder:       "/v2/trade/batchOrder",         // POST	批量下单 								支持普通订单和止盈止损订单
	PrivateRestTradeCancelOrder:      "/v1/trade/cancelOrder",        // POST	取消订单 								这里进行撤单下单操作
	PrivateRestTradeBatchCancelOrder: "/v1/trade/batchCancelOrder",   // POST	批量撤单 								撤单请求参数为List对象，每个对象定义如下，List对象最大个数为20
	PrivateRestTradeCancelAllOrder:   "/v1/trade/cancelAllOrder",     // POST	撤销所有订单 						批量撤所有订单
	PrivateRestTradeOpenOrders:       "/v2/trade/openOrders",         // GET	获取当前挂单 						该接口仅查询未成交订单列表，已成交或已撤单订单通过历史订单接口查询
	PrivateRestTradeOrderInfo:        "/v2/trade/order/info",         // GET	获取订单信息 						基于订单id获取完整订单信息
	PrivateRestTradeHistoryOrders:    "/v2/history/orders",           // GET	获取历史订单 						该接口仅查询完结态订单，未完结态通过当前订单列表来查询，仅获取最近90天已完成的订单数据，返回数据按照订单创建时间倒序排序。
	PrivateRestTradeOrderOperations:  "/v2/history/order/operations", // GET	获取订单操作明细记录		 获取最近90天具体订单操作的记录，返回数据按照订单操作时间倒序排序。
	PrivateRestTradeHistoryTrades:    "/v2/history/trades",           // GET	账户成交历史 						获取最近90天已成交订单明细数据，返回数据按照成交时间倒序排序。
	PrivateRestTradePosition:         "/v2/trade/positions",          // GET	获取交易账户仓位 				 仅适用于合约业务线
	PrivateRestTradeLeverPost:        "/v1/trade/lever",              // POST	设置杠杆倍数						 设置仓位和订单的币种维度交易杠杆
	PrivateRestTradeLeverGet:         "/v1/trade/lever",              // GET	获取杠杆倍数						 获取仓位和订单的币种维度交易杠杆
	PrivateRestTradeStopPosition:     "/v2/trade/stopPosition",       // POST	设置仓位止盈止损 				该接口可以设置仓位止盈止损参数，支持全部仓位设置止盈止损和部分仓位止盈止损

	// Account
	PrivateRestAccountBalance:          "/v1/account/balance",          // GET		获取账户余额 								获取账户余额
	PrivateRestAccountTransferBalance:  "/v1/account/transferBalance",  // GET		获取交易账户最大可转余额 		 查询交易账户指定币种的可划转余额
	PrivateRestAccountAvailableBalance: "/v1/account/availableBalance", // GET		获取交易账户可用余额 		 查询交易账户指定币种的可用余额

	// Asset
	PrivateRestAssetAccountInfo:     "/v1/asset/account/info",     // GET 	获取账户配置信息 		该接口可以查询账户保证金模式，账户状态等信息
	PrivateRestAssetBalances:        "/v1/asset/balances",         // GET 	获取资金账户余额		该接口可以查询资金账户余额信息
	PrivateRestAssetDepositAddress:  "/v1/asset/deposit/address",  // GET 	获取充值地址				仅支持有充值权限用户进行调用
	PrivateRestAssetTransfer:        "/v1/asset/transfer",         // POST 	账户内划转申请			账户内账户类型间划转，不支持跨账户划转，额度受权限配置控制
	PrivateRestAssetTransferHistory: "/v1/asset/transfer/history", // GET 	获取账户内划转记录	获取最近三个月资金划转记录，返回数据按照划转创建时间倒序排序
}
