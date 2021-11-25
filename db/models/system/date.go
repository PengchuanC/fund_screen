package system

import "github.com/golang-module/carbon"

type TradingCalendar struct {
	Id int64
	Date carbon.ToDateString
	Market int
	WeekEnd int
	MonthEnd int
	QuarterEnd int
	YearEnd int
	Jsid int64
}

func (t TradingCalendar) TableName() string {
	return "sc_system_trading_calendar"
}