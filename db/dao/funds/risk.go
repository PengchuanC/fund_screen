package funds

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/dao/tradingday"
	"fund_screen/db/models/funds"
	"github.com/golang-module/carbon"
	"reflect"
)

// FundShortRisk 基金短期风控指标同类排名
func FundShortRisk(ctx context.Context, secucode string) map[string]interface{} {
	var (
		response map[string]interface{}
		cycles map[string]int
		models []funds.Model
	)
	cycles = map[string]int{"近3月": 3, "近6月": 6}
	models = []funds.Model{funds.RankAbnormal{}, funds.RankPerformance{}, funds.RankSharpe{}, funds.RankSD{}, funds.RankDrawback{}}
	response = risk(ctx, secucode, cycles, models)
	data := map[string]interface{}{
		"legend": []string{"超额收益", "绝对收益", "夏普比", "波动率", "最大回撤"},
		"value": response,
	}
	return data
}

// FundLongRisk 基金短期风控指标同类排名
func FundLongRisk(ctx context.Context, secucode string) map[string]interface{} {
	var (
		response map[string]interface{}
		cycles map[string]int
		models []funds.Model
	)
	cycles = map[string]int{"近1年": 12, "近2年": 24, "近3年": 36}
	models = []funds.Model{funds.RankAbnormal{}, funds.RankPerformance{}, funds.RankSelection{}, funds.RankTiming{}, funds.RankSharpe{}, funds.RankSD{}, funds.RankDrawback{}}
	response = risk(ctx, secucode, cycles, models)
	data := map[string]interface{}{
		"legend": []string{"超额收益", "绝对收益", "选股能力", "择时能力", "夏普比", "波动率", "最大回撤"},
		"value": response,
	}
	return data
}

func queryTemplate(ctx context.Context, secucode string, date carbon.ToDateString, cycle int, model funds.Model) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	dtype := reflect.TypeOf(model)
	result := reflect.New(dtype).Interface()
	sub := tx.Table(model.TableName()).Select("max(date)").Where("secucode = ? and date <= ?", secucode, date)
	tx.Table(model.TableName()).Where("secucode = ? and cycle = ? and date = (?)", secucode, cycle, sub).Find(result)
	return result
}


// 基金风控指标同类排名
func risk(ctx context.Context, secucode string, cycles map[string]int, models []funds.Model) map[string]interface{} {
	var (
		response map[string]interface{}
		lastFriday carbon.ToDateString
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	response = make(map[string]interface{})
	lastFriday = tradingday.NearestFriday(ctx, carbon.Now())
	for k, v := range cycles {
		var array []float64
		for _, ins := range models {
			resp := queryTemplate(ctx, secucode, lastFriday, v, ins)
			t := reflect.ValueOf(resp).Elem()
			array = append(array, t.FieldByName("Pct").Float() / 100)
		}
		response[k] = array
	}
	return response
}