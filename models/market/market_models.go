package market

import (
	"encoding/json"
	"fmt"
	"github.com/amir-the-h/okex"
	"strconv"
	"time"
)

type (
	Ticker struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID    string              `json:"instId"`
		// last String 最新成交价
		Last      okex.JSONFloat64    `json:"last"`
		// lastSz String 最新成交的数量
		LastSz    okex.JSONFloat64    `json:"lastSz"`
		// askPx String 卖一价
		AskPx     okex.JSONFloat64    `json:"askPx"`
		// askSz String 卖一价的挂单数数量
		AskSz     okex.JSONFloat64    `json:"askSz"`
		// bidPx String 买一价
		BidPx     okex.JSONFloat64    `json:"bidPx"`
		// bidSz String 买一价的挂单数量
		BidSz     okex.JSONFloat64    `json:"bidSz"`
		// open24h String 24小时开盘价
		Open24h   okex.JSONFloat64    `json:"open24h"`
		// high24h String 24小时最高价
		High24h   okex.JSONFloat64    `json:"high24h"`
		// low24h String 24小时最低价
		Low24h    okex.JSONFloat64    `json:"low24h"`
		// volCcy24h String 24小时成交量，以币为单位如果是衍生品合约，数值为交易货币的数量。如果是币币/币币杠杆，数值为计价货币的数量。
		VolCcy24h okex.JSONFloat64    `json:"volCcy24h"`
		// vol24h String 24小时成交量，以张为单位如果是衍生品合约，数值为合约的张数。如果是币币/币币杠杆，数值为交易货币的数量。
		Vol24h    okex.JSONFloat64    `json:"vol24h"`
		// sodUtc0 String UTC 0 时开盘价
		SodUtc0   okex.JSONFloat64    `json:"sodUtc0"`
		// sodUtc8 String UTC+8 时开盘价
		SodUtc8   okex.JSONFloat64    `json:"sodUtc8"`
		// instType String 产品类型SPOT：币币MARGIN：币币杠杆SWAP：永续合约 FUTURES：交割合约  OPTION：期权 交易产品类型不匹配指数（instType和uly不匹配）
		InstType  okex.InstrumentType `json:"instType,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS        okex.JSONTime       `json:"ts"`
	}
	IndexTicker struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string           `json:"instId"`
		// idxPx String 最新指数价格
		IdxPx   okex.JSONFloat64 `json:"idxPx"`
		// high24h String 24小时最高价
		High24h okex.JSONFloat64 `json:"high24h"`
		// low24h String 24小时最低价
		Low24h  okex.JSONFloat64 `json:"low24h"`
		// open24h String 24小时开盘价
		Open24h okex.JSONFloat64 `json:"open24h"`
		// sodUtc0 String UTC 0 时开盘价
		SodUtc0 okex.JSONFloat64 `json:"sodUtc0"`
		// sodUtc8 String UTC+8 时开盘价
		SodUtc8 okex.JSONFloat64 `json:"sodUtc8"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS      okex.JSONTime    `json:"ts"`
	}
	OrderBook struct {
		// asks Array 卖方深度
		Asks []*OrderBookEntity `json:"asks"`
		// bids Array 买方深度
		Bids []*OrderBookEntity `json:"bids"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS   okex.JSONTime      `json:"ts"`
	}
	OrderBookWs struct {
		// asks Array 卖方深度
		Asks     []*OrderBookEntity `json:"asks"`
		// bids Array 买方深度
		Bids     []*OrderBookEntity `json:"bids"`
		// checksum Integer 检验和 （下方注解）
		Checksum int                `json:"checksum,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS       okex.JSONTime      `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O      float64
		H      float64
		L      float64
		C      float64
		Vol    float64
		VolCcy float64
		TS     okex.JSONTime
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		TS okex.JSONTime
	}
	Trade struct {
		// instId String 是 产品ID，如 BTC-USD-190927-5000-C instId和instType不匹配 查询条件中的instId的交易产品当前不是可交易状态，请填写单个instid逐个查询状态详情 instId {0} 报价不可以超过你预设的价格限制
		InstID  string           `json:"instId"`
		// tradeId String 最新成交ID
		TradeID okex.JSONFloat64 `json:"tradeId"`
		// px String 可选               委托价格，仅适用于limit、post_only、fok、ioc类型的订单
		Px      okex.JSONFloat64 `json:"px"`
		// sz String 是 委托数量
		Sz      okex.JSONFloat64 `json:"sz"`
		// side String 是 订单方向 buy：买， sell：卖
		Side    okex.TradeSide   `json:"side,string"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS      okex.JSONTime    `json:"ts"`
	}
	TotalVolume24H struct {
		// volUsd String 24小时平台总成交量，以美元为单位
		VolUsd okex.JSONFloat64 `json:"volUsd"`
		// volCny String 24小时平台总成交量，以人民币为单位
		VolCny okex.JSONFloat64 `json:"volCny"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS     okex.JSONTime    `json:"ts"`
	}
	IndexComponent struct {
		// index String 是 指数，如 BTC-USDT
		Index      string           `json:"index"`
		// last String 最新成交价
		Last       okex.JSONFloat64 `json:"last"`
		// components String 成分
		Components []*Component     `json:"components"`
		// ts String 成交明细产生时间，Unix时间戳的毫秒数格式，如1597026383085
		TS         okex.JSONTime    `json:"ts"`
	}
	Component struct {
		// exch String 交易所名称
		Exch   string           `json:"exch"`
		// symbol String 采集的币对名称
		Symbol string           `json:"symbol"`
		// symPx String 采集的币对价格
		SymPx  okex.JSONFloat64 `json:"symPx"`
		// wgt String 权重
		Wgt    okex.JSONFloat64 `json:"wgt"`
		// cnvPx String 换算成指数后的价格
		CnvPx  okex.JSONFloat64 `json:"cnvPx"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice, err = strconv.ParseFloat(dp, 64)
	if err != nil {
		return err
	}
	o.Size, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts string
		err                          error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	return nil
}
