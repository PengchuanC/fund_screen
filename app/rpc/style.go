package rpc

import (
	"context"
	"fund_screen/db/connector"
	funds2 "fund_screen/db/models/funds"
	"fund_screen/services"
)

func (s *ScreenRpcServer) FundScaleNature(ctx context.Context, req *services.StyleReq) (*services.StyleResp, error) {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		funds   []string
		many    bool
		results []funds2.ScaleNature
		ret services.StyleResp
	)

	funds = req.Funds
	many = req.Many
	if many {
		tx.Model(funds2.ScaleNature{}).Where("secucode in (?)", funds).Find(&results)
	} else {
		tx.Raw("SELECT distinct a.secucode, a.date, a.style from sc_derivative_fund_scale_nature a join (select secucode, max(`date`) as `date` from sc_derivative_fund_scale_nature group by secucode) b on a.secucode=b.secucode and a.date=b.date where a.secucode in (?) order by a.secucode, a.date", funds).Find(&results)
	}
	var data = map[string]*services.Styles{}
	for _, r := range results {
		data[r.Secucode] = &services.Styles{Styles: []*services.Styles_Style{}}
	}
	for s := range data {
		var styles []*services.Styles_Style
		for _, r := range results {
			if s == r.Secucode {
				styles = append(styles, &services.Styles_Style{
					Secucode: r.Secucode,
					Date:     r.Date.ToDateString(),
					Style:    r.Style,
				})
			}
		}
		data[s] = &services.Styles{Styles: styles}
	}
	ret = services.StyleResp{Data: data}
	return &ret, nil
}

func (s *ScreenRpcServer) FundStyleNature(ctx context.Context, req *services.StyleReq) (*services.StyleResp, error) {
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	var (
		funds   []string
		many    bool
		results []funds2.StyleNature
		ret services.StyleResp
	)

	funds = req.Funds
	many = req.Many
	if many {
		tx.Model(funds2.StyleNature{}).Where("secucode in (?)", funds).Find(&results)
	} else {
		tx.Raw("SELECT distinct a.secucode, a.date, a.style from sc_derivative_fund_style_nature a join (select secucode, max(`date`) as `date` from sc_derivative_fund_style_nature group by secucode) b on a.secucode=b.secucode and a.date=b.date where a.secucode in (?) order by a.secucode, a.date", funds).Find(&results)
	}
	var data = map[string]*services.Styles{}
	for _, r := range results {
		data[r.Secucode] = &services.Styles{Styles: []*services.Styles_Style{}}
	}
	for s := range data {
		var styles []*services.Styles_Style
		for _, r := range results {
			if s == r.Secucode {
				styles = append(styles, &services.Styles_Style{
					Secucode: r.Secucode,
					Date:     r.Date.ToDateString(),
					Style:    r.Style,
				})
			}
		}
		data[s] = &services.Styles{Styles: styles}
	}
	ret = services.StyleResp{Data: data}
	return &ret, nil
}

func (s *ScreenRpcServer) FundRelatedIndex(ctx context.Context, req *services.IndexCorrReq) (*services.IndexCorrResp, error) {
	db = connector.GetDB()
	tx = db.WithContext(ctx)

	var (
		indexes []string
		results []funds2.IndexCorrelation
		data map[string]string
	)

	indexes = req.Indexes
	tx.Raw(" SELECT a.* FROM sc_derivative_index_correlation a join (SELECT secucode, MAX(`date`) as `date` FROM sc_derivative_index_correlation group by secucode) b on a.`date`  = b.`date` and a.secucode = b.secucode where a.secucode in (?)", indexes).Find(&results)
	data = map[string]string{}
	for _, r := range results {
		data[r.Secucode] = r.Index
	}
	return &services.IndexCorrResp{Data: data}, nil
}
