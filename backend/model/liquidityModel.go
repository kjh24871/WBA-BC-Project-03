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

	// privatekey로부터 publickey를 거쳐 자신의 address 변환
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	fmt.Println("fail convert, publickey")
	// }
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	
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

// func (p *Model)AddLiquidity(amount int64) string {
// 	privateKey, err := crypto.HexToECDSA("f7b0033d5c91b7258b2557a66b1743195ffd77fc285b4cbba2ecd3f94d9c5939")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// privatekey로부터 publickey를 거쳐 자신의 address 변환
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		fmt.Println("fail convert, publickey")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	// 현재 계정의 nonce를 가져옴. 다음 트랜잭션에서 사용할 nonce
// 	nonce, err := p.client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// 전송할 양, gasLimit, gasPrice 설정. 추천되는 gasPrice를 가져옴
// 	value := big.NewInt(amount * 1000000000000000000)
// 	fmt.Println("value : " ,value)
// 	// value := big.NewInt(700000000000000000)
// 	gasPrice, err := p.client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 보낼 주소
// 	// toAddress := common.HexToAddress(dstAddress)

// 	// 컨트랙트 전송시 사용할 함수명
// 	transferFnSignature := []byte("addLiquidity(uint256)")
// 	hash := sha3.NewLegacyKeccak256()
// 	hash.Write(transferFnSignature)
// 	methodID := hash.Sum(nil)[:4]
// 	fmt.Println(hexutil.Encode(methodID))

// 	// paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
// 	// fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

// 	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
// 	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
// 	zvalue := big.NewInt(0)
// 	//컨트랙트 전송 정보 입력
// 	var pdata []byte
// 	pdata = append(pdata, methodID...)
// 	// pdata = append(pdata, paddedAddress...)
// 	pdata = append(pdata, paddedAmount...)

// 	gasLimit := uint64(200000)
// 	fmt.Println(gasLimit)

// 	// 트랜잭션 생성
// 	tx := types.NewTransaction(nonce, p.liquidityAddress, zvalue, gasLimit, gasPrice, pdata)
// 	chainID, err := p.client.NetworkID(context.Background())
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 트랜잭션 서명
// 	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 트랜잭션 전송
// 	err = p.client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//tx.hash를 이용해 전송결과를 확인
// 	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
// 	return signedTx.Hash().Hex()
// }

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