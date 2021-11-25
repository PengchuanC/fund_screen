package funds

import (
	"context"
	"fund_screen/db/connector"
	manager2 "fund_screen/db/dao/manager"
	"fund_screen/db/models/funds"
	"fund_screen/db/types"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"math"
	"reflect"
)

var (
	db *gorm.DB
	tx *gorm.DB
)

func FundLatestPerformance(ctx context.Context, secucode string) map[string]interface{} {
	var (
		fund        funds.Fund
		performance funds.FPerformance
		netValue    funds.FNetValue
		resp        map[string]interface{}
		maxDate     carbon.ToDateString
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	fund = funds.Fund{Secucode: secucode}
	tx.Take(&fund)
	performance = funds.FPerformance{Secucode: fund}
	netValue = funds.FNetValue{Secucode: fund}
	tx.Select("MAX(date)").Where("secucode = ?", secucode).Table(performance.TableName()).Find(&maxDate)
	tx.Where("date = ?", maxDate).Find(&performance, "secucode = ?", secucode)
	tx.Where("date = ?", maxDate).Find(&netValue, "secucode = ?", secucode)
	resp = make(map[string]interface{})
	resp["secucode"] = fund.Secucode
	resp["sec_name"] = fund.Secuabbr
	resp["nav"] = netValue.Unitnv
	resp["nav_acc"] = netValue.Accumulatedunitnv
	resp["return_1w"] = performance.Week
	resp["return_1m"] = performance.Month
	resp["return_3m"] = performance.Month3
	resp["return_6m"] = performance.Month6
	resp["return_1y"] = performance.Year
	resp["return_3y"] = performance.Year3
	resp["return_ytd"] = performance.Ytd
	resp["return_std"] = performance.Ftd
	resp["update_date"] = performance.Date
	return resp
}

func FundManagerInfo(ctx context.Context, secucode string) map[string]interface{} {
	var (
		manager  manager2.Manager
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	manager = manager2.Manager{Secucode: secucode}
	resp := manager.Info()
	return resp
}

func queryLatestValue(ctx context.Context, secucode string, model funds.Model) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	dtype := reflect.TypeOf(model)
	result := reflect.New(dtype).Interface()
	subQuery := tx.Table(model.TableName()).Select("max(date)").Where("secucode = ?", secucode)
	tx.Table(model.TableName()).Where("secucode = ? and date = (?)", secucode, subQuery).Find(result)
	return result
}

func queryLatestValues(ctx context.Context, secucode string, model funds.Model) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	dtype := reflect.TypeOf(model)
	results := reflect.New(reflect.SliceOf(dtype)).Interface()
	subQuery := tx.Table(model.TableName()).Select("max(date)").Where("secucode = ?", secucode)
	tx.Table(model.TableName()).Where("secucode = ? and date = (?)", secucode, subQuery).Find(results)
	return results
}

// FundPerformance 区间业绩表现-主要是风控指标
func FundPerformance(ctx context.Context, secucode string) map[int]types.PerformanceType {
	var (
		performance map[int]types.PerformanceType
		tmp         interface{}
		absolute    funds.FPerformance
		abnormal    []funds.AnalysisAbnormal
		sd          []funds.AnalysisSD
		drawback    []funds.AnalysisDrawback
		sharpe      []funds.AnalysisSharpe
		calmar      []funds.AnalysisCalmar
		tmff        []funds.AnalysisTMFF3
	)
	performance = types.NewPerformance()

	tmp = queryLatestValue(ctx, secucode, funds.FPerformance{})
	absolute = *tmp.(*funds.FPerformance)
	for key, value := range performance {
		switch key {
		case 1:
			value.Absolute = absolute.Month / 100
			value.Annual = math.Pow(1+absolute.Month/100, 12) - 1
		case 3:
			value.Absolute = absolute.Month3 / 100
			value.Annual = math.Pow(1+absolute.Month3/100, 4) - 1
		case 6:
			value.Absolute = absolute.Month6 / 100
			value.Annual = math.Pow(1+absolute.Month6/100, 2) - 1
		case 12:
			value.Absolute = absolute.Year / 100
			value.Annual = absolute.Year / 100
		case 36:
			value.Absolute = absolute.Year3 / 100
			value.Annual = absolute.Year3Annual / 100
		case 60:
			value.Absolute = absolute.Year5 / 100
			value.Annual = absolute.Year5Annual / 100
		}
		performance[key] = value
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisAbnormal{})
	abnormal = *tmp.(*[]funds.AnalysisAbnormal)
	for _, a := range abnormal {
		if value, ok := performance[a.Cycle]; ok {
			value.Abnormal = a.Value
			performance[a.Cycle] = value
		}
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisSD{})
	sd = *tmp.(*[]funds.AnalysisSD)
	for _, a := range sd {
		if value, ok := performance[a.Cycle]; ok {
			value.Vol = a.Value
			performance[a.Cycle] = value
		}
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisDrawback{})
	drawback = *tmp.(*[]funds.AnalysisDrawback)
	for _, a := range drawback {
		if value, ok := performance[a.Cycle]; ok {
			value.Drawback = a.Value
			performance[a.Cycle] = value
		}
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisSharpe{})
	sharpe = *tmp.(*[]funds.AnalysisSharpe)
	for _, a := range sharpe {
		if value, ok := performance[a.Cycle]; ok {
			value.Sharpe = a.Value
			performance[a.Cycle] = value
		}
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisCalmar{})
	calmar = *tmp.(*[]funds.AnalysisCalmar)
	for _, a := range calmar {
		if value, ok := performance[a.Cycle]; ok {
			value.Calmar = a.Value
			performance[a.Cycle] = value
		}
	}

	tmp = queryLatestValues(ctx, secucode, funds.AnalysisTMFF3{})
	tmff = *tmp.(*[]funds.AnalysisTMFF3)
	for _, a := range tmff {
		if value, ok := performance[a.Cycle]; ok {
			value.Selection = a.Selection
			value.Timing = a.Timing
			performance[a.Cycle] = value
		}
	}

	return performance
}

func FundTurnover(ctx context.Context, secucode string) []map[string]interface{} {
	var (
		turnover []funds.FTurnover
		resp     []map[string]interface{}
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&funds.FTurnover{}).Find(&turnover, "secucode = ?", secucode)
	for _, t := range turnover {
		r := map[string]interface{}{"date": t.Date, "ratio": t.Ratio}
		resp = append(resp, r)
	}
	return resp
}

func FundHolderStruct(ctx context.Context, secucode string) []funds.FHolder {
	var (
		holder []funds.FHolder
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&funds.FHolder{}).Find(&holder, "secucode = ?", secucode)
	return holder
}

func FundScaleChange(ctx context.Context, secucode string) []funds.AnalysisScale {
	var (
		scale []funds.AnalysisScale
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&funds.AnalysisScale{}).Find(&scale, "secucode = ?", secucode)
	return scale
}