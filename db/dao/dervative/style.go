package dervative

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"github.com/golang-module/carbon"
)

// HistoryStyle 基金与指数风格的相似度
func HistoryStyle(ctx context.Context, secucode string) []map[string]interface{}{
	var (
		style []funds.IndexCorrelation
		resp []map[string]interface{}
		date carbon.Carbon
	)
	date = carbon.CreateFromTime(carbon.Now().Year(), 1, 1).SubYears(5)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(funds.IndexCorrelation{}).Where("secucode = ? and date > ?", secucode, date.ToDateString()).Find(&style)
	for _, s := range style {
		var value int
		switch s.Index {
		case "000300":
			value = 0
		case "000905":
			value = 1
		case "000852":
			value = 2
		default:
			value = 3
		}
		resp = append(resp, map[string]interface{}{"date": s.Date, "value": value})
	}
	return resp
}


func RBSAStyle(ctx context.Context, secucode string) []funds.RBSA {
	var (
		style []funds.RBSA
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(funds.RBSA{}).Where("secucode = ?", secucode).Order("date").Find(&style)
	return style
}