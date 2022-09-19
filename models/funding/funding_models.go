package funding

import "github.com/amir-the-h/okex"

type (
	Currency struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy         string `json:"ccy"`
		// name String 币种中文名称，不显示则无对应名称
		Name        string `json:"name"`
		// chain String 币种链信息有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20，USDT-Omni多个链
		Chain       string `json:"chain"`
		// minWd String 币种单笔最小提币量
		MinWd       string `json:"minWd"`
		// minFee String 最小提币手续费数量
		MinFee      string `json:"minFee"`
		// maxFee String 最大提币手续费数量
		MaxFee      string `json:"maxFee"`
		// canDep Boolean 是否可充值，false表示不可链上充值，true表示可以链上充值
		CanDep      bool   `json:"canDep"`
		// canWd Boolean 是否可提币，false表示不可链上提币，true表示可以链上提币
		CanWd       bool   `json:"canWd"`
		// canInternal Boolean 是否可内部转账，false表示不可内部转账，true表示可以内部转账
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy       string `json:"ccy"`
		// bal String 余额
		Bal       string `json:"bal"`
		// frozenBal String 冻结（不可用）
		FrozenBal string `json:"frozenBal"`
		// availBal String 可用余额
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		// transId String 划转 ID
		TransID string           `json:"transId"`
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy     string           `json:"ccy"`
		// amt String 是 划转数量
		Amt     okex.JSONFloat64 `json:"amt"`
		// from String 是 转出账户6：资金账户 18：交易账户 from和to不可相同
		From    okex.AccountType `json:"from,string"`
		// to String 是 转入账户6：资金账户 18：交易账户
		To      okex.AccountType `json:"to,string"`
	}
	Bill struct {
		// billId String 账单 ID
		BillID string           `json:"billId"`
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy    string           `json:"ccy"`
		// bal String 余额
		Bal    okex.JSONFloat64 `json:"bal"`
		// balChg String 账户层面的余额变动数量
		BalChg okex.JSONFloat64 `json:"balChg"`
		// type String 报价方类型（当前未生效，将返回 "" ）
		Type   okex.BillType    `json:"type,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS     okex.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		// addr String 充值地址
		Addr     string           `json:"addr"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag      string           `json:"tag,omitempty"`
		// memo String 部分币种充值需要 memo，若不需要则不返回此字段
		Memo     string           `json:"memo,omitempty"`
		// pmtId String 部分币种充值需要此字段（payment_id），若不需要则不返回此字段
		PmtID    string           `json:"pmtId,omitempty"`
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy      string           `json:"ccy"`
		// chain String 币种链信息有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20，USDT-Omni多个链
		Chain    string           `json:"chain"`
		// ctAddr String 合约地址后6位
		CtAddr   string           `json:"ctAddr"`
		// selected Boolean 该地址是否为页面选中的地址
		Selected bool             `json:"selected"`
		// to String 是 转入账户6：资金账户 18：交易账户
		To       okex.AccountType `json:"to,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS       okex.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy   string            `json:"ccy"`
		// chain String 币种链信息有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20，USDT-Omni多个链
		Chain string            `json:"chain"`
		// txId String 否 区块转账哈希记录
		TxID  string            `json:"txId"`
		// from String 是 转出账户6：资金账户 18：交易账户 from和to不可相同
		From  string            `json:"from"`
		// to String 是 转入账户6：资金账户 18：交易账户
		To    string            `json:"to"`
		// depId String 否 充值记录 ID
		DepId string            `json:"depId"`
		// amt String 是 划转数量
		Amt   okex.JSONFloat64  `json:"amt"`
		// state String 订单状态  canceled：撤单成功live：等待成交partially_filled：部分成交filled：完全成交
		State okex.DepositState `json:"state,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS    okex.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy   string           `json:"ccy"`
		// chain String 币种链信息有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20，USDT-Omni多个链
		Chain string           `json:"chain"`
		// wdId String 提币申请ID
		WdID  okex.JSONInt64   `json:"wdId"`
		// amt String 是 划转数量
		Amt   okex.JSONFloat64 `json:"amt"`
	}
	WithdrawalHistory struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy   string               `json:"ccy"`
		// chain String 币种链信息有的币种下有多个链，必须要做区分，如USDT下有USDT-ERC20，USDT-TRC20，USDT-Omni多个链
		Chain string               `json:"chain"`
		// txId String 否 区块转账哈希记录
		TxID  string               `json:"txId"`
		// from String 是 转出账户6：资金账户 18：交易账户 from和to不可相同
		From  string               `json:"from"`
		// to String 是 转入账户6：资金账户 18：交易账户
		To    string               `json:"to"`
		// tag String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
		Tag   string               `json:"tag,omitempty"`
		// pmtId String 部分币种充值需要此字段（payment_id），若不需要则不返回此字段
		PmtID string               `json:"pmtId,omitempty"`
		// memo String 部分币种充值需要 memo，若不需要则不返回此字段
		Memo  string               `json:"memo,omitempty"`
		// amt String 是 划转数量
		Amt   okex.JSONFloat64     `json:"amt"`
		// fee String 订单交易手续费，平台向用户收取的交易手续费，手续费扣除为负数。如： -0.01
		Fee   okex.JSONInt64       `json:"fee"`
		// wdId String 提币申请ID
		WdID  okex.JSONInt64       `json:"wdId"`
		// state String 订单状态  canceled：撤单成功live：等待成交partially_filled：部分成交filled：完全成交
		State okex.WithdrawalState `json:"state,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS    okex.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy  string           `json:"ccy"`
		// amt String 是 划转数量
		Amt  okex.JSONFloat64 `json:"amt"`
		// side String 是 订单方向 buy：买， sell：卖
		Side okex.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		// ccy String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
		Ccy      string           `json:"ccy"`
		// amt String 是 划转数量
		Amt      okex.JSONFloat64 `json:"amt"`
		// earnings String 币种持仓收益
		Earnings okex.JSONFloat64 `json:"earnings"`
	}
)
