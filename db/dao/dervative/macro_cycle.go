package dervative

import (
	"context"
	"errors"
	"fmt"
	"fund_screen/db/connector"
	"fund_screen/db/dao/tradingday"
	"fund_screen/db/models/funds"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	tx *gorm.DB
)

// 全部宏观周期
func macroCycle(ctx context.Context) []funds.MacroCycle {
	var cycles []funds.MacroCycle
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(funds.MacroCycle{}).Find(&cycles)
	return cycles
}

func periodPerformance(ctx context.Context, secucode string, start, end carbon.ToDateString) (float64, error) {
	var (
		netvalues []funds.FNetValueRestore
		dates []carbon.ToDateString
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	start = tradingday.NearestTradingDay(ctx, start.ToDateString())
	end = tradingday.NearestTradingDay(ctx, end.ToDateString())
	dates = []carbon.ToDateString{start, end}
	tx.Model(funds.FNetValueRestore{}).Where("secucode = ? and date in (?)", secucode, dates).Find(&netvalues).Order("date")
	if len(netvalues) != 2 {
		return 0, errors.New("基金尚未成立，暂无数据")
	}
	return netvalues[1].Unitnvrestored / netvalues[0].Unitnvrestored - 1, nil
}

// FundMacroCyclePerformance 基金在周期内的表现
func FundMacroCyclePerformance(ctx context.Context, secucode string) []map[string]interface{} {
	var (
		cycle []funds.MacroCycle
		perf map[string]interface{}
		performance []map[string]interface{}
	)
	cycle = macroCycle(ctx)
	for _, c := range cycle {
		perf = make(map[string]interface{})
		value, err := periodPerformance(ctx, secucode, c.Start, c.End)
		if err != nil {
			continue
		}
		perf["date"] = fmt.Sprintf("%s(%s至%s)", c.Flag, c.Start.Format("Ymd"), c.End.Format("Ymd"))
		perf["value"] = value
		performance = append(performance, perf)
	}
	return performance
}