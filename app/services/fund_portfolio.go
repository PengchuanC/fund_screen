package services

import (
	"fund_screen/db/dao/funds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Allocate(c *gin.Context)  {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundAllocate(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func AllocateHistory(c *gin.Context)  {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundAllocateHistory(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func KeyStock(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
	return
	}
	resp := funds.FundKeyStock(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func KeyBond(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundKeyBond(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func IndustryCsi(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundIndustry(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func Concentrate(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundConcentrate(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}

func IndustryStyle(c *gin.Context) {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	resp := funds.FundIndustryStyle(c, request.Secucode)
	c.JSON(http.StatusOK, resp)
}