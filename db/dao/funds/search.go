package funds

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"strconv"
)

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// FundSearchFundList 根据关键字查找基金代码和基金简称
func FundSearchFundList(ctx context.Context, search string) [][2]string {
	var (
		results []funds.Fund
		ret [][2]string
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	search1 := search + "%"
	search2 := "%" + search + "%"
	if IsNum(search) {
		tx.Model(funds.Fund{}).Where("secucode like ?", search1).Find(&results)
	} else {
		tx.Model(funds.Fund{}).Where("secuabbr like ?", search2).Find(&results)
	}
	for i, f := range results{
		if i >= 10 { break }
		r := [2]string{f.Secucode, f.Secuabbr}
		ret = append(ret, r)
	}
	return ret
}
