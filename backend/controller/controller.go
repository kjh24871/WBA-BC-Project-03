package controller

import (
	"lecture/WBA-BC-Project-03/backend/model"
	"lecture/WBA-BC-Project-03/backend/protocol"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

// GetLiquidityList godoc
// @Summary call GetLiquidityList, return ok by json.
// @현재 풀 목록 조회
// @name GetLiquidityList
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Router /all [GET]
// @Success 200 {object} Controller
func (p *Controller) GetLiquidityList(c *gin.Context) {
	p.md.GetLiquidityList().Response(c)
}

func (p *Controller) GetMyPools(c *gin.Context) {
	// pk := c.GetHeader("privateKey")
}

// CreateLiquidity godoc
// @Summary call CreateLiquidity, return ok by json.
// @새로운 풀 추가
// @name CreateLiquidity
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Param address body CreateLiquidityInput true "Token address"
// @Router /pool [POST]
// @Success 200 {object} Controller
func (p *Controller) CreateLiquidity(c *gin.Context) {
	pk := c.GetHeader("pk")
	var input CreateLiquidityInput
	c.BindJSON(&input)
	p.md.NewLiquidityPool(input.Token1, input.Token2, pk).Response(c)
}

type CreateLiquidityInput struct {
	Token1 string `json:"token1"`
	Token2 string `json:"token2"`
}

// BalanceOf godoc
// @Summary call BalanceOf, return ok by json.
// @풀 지분 확인
// @name BalanceOf
// @Accept  json
// @Produce  json
// @Param name query string true "pool name"
// @Param address query string true "my address"
// @Router /pool/balance [GET]
// @Success 200 {object} Controller
func (p *Controller) BalanceOf(c *gin.Context) {
	name := c.Query("name")
	address := c.Query("address")
	p.md.BalanceOf(name, address).Response(c)
}

// AddLiquidity godoc
// @Summary call AddLiquidity, return ok by json.
// @유동성 추가
// @name AddLiquidity
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Param input body AddLiquidityInput true "Token address and amount"
// @Router /pool [PUT]
// @Success 200 {object} Controller
func (p *Controller) AddLiquidity(c *gin.Context) {
	pk := c.GetHeader("pk")
	var input AddLiquidityInput
	c.BindJSON(&input)
	intAmountA, _ := strconv.ParseInt(input.AmountA, 10, 64)
	intAmountB, _ := strconv.ParseInt(input.AmountB, 10, 64)
	c.JSON(200, p.md.AddLiquidity(intAmountA, intAmountB, pk, input.SymA, input.SymB))
}

type AddLiquidityInput struct {
	AmountA string `json:"amount_a"`
	AmountB string `json:"amount_b"`
	SymA    string `json:"sym_a"`
	SymB    string `json:"sym_b"`
}

// RemoveLiquidity godoc
// @Summary call RemoveLiquidity, return ok by json.
// @유동성 제거
// @name RemoveLiquidity
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Param input body RemoveLiquidityInput true "Token address and amount"
// @Router /pool [DELETE]
// @Success 200 {object} Controller
func (p *Controller) RemoveLiquidity(c *gin.Context) {
	var input RemoveLiquidityInput
	pk := c.GetHeader("pk")
	c.BindJSON(&input)
	intAmount, _ := strconv.ParseInt(input.Amount, 10, 64)
	p.md.RemoveLiquidity(input.PoolName, intAmount, pk).Response(c)
}

type RemoveLiquidityInput struct {
	Amount   string `json:"amount"`
	PoolName string `json:"name"`
}

// Swap godoc
// @Summary call Swap, return ok by json.
// @교환하기
// @name Swap
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Param input body SwapInput true "Token symbol and amount"
// @Router /swap [POST]
// @Success 200 {object} Controller
func (p *Controller) Swap(c *gin.Context) {
	var input SwapInput
	c.BindJSON(&input)
	address := p.md.GetAddressWithName(input.Input + "-" + input.Output)
	if address == common.HexToAddress("0x0") {
		address = p.md.GetAddressWithName(input.Output + "-" + input.Input)
		if address == common.HexToAddress("0x0") {
			protocol.Fail(nil, 500)
		}
	}
	pk := c.GetHeader("pk")
	intAmount, _ := strconv.ParseInt(input.Amount, 10, 64)
	p.md.SwapLiquidity(address, input.Input, intAmount, pk).Response(c)
}

type SwapInput struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Amount string `json:"amount"`
}

// func (p *Controller) CheckTx(c *gin.Context) {
// 	pk := c.GetHeader("pk")
// 	tx := c.Query("tx")
// }

// PullReward godoc
// @Summary call PullReward, return ok by json.
// @교환하기
// @name PullReward
// @Accept  json
// @Produce  json
// @Param pk header string true "private key"
// @Router /reward [GET]
// @Success 200 {object} Controller
// func (p *Controller) GetReward(c *gin.Context) {
// 	pk := c.GetHeader("pk")
// 	p.md.GetReward(pk).Response(c)
// }
