package sdksource

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go/token"
	"log"
	"math"
	"math/big"
)

func GetAccount(node, addr string) string{
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//convert the addr string to common.Address type
	address := common.HexToAddress(addr)

	//get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(header.Number.Int64())

	balance, err := client.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	accountStr := ethValue.String() + "ETH"
	return accountStr
}

func GetAccountERC20(node, addr string) string {
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//ERC20 Token QT Address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}


	//convert the addr string to common.Address type
	address := common.HexToAddress(addr)

	//get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(header.Number.Int64())

	balance, err := client.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	accountStr := ethValue.String() + "ETH"
	return accountStr
}