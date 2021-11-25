package funds

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"fund_screen/db/models/system"
	"github.com/golang-module/carbon"
	"sort"
)

var styleMap = map[string]string{
	"电子": "TMT", "计算机": "TMT", "传媒": "TMT", "通信": "TMT", "农林牧渔": "周期", "采掘": "周期", "化工": "周期",
	"钢铁": "周期", "有色金属": "周期", "家用电器": "消费", "食品饮料": "消费", "纺织服装": "消费", "轻工制造": "消费",
	"医药生物": "消费", "公用事业": "消费", "商业贸易": "消费", "休闲服务": "消费", "交通运输": "中游制造",
	"综合": "中游制造", "建筑材料": "中游制造", "建筑装饰": "中游制造", "电气设备": "中游制造", "汽车": "中游制造",
	"机械设备": "中游制造", "国防军工": "中游制造", "房地产": "金融地产", "银行": "金融地产", "非银金融": "金融地产",
}

// FundAllocate 获取基金资产配置比例及与上期比较
func FundAllocate(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		allocate, prevAllocate funds.FAllocate
		//allocates []funds.FAllocate
		last carbon.ToDateString
		lastStr string
		chart []map[string]interface{}
		table []map[string]interface{}
	)

	tx.Model(&allocate).Select("max(date)").Where("secucode = ?", secucode).Find(&last)
	lastStr = last.ToDateString()
	tx.Model(&allocate).Where("secucode = ? and date = ?", secucode, lastStr).Find(&allocate)
	subQuery := tx.Model(&prevAllocate).Select("id").Where("secucode = ? and date < ?", secucode, lastStr).Order("-date").Limit(1)
	tx.Model(&prevAllocate).Where("id = (?)", subQuery).Find(&prevAllocate)
	chart = []map[string]interface{}{
		{"name": "股票", "ratio": allocate.Stock}, {"name": "债券", "ratio": allocate.Bond},
		{"name": "基金", "ratio": allocate.Fund}, {"name": "商品", "ratio": allocate.Metals},
		{"name": "现金", "ratio": allocate.Monetary}, {"name": "其他", "ratio": allocate.Other},
	}
	table = []map[string]interface{}{
		{"asset": "股票", "current": allocate.Stock, "change": allocate.Stock - prevAllocate.Stock},
		{"asset": "债券", "current": allocate.Bond, "change": allocate.Bond - prevAllocate.Bond},
		{"asset": "基金", "current": allocate.Fund, "change": allocate.Fund - prevAllocate.Fund},
		{"asset": "商品", "current": allocate.Metals, "change": allocate.Metals - prevAllocate.Metals},
		{"asset": "现金", "current": allocate.Monetary, "change": allocate.Monetary - prevAllocate.Monetary},
		{"asset": "其他", "current": allocate.Other, "change": allocate.Other - prevAllocate.Other},
	}
	return map[string]interface{}{"chart": chart, "table": table, "date": lastStr}
}

// FundAllocateHistory 获取基金历史报告期资产配置比例
func FundAllocateHistory(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		dates []carbon.ToDateString
		allocate funds.FAllocate
		allocates []funds.FAllocate
		ret []map[string]interface{}
	)

	subQuery := tx.Model(&system.TradingCalendar{}).Select("date").Where("quarter_end = 1").Find(&dates)
	tx.Model(&allocate).Where("secucode = ? and date in (?)", secucode, subQuery).Order("date").Find(&allocates)

	ret = []map[string]interface{}{}
	for _, a := range allocates {
		r := map[string]interface{}{
			"日期": a.Date, "股票": a.Stock, "债券": a.Bond, "基金": a.Fund,
			"商品": a.Metals, "现金": a.Monetary, "其他": a.Other,
		}
		ret = append(ret, r)
	}
	return map[string]interface{}{"history": ret}
}

// FundKeyStock 重仓持股
func FundKeyStock(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		last carbon.ToDateString
		now funds.FStockKey
		prev funds.FStockKey
		nh []funds.FStockKey
		ph []funds.FStockKey
		stocks []string
		valuation []funds.StockValuation
		resp []map[string]interface{}
	)

	tx.Model(&now).Select("max(date)").Where("secucode = ?", secucode).Find(&last)
	tx.Where("secucode = ? and date = ?", secucode, last.ToDateString()).Find(&nh)
	subQuery := tx.Model(&prev).Select("max(date)").Where("secucode = ? and date < ?", secucode, last)
	tx.Where("secucode = ? and date = (?)", secucode, subQuery).Find(&ph)
	for _, r := range nh {
		stocks = append(stocks, r.Stockcode)
	}
	subQuery = tx.Model(&funds.StockValuation{}).Select("max(date)")
	tx.Where("secucode in (?) and date = (?)", stocks, subQuery).Find(&valuation)

	resp = []map[string]interface{}{}
	for _, n := range nh {
		var r = map[string]interface{}{
			"stockcode": n.Stockcode,
			"stockname": n.Stockabbr,
			"ratio": n.Ratio,
			"ratio2": 0,
			"change": nil,
			"pe_ttm": nil,
			"pb": nil,
		}

		for _, p := range ph {
			if p.Stockcode == n.Stockcode {
				r["ratio2"] = p.Ratio
				r["change"] = r["ratio"].(float64) - r["ratio2"].(float64)
				break
			}
		}

		for _, v := range valuation {
			if v.Secucode == n.Stockcode {
				r["pe_ttm"] = v.PeTtm
				r["pb"] = v.Pb
				break
			}
		}

		resp = append(resp, r)
	}
	return map[string]interface{}{"stock": resp, "date": last.ToDateString()}
}

// FundKeyBond 重仓持债
func FundKeyBond(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		last carbon.ToDateString
		now funds.FBondDetail
		prev funds.FBondDetail
		nh []funds.FBondDetail
		ph []funds.FBondDetail
		resp []map[string]interface{}
	)

	tx.Model(&now).Select("max(date)").Where("secucode = ?", secucode).Find(&last)
	tx.Where("secucode = ? and date = ?", secucode, last.ToDateString()).Order("serial").Limit(10).Find(&nh)
	subQuery := tx.Model(&prev).Select("max(date)").Where("secucode = ? and date < ?", secucode, last)
	tx.Where("secucode = ? and date = (?)", secucode, subQuery).Find(&ph)

	resp = []map[string]interface{}{}
	for _, n := range nh {
		var r = map[string]interface{}{
			"bondcode": n.Bondcode,
			"bondname": n.Bondabbr,
			"ratio": n.Ratio,
			"ratio2": 0,
			"change": nil,
		}

		for _, p := range ph {
			if p.Bondcode == n.Bondcode {
				r["ratio2"] = p.Ratio
				r["change"] = r["ratio"].(float64) - r["ratio2"].(float64)
				break
			}
		}

		resp = append(resp, r)
	}
	return map[string]interface{}{"bond": resp, "date": last.ToDateString()}
}

// FundIndustry 行业占比
func FundIndustry(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		last carbon.ToDateString
		now funds.FIndustryCsi
		prev funds.FIndustryCsi
		nh []funds.FIndustryCsi
		ph []funds.FIndustryCsi
		resp []map[string]interface{}
	)

	tx.Model(&now).Select("max(date)").Where("secucode = ?", secucode).Find(&last)
	tx.Where("secucode = ? and date = ?", secucode, last.ToDateString()).Find(&nh)
	subQuery := tx.Model(&prev).Select("max(date)").Where("secucode = ? and date < ?", secucode, last)
	tx.Where("secucode = ? and date = (?)", secucode, subQuery).Find(&ph)

	resp = []map[string]interface{}{}
	for _, n := range nh {
		var r = map[string]interface{}{
			"industry_code": n.Industry,
			"industry": n.Name,
			"ratio": n.Ratio,
			"ratio2": 0,
			"change": nil,
		}

		for _, p := range ph {
			if p.Industry == n.Industry {
				r["ratio2"] = p.Ratio
				r["change"] = r["ratio"].(float64) - r["ratio2"].(float64)
				break
			}
		}

		resp = append(resp, r)
	}
	return map[string]interface{}{"industry": resp, "date": last.ToDateString()}
}

// FundConcentrate 股票集中度
func FundConcentrate(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	type Concentrate struct {
		Date  carbon.ToDateString `json:"date"`
		Ratio float64             `json:"ratio,omitempty"`
	}

	var (
		key funds.FStockKey
		cs []Concentrate
	)

	tx.Model(&key).Select(
		"date, sum(ratio) as ratio").Where(
			"secucode = ?", secucode).Group("date").Order("date").Find(&cs)
	return map[string]interface{}{"concentrate": cs}
}

// FundIndustryStyle 持仓风格
func FundIndustryStyle(ctx context.Context, secucode string) interface{} {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		detail funds.FStockDetail
		details []funds.FStockDetail
		data map[string]map[string]interface{}
		resp []map[string]interface{}
		stocks []string
		stocksMap map[string]string
		stockSW []funds.StockSW
		stockSWMap map[string]funds.StockSW
	)

	tx.Model(&detail).Where("secucode = ?", secucode).Find(&details)

	data = map[string]map[string]interface{}{}
	stocksMap = map[string]string{}
	for _, d := range details {
		if _, err := data[d.Date.ToDateString()]; !err {
			data[d.Date.ToDateString()] = map[string]interface{}{
				"date": d.Date,
				"TMT": .0,
				"中游制造": .0,
				"周期": .0,
				"消费": .0,
				"金融地产": .0,
				"其他": .0,
			}
		}
		stocksMap[d.Stockcode] = d.Stockcode
	}

	for _, s := range stocksMap {
		stocks = append(stocks, s)
	}

	tx.Model(&funds.StockSW{}).Where("secucode in (?)", stocks).Find(&stockSW)
	stockSWMap = map[string]funds.StockSW{}
	for _, sw := range stockSW {
		stockSWMap[sw.Secucode] = sw
	}

	for _, d := range details {
		is, _ := stockSWMap[d.Stockcode]
		industrySw := is.First
		style := styleMap[industrySw]
		r, _ := data[d.Date.ToDateString()][style]
		var ratio float64
		if r != nil {
			ratio = r.(float64) + d.Ratio
		} else {
			r, _ = data[d.Date.ToDateString()]["其他"]
			ratio = r.(float64) + d.Ratio
		}
		data[d.Date.ToDateString()][style] = ratio
	}

	var dates []string
	for date, _ := range data {
		dates = append(dates, date)
	}
	sort.Strings(dates)
	names := []string{"TMT", "中游制造", "周期", "消费", "金融地产", "其他"}
	for _, date := range dates {
		d := data[date]
		var sum = .0
		for _, name := range names {
			sum += d[name].(float64)
		}
		for _, name := range names {
			d[name] = d[name].(float64) / sum
		}
		resp = append(resp, d)
	}
	return map[string]interface{}{"data": resp, "names": names}
}
