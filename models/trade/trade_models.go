package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string           `json:"clOrdId"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag     string           `json:"tag"`
		// sMsg String 事件执行失败时的msg
		SMsg    string           `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode   okex.JSONInt64   `json:"sCode"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID   okex.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID   string           `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string           `json:"clOrdId"`
		// sMsg String 事件执行失败时的msg
		SMsg    string           `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID   string           `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID string           `json:"clOrdId"`
		// reqId String 否 用户自定义修改事件ID字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
		ReqID   string           `json:"reqId"`
		// sMsg String 事件执行失败时的msg
		SMsg    string           `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string            `json:"instId"`
		// posSide String 可选 持仓方向 在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID      string              `json:"instId"`
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy         string              `json:"ccy"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID       string              `json:"ordId"`
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID     string              `json:"clOrdId"`
		// tradeId String 最新成交ID
		TradeID     string              `json:"tradeId"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag         string              `json:"tag"`
		// category String 订单种类 normal：普通委托twap：TWAP自动换币adl：ADL自动减仓 full_liquidation：强制平仓partial_liquidation：强制减仓  delivery：交割 ddh：对冲减仓类型订单
		Category    string              `json:"category"`
		// feeCcy String 交易手续费币种
		FeeCcy      string              `json:"feeCcy"`
		// rebateCcy String 返佣金币种
		RebateCcy   string              `json:"rebateCcy"`
		// px String 可选               委托价格，仅适用于limit、post_only、fok、ioc类型的订单
		Px          okex.JSONFloat64    `json:"px"`
		// sz String 是 委托数量
		Sz          okex.JSONInt64      `json:"sz"`
		// pnl String 收益
		Pnl         okex.JSONFloat64    `json:"pnl"`
		// accFillSz String 累计成交数量
		AccFillSz   okex.JSONInt64      `json:"accFillSz"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为0
		FillPx      okex.JSONFloat64    `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz      okex.JSONInt64      `json:"fillSz"`
		// fillTime String 最新成交时间
		FillTime    okex.JSONFloat64    `json:"fillTime"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为0
		AvgPx       okex.JSONFloat64    `json:"avgPx"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
		Lever       okex.JSONFloat64    `json:"lever"`
		// tpTriggerPx String 止盈触发价
		TpTriggerPx okex.JSONFloat64    `json:"tpTriggerPx"`
		// tpOrdPx String 止盈委托价
		TpOrdPx     okex.JSONFloat64    `json:"tpOrdPx"`
		// slTriggerPx String 止损触发价
		SlTriggerPx okex.JSONFloat64    `json:"slTriggerPx"`
		// slOrdPx String 止损委托价
		SlOrdPx     okex.JSONFloat64    `json:"slOrdPx"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为负数。如： -0.01
		Fee         okex.JSONFloat64    `json:"fee"`
		// rebate String 返佣金额，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为正数，如：0.01
		Rebate      okex.JSONFloat64    `json:"rebate"`
		// state String 订单状态  canceled：撤单成功live：等待成交partially_filled：部分成交filled：完全成交
		State       okex.OrderState     `json:"state"`
		// tdMode String 是 交易模式保证金模式：isolated：逐仓 ；cross：全仓              非保证金模式：cash：非保证金
		TdMode      okex.TradeMode      `json:"tdMode"`
		// posSide String 可选 持仓方向 在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
		PosSide     okex.PositionSide   `json:"posSide"`
		// side String 是 订单方向 buy：买， sell：卖
		Side        okex.OrderSide      `json:"side"`
		// ordType String 是 订单类型 market：市价单limit：限价单              post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余              optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType     okex.OrderType      `json:"ordType"`
		// instType String 产品类型SPOT：币币MARGIN：币币杠杆SWAP：永续合约 FUTURES：交割合约  OPTION：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType    okex.InstrumentType `json:"instType"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于币币市价订单base_ccy: 交易货币              ；quote_ccy：计价货币买单默认quote_ccy， 卖单默认base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy      okex.QuantityType   `json:"tgtCcy"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如：1597026383085
		UTime       okex.JSONTime       `json:"uTime"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ：1597026383085
		CTime       okex.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID    string              `json:"ordId"`
		// tradeId String 最新成交ID
		TradeID  string              `json:"tradeId"`
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID  string              `json:"clOrdId"`
		// billId String 账单 ID
		BillID   string              `json:"billId"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag      okex.JSONFloat64    `json:"tag"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为0
		FillPx   okex.JSONFloat64    `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz   okex.JSONFloat64    `json:"fillSz"`
		// feeCcy String 交易手续费币种
		FeeCcy   okex.JSONFloat64    `json:"feeCcy"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为负数。如： -0.01
		Fee      okex.JSONFloat64    `json:"fee"`
		// instType String 产品类型SPOT：币币MARGIN：币币杠杆SWAP：永续合约 FUTURES：交割合约  OPTION：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// side String 是 订单方向 buy：买， sell：卖
		Side     okex.OrderSide      `json:"side"`
		// posSide String 可选 持仓方向 在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
		PosSide  okex.PositionSide   `json:"posSide"`
		// execType String 流动性方向 T：taker M：maker
		ExecType okex.OrderFlowType  `json:"execType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		// algoId String 策略委托单ID
		AlgoID string         `json:"algoId"`
		// sMsg String 事件执行失败时的msg
		SMsg   string         `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		// algoId String 策略委托单ID
		AlgoID string         `json:"algoId"`
		// sMsg String 事件执行失败时的msg
		SMsg   string         `json:"sMsg"`
		// sCode String 事件执行结果的code，0代表成功
		SCode  okex.JSONInt64 `json:"sCode"`
	}
	AlgoOrder struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID       string              `json:"instId"`
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy          string              `json:"ccy"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID        string              `json:"ordId"`
		// algoId String 策略委托单ID
		AlgoID       string              `json:"algoId"`
		// clOrdId String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 clOrdId重复
		ClOrdID      string              `json:"clOrdId"`
		// tradeId String 最新成交ID
		TradeID      string              `json:"tradeId"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag          string              `json:"tag"`
		// category String 订单种类 normal：普通委托twap：TWAP自动换币adl：ADL自动减仓 full_liquidation：强制平仓partial_liquidation：强制减仓  delivery：交割 ddh：对冲减仓类型订单
		Category     string              `json:"category"`
		// feeCcy String 交易手续费币种
		FeeCcy       string              `json:"feeCcy"`
		// rebateCcy String 返佣金币种
		RebateCcy    string              `json:"rebateCcy"`
		// timeInterval String 是 下单间隔
		TimeInterval string              `json:"timeInterval"`
		// px String 可选               委托价格，仅适用于limit、post_only、fok、ioc类型的订单
		Px           okex.JSONFloat64    `json:"px"`
		// pxVar String 可选 挂单价距离盘口的比例 pxVar和pxSpread只能传入一个
		PxVar        okex.JSONFloat64    `json:"pxVar"`
		// pxSpread String 可选 挂单价距离盘口的价距
		PxSpread     okex.JSONFloat64    `json:"pxSpread"`
		// pxLimit String 是 挂单限制价
		PxLimit      okex.JSONFloat64    `json:"pxLimit"`
		// sz String 是 委托数量
		Sz           okex.JSONInt64      `json:"sz"`
		// szLimit String 是 单笔数量
		SzLimit      okex.JSONInt64      `json:"szLimit"`
		// actualSz String 实际委托量
		ActualSz     okex.JSONFloat64    `json:"actualSz"`
		// actualPx String 实际委托价
		ActualPx     okex.JSONFloat64    `json:"actualPx"`
		// pnl String 收益
		Pnl          okex.JSONFloat64    `json:"pnl"`
		// accFillSz String 累计成交数量
		AccFillSz    okex.JSONInt64      `json:"accFillSz"`
		// fillPx String 最新成交价格，如果成交数量为0，该字段也为0
		FillPx       okex.JSONFloat64    `json:"fillPx"`
		// fillSz String 最新成交数量
		FillSz       okex.JSONInt64      `json:"fillSz"`
		// fillTime String 最新成交时间
		FillTime     okex.JSONFloat64    `json:"fillTime"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为0
		AvgPx        okex.JSONFloat64    `json:"avgPx"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
		Lever        okex.JSONFloat64    `json:"lever"`
		// tpTriggerPx String 止盈触发价
		TpTriggerPx  okex.JSONFloat64    `json:"tpTriggerPx"`
		// tpOrdPx String 止盈委托价
		TpOrdPx      okex.JSONFloat64    `json:"tpOrdPx"`
		// slTriggerPx String 止损触发价
		SlTriggerPx  okex.JSONFloat64    `json:"slTriggerPx"`
		// slOrdPx String 止损委托价
		SlOrdPx      okex.JSONFloat64    `json:"slOrdPx"`
		// ordPx String 计划委托委托价格
		OrdPx        okex.JSONFloat64    `json:"ordPx"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为负数。如： -0.01
		Fee          okex.JSONFloat64    `json:"fee"`
		// rebate String 返佣金额，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为正数，如：0.01
		Rebate       okex.JSONFloat64    `json:"rebate"`
		// state String 订单状态  canceled：撤单成功live：等待成交partially_filled：部分成交filled：完全成交
		State        okex.OrderState     `json:"state"`
		// tdMode String 是 交易模式保证金模式：isolated：逐仓 ；cross：全仓              非保证金模式：cash：非保证金
		TdMode       okex.TradeMode      `json:"tdMode"`
		// actualSide String 实际触发方向， tp：止盈 sl： 止损
		ActualSide   okex.PositionSide   `json:"actualSide"`
		// posSide String 可选 持仓方向 在双向持仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
		PosSide      okex.PositionSide   `json:"posSide"`
		// side String 是 订单方向 buy：买， sell：卖
		Side         okex.OrderSide      `json:"side"`
		// ordType String 是 订单类型 market：市价单limit：限价单              post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余              optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） 逐仓自主划转保证金模式不支持ordType为iceberg、twap的策略委托单
		OrdType      okex.AlgoOrderType  `json:"ordType"`
		// instType String 产品类型SPOT：币币MARGIN：币币杠杆SWAP：永续合约 FUTURES：交割合约  OPTION：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType     okex.InstrumentType `json:"instType"`
		// tgtCcy String 否 市价单委托数量的类型，仅适用于币币市价订单base_ccy: 交易货币              ；quote_ccy：计价货币买单默认quote_ccy， 卖单默认base_ccy 计划委托不支持使用tgtCcy参数
		TgtCcy       okex.QuantityType   `json:"tgtCcy"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ：1597026383085
		CTime        okex.JSONTime       `json:"cTime"`
		// triggerTime String 策略委托触发时间，Unix时间戳的毫秒数格式，如 1597026383085
		TriggerTime  okex.JSONTime       `json:"triggerTime"`
	}
)
