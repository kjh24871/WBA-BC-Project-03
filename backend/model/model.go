package model

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	contracts "final/backend/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Model struct {
	client       *ethclient.Client
	tokenAddress common.Address //컨트랙트 어드레스
	ownerAddress common.Address //내 지갑 주소
}

func NewModel() (*Model, error) {
	r := &Model{}
	var err error
	r.client, err = ethclient.Dial("https://api.test.wemix.com")
	if err != nil {
		fmt.Println("client error")
	}
	r.tokenAddress = common.HexToAddress("0xE42f6266a0E88b4160413d2c182D25F3a74F5472")
	r.ownerAddress = common.HexToAddress("0x15A67B0bB392b2978bFeDBC67809A562d7045767")
	return r, err
}

func (p *Model) BlockchainCreateTransferTx(pk string, dstAddress string, amount int64) string {
	// metamask에서 뽑아낸 privatekey를 변환
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Println(err)
	}

	// privatekey로부터 publickey를 거쳐 자신의 address 변환
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("fail convert, publickey")
	}
	// 보낼 address 설정
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 현재 계정의 nonce를 가져옴. 다음 트랜잭션에서 사용할 nonce
	nonce, err := p.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	// 전송할 양, gasLimit, gasPrice 설정. 추천되는 gasPrice를 가져옴
	value := big.NewInt(amount)
	gasLimit := uint64(21000)
	gasPrice, err := p.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// 전송받을 상대방 address 설정
	toAddress := common.HexToAddress(dstAddress)
	// 트랜잭션 생성
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := p.client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// 트랜잭션 서명
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
	}

	// RLP 인코딩 전 트랜잭션 묶음. 현재는 1개의 트랜잭션
	ts := types.Transactions{signedTx}
	// RLP 인코딩
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
	rawTxHex := hex.EncodeToString(rawTxBytes)
	rTxBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {
		fmt.Println(err.Error())
	}

	// RLP 디코딩
	rlp.DecodeBytes(rTxBytes, &tx)
	// 트랜잭션 전송
	err = p.client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Println(err)
	}
	//출력된 tx.hash를 익스플로러에 조회 가능
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (p *Model) ContractCreateTransferTx(pk string, dstAddress string, amount int64) string {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Println(err)
	}

	// privatekey로부터 publickey를 거쳐 자신의 address 변환
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("fail convert, publickey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 현재 계정의 nonce를 가져옴. 다음 트랜잭션에서 사용할 nonce
	nonce, err := p.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}
	// 전송할 양, gasLimit, gasPrice 설정. 추천되는 gasPrice를 가져옴
	value := big.NewInt(amount * 1000000000000000000)
	// value := big.NewInt(700000000000000000)
	gasPrice, err := p.client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// 보낼 주소
	toAddress := common.HexToAddress(dstAddress)

	// 컨트랙트 전송시 사용할 함수명
	transferFnSignature := []byte("balanceOf(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	zvalue := big.NewInt(0)
	//컨트랙트 전송 정보 입력
	var pdata []byte
	pdata = append(pdata, methodID...)
	pdata = append(pdata, paddedAddress...)
	pdata = append(pdata, paddedAmount...)

	gasLimit := uint64(200000)
	fmt.Println(gasLimit)

	// 트랜잭션 생성
	tx := types.NewTransaction(nonce, p.tokenAddress, zvalue, gasLimit, gasPrice, pdata)
	chainID, err := p.client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// 트랜잭션 서명
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
	}

	// 트랜잭션 전송
	err = p.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}

	//tx.hash를 이용해 전송결과를 확인
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex()
}

func (p *Model) GetSymbolByToken(token string) string {
	instance, err := contracts.NewContracts(p.tokenAddress, p.client)
	if err != nil {
		panic(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	if name == token {
		symbol, err := instance.Symbol(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		return symbol
	}
	return "no symbol"
}

func (p *Model) GetTokenWithAddress(address string) *big.Int {
	instance, err := contracts.NewContracts(p.tokenAddress, p.client)
	if err != nil {
		panic(err)
	}
	dstAddress := common.HexToAddress(address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, dstAddress)
	if err != nil {
		panic(err)
	}
	return bal
}

func (p *Model) TransferCoinWithAddress(address string, value int64) string {
	return p.BlockchainCreateTransferTx("유저의 pk", address, value)
}

func (p *Model) TransferCoinWithPK(address string, pk string, value int64) string {
	return p.BlockchainCreateTransferTx(pk, address, value)
}
func (p *Model) TransferTokenWithAddress(address string, value int64) string {
	return p.ContractCreateTransferTx("유저의 pk", address, value)
}

func (p *Model) TransferTokenWithPK(address string, pk string, value int64) string {
	return p.ContractCreateTransferTx(pk, address, value)
}
