package index

import "github.com/golang-module/carbon"

type IYearlyPerf struct {
	Id       int64               `json:"id,omitempty"`
	Secucode string              `json:"secucode,omitempty"`
	Date     carbon.ToDateString `json:"date"`
	Year     int                 `json:"year,omitempty"`
	Value    float64             `json:"value,omitempty"`
}

func (f IYearlyPerf) TableName() string {
	return "sc_derivative_index_yearly"
}


type IPerformance struct {
	Id          int64
	Secucode    string
	Date        carbon.ToDateString
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
	Year3       float64
	Year5       float64
}

func (p IPerformance) TableName() string {
	return "sc_index_performance"
}