package services

import (
	"fund_screen/db/dao/dervative"
	"fund_screen/db/dao/manager"
	"github.com/gin-gonic/gin"
	"net/http"
)


// TopHolding 基金经理任职以来持仓前20的股票
func TopHolding(c *gin.Context){
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	ret := dervative.FundTopHolding(c, request.Secucode)
	c.JSON(http.StatusOK, ret)
}

// Managed 同一基金经理在管基金
func Managed(c *gin.Context)  {
	var request Request
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	ret := manager.Managed(c, request.Secucode)
	c.JSON(http.StatusOK, ret)
}