// auth: kunlun
// date: 2019-01-07
package domain

const (
	BID  = "bid"
	ASK  = "ask"
	PING = "ping"
	PONG = "pong"
)

// tick
type Tick struct {
	Datas []ResponseData `json:"datas"`
}

type ResponseData struct {
	Dir       string  `json:"dir"`       //bid 卖  ask 买
	Symbol    string  `json:"symbol"`    // 交易对
	Ts        int64   `json:"ts"`        // 成交时间戳
	Amount    float64 `json:"amount"`    //成交量
	Price     float64 `json:"price"`     //成交价
	DayVolume float64 `json:"dayVolume"` //24小时成交总量
	DayPrice  float64 `json:"dayPrice"`  //24小时成交价
	DayHigh   float64 `json:"dayHigh"`   //最高价
	DayLow    float64 `json:"dayLow"`    //最低价
}
