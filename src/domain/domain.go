// auth: kunlun
// date: 2019-01-06
// BUG(kunlun): #1: 数据类型对应Java中的BigDecimal 类型待商榷！
package domain

import "time"

// entrust entity
type Entrust struct {
	Id             int64     `json:"id"`
	UserId         int64     `json:"userId"`
	CoinCode       string    `json:"coinCode"`
	TradeType      int       `json:"tradeType"`
	Status         int       `json:"status"`
	TradeAmount    string    `json:"tradeAmount"`
	TradePrice     string    `json:"tradePrice"`
	DealAmount     string    `json:"dealAmount"`
	DealPrice      string    `json:"dealPrice"`
	CreateTime     time.Time `json:"createTime"`
	PoundageAmount string    `json:"poundageAmount"`
	Position       string    `json:"position"`
}

// leverage entity
type Leverage struct {
	Id             int64     `json:"id"`
	UserId         int64     `json:"userId"`
	MatUserId      int64     `json:"matUserId"`
	CoinCode       string    `json:"coinCode"`
	TargetCoin     string    `json:targetCoin`
	SourceCoin     string    `json:"sourceCoin"`
	TradePrice     string    `json:"tradePrice"`
	TradeAmount    string    `json:"tradeAmount"`
	Deposit        string    `json:"deposit"`
	Position       int       `json:"position"`
	TradeType      int       `json:"tradeType"`
	DealAmount     string    `json:"dealAmount"`
	DealPrice      string    `json:"dealPrice"`
	DealTime       time.Time `json:"dealTime"`
	Status         int       `json:"status"`
	OrderStatus    int       `json:"orderStatus"`
	CreateTime     time.Time `json:"createTime"`
	UpdateTime     time.Time `json:"updateTime"`
	CheckStr       string    `json:"checkStr"`
	PoundageAmount string    `json:"poundageAmount"`
	HoldPosition   string    `json:"holdPosition"`
}
