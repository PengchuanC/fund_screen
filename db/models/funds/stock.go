package funds

import "github.com/golang-module/carbon"

type Stock struct {
	Secucode string `gorm:"primaryKey" json:"secucode,omitempty"`
	Secuabbr string `json:"secuabbr,omitempty"`
	Sector   int    `json:"sector,omitempty"`
	State    int    `json:"state,omitempty"`
	Market   int    `json:"market,omitempty"`
	Jsid     int64  `json:"jsid,omitempty"`
}

func (s Stock) TableName() string {
	return "sc_stock_secumain"
}

// StockValuation 股票估值数据
type StockValuation struct {
	Id           int64         `json:"id" json:"id,omitempty"`
	Secucode     string        `json:"secucode"`
	Category     int           `json:"category,omitempty"`
	Date         carbon.Carbon `json:"date"`
	Mv           float64       `json:"mv,omitempty"`
	NegotiableMv float64       `json:"negotiable_mv,omitempty"`
	PeLyr        float32       `json:"pe_lyr,omitempty"`
	PeTtm        float32       `json:"pe_ttm,omitempty"`
	Pb           float32       `json:"pb,omitempty"`
	Pcf          float32       `json:"pcf,omitempty"`
	PcfTtm       float32       `json:"pcf_ttm,omitempty"`
	Ps           float32       `json:"ps,omitempty"`
	Jsid         int64         `json:"jsid,omitempty"`
}

func (s StockValuation) TableName() string  {
	return "sc_stock_valuation"
}

type StockSW struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Secuabbr string              `json:"secuabbr,omitempty"`
	Category int8                `json:"category,omitempty"`
	First    string              `json:"first,omitempty"`
	Second   string              `json:"second,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Jsid     int64               `json:"jsid,omitempty"`
}

func (s StockSW) TableName() string {
	return "sc_stock_industry_sw"
}