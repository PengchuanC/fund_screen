package utils

import (
	"context"
	"fmt"
	"fund_screen/db/connector"
	"fund_screen/db/dao/tradingday"
	"fund_screen/db/models/funds"
	"github.com/golang-module/carbon"
)

// PeriodPerformanceRank 区间回报同类排名
func PeriodPerformanceRank(ctx context.Context, secucode string, start string) string {
	var (
		s string
		e carbon.ToDateString
		n funds.FNetValueRestore
		c funds.FClassifyNOI
		same []string
		pct float64
	)
	type result struct {
		Secucode string
		Pct float64
	}
	var changesS, changesE, change []result
	db = connector.GetDB()
	tx := db.WithContext(ctx)
	s = tradingday.NearestTradingDayAfter(ctx, start)
	tx.Model(&n).Select("max(date)").Where("secucode = ?", secucode).Find(&e)
	tx.Model(&c).Select("secucode").Where("second = (?)", db.Model(&c).Select("second").Where("secucode = ?", secucode)).Find(&same)
	tx.Model(&n).Select("secucode, unitnvrestored as pct").Where("date = ? and secucode in ?", s, same).Find(&changesS)
	tx.Model(&n).Select("secucode, unitnvrestored as pct").Where("date = ? and secucode in ?", e.ToDateString(), same).Find(&changesE)
	for _, ce := range changesE {
		for _, cs := range changesS {
			if ce.Secucode == cs.Secucode {
				change = append(change, result{
					Secucode: ce.Secucode,
					Pct:      ce.Pct / cs.Pct - 1,
				})
				if ce.Secucode == secucode {
					pct = ce.Pct / cs.Pct - 1
				}
				break
			}
		}
	}
	var rank = 1
	for _, c := range change {
		if c.Pct > pct {
			rank++
		}
	}
	return fmt.Sprintf("%d/%d",rank, len(change))
}
