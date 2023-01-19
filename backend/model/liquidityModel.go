package model

import (
	"lecture/WBA-BC-Project-03/backend/protocol"
	// "context"
	// "crypto/ecdsa"
	// "encoding/hex"

	"math/big"

	tokenContracts "lecture/WBA-BC-Project-03/backend/model/wemex/ERC20"
	contracts "lecture/WBA-BC-Project-03/backend/model/wemex/liquidity"
	liquidityFactoryContracts "lecture/WBA-BC-Project-03/backend/model/wemex/liquidityFactory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rlp"
	// "golang.org/x/crypto/sha3"
)

func (p *Model) GetLiquidityList() *protocol.ApiResponse[any] {
	contracts, err := liquidityFactoryContracts.NewLiquidityFactory(p.liquidityFactoryAddresss, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	list, err := contracts.GetLiquidityList(&bind.CallOpts{})
	if err != nil {
		return protocol.Fail(err, 500)
	}
	return protocol.SuccessData(list, protocol.OK)
}

func (p *Model) SwapLiquidity(poolAddress common.Address, input string, amount int64, pk string) *protocol.ApiResponse[any] {
	contracts, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	tokenAAddress, tokenBAddress, err := contracts.GetTokenAddress(&bind.CallOpts{})
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tokenAContract, err := tokenContracts.NewERC20(tokenAAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	inputTokenAddress := tokenAAddress
	tokenASymbol, err := tokenAContract.Symbol(&bind.CallOpts{})
	if err != nil {
		return protocol.Fail(err, 500)
	}
	if tokenASymbol != input {
		inputTokenAddress = tokenBAddress
	}
	inputAmount := big.NewInt(amount)

	inputTokenContract, err := tokenContracts.NewERC20(inputTokenAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	transactorOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1112))
	if err != nil {
		return protocol.Fail(err, 500)
	}
	_, err = inputTokenContract.Approve(transactorOpts, poolAddress, big.NewInt(amount))
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tx, err := contracts.Swap(transactorOpts, inputTokenAddress, inputAmount)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	return protocol.SuccessData(tx.Hash().Hex(), protocol.OK)
}

func (p *Model) AddLiquidity(_desiredAmountA int64, _desiredAmountB int64, pk string, symA string, symB string) *protocol.ApiResponse[any] {

	//  심볼 2개 합쳐서 pool 이름 구하기 -> pool 이름을 주소 구하기 -> 두개 token Address도 나온다
	// 심볼1-심볼2
	poolContracts, err := liquidityFactoryContracts.NewLiquidityFactory(p.liquidityFactoryAddresss, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	poolAddress, err := poolContracts.GetAddressWithName(&bind.CallOpts{}, symA+"-"+symB)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	contracts, err := contracts.NewLiquidity(poolAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tokenAAddress, tokenBAddress, err := contracts.GetTokenAddress(&bind.CallOpts{})
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tokenAContracts, err := tokenContracts.NewERC20(tokenAAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tokenBContracts, err := tokenContracts.NewERC20(tokenBAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	inputAmountA := big.NewInt(_desiredAmountA)
	inputAmountB := big.NewInt(_desiredAmountB)
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	transactorOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1112))
	if err != nil {
		return protocol.Fail(err, 500)
	}
	// TODO : Approve 2 토큰하기
	_, err = tokenAContracts.Approve(transactorOpts, poolAddress, inputAmountA)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	_, err = tokenBContracts.Approve(transactorOpts, poolAddress, inputAmountB)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	tx, err := contracts.AddLiquidity(transactorOpts, inputAmountA, inputAmountB)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	return protocol.SuccessData(tx.Hash().Hex(), protocol.OK)
}

func (p *Model) BalanceOf(poolName string, strAddress string) *protocol.ApiResponse[any] {

	liquidityFactoryContracts, err := liquidityFactoryContracts.NewLiquidityFactory(p.liquidityFactoryAddresss, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	poolAddress, err := liquidityFactoryContracts.GetAddressWithName(&bind.CallOpts{},poolName)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	instance, err := contracts.NewLiquidity(poolAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	address := common.HexToAddress(strAddress)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	return protocol.SuccessData(balance, protocol.OK)
}

func (p *Model) RemoveLiquidity(lpAmount int64, pk string) *protocol.ApiResponse[any] {
	contracts, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	inputAmount := big.NewInt(lpAmount)

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return protocol.Fail(err, 500)
	}

	transactorOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1112))
	if err != nil {
		return protocol.Fail(err, 500)
	}

	tx, err := contracts.RemoveLiquidity(transactorOpts, inputAmount)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	return protocol.SuccessData(tx.Hash().Hex(), protocol.OK)
}
