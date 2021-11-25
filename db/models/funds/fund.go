package funds

import (
	"github.com/golang-module/carbon"
)

// Fund 基金主表
type Fund struct {
	Secucode string `gorm:"primaryKey"`
	Secuabbr string
	Category int
	Jsid     int64
}

func (f *Fund) TableName() string {
	return "sc_fund_secumain"
}

// FMainCode 基金主代码
type FMainCode struct {
	Id              int64
	Secucode        string
	Secuabbr        string
	Maincode        string
	Company         string
	Category        int
	Nature          int
	Fof             int
	Initiating      int
	LaunchDate      carbon.ToDateString
	ListedDate      carbon.ToDateString
	EstablishDateii carbon.ToDateString
	Benchmark       string
	Jsid            int64
}

func (m *FMainCode) TableName() string {
	return "sc_fund_maincode"
}

// FManager 基金经理
type FManager struct {
	Id             int64
	Secucode       Fund `gorm:"foreignKey:Secucode"`
	Personalcode   int64
	Name           string
	Postname       int
	Accessiondate  carbon.ToDateString
	Dimissiondate  carbon.ToDateString
	Managementtime int
	Performance    float64
	Incumbent      int
	Jsid           int64
}

func (m FManager) TableName() string {
	return "sc_fund_manager"
}

// FPerformance 基金业绩表现
type FPerformance struct {
	Id          int64
	Secucode    Fund `gorm:"foreignKey:Secucode"`
	Date        carbon.ToDateString
	Unitnv      float64
	Daily       float64
	Wtd         float64
	Week        float64
	Mtd         float64
	Month       float64
	Month3      float64
	Month6      float64
	Ytd         float64
	Year        float64
	Year2       float64
	Year2Annual float64
	Year3       float64
	Year3Annual float64
	Year5       float64
	Year5Annual float64
	Ftd         float64
	FtdAnnual   float64
	Jsid        int64
}

func (p FPerformance) TableName() string {
	return "sc_fund_performance"
}

// FAssociate 代码关联
type FAssociate struct {
	Id       int64
	Secucode string
	Relate   string
	Define   int
	Ms       string
	Jsid     int64
}

func (a FAssociate) TableName() string {
	return "sc_fund_associate"
}

// FNetValue 单位净值
type FNetValue struct {
	Id                int64
	Secucode          Fund `gorm:"foreignKey:Secucode"`
	Date              carbon.ToDateString
	Nv                float64
	Unitnv            float64
	Accumulatedunitnv float64
	Dailyprofit       float64
	Latestweeklyyield float64
	Jsid              int64
}

func (p FNetValue) TableName() string {
	return "sc_fund_net_value"
}

type FNetValueRestore struct {
	Id             int64
	Secucode       Fund `gorm:"foreignKey:Secucode"`
	Date           carbon.ToDateString
	Unitnv         float64
	Unitnvrestored float64
	Jsid           int64
}

func (n FNetValueRestore) TableName() string {
	return "sc_fund_net_value_restored"
}

type FClassifyNOI struct {
	Id       int64
	Secucode string
	First    string
	Second   string
	Jsid     int64
}

func (n FClassifyNOI) TableName() string {
	return "sc_fund_classify_nomura"
}

// FTurnover 换手率
type FTurnover struct {
	Id       int64
	Secucode Fund `gorm:"foreignKey:Secucode"`
	Date     carbon.ToDateString
	Ratio    float64
	Jsid     int64
}

func (n FTurnover) TableName() string {
	return "sc_fund_turnover"
}

// FHolder 持有人结构
type FHolder struct {
	Id          int64               `json:"id,omitempty"`
	Secucode    Fund                `gorm:"foreignKey:Secucode" json:"secucode"`
	Date        carbon.ToDateString `json:"date"`
	Individual  float64             `json:"individual,omitempty"`
	Institution float64             `json:"institution,omitempty"`
	Undefined   float64             `json:"undefined,omitempty"`
	Jsid        int64               `json:"jsid,omitempty"`
}

func (n FHolder) TableName() string {
	return "sc_fund_holder"
}
