package dervative

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
)

type FTopHolding struct {
	funds.TopHolding
	funds.StockValuation
}

func FundTopHolding(ctx context.Context, secucode string) []FTopHolding{
	var (
		holdings []funds.TopHolding
		valuation []funds.StockValuation
		stocks []string
		results []FTopHolding
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	subQuery := tx.Model(funds.TopHolding{}).Select("max(date)").Where("secucode = ?", secucode)
	tx.Model(funds.TopHolding{}).Where("secucode = ? and date = (?)", secucode, subQuery).Find(&holdings)
	for _, h := range holdings {
		stocks = append(stocks, h.Stockcode)
	}
	subQuery = tx.Model(funds.StockValuation{}).Select("max(id)")
	subQuery = tx.Model(funds.StockValuation{}).Select("date").Where("id = (?)",subQuery)
	tx.Model(funds.StockValuation{}).Where("date = (?) and secucode in ?", subQuery, stocks).Find(&valuation)
	for _, h := range holdings {
		for _, v := range valuation {
			if h.Stockcode == v.Secucode {
				fth := FTopHolding{
					TopHolding:     h,
					StockValuation: v,
				}
				results = append(results, fth)
				break
			}
		}
	}
	return results
}
