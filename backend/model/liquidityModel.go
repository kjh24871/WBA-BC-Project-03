package model

import (
	// "context"
	// "crypto/ecdsa"
	// "encoding/hex"
	"fmt"
	"math/big"

	contracts "final/backend/contract2"

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
	contracts, err := contracts.NewContracts(p.liquidityAddress, p.client)
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
		panic(eff)
	}

	fmt.Println("Transaction sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (p *Model)AddLiquidity(amount int64) string {
	contracts, err := contracts.NewContracts(p.liquidityAddress, p.client)
	if err != nil {
		panic(err)
	}
	maxToken := big.NewInt(amount * 1000000000000000000)

	privateKey, err := crypto.HexToECDSA("f7b0033d5c91b7258b2557a66b1743195ffd77fc285b4cbba2ecd3f94d9c5939")
	if err != nil {
		fmt.Println(err)
	}
	
	transactorOpts,err := bind.NewKeyedTransactorWithChainID(privateKey,big.NewInt(1112))
	if err != nil {
		panic(err)
	}

	tx, err := contracts.AddLiquidity(transactorOpts, maxToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (p *Model)BalanceOf(strAddress string) *big.Int{
	instance, err := contracts.NewContracts(p.liquidityAddress, p.client)
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

// func (p *Model) Approve(strAddress string, amount int64) {
// 	instance, err := contracts.NewContracts(p.liquidityAddress, p.client)
// 	if err != nil {
// 		panic(err)
// 	}
// 	address := common.HexToAddress(strAddress)
// 	instance.Appro 
// }