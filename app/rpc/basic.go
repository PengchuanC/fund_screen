package rpc

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"fund_screen/services"
)

func (s *ScreenRpcServer) FundBasicInfoHandler(ctx context.Context, request *services.FundBasicInfoRequest) (*services.FundBasicInfoResponse, error) {
	var db = connector.GetDB()
	var tx = db.WithContext(ctx)
	var info []funds.FMainCode
	var ret []*services.BasicInfo

	tx.Model(funds.FMainCode{}).Find(&info)
	for _, i := range info {
		ret = append(ret, &services.BasicInfo{
			Secucode:   i.Secucode,
			LaunchDate: i.LaunchDate.ToDateString(),
		})
	}
	return &services.FundBasicInfoResponse{Data: ret}, nil
}
