package rpc

import (
	"context"
	"fund_screen/db/connector"
	funds2 "fund_screen/db/models/funds"
)
import "fund_screen/services"

func (s *ScreenRpcServer) FundCategory(ctx context.Context, request *services.ClassifyReq) (*services.ClassifyResp, error) {
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	var (
		category []funds2.FClassifyNOI
		resp []*services.Classify
		funds []string
	)
	funds = request.Secucode
	if len(funds) == 0 {
		tx.Model(&funds2.FClassifyNOI{}).Find(&category)
	} else {
		tx.Model(&funds2.FClassifyNOI{}).Where("secucode in (?)", funds).Find(&category)
	}
	for _, c := range category {
		r := &services.Classify{
			Secucode: c.Secucode,
			First:    c.First,
			Second:   c.Second,
		}
		resp = append(resp, r)
	}
	return &services.ClassifyResp{Data: resp}, nil
}
