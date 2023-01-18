package model

import (
	// "context"
	// "crypto/ecdsa"
	// "encoding/hex"
	"fmt"
	"math/big"

	contracts "final/backend/model/wemex/liquidity"
	tokenContracts "final/backend/model/wemex/liquidityFactory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rlp"
	// "golang.org/x/crypto/sha3"
)

func (p *Model)SwapLiquidity(inputTokenAddress string, amount int64 ) string {
	contracts, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	if err != nil {
		panic(err)
	}
	tokenAddress := common.HexToAddress(inputTokenAddress)
	inputAmount := big.NewInt(amount)

	privateKey, err := crypto.HexToECDSA("f7b0033d5c91b7258b2557a66b1743195ffd77fc285b4cbba2ecd3f94d9c5939")
	if err != nil {
		panic(err)
	}

	transactorOpts,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1112))
	if err != nil {
		panic(err)
	}
	tx, err := contracts.Swap(transactorOpts, tokenAddress, inputAmount)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (p *Model)AddLiquidity(_desiredAmountA int64, _desiredAmountB int64 ) string {
	contracts, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	token1Contracts, err := tokenContracts.NewLiquidity(p.tokenAddress, p.client)
	token2Contracts, err := tokenContracts.NewLiquidity(p.token2Address, p.client)
	if err != nil {
		panic(err)
	}
	inputAmountA := big.NewInt(_desiredAmountA)
	inputAmountB := big.NewInt(_desiredAmountB)
	privateKey, err := crypto.HexToECDSA("f7b0033d5c91b7258b2557a66b1743195ffd77fc285b4cbba2ecd3f94d9c5939")
	if err != nil {
		fmt.Println(err)
	}
	
	transactorOpts,err := bind.NewKeyedTransactorWithChainID(privateKey,big.NewInt(1112))
	if err != nil {
		panic(err)
	}
	// TODO : Approve 2 토큰하기
	tx, err := token1Contracts.Approve(transactorOpts, p.liquidityAddress, inputAmountA)
	if err != nil {
		panic(err)
	}
	tx, err = token2Contracts.Approve(transactorOpts, p.liquidityAddress, inputAmountB)
	if err != nil {
		panic(err)
	}
	tx, err = contracts.AddLiquidity(transactorOpts, inputAmountA, inputAmountB)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (p *Model)BalanceOf(strAddress string) *big.Int{
	instance, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress(strAddress)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		panic(err)
	}
	return balance
}

func (p *Model) RemoveLiquidity(lpAmount int64) string{
	contracts, err := contracts.NewLiquidity(p.liquidityAddress, p.client)
	inputAmount := big.NewInt(lpAmount)

	privateKey, err := crypto.HexToECDSA("f7b0033d5c91b7258b2557a66b1743195ffd77fc285b4cbba2ecd3f94d9c5939")
	if err != nil {
		fmt.Println(err)
	}
	
	transactorOpts,err := bind.NewKeyedTransactorWithChainID(privateKey,big.NewInt(1112))
	if err != nil {
		panic(err)
	}

	tx, err := contracts.RemoveLiquidity(transactorOpts, inputAmount)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}