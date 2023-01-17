package controller

import (
	// "final/backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *Controller) AddLiquidity(c *gin.Context){
	amount := c.Query("amount")
	intAmount, _ := strconv.ParseInt(amount, 10, 64)
	c.JSON(200, p.md.AddLiquidity(intAmount))
}

func (p *Controller) BalanceOf(c *gin.Context){
	address := c.Query("address")
	c.JSON(200, p.md.BalanceOf(address))
}