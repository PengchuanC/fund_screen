package services

import (
	"fund_screen/db/dao/dervative"
	funds2 "fund_screen/db/dao/funds"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LatestPerformance 基金最新业绩表现
func LatestPerformance(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds2.FundLatestPerformance(c, request.Secucode)
	c.JSON(200, resp)
}

// ManagerInfo 基金经理信息
func ManagerInfo(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds2.FundManagerInfo(c, request.Secucode)
	c.JSON(200, resp)
}

// Performance 区间业绩评价
func Performance(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	var performance = funds2.FundPerformance(c, request.Secucode)

	c.JSON(200, performance)
}

// Turnover 换手率
func Turnover(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	var resp = funds2.FundTurnover(c, request.Secucode)
	c.JSON(200, resp)
}

// HolderStruct 持有人结构
func HolderStruct(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	var resp = funds2.FundHolderStruct(c, request.Secucode)
	c.JSON(200, resp)
}

// ScaleChange 规模变化
func ScaleChange(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	var resp = funds2.FundScaleChange(c, request.Secucode)
	c.JSON(200, resp)
}

// MacroCyclePerformance 周期表现
func MacroCyclePerformance(c *gin.Context){
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	var resp = dervative.FundMacroCyclePerformance(c, request.Secucode)
	c.JSON(200, resp)
}

// HistoryStyle 历史风格变化
func HistoryStyle(c *gin.Context)  {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := dervative.HistoryStyle(c, request.Secucode)
	c.JSON(200, resp)
}

// RiskShortPerformance 风险指标分位数
func RiskShortPerformance(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds2.FundShortRisk(c, request.Secucode)
	c.JSON(200, resp)
}

// RiskLongPerformance 风险指标分位数
func RiskLongPerformance(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds2.FundLongRisk(c, request.Secucode)
	c.JSON(200, resp)
}

// RBSAStyle 基金风格分析
func RBSAStyle(c *gin.Context)  {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := dervative.RBSAStyle(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

// StyleAndBenchmarkIndex 风格指数与宽基指数
func StyleAndBenchmarkIndex(c *gin.Context){
	resp := funds2.StyleAndBenchmarkIndex(c)
	c.JSON(http.StatusOK, resp)
}

// NetValueSeries 基金单位资产净值
func NetValueSeries(c *gin.Context) {
	var r NetValueRequest
	_ = c.ShouldBind(&r)
	if r.Style == "" {
		r.Style = "H11020"
	}
	if r.Benchmark== "" {
		r.Benchmark = "000300"
	}
	resp := funds2.NetValueSeries(c, r.Secucode, r.Style, r.Benchmark)
	c.JSON(http.StatusOK, resp)
}

// RecentPerformance 基金及指数近期业绩表现
func RecentPerformance(c *gin.Context) {
	var r NetValueRequest
	_ = c.ShouldBind(&r)
	if r.Style == "" {
		r.Style = "H11020"
	}
	if r.Benchmark== "" {
		r.Benchmark = "000300"
	}
	resp := funds2.RecentPerformance(c, r.Secucode, r.Style, r.Benchmark)
	c.JSON(http.StatusOK, resp)
}

// YearlyPerformance 基金及指数分年度业绩表现
func YearlyPerformance(c *gin.Context) {
	var r NetValueRequest
	_ = c.ShouldBind(&r)
	if r.Style == "" {
		r.Style = "H11020"
	}
	if r.Benchmark== "" {
		r.Benchmark = "000300"
	}
	resp := funds2.YearlyPerformance(c, r.Secucode, r.Style, r.Benchmark)
	c.JSON(http.StatusOK, resp)
}