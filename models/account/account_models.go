package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		// totalEq String 美金层面权益
		TotalEq     okex.JSONFloat64  `json:"totalEq"`
		// isoEq String 美金层面逐仓仓位权益 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoEq       okex.JSONFloat64  `json:"isoEq"`
		// adjEq String 美金层面有效保证金 适用于 跨币种保证金模式 和 组合保证金模式
		AdjEq       okex.JSONFloat64  `json:"adjEq,omitempty"`
		// ordFroz String 美金层面全仓挂单占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		OrdFroz     okex.JSONFloat64  `json:"ordFroz,omitempty"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr         okex.JSONFloat64  `json:"imr,omitempty"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr         okex.JSONFloat64  `json:"mmr,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio    okex.JSONFloat64  `json:"mgnRatio,omitempty"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okex.JSONFloat64  `json:"notionalUsd,omitempty"`
		// details Array 各个账户的资产估值
		Details     []*BalanceDetails `json:"details,omitempty"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime       okex.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy           string           `json:"ccy"`
		// eq String 币种总权益
		Eq            okex.JSONFloat64 `json:"eq"`
		// cashBal String 币种余额
		CashBal       okex.JSONFloat64 `json:"cashBal"`
		// isoEq String 美金层面逐仓仓位权益 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoEq         okex.JSONFloat64 `json:"isoEq,omitempty"`
		// availEq String 可用保证金 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		AvailEq       okex.JSONFloat64 `json:"availEq,omitempty"`
		// disEq String 美金层面币种折算权益
		DisEq         okex.JSONFloat64 `json:"disEq"`
		// availBal String 可用余额
		AvailBal      okex.JSONFloat64 `json:"availBal"`
		// frozenBal String 冻结（不可用）
		FrozenBal     okex.JSONFloat64 `json:"frozenBal"`
		// ordFrozen String 挂单冻结数量
		OrdFrozen     okex.JSONFloat64 `json:"ordFrozen"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab          okex.JSONFloat64 `json:"liab,omitempty"`
		// upl String 未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		Upl           okex.JSONFloat64 `json:"upl,omitempty"`
		UplLib        okex.JSONFloat64 `json:"uplLib,omitempty"`
		// crossLiab String 币种全仓负债额 适用于 跨币种保证金模式 和 组合保证金模式
		CrossLiab     okex.JSONFloat64 `json:"crossLiab,omitempty"`
		// isoLiab String 币种逐仓负债额 适用于 跨币种保证金模式 和 组合保证金模式
		IsoLiab       okex.JSONFloat64 `json:"isoLiab,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio      okex.JSONFloat64 `json:"mgnRatio,omitempty"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest      okex.JSONFloat64 `json:"interest,omitempty"`
		// twap String 当前负债币种触发系统自动换币的风险 0、1、2、3、4、5其中之一，数字越大代表您的负债币种触发自动换币概率越高 适用于 跨币种保证金模式 和 组合保证金模式 
		Twap          okex.JSONFloat64 `json:"twap,omitempty"`
		// maxLoan String 币种最大可借 适用于 跨币种保证金模式 和 组合保证金模式
		MaxLoan       okex.JSONFloat64 `json:"maxLoan,omitempty"`
		// eqUsd String 币种权益美金价值
		EqUsd         okex.JSONFloat64 `json:"eqUsd"`
		// notionalLever String 币种杠杆倍数 适用于 单币种保证金模式
		NotionalLever okex.JSONFloat64 `json:"notionalLever,omitempty"`
		// stgyEq String 策略权益
		StgyEq        okex.JSONFloat64 `json:"stgyEq"`
		// isoUpl String 逐仓未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		IsoUpl        okex.JSONFloat64 `json:"isoUpl,omitempty"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime         okex.JSONTime    `json:"uTime"`
	}
	Position struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID      string              `json:"instId"`
		// posCcy String 仓位资产币种，仅适用于 币币杠杆 仓位
		PosCcy      string              `json:"posCcy,omitempty"`
		// liabCcy String 负债币种，仅适用于 币币杠杆
		LiabCcy     string              `json:"liabCcy,omitempty"`
		// optVal String 期权市值，仅适用于 期权
		OptVal      string              `json:"optVal,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy         string              `json:"ccy"`
		// posId String 否 持仓ID 支持多个 posId 查询（不超过20个），半角逗号分割
		PosID       string              `json:"posId"`
		// tradeId String 最新成交ID
		TradeID     string              `json:"tradeId"`
		// pos String 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为 0 的仓位
		Pos         okex.JSONFloat64    `json:"pos"`
		// availPos String 可平仓数量，适用于币币杠杆 , 交割/永续 （开平仓模式）， 期权 （交易账户及保证金账户逐仓）。
		AvailPos    okex.JSONFloat64    `json:"availPos,omitempty"`
		// avgPx String 成交均价，如果成交数量为0，该字段也为 0
		AvgPx       okex.JSONFloat64    `json:"avgPx"`
		// upl String 未实现盈亏 适用于 单币种保证金模式 和 跨币种保证金模式 和 组合保证金模式
		Upl         okex.JSONFloat64    `json:"upl"`
		// uplRatio String 未实现收益率
		UplRatio    okex.JSONFloat64    `json:"uplRatio"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever       okex.JSONFloat64    `json:"lever"`
		// liqPx String 预估强平价  不适用于 期权
		LiqPx       okex.JSONFloat64    `json:"liqPx,omitempty"`
		// imr String 美金层面占用保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Imr         okex.JSONFloat64    `json:"imr,omitempty"`
		// margin String 保证金余额，可增减，仅适用于 逐仓
		Margin      okex.JSONFloat64    `json:"margin,omitempty"`
		// mgnRatio String 美金层面保证金率 适用于 跨币种保证金模式  和 组合保证金模式
		MgnRatio    okex.JSONFloat64    `json:"mgnRatio"`
		// mmr String 美金层面维持保证金 适用于 跨币种保证金模式 和 组合保证金模式
		Mmr         okex.JSONFloat64    `json:"mmr"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab        okex.JSONFloat64    `json:"liab,omitempty"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest    okex.JSONFloat64    `json:"interest"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`
		// adl String 信号区 分为5档，从1到5，数字越小代表adl强度越弱
		ADL         okex.JSONFloat64    `json:"adl"`
		// last String 最新成交价
		Last        okex.JSONFloat64    `json:"last"`
		// deltaBS String 美金本位持仓仓位delta，仅适用于 期权
		DeltaBS     okex.JSONFloat64    `json:"deltaBS"`
		// deltaPA String 币本位持仓仓位delta，仅适用于 期权
		DeltaPA     okex.JSONFloat64    `json:"deltaPA"`
		// gammaBS String 美金本位持仓仓位gamma，仅适用于 期权
		GammaBS     okex.JSONFloat64    `json:"gammaBS"`
		// gammaPA String 币本位持仓仓位gamma，仅适用于 期权
		GammaPA     okex.JSONFloat64    `json:"gammaPA"`
		// thetaBS String 美金本位持仓仓位theta，仅适用于 期权
		ThetaBS     okex.JSONFloat64    `json:"thetaBS"`
		// thetaPA String 币本位持仓仓位theta，仅适用于 期权
		ThetaPA     okex.JSONFloat64    `json:"thetaPA"`
		// vegaBS String 美金本位持仓仓位vega，仅适用于 期权
		VegaBS      okex.JSONFloat64    `json:"vegaBS"`
		// vegaPA String 币本位持仓仓位vega，仅适用于 期权
		VegaPA      okex.JSONFloat64    `json:"vegaPA"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide     okex.PositionSide   `json:"posSide"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType    okex.InstrumentType `json:"instType"`
		// cTime String 订单创建时间，Unix时间戳的毫秒数格式， 如 ： 1597026383085
		CTime       okex.JSONTime       `json:"cTime"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		// eventType String 事件类型，枚举值： snapshot ：首推快照  delivered ：交割      exercised ：行权  transferred ：划转 filled ：成交      liquidation ：强平  claw_back ：穿仓补偿   adl ：ADL自动减仓      funding_fee ：资金费  adjust_margin ：调整保证金      set_leverage ：设置杠杆  interest_deduction ：扣息
		EventType okex.EventType    `json:"eventType"`
		// pTime String 持仓信息的推送时间，Unix时间戳的毫秒数格式，如  1597026383085
		PTime     okex.JSONTime     `json:"pTime"`
		// uTime String 订单状态更新时间，Unix时间戳的毫秒数格式，如： 1597026383085
		UTime     okex.JSONTime     `json:"uTime"`
		// posData Array 持仓详细信息
		PosData   []*Position       `json:"posData"`
		// balData Array 币种资产信息
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		// adjEq String 美金层面有效保证金 适用于 跨币种保证金模式 和 组合保证金模式
		AdjEq   okex.JSONFloat64                     `json:"adjEq,omitempty"`
		// balData Array 币种资产信息
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		// posData Array 持仓详细信息
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy   string           `json:"ccy"`
		// eq String 币种总权益
		Eq    okex.JSONFloat64 `json:"eq"`
		// disEq String 美金层面币种折算权益
		DisEq okex.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID      string              `json:"instId"`
		// posCcy String 仓位资产币种，仅适用于 币币杠杆 仓位
		PosCcy      string              `json:"posCcy,omitempty"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy         string              `json:"ccy"`
		// notionalCcy String 以 币 为单位的持仓数量
		NotionalCcy okex.JSONFloat64    `json:"notionalCcy"`
		// pos String 持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为 0 的仓位
		Pos         okex.JSONFloat64    `json:"pos"`
		// notionalUsd String 以美金价值为单位的持仓数量，即仓位美金价值 适用于 跨币种保证金模式 和 组合保证金模式
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide     okex.PositionSide   `json:"posSide"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType    okex.InstrumentType `json:"instType"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy       string              `json:"ccy"`
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID    string              `json:"instId"`
		// notes String 备注 仅适用于 资金划转 ，不是 资金划转 时，返回 ""
		Notes     string              `json:"notes"`
		// billId String 账单 ID
		BillID    string              `json:"billId"`
		// ordId String 订单ID ordId或clOrdId至少填一个 ordId重复
		OrdID     string              `json:"ordId"`
		// balChg String 账户层面的余额变动数量
		BalChg    okex.JSONFloat64    `json:"balChg"`
		// posBalChg String 仓位层面的余额变动数量
		PosBalChg okex.JSONFloat64    `json:"posBalChg"`
		// bal String 余额
		Bal       okex.JSONFloat64    `json:"bal"`
		// posBal String 仓位层面的余额数量
		PosBal    okex.JSONFloat64    `json:"posBal"`
		// sz String 是 委托数量
		Sz        okex.JSONFloat64    `json:"sz"`
		// pnl String 收益
		Pnl       okex.JSONFloat64    `json:"pnl"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为 负数 。如：  -0.01
		Fee       okex.JSONFloat64    `json:"fee"`
		// from String 是 转出账户 6 ：资金账户  18 ：交易账户 from和to不可相同
		From      okex.AccountType    `json:"from,string"`
		// to String 是 转入账户 6 ：资金账户  18 ：交易账户
		To        okex.AccountType    `json:"to,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type      okex.BillType       `json:"type,string"`
		// subType String 否 账单子类型 1 ：买入  2 ：卖出  3 ：开多4 ：开空  5 ：平多  6 ：平空  9 ：市场借币扣息  11 ：转入12 ：转出  14 ：尊享借币扣息  160 ：手动追加保证金  161 ：手动减少保证金162 ：自动追加保证金  114 ：自动换币买入  115 ：自动换币卖出  118 ：系统换币转入119 ：系统换币转出  100 ：强减平多  101 ：强减平空  102 ：强减买入103 ：强减卖出  104 ：强平平多  105 ：强平平空  106 ：强平买入107 ：强平卖出  110 ：强平换币转入  111 ：强平换币转出  125 ：自动减仓平多126 ：自动减仓平空  127 ：自动减仓买入  128 ：自动减仓卖出  131 ：对冲买入132 ：对冲卖出  170 ：到期行权  171 ：到期被行权  172 ：到期作废112 ：交割平多  113 ：交割平空  117 ：交割/期权穿仓补偿  173 ：资金费支出174 ：资金费收入  200 :系统转入  201 :手动转入  202 :系统转出203 :手动转出  204 : 大宗交易买  205 : 大宗交易卖  206 : 大宗交易开多207 : 大宗交易开空  208 : 大宗交易平多  209 : 大宗交易平空  224 : 还债转入225 : 还债转出
		SubType   okex.BillSubType    `json:"subType,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level      string            `json:"level"`
		// levelTmp String 特约用户的临时体验用户等级，例如  lv3
		LevelTmp   string            `json:"levelTmp"`
		// acctLv String 账户层级 1 ：简单交易模式， 2 ：单币种保证金模式， 3 ：跨币种保证金模式 ， 4 ：组合保证金模式
		AcctLv     string            `json:"acctLv"`
		// autoLoan Boolean 是否自动借币 true ：自动借币  false ：非自动借币
		AutoLoan   bool              `json:"autoLoan"`
		// uid String 账户ID，账户uid和app上的一致
		UID        string            `json:"uid"`
		// greeksType String 当前希腊字母展示方式 PA ：币本位  BS ：美元本位
		GreeksType okex.GreekType    `json:"greeksType"`
		// posMode String 持仓方式 long_short_mode ：双向持仓net_mode ：单向持仓 仅适用 交割/永续
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		// posMode String 持仓方式 long_short_mode ：双向持仓net_mode ：单向持仓 仅适用 交割/永续
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string            `json:"instId"`
		// lever String 杠杆倍数，0.01到125之间的数值，仅适用于  币币杠杆/交割/永续
		Lever   okex.JSONFloat64  `json:"lever"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okex.MarginMode   `json:"mgnMode"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string           `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy     string           `json:"ccy"`
		// maxBuy String 币币/币币杠杆 ：最大可买的交易币数量 单币种保证金模式下的全仓杠杆订单，为交易币数量 交割 / 永续 / 期权 ：最大可开多的合约张数
		MaxBuy  okex.JSONFloat64 `json:"maxBuy"`
		// maxSell String 币币/币币杠杆 ：最大可卖的计价币数量 单币种保证金模式下的全仓杠杆订单，为交易币数量 交割 / 永续 / 期权 ：最大可开空的合约张数
		MaxSell okex.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID    string           `json:"instId"`
		// availBuy String 最大买入可用数量
		AvailBuy  okex.JSONFloat64 `json:"availBuy"`
		// availSell String 最大卖出可用数量
		AvailSell okex.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string            `json:"instId"`
		// amt String 是 划转数量
		Amt     okex.JSONFloat64  `json:"amt"`
		// posSide String 可选 持仓方向  在双向持仓模式下必填，且仅可选择  long  或  short 。 仅适用交割、永续。
		PosSide okex.PositionSide `json:"posSide,string"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string           `json:"instId"`
		// mgnCcy String 可选 保证金币种，如  BTC 币币杠杆单币种全仓情况下必须指定保证金币种
		MgnCcy  string           `json:"mgnCcy"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy     string           `json:"ccy"`
		// maxLoan String 币种最大可借 适用于 跨币种保证金模式 和 组合保证金模式
		MaxLoan okex.JSONFloat64 `json:"maxLoan"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode okex.MarginMode  `json:"mgnMode"`
		// side String 是 订单方向   buy ：买，  sell ：卖
		Side    okex.OrderSide   `json:"side,string"`
	}
	Fee struct {
		// level String 当前在平台上真实交易量的用户等级，例如  lv1
		Level    string              `json:"level"`
		// taker String USDT&USDⓈ&Crypto 交易区的吃单手续费率，永续和交割合约时，为币本位合约费率
		Taker    okex.JSONFloat64    `json:"taker"`
		// maker String USDT&USDⓈ&Crypto 交易区挂单手续费率，永续和交割合约时，为币本位合约费率
		Maker    okex.JSONFloat64    `json:"maker"`
		// delivery String 交割手续费率
		Delivery okex.JSONFloat64    `json:"delivery,omitempty"`
		// exercise String 行权手续费率
		Exercise okex.JSONFloat64    `json:"exercise,omitempty"`
		// category String 订单种类  normal ：普通委托 twap ：TWAP自动换币adl ：ADL自动减仓   full_liquidation ：强制平仓 partial_liquidation ：强制减仓    delivery ：交割   ddh ：对冲减仓类型订单
		Category okex.FeeCategory    `json:"category,string"`
		// instType String 产品类型 SPOT ：币币 MARGIN ：币币杠杆 SWAP ：永续合约  FUTURES ：交割合约    OPTION ：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType okex.InstrumentType `json:"instType"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		// instId String 是 产品ID，如  BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID       string           `json:"instId"`
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy          string           `json:"ccy"`
		// interest String 计息 适用于 跨币种保证金模式 和 组合保证金模式
		Interest     okex.JSONFloat64 `json:"interest"`
		// interestRate String 计息利率(小时)
		InterestRate okex.JSONFloat64 `json:"interestRate"`
		// liab String 币种负债额 适用于 跨币种保证金模式 和 组合保证金模式
		Liab         okex.JSONFloat64 `json:"liab"`
		// mgnMode String 是 保证金模式   cross ：全仓 ；  isolated ：逐仓
		MgnMode      okex.MarginMode  `json:"mgnMode"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
		TS           okex.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy          string           `json:"ccy"`
		// interestRate String 计息利率(小时)
		InterestRate okex.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		// greeksType String 当前希腊字母展示方式 PA ：币本位  BS ：美元本位
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		// ccy String 否 保证金币种，仅适用于 单币种保证金模式 下的 全仓杠杆 订单
		Ccy   string           `json:"ccy"`
		// maxWd String 币种单笔最大提币量
		MaxWd okex.JSONFloat64 `json:"maxWd"`
	}
)
