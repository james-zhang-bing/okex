package publicdata
import "fmt"

import (
	"github.com/amir-the-h/okex"
)

type (
	Instrument struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID    string               `json:"instId"`
		// uly String 否 标的指数
		Uly       string               `json:"uly,omitempty"`
		// baseCcy String 交易货币币种，如  BTC-USDT 中的 BTC
		BaseCcy   string               `json:"baseCcy,omitempty"`
		// quoteCcy String 计价货币币种，如  BTC-USDT 中的 USDT
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		// settleCcy String 盈亏结算和保证金币种，如  BTC  仅适用于 交割/永续/期权
		SettleCcy string               `json:"settleCcy,omitempty"`
		// ctValCcy String 合约面值计价币种，仅适用于 交割/永续/期权
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		// ctVal String 合约面值 仅支持 FUTURES/SWAP
		CtVal     okex.JSONFloat64     `json:"ctVal,omitempty"`
		// ctMult String 合约乘数，仅适用于 交割/永续/期权
		CtMult    okex.JSONFloat64     `json:"ctMult,omitempty"`
		// stk String 行权价格，仅适用于 期权
		Stk       okex.JSONFloat64     `json:"stk,omitempty"`
		// tickSz String 下单价格精度，如  0.0001
		TickSz    okex.JSONFloat64     `json:"tickSz,omitempty"`
		// lotSz String 下单数量精度，如 BTC-USDT-SWAP： 1
		LotSz     okex.JSONFloat64     `json:"lotSz,omitempty"`
		// minSz String 最小下单数量
		MinSz     okex.JSONFloat64     `json:"minSz,omitempty"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever     okex.JSONFloat64     `json:"lever"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType  okex.InstrumentType  `json:"instType"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category  okex.FeeCategory     `json:"category,string"`
		// optType String 期权类型， C 或 P  仅适用于 期权
		OptType   okex.OptionType      `json:"optType,omitempty"`
		// listTime String 上线日期  Unix时间戳的毫秒数格式，如  1597026383085
		ListTime  okex.JSONTime        `json:"listTime"`
		// expTime String 否 请求有效截止时间。Unix时间戳的毫秒数格式，如  1597026383085 无效的expTime
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		// ctType String 否 linear ： 正向合约 inverse ：反向合约 仅 交割/永续 有效
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		// alias String 合约日期别名 this_week ：本周  next_week ：次周quarter ：季度 next_quarter ：次季度  仅适用于 交割
		Alias     okex.AliasType       `json:"alias,omitempty"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State     okex.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		// details Array 各个账户的资产估值
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID string                    `json:"instId"`
		// px String 可选  委托价格，仅适用于 limit 、 post_only 、 fok 、 ioc 类型的订单
		Px     okex.JSONFloat64          `json:"px"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// oi String 持仓量（按 张 折算）
		Oi       okex.JSONFloat64    `json:"oi"`
		// oiCcy String 持仓量（按 币 折算）
		OiCcy    okex.JSONFloat64    `json:"oiCcy"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID          string              `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType        okex.InstrumentType `json:"instType"`
		// fundingRate String 资金费率
		FundingRate     okex.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okex.JSONFloat64    `json:"NextFundingRate"`
		// fundingTime String 资金费时间 ，Unix时间戳的毫秒数格式，如  1597026383085
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		// nextFundingTime String 下一期资金费时间 ，Unix时间戳的毫秒数格式，如  1622851200000
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// buyLmt String 最高买价
		BuyLmt   okex.JSONFloat64    `json:"buyLmt"`
		// sellLmt String 最低卖价
		SellLmt  okex.JSONFloat64    `json:"sellLmt"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// settlePx String 预估交割、行权价格
		SettlePx okex.JSONFloat64    `json:"settlePx"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// uly String 否 标的指数
		Uly      string              `json:"uly"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// delta String 期权价格对uly价格的敏感度
		Delta    okex.JSONFloat64    `json:"delta"`
		// gamma String delta对uly价格的敏感度
		Gamma    okex.JSONFloat64    `json:"gamma"`
		// vega String 权价格对隐含波动率的敏感度
		Vega     okex.JSONFloat64    `json:"vega"`
		// theta String 期权价格对剩余期限的敏感度
		Theta    okex.JSONFloat64    `json:"theta"`
		// deltaBS String 美金本位持仓仓位delta，仅适用于 期权
		DeltaBS  okex.JSONFloat64    `json:"deltaBS"`
		// gammaBS String 美金本位持仓仓位gamma，仅适用于 期权
		GammaBS  okex.JSONFloat64    `json:"gammaBS"`
		// vegaBS String 美金本位持仓仓位vega，仅适用于 期权
		VegaBS   okex.JSONFloat64    `json:"vegaBS"`
		// thetaBS String 美金本位持仓仓位theta，仅适用于 期权
		ThetaBS  okex.JSONFloat64    `json:"thetaBS"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever    okex.JSONFloat64    `json:"lever"`
		// markVol String 标记波动率
		MarkVol  okex.JSONFloat64    `json:"markVol"`
		// bidVol String bid波动率
		BidVol   okex.JSONFloat64    `json:"bidVol"`
		// askVol String ask波动率
		AskVol   okex.JSONFloat64    `json:"askVol"`
		// realVol String 已实现波动率（目前该字段暂未启用）
		RealVol  okex.JSONFloat64    `json:"realVol"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy          string           `json:"ccy"`
		// amt String 是 划转数量
		Amt          okex.JSONFloat64 `json:"amt"`
		// discountLv String 否 折算率等级 1 :第一档  2 :第二档  3 :第三档4 :第四档  5 :第五档
		DiscountLv   okex.JSONInt64   `json:"discountLv"`
		// discountInfo Array 币种折算率详情
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		// discountRate String 折算率
		DiscountRate okex.JSONInt64 `json:"discountRate"`
		// maxAmt String 最多可调整的保证金数量
		MaxAmt       okex.JSONInt64 `json:"maxAmt"`
		// minAmt String 最小申购量
		MinAmt       okex.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID    string                    `json:"instId"`
		// uly String 否 标的指数
		Uly       string                    `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType  okex.InstrumentType       `json:"instType"`
		// totalLoss String 当前 underlying 下，当期内的总穿仓亏损 以 USDT 为单位，每天16:00（UTC+8）清零
		TotalLoss okex.JSONFloat64          `json:"totalLoss"`
		// details Array 各个账户的资产估值
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy     string            `json:"ccy,omitempty"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side    okex.OrderSide    `json:"side"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		OosSide okex.PositionSide `json:"posSide"`
		// bkPx String 破产价格，与系统爆仓账号委托成交的价格。
		BkPx    okex.JSONFloat64  `json:"bkPx"`
		// sz String 是 委托数量
		Sz      okex.JSONFloat64  `json:"sz"`
		// bkLoss String 穿仓亏损数量
		BkLoss  okex.JSONFloat64  `json:"bkLoss"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID   string              `json:"instId"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// markPx String 标记价格
		MarkPx   okex.JSONFloat64    `json:"markPx"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID       string              `json:"instId"`
		// uly String 否 标的指数
		Uly          string              `json:"uly,omitempty"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType     okex.InstrumentType `json:"instType"`
		// tier String 否 查指定档位
		Tier         okex.JSONInt64      `json:"tier"`
		// minSz String 最小下单数量
		MinSz        okex.JSONFloat64    `json:"minSz"`
		// maxSz String 最大持仓量
		MaxSz        okex.JSONFloat64    `json:"maxSz"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr          okex.JSONFloat64    `json:"mmr"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr          okex.JSONFloat64    `json:"imr"`
		// optMgnFactor String 期权保证金系数 （仅适用于期权）
		OptMgnFactor okex.JSONFloat64    `json:"optMgnFactor,omitempty"`
		// quoteMaxLoan String 计价货币 最大借币量（仅适用于杠杆，且 instId 参数生效时），例如 BTC-USDT 里的 USDT最大借币量
		QuoteMaxLoan okex.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		// baseMaxLoan String 交易货币 最大借币量（仅适用于杠杆，且 instId 参数生效时），例如 BTC-USDT 里的 BTC最大借币量
		BaseMaxLoan  okex.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		// maxLever String 最高可用杠杆倍数
		MaxLever     okex.JSONFloat64    `json:"maxLever"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		// basic Array 基础利率和借币限额
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		// vip Array 专业用户
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		// regular Array 普通用户
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy   string           `json:"ccy"`
		// rate String 最新出借利率
		Rate  okex.JSONFloat64 `json:"rate"`
		// quota String 基础借币限额
		Quota okex.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level         string           `json:"level"`
		// irDiscount String 利率的折扣率
		IrDiscount    okex.JSONFloat64 `json:"irDiscount"`
		// loanQuotaCoef String 借币限额系数
		LoanQuotaCoef int              `json:"loanQuotaCoef,string"`
	}
	State struct {
		// title String 系统维护说明的标题
		Title       string        `json:"title"`
		// state String 订单状态    canceled ：撤单成功 live ：等待成交partially_filled ：部分成交 filled ：完全成交
		State       string        `json:"state"`
		// href String 系统维护详情的超级链接,若无返回值，默认值为空，如： “”
		Href        string        `json:"href"`
		// serviceType String 服务类型，  0 ：WebSocket ;5 ：交易服务； 99 ：其他（如：停止部分产品交易）
		ServiceType string        `json:"serviceType"`
		// system String 系统， unified ：交易账户
		System      string        `json:"system"`
		// scheDesc String 改期进度说明，如：由 2021-01-26T16:30:00.000Z 改期到 2021-01-28T16:30:00.000Z
		ScheDesc    string        `json:"scheDesc"`
		// begin String 否 筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
		Begin       okex.JSONTime `json:"begin"`
		// end String 否 筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
		End         okex.JSONTime `json:"end"`
	}
)
func (m *Instrument)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n标的指数: %v",str,m.Uly)
	str=fmt.Sprintf("%s\n交易货币币种: %v",str,m.BaseCcy)
	str=fmt.Sprintf("%s\n计价货币币种: %v",str,m.QuoteCcy)
	str=fmt.Sprintf("%s\n盈亏结算和保证金币种: %v",str,m.SettleCcy)
	str=fmt.Sprintf("%s\n合约面值计价币种: %v",str,m.CtValCcy)
	str=fmt.Sprintf("%s\n合约面值: %v",str,m.CtVal)
	str=fmt.Sprintf("%s\n合约乘数: %v",str,m.CtMult)
	str=fmt.Sprintf("%s\n行权价格: %v",str,m.Stk)
	str=fmt.Sprintf("%s\n下单价格精度: %v",str,m.TickSz)
	str=fmt.Sprintf("%s\n下单数量精度: %v",str,m.LotSz)
	str=fmt.Sprintf("%s\n最小下单数量: %v",str,m.MinSz)
	str=fmt.Sprintf("%s\n杠杆倍数: %v",str,m.Lever)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n订单种类: %v",str,m.Category)
	str=fmt.Sprintf("%s\n期权类型: %v",str,m.OptType)
	str=fmt.Sprintf("%s\n上线日期: %v",str,m.ListTime)
	str=fmt.Sprintf("%s\n请求有效截止时间。Unix时间戳的毫秒数格式: %v",str,m.ExpTime)
	str=fmt.Sprintf("%s\n正向合约: %v",str,m.CtType)
	str=fmt.Sprintf("%s\n合约日期别名: %v",str,m.Alias)
	str=fmt.Sprintf("%s\n订单状态: %v",str,m.State)
	return str
}
func (m *DeliveryExerciseHistory)String()string{
	var str string
	str=fmt.Sprintf("%s\n各个账户的资产估值: %v",str,m.Details)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *DeliveryExerciseHistoryDetails)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n委托价格: %v",str,m.Px)
	str=fmt.Sprintf("%s\n报价方类型（当前未生效: %v",str,m.Type)
	return str
}
func (m *OpenInterest)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n持仓量（按: %v",str,m.Oi)
	str=fmt.Sprintf("%s\n持仓量（按: %v",str,m.OiCcy)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *FundingRate)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n资金费率: %v",str,m.FundingRate)
	str=fmt.Sprintf("%s\n: %v",str,m.NextFundingRate)
	str=fmt.Sprintf("%s\n资金费时间: %v",str,m.FundingTime)
	str=fmt.Sprintf("%s\n下一期资金费时间: %v",str,m.NextFundingTime)
	return str
}
func (m *LimitPrice)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n最高买价: %v",str,m.BuyLmt)
	str=fmt.Sprintf("%s\n最低卖价: %v",str,m.SellLmt)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *EstimatedDeliveryExercisePrice)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n预估交割、行权价格: %v",str,m.SettlePx)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *OptionMarketData)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n标的指数: %v",str,m.Uly)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n期权价格对uly价格的敏感度: %v",str,m.Delta)
	str=fmt.Sprintf("%s\ndelta对uly价格的敏感度: %v",str,m.Gamma)
	str=fmt.Sprintf("%s\n权价格对隐含波动率的敏感度: %v",str,m.Vega)
	str=fmt.Sprintf("%s\n期权价格对剩余期限的敏感度: %v",str,m.Theta)
	str=fmt.Sprintf("%s\n美金本位持仓仓位delta: %v",str,m.DeltaBS)
	str=fmt.Sprintf("%s\n美金本位持仓仓位gamma: %v",str,m.GammaBS)
	str=fmt.Sprintf("%s\n美金本位持仓仓位vega: %v",str,m.VegaBS)
	str=fmt.Sprintf("%s\n美金本位持仓仓位theta: %v",str,m.ThetaBS)
	str=fmt.Sprintf("%s\n杠杆倍数: %v",str,m.Lever)
	str=fmt.Sprintf("%s\n标记波动率: %v",str,m.MarkVol)
	str=fmt.Sprintf("%s\nbid波动率: %v",str,m.BidVol)
	str=fmt.Sprintf("%s\nask波动率: %v",str,m.AskVol)
	str=fmt.Sprintf("%s\n已实现波动率（目前该字段暂未启用）: %v",str,m.RealVol)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *GetDiscountRateAndInterestFreeQuota)String()string{
	var str string
	str=fmt.Sprintf("%s\n保证金币种: %v",str,m.Ccy)
	str=fmt.Sprintf("%s\n划转数量: %v",str,m.Amt)
	str=fmt.Sprintf("%s\n折算率等级: %v",str,m.DiscountLv)
	str=fmt.Sprintf("%s\n币种折算率详情: %v",str,m.DiscountInfo)
	return str
}
func (m *DiscountInfo)String()string{
	var str string
	str=fmt.Sprintf("%s\n折算率: %v",str,m.DiscountRate)
	str=fmt.Sprintf("%s\n最多可调整的保证金数量: %v",str,m.MaxAmt)
	str=fmt.Sprintf("%s\n最小申购量: %v",str,m.MinAmt)
	return str
}
func (m *SystemTime)String()string{
	var str string
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *LiquidationOrder)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n标的指数: %v",str,m.Uly)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n当前: %v",str,m.TotalLoss)
	str=fmt.Sprintf("%s\n各个账户的资产估值: %v",str,m.Details)
	return str
}
func (m *LiquidationOrderDetail)String()string{
	var str string
	str=fmt.Sprintf("%s\n保证金币种: %v",str,m.Ccy)
	str=fmt.Sprintf("%s\n订单方向: %v",str,m.Side)
	str=fmt.Sprintf("%s\n持仓方向: %v",str,m.OosSide)
	str=fmt.Sprintf("%s\n破产价格: %v",str,m.BkPx)
	str=fmt.Sprintf("%s\n委托数量: %v",str,m.Sz)
	str=fmt.Sprintf("%s\n穿仓亏损数量: %v",str,m.BkLoss)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *MarkPrice)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n标记价格: %v",str,m.MarkPx)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *PositionTier)String()string{
	var str string
	str=fmt.Sprintf("%s\n产品ID: %v",str,m.InstID)
	str=fmt.Sprintf("%s\n标的指数: %v",str,m.Uly)
	str=fmt.Sprintf("%s\n产品类型: %v",str,m.InstType)
	str=fmt.Sprintf("%s\n查指定档位: %v",str,m.Tier)
	str=fmt.Sprintf("%s\n最小下单数量: %v",str,m.MinSz)
	str=fmt.Sprintf("%s\n最大持仓量: %v",str,m.MaxSz)
	str=fmt.Sprintf("%s\n美金层面维持保证金: %v",str,m.Mmr)
	str=fmt.Sprintf("%s\n美金层面占用保证金: %v",str,m.Imr)
	str=fmt.Sprintf("%s\n期权保证金系数: %v",str,m.OptMgnFactor)
	str=fmt.Sprintf("%s\n计价货币: %v",str,m.QuoteMaxLoan)
	str=fmt.Sprintf("%s\n交易货币: %v",str,m.BaseMaxLoan)
	str=fmt.Sprintf("%s\n最高可用杠杆倍数: %v",str,m.MaxLever)
	str=fmt.Sprintf("%s\n成交明细产生时间: %v",str,m.TS)
	return str
}
func (m *InterestRateAndLoanQuota)String()string{
	var str string
	str=fmt.Sprintf("%s\n基础利率和借币限额: %v",str,m.Basic)
	str=fmt.Sprintf("%s\n专业用户: %v",str,m.Vip)
	str=fmt.Sprintf("%s\n普通用户: %v",str,m.Regular)
	return str
}
func (m *InterestRateAndLoanBasic)String()string{
	var str string
	str=fmt.Sprintf("%s\n保证金币种: %v",str,m.Ccy)
	str=fmt.Sprintf("%s\n最新出借利率: %v",str,m.Rate)
	str=fmt.Sprintf("%s\n基础借币限额: %v",str,m.Quota)
	return str
}
func (m *InterestRateAndLoanUser)String()string{
	var str string
	str=fmt.Sprintf("%s\n当前在平台上真实交易量的用户等级: %v",str,m.Level)
	str=fmt.Sprintf("%s\n利率的折扣率: %v",str,m.IrDiscount)
	str=fmt.Sprintf("%s\n借币限额系数: %v",str,m.LoanQuotaCoef)
	return str
}
func (m *State)String()string{
	var str string
	str=fmt.Sprintf("%s\n系统维护说明的标题: %v",str,m.Title)
	str=fmt.Sprintf("%s\n订单状态: %v",str,m.State)
	str=fmt.Sprintf("%s\n系统维护详情的超级链接,若无返回值: %v",str,m.Href)
	str=fmt.Sprintf("%s\n服务类型: %v",str,m.ServiceType)
	str=fmt.Sprintf("%s\n系统: %v",str,m.System)
	str=fmt.Sprintf("%s\n改期进度说明: %v",str,m.ScheDesc)
	str=fmt.Sprintf("%s\n筛选的开始时间戳: %v",str,m.Begin)
	str=fmt.Sprintf("%s\n筛选的结束时间戳: %v",str,m.End)
	return str
}
