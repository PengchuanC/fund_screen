package funds

import (
	"context"
	"fmt"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"fund_screen/db/models/index"
	"github.com/golang-module/carbon"
	"strconv"
)

func round2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func round4(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

// StyleAndBenchmarkIndex 风格指数与宽基指数
func StyleAndBenchmarkIndex(ctx context.Context) map[string]interface{} {
	var (
		style     []index.Index
		benchmark []index.Index
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(index.Index{}).Select([]string{"secucode", "chiname"}).Where("secucode like '000%'").Find(&benchmark)
	tx.Model(index.Index{}).Select([]string{"secucode", "chiname"}).Where("secucode like 'H11%'").Find(&style)
	return map[string]interface{}{
		"style":     style,
		"benchmark": benchmark,
	}
}

func indicesNames(fund, style, benchmark string) [3]string {
	var (
		fName funds.Fund
		sName index.Index
		bName index.Index
	)
	db = connector.GetDB()
	db.Model(fName).Find(&fName, funds.Fund{Secucode: fund})
	db.Model(sName).Select([]string{"secucode", "chiname"}).Find(&sName, index.Index{Secucode: style})
	db.Model(bName).Select([]string{"secucode", "chiname"}).Find(&bName, index.Index{Secucode: benchmark})
	return [3]string{fName.Secuabbr, sName.Chiname, bName.Chiname}
}

// NetValueSeries 净值序列
func NetValueSeries(ctx context.Context, secucode, style, benchmark string) map[string]interface{} {
	var (
		nav []funds.FNetValueRestore
		f   funds.FNetValueRestore
		i   index.IQuote
		st  []index.IQuote
		bc  []index.IQuote
	)
	type retType struct {
		Date      carbon.ToDateString `json:"date"`
		Fund      float64             `json:"fund,omitempty"`
		Style     float64             `json:"style,omitempty"`
		Benchmark float64             `json:"benchmark,omitempty"`
	}
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(f).Select([]string{"date", "unitnvrestored"}).Where("secucode = ?", secucode).Order("date").Find(&nav)
	tx.Model(i).Select([]string{"date", "closeprice"}).Where("secucode = ?", style).Order("date").Find(&st)
	tx.Model(i).Select([]string{"date", "closeprice"}).Where("secucode = ?", benchmark).Order("date").Find(&bc)

	var data []retType
	sInit := round4(st[0].Closeprice)
	bInit := round4(bc[0].Closeprice)
	for _, n := range nav {
		var d = retType{Date: n.Date, Fund: round4(n.Unitnvrestored), Style: sInit, Benchmark: bInit}
		var length = len(data)
		if length > 0 {
			d.Style = data[length-1].Style
			d.Benchmark = data[length-1].Benchmark
		}
		for _, s := range st {
			if s.Date == n.Date {
				d.Style = round4(s.Closeprice)
				break
			} else if s.Date.Carbon.Gt(n.Date.Carbon) {
				break
			}
		}
		for _, b := range bc {
			if b.Date == n.Date {
				d.Benchmark = round4(b.Closeprice)
				break
			} else if b.Date.Carbon.Gt(n.Date.Carbon) {
				break
			}
		}
		data = append(data, d)
	}
	dInit := data[0]
	var data_ []retType
	for _, d := range data {
		d.Fund /= dInit.Fund
		d.Style /= dInit.Style
		d.Benchmark /= dInit.Benchmark
		data_ = append(data_, d)
	}

	names := indicesNames(secucode, style, benchmark)
	return map[string]interface{}{
		"names": names,
		"data":  data_,
	}
}

// RecentPerformance 近期业绩表现
func RecentPerformance(ctx context.Context, fund, style, benchmark string) []map[string]interface{} {
	var (
		performance funds.FPerformance
		rank []funds.RankPerformance
		r funds.RankPerformance
		st index.IPerformance
		bc index.IPerformance
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	performance = funds.FPerformance{Secucode: funds.Fund{Secucode: fund}}
	subQuery := tx.Model(performance).Select("max(date)").Where("secucode = ?", fund)
	tx.Model(performance).Where("secucode = ? and date = (?)", fund, subQuery).Find(&performance)
	subQuery = tx.Model(st).Select("max(date)").Where("secucode= ?", style)
	tx.Model(st).Where("secucode = ? and date = (?)", style, subQuery).Find(&st)
	tx.Model(bc).Where("secucode = ? and date = (?)", benchmark, subQuery).Find(&bc)
	names := indicesNames(fund, style, benchmark)
	var table []map[string]interface{}
	f := make(map[string]interface{})
	f = map[string]interface{}{
		"name": names[0],
		"m3":   round2(performance.Month3),
		"m6":   round2(performance.Month6),
		"y1":   round2(performance.Year),
		"y3":   round2(performance.Year3),
		"y5":   round2(performance.Year5),
		"ytd":  round2(performance.Ytd),
	}
	s := make(map[string]interface{})
	s = map[string]interface{}{
		"name": names[1],
		"m3":   round2(st.Month3),
		"m6":   round2(st.Month6),
		"y1":   round2(st.Year),
		"y3":   round2(st.Year3),
		"y5":   round2(st.Year5),
		"ytd":  round2(st.Ytd),
	}
	b := make(map[string]interface{})
	b = map[string]interface{}{
		"name": names[2],
		"m3":   round2(bc.Month3),
		"m6":   round2(bc.Month6),
		"y1":   round2(bc.Year),
		"y3":   round2(bc.Year3),
		"y5":   round2(bc.Year5),
		"ytd":  round2(bc.Ytd),
	}

	fr := make(map[string]interface{})
	var mapping = map[int]string{3: "m3", 6: "m6", 12: "y1", 36: "y3", 60: "y5", 998: "ytd"}
	subQuery = tx.Model(r).Select("max(date)").Where("secucode = ?", fund)
	tx.Model(r).Where("secucode = ? and date = (?)", fund, subQuery).Find(&rank)
	for _, v := range rank{
		for k, m := range mapping {
			if k == v.Cycle {
				fr[m] = v.Rank
				break
			}
		}
	}
	fr["name"] = "同类排名"

	table = []map[string]interface{}{f, fr, s, b}
	return table
}

// YearlyPerformance 分年度业绩表现
func YearlyPerformance(ctx context.Context, fund, style, benchmark string) map[string]interface{} {
	var (
		performance []funds.FYearlyPerf
		fyp         funds.FYearlyPerf
		idx         index.IYearlyPerf
		rk funds.RankYearlyPerformance
		st          []index.IYearlyPerf
		bc          []index.IYearlyPerf
		rank []funds.RankYearlyPerformance
		start       int
	)
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	start = carbon.Now().SubYears(5).Year()
	tx.Model(fyp).Where("secucode = ? and year >= ?", fund, start).Find(&performance)
	tx.Model(idx).Where("secucode = ? and year >= ?", style, start).Find(&st)
	tx.Model(idx).Where("secucode = ? and year >= ?", benchmark, start).Find(&bc)
	tx.Model(rk).Where("secucode = ? and year >= ?", fund, start).Find(&rank)
	names := indicesNames(fund, style, benchmark)

	var table []map[string]interface{}
	var p, s, b, r map[string]interface{}
	p = make(map[string]interface{})
	s = make(map[string]interface{})
	b = make(map[string]interface{})
	r = make(map[string]interface{})
	for i := 0; i < len(performance); i++ {
		p[strconv.Itoa(performance[i].Year)] = round2(performance[i].Value*100)
		s[strconv.Itoa(st[i].Year)] = round2(st[i].Value*100)
		b[strconv.Itoa(bc[i].Year)] = round2(bc[i].Value*100)
	}
	for _, v := range rank {
		r[strconv.Itoa(v.Year)] = v.Value
	}
	p["name"] = names[0]
	s["name"] = names[1]
	b["name"] = names[2]
	r["name"] = "同类排名"
	table = []map[string]interface{}{p, r, s, b}

	var chart map[string]interface{}
	var fundC, styleC, benchmarkC []float64
	var years []int
	for i := 0; i < len(performance); i++ {
		fundC = append(fundC, round2(performance[i].Value*100))
		styleC = append(styleC, round2(st[i].Value*100))
		benchmarkC = append(benchmarkC, round2(bc[i].Value*100))
		years = append(years, performance[i].Year)
	}
	chart = map[string]interface{}{
		"names":     names,
		"fund":      fundC,
		"style":     styleC,
		"benchmark": benchmarkC,
		"date":      years,
	}

	return map[string]interface{}{
		"table": table,
		"chart": chart,
	}
}
