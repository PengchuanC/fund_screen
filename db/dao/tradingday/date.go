package tradingday

import (
	"context"
	"fund_screen/db/connector"
	"fund_screen/db/models/system"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	tx *gorm.DB
)

// NearestTradingDay 最接近的交易日
func NearestTradingDay(ctx context.Context, date string) carbon.ToDateString {
	var calendar system.TradingCalendar
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&calendar).Where("date <= ?", date).Last(&calendar)
	return calendar.Date
}

// NearestFriday 上周最后一个交易日，一般为周五
func NearestFriday(ctx context.Context, date carbon.Carbon) carbon.ToDateString {
	var calendar system.TradingCalendar
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&calendar).Where("date < ? and week_end = 1", date.Format("Y-m-d")).Last(&calendar)
	return calendar.Date
}

// NearestTradingDayAfter 最接近的交易日
func NearestTradingDayAfter(ctx context.Context, date string) string {
	var calendar carbon.ToDateString
	db = connector.GetDB()
	tx = db.WithContext(ctx)
	tx.Model(&system.TradingCalendar{}).Select("min(date)").Where("date >= ?", date).Find(&calendar)
	return calendar.ToDateString()
}
