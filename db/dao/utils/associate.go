package utils

import (
	"fund_screen/db/connector"
	"fund_screen/db/models/funds"
	"gorm.io/gorm"
)

var db *gorm.DB

// FundAssociate 寻找基金的关联基金，优先寻找联接基金对应的场内基金，其次寻找其他份额对应的A份额，
func FundAssociate(relate string) string {
	var (
		//secucode string
		am funds.FAssociate
		as []funds.FAssociate
	)
	db = connector.GetDB()
	db.Model(&am).Where("relate = ?", relate).Order("-define").Find(&as)
	if len(as) == 0 {
		return relate
	}
	am = as[0]
	return am.Secucode[0:6]
}
