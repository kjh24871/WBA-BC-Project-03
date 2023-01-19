package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client      *mongo.Client
	colBlock    *mongo.Collection
	colContract *mongo.Collection
}

type Block struct {
	BlockHash    string        `bson:"blockHash"`
	BlockNumber  uint64        `bson:"blockNumber"`
	GasLimit     uint64        `bson:"gasLimit"`
	GasUsed      uint64        `bson:"gasUsed"`
	Time         uint64        `bson:"timestamp"`
	Nonce        uint64        `bson:"nonce"`
	Transactions []Transaction `bson:"transactions"`
}

type Transaction struct {
	TxHash      string `bson:"hash"`
	From        string `bson:"from"`
	To          string `bson:"to"` // 컨트랙트의 경우 nil 반환
	Nonce       uint64 `bson:"nonce"`
	GasPrice    uint64 `bson:"gasPrice"`
	GasLimit    uint64 `bson:"gasLimit"`
	Amount      uint64 `bson:"amount"`
	BlockHash   string `bson:"blockHash"`
	BlockNumber uint64 `bson:"blockNumber"`
}

type Pool struct {
	PoolAddress string `bson:"address"`
}

func NewModel(mgUrl string) (*Model, error) {
	r := &Model{}

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("daemon")
		r.colBlock = db.Collection("block")
		r.colContract = db.Collection("pool")
	}
	return r, nil
}

func (p *Model) CheckPool() []Pool {
	cursor, err := p.colContract.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var pools []Pool
	if err = cursor.All(context.TODO(), &pools); err != nil {
		panic(err)
	}
	return pools
}

func (p *Model) SaveBlock(block *Block) error {
	result, err := p.colBlock.InsertOne(context.TODO(), block)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Insert succeed: %s\n", result.InsertedID)
	return nil
}
