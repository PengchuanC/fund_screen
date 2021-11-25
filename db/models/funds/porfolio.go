package funds

import "github.com/golang-module/carbon"

// FAllocate 资产配置
type FAllocate struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Stock    float64             `json:"stock,omitempty"`
	Bond     float64             `json:"bond,omitempty"`
	Fund     float64             `json:"fund,omitempty"`
	Metals   float64             `json:"metals,omitempty"`
	Monetary float64             `json:"monetary,omitempty"`
	Other    float64             `json:"other,omitempty"`
	Jsid     int64               `json:"jsid,omitempty"`
}

func (a FAllocate) TableName() string {
	return "sc_fund_portfolio_allocate"
}

// FStockKey 基金重仓股票，含普通基金和QDII基金
type FStockKey struct {
	Id        int64               `json:"id,omitempty"`
	Secucode  string              `json:"secucode,omitempty"`
	Date      carbon.ToDateString `json:"date"`
	Category  string              `json:"category,omitempty"`
	Serial    int32               `json:"serial,omitempty"`
	Stockcode string              `json:"stockcode,omitempty"`
	Stockabbr string              `json:"stockabbr,omitempty"`
	Shares    float64             `json:"shares,omitempty"`
	MktCap    float64             `json:"mkt_cap,omitempty"`
	Ratio     float64             `json:"ratio,omitempty"`
	Jsid      int64               `json:"jsid,omitempty"`
}

func (f FStockKey) TableName() string {
	return "sc_fund_portfolio_stock_key"
}

// FStockDetail 基金持股明细，仅包含中报和年报
type FStockDetail struct {
	Id        int64               `json:"id,omitempty"`
	Secucode  string              `json:"secucode,omitempty"`
	Date      carbon.ToDateString `json:"date"`
	Category  string              `json:"category,omitempty"`
	Serial    int32               `json:"serial,omitempty"`
	Stockcode string              `json:"stockcode,omitempty"`
	Stockabbr string              `json:"stockabbr,omitempty"`
	Shares    float64             `json:"shares,omitempty"`
	MktCap    float64             `json:"mkt_cap,omitempty"`
	Ratio     float64             `json:"ratio,omitempty"`
	Jsid      int64               `json:"jsid,omitempty"`
}

func (f FStockDetail) TableName() string {
	return "sc_fund_portfolio_stock_detail"
}

// FBondDetail 基金持债明细
type FBondDetail struct {
	Id        int64               `json:"id,omitempty"`
	Secucode  string              `json:"secucode,omitempty"`
	Date      carbon.ToDateString `json:"date"`
	Category  string              `json:"category,omitempty"`
	Serial    int32               `json:"serial,omitempty"`
	Bondcode string              `json:"bondcode,omitempty"`
	Bondabbr string              `json:"bondabbr,omitempty"`
	Shares    float64             `json:"shares,omitempty"`
	MktCap    float64             `json:"mkt_cap,omitempty"`
	Ratio     float64             `json:"ratio,omitempty"`
	Jsid      int64               `json:"jsid,omitempty"`
}

func (f FBondDetail) TableName() string {
	return "sc_fund_portfolio_bond_detail"
}

// FIndustryCsi 基金证监会行业分类占比
type FIndustryCsi struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Industry string              `json:"industry,omitempty"`
	Name     string              `json:"name,omitempty"`
	MktCap   float64             `json:"mkt_cap,omitempty"`
	Ratio    float64             `json:"ratio,omitempty"`
	Jsid     int64               `json:"jsid,omitempty"`
}

func (f FIndustryCsi) TableName() string {
	return "sc_fund_portfolio_industry_csi"
}