package funds

import "github.com/golang-module/carbon"

type RankAbnormal struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankAbnormal) TableName() string {
	return "sc_rank_analysis_abnormal"
}

type RankAbsolute struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankAbsolute) TableName() string {
	return "sc_rank_analysis_absolute"
}

type RankCalmar struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankCalmar) TableName() string {
	return "sc_rank_analysis_calmar"
}

type RankDrawback struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankDrawback) TableName() string {
	return "sc_rank_analysis_drawback"
}

type RankSelection struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankSelection) TableName() string {
	return "sc_rank_analysis_selection"
}

type RankSharpe struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankSharpe) TableName() string {
	return "sc_rank_analysis_sharpe"
}

type RankSortino struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankSortino) TableName() string {
	return "sc_rank_analysis_sortino"
}

type RankSD struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankSD) TableName() string {
	return "sc_rank_analysis_standard_deviation"
}

type RankTiming struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
}

func (r RankTiming) TableName() string {
	return "sc_rank_analysis_timing"
}

type RankPerformance struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `gorm:"foreignKey:Secucode" json:"secucode"`
	Cycle    int                 `json:"cycle,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Rank     string              `json:"rank,omitempty"`
	Pct      float64             `json:"pct,omitempty"`
	UpdateAt carbon.ToDateString `json:"update_at" gorm:"column:updateAt"`
}

func (r RankPerformance) TableName() string {
	return "sc_rank_analysis_performance"
}

type RankYearlyPerformance struct {
	Id int64
	Secucode string
	Date carbon.ToDateString
	Year int
	Value string
	Classify string
}

func (r RankYearlyPerformance) TableName() string {
	return "sc_rank_fund_yearly"
}