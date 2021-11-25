package funds

import "github.com/golang-module/carbon"

// MacroCycle 宏观周期
type MacroCycle struct {
	Id    int64
	Flag  string
	Start carbon.ToDateString
	End   carbon.ToDateString
}

func (m MacroCycle) TableName() string {
	return "sc_derivative_macro_cycle"
}

// IndexCorrelation 与指数相关系数
type IndexCorrelation struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Index    string              `json:"index,omitempty"`
	Date     carbon.ToDateString `json:"date"`
}

func (i IndexCorrelation) TableName() string {
	return "sc_derivative_index_correlation"
}

// TopHolding 基金经理任职以来持仓次数靠前的股票
type TopHolding struct {
	Id                 int64               `json:"id,omitempty"`
	Secucode           Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Stockcode          string              `json:"stockcode,omitempty"`
	Stockabbr          string              `json:"stockabbr,omitempty"`
	Industry           string              `json:"industry,omitempty"`
	Earliest           carbon.ToDateString `json:"earliest"`
	Latest             carbon.ToDateString `json:"latest"`
	Count              int                 `json:"count,omitempty"`
	AvgMktCapRatio     float64             `json:"avg_mkt_cap_ratio,omitempty"`
	PeriodReturn       float64             `json:"period_return,omitempty"`
	PeriodAnnualReturn float64             `json:"period_annual_return,omitempty"`
	PeriodIndexReturn  float64             `json:"period_index_return,omitempty"`
	IndustryReturn     float64             `json:"industry_return,omitempty"`
	Date               carbon.ToDateString `json:"date"`
}

func (t TopHolding) TableName() string {
	return "sc_derivative_holding_top20"
}

type RBSA struct {
	Id          int64               `json:"id,omitempty"`
	Secucode    string              `json:"secucode,omitempty"`
	Date        carbon.ToDateString `json:"date"`
	SmallValue  float64             `json:"small_value,omitempty"`
	SmallGrowth float64             `json:"small_growth,omitempty"`
	MidValue    float64             `json:"mid_value,omitempty"`
	MidGrowth   float64             `json:"mid_growth,omitempty"`
	LargeValue  float64             `json:"large_value,omitempty"`
	LargeGrowth float64             `json:"large_growth,omitempty"`
	Bond        float64             `json:"bond,omitempty"`
	RSquare     float64             `json:"r_square,omitempty"`
}

func (r RBSA) TableName() string {
	return "sc_derivative_fund_rbsa"
}

type FYearlyPerf struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Year     int                 `json:"year,omitempty"`
	Value    float64             `json:"value,omitempty"`
}

func (f FYearlyPerf) TableName() string {
	return "sc_derivative_fund_yearly"
}

type ScaleNature struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Style    string              `json:"style,omitempty"`
	Value    float64             `json:"value,omitempty"`
}

func (s ScaleNature) TableName() string {
	return "sc_derivative_fund_scale_nature"
}

type StyleNature struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Style    string              `json:"style,omitempty"`
	Value    float64             `json:"value,omitempty"`
}

func (s StyleNature) TableName() string {
	return "sc_derivative_fund_style_nature"
}


type IndustryNature struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Style    string              `json:"style,omitempty"`
	Value    float64             `json:"value,omitempty"`
}

func (s IndustryNature) TableName() string {
	return "sc_derivative_fund_industry_nature"
}

type IndustryStyle struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Style    string              `json:"style,omitempty"`
}

func (s IndustryStyle) TableName() string {
	return "sc_derivative_fund_industry_style"
}