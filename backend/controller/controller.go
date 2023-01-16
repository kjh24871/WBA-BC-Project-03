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

func (p *Controller) GetSymbolByToken(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, p.md.GetSymbolByToken(name))
}

func (p *Controller) GetAddressToken(c *gin.Context) {
	address := c.Query("address")
	c.JSON(200, p.md.GetTokenWithAddress(address))
}

func (p *Controller) TransferCoinWithAddress(c *gin.Context) {
	address := c.Query("address")
	value := c.Query("value")
	intValue, _ := strconv.ParseInt(value, 10, 64)
	c.JSON(200, p.md.TransferCoinWithAddress(address, intValue))
}

func (p *Controller) TransferCoinWithPK(c *gin.Context) {
	address := c.Query("address")
	pk := c.Query("pk")
	value := c.Query("value")
	intValue, _ := strconv.ParseInt(value, 10, 64)
	c.JSON(200, p.md.TransferCoinWithPK(address, pk, intValue))
}

func (p *Controller) TransferTokenWithAddress(c *gin.Context) {
	address := c.Query("address")
	value := c.Query("value")
	intValue, _ := strconv.ParseInt(value, 10, 64)
	c.JSON(200, p.md.TransferTokenWithAddress(address, intValue))
}

func (p *Controller) TransferTokenWithPK(c *gin.Context) {
	address := c.Query("address")
	pk := c.Query("pk")
	value := c.Query("value")
	intValue, _ := strconv.ParseInt(value, 10, 64)
	c.JSON(200, p.md.TransferTokenWithPK(address, pk, intValue))
}
