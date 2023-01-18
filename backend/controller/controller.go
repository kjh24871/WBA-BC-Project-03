package controller

import (
	"final/backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) GetAllPools(c *gin.Context) {

}

func (p *Controller) GetMyPools(c *gin.Context) {

}

func (p *Controller) CreateLiquidity(c *gin.Context) {

}

func (p *Controller) BalanceOf(c *gin.Context) {
	address := c.Query("address")
	// c.JSON(200, p.md.BalanceOf(address)) ==
	// protocl.Success(200).Response(c)
	c.JSON(200,p.md.BalanceOf(address))
}

func (p *Controller) AddLiquidity(c *gin.Context) {
	amountA := c.Query("amountA")
	amountB := c.Query("amountB")
	intAmountA, _ := strconv.ParseInt(amountA, 10, 64)
	intAmountB, _ := strconv.ParseInt(amountB, 10, 64)
	c.JSON(200, p.md.AddLiquidity(intAmountA, intAmountB))
}

// func (p *Controller) RemoveLiquidity(c *gin.Context) {
// 	amount := c.Query("amount")
// 	intAmount, _ := strconv.ParseInt(amount, 10, 64)
// 	c.JSON(200, p.md.AddLiquidity(intAmount))
// }

func (p *Controller) Swap(c *gin.Context) {
	amount := c.Query("amount")
	intAmount, _ := strconv.ParseInt(amount, 10, 64)
	c.JSON(200, p.md.SwapLiquidity("0xB451E9Fd9114611e88B257D04A62D22a86FFA1c2", intAmount))
}
