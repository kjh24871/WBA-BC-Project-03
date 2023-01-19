package model

import (
	"lecture/WBA-BC-Project-03/backend/protocol"

	"github.com/ethereum/go-ethereum/common"

	// "context"
	// "crypto/ecdsa"
	// "encoding/hex"
	// "fmt"
	"math/big"

	// tokenContracts "lecture/WBA-BC-Project-03/backend/model/wemex/ERC20"
	// contracts "lecture/WBA-BC-Project-03/backend/model/wemex/liquidity"
	liquidityFactoryContracts "lecture/WBA-BC-Project-03/backend/model/wemex/liquidityFactory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rlp"
	// "golang.org/x/crypto/sha3"
)

func (p *Model) GetAddressWithName(name string) common.Address {
	contracts, err := liquidityFactoryContracts.NewLiquidityFactory(p.liquidityFactoryAddresss, p.client)
	if err != nil {
		panic(err)
	}
	contractAddress, err := contracts.GetAddressWithName(&bind.CallOpts{}, name)
	if err != nil {
		panic(err)
	}
	return contractAddress
}

func (p *Model) NewLiquidityPool(inputToken1Address string, inputToken2Address string, ownerPK string) *protocol.ApiResponse[any] {
	contracts, err := liquidityFactoryContracts.NewLiquidityFactory(p.liquidityFactoryAddresss, p.client)
	if err != nil {
		return protocol.Fail(err, 500)
	}
	privateKey, err := crypto.HexToECDSA(ownerPK)
	if err != nil {
		return protocol.Fail(err, 501)
	}
	transactorOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1112))
	if err != nil {
		return protocol.Fail(err, 502)
	}
	token1Address := common.HexToAddress(inputToken1Address)
	token2Address := common.HexToAddress(inputToken2Address)
	tx, err := contracts.NewLiquidity(transactorOpts, token1Address, token2Address)
	if err != nil {
		return protocol.Fail(err, 503)
	}
	return protocol.SuccessData(tx.Hash().Hex(), protocol.OK)
}
