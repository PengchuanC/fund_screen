package routers

import (
	"fund_screen/app/services"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, middlewares ...gin.HandlerFunc){
	for _, middleware := range middlewares {
		r.Use(middleware)
	}

	v2 := r.Group("/api/v2")

	search := v2.Group("/search")
	{
		search.GET("/fundList", services.SearchFundList)
	}

	fundinfo := v2.Group("/fundinfo")
	{
		fundinfo.GET("", services.LatestPerformance)
		fundinfo.GET("/performance", services.Performance)
		fundinfo.GET("/turnover", services.Turnover)
		fundinfo.GET("/holder", services.HolderStruct)
		fundinfo.GET("/scale", services.ScaleChange)
		fundinfo.GET("/cycle", services.MacroCyclePerformance)
		fundinfo.GET("his_style", services.HistoryStyle)
		fundinfo.GET("/risk/short", services.RiskShortPerformance)
		fundinfo.GET("/risk/long", services.RiskLongPerformance)
		fundinfo.GET("/style", services.RBSAStyle)
		fundinfo.GET("/style&benchmark", services.StyleAndBenchmarkIndex)
		fundinfo.GET("/nav", services.NetValueSeries)
		fundinfo.GET("/nav/recent", services.RecentPerformance)
		fundinfo.GET("/nav/yearly", services.YearlyPerformance)
	}

	manager := v2.Group("/manager")
	{
		manager.GET("", services.ManagerInfo)
		manager.GET("/managed", services.Managed)
		manager.GET("/topHolding", services.TopHolding)
	}

	fundPortfolio := v2.Group("/asset")
	{
		fundPortfolio.GET("", services.Allocate)
		fundPortfolio.GET("/history", services.AllocateHistory)
		fundPortfolio.GET("/keyStock", services.KeyStock)
		fundPortfolio.GET("/keyBond", services.KeyBond)
		fundPortfolio.GET("/industryCsi", services.IndustryCsi)
		fundPortfolio.GET("/concentrate", services.Concentrate)
		fundPortfolio.GET("/style", services.IndustryStyle)
	}
}