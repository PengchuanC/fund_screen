package funds

import "github.com/golang-module/carbon"

type Model interface {
	TableName() string
}

// AnalysisScale 基金规模
type AnalysisScale struct {
	Id       int64               `json:"id,omitempty"`
	Secucode Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Date     carbon.ToDateString `json:"date"`
	Nvi      float64             `json:"nvi,omitempty"`
	CombNvi  float64             `json:"comb_nvi,omitempty"`
	Jsid     int64               `json:"jsid,omitempty"`
}

func (s AnalysisScale) TableName() string {
	return "sc_fund_analysis_scale"
}

// AnalysisAbnormal 超额收益
type AnalysisAbnormal struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisAbnormal) TableName() string {
	return "sc_fund_analysis_abnormal"
}

// AnalysisAbsolute 绝对收益
type AnalysisAbsolute struct {
	Id         int64
	Secucode   Fund `gorm:"foreignKey:Secucode"`
	Cycle      int
	Date       carbon.ToDateString
	RiseToFall string
	Compound   float64
	Average    float64
	Jsid       int64
}

func (s AnalysisAbsolute) TableName() string {
	return "sc_fund_analysis_absolute"
}

type AnalysisCalmar struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisCalmar) TableName() string {
	return "sc_fund_analysis_calmar"
}

type AnalysisCapture struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Upside   float64
	Downside float64
	Jsid     int64
}

func (s AnalysisCapture) TableName() string {
	return "sc_fund_analysis_capture"
}

type AnalysisDrawback struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisDrawback) TableName() string {
	return "sc_fund_analysis_drawback"
}

type AnalysisDrawbackRestore struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisDrawbackRestore) TableName() string {
	return "sc_fund_analysis_drawback_restore"
}

type AnalysisSharpe struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisSharpe) TableName() string {
	return "sc_fund_analysis_sharpe"
}

type AnalysisSortino struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisSortino) TableName() string {
	return "sc_fund_analysis_sortino"
}

type AnalysisSD struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Cycle    int
	Date     carbon.ToDateString
	Value    float64
	Jsid     int64
}

func (s AnalysisSD) TableName() string {
	return "sc_fund_analysis_standard_deviation"
}

type AnalysisTMFF3 struct {
	Id        int64
	Secucode  Fund `gorm:"foreignKey:Secucode"`
	Cycle     int
	Date      carbon.ToDateString
	Selection float64
	Timing    float64
	Market    float64
	Smb       float64
	Hml       float64
	Jsid      int64
}

func (s AnalysisTMFF3) TableName() string {
	return "sc_fund_analysis_tm_ff3"
}
