package services

import (
	"fund_screen/db/dao/funds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchFundList(c *gin.Context)  {
	var request SearchRequest
	err := c.ShouldBind(&request)
	if err != nil {
		return
	}
	ret := funds.FundSearchFundList(c, request.Search)
	c.JSON(http.StatusOK, ret)
}
