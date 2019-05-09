package sdksource

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func TransferETH(rootDir, node, fromName, password, toAddr string, amount int64, gasLimit uint64) string {
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//get the nonce from the fromAddress to be dumped into tx
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//amount convertion to wei
	value := big.NewInt(amount)

	//gaslimit
	//gasLimit := uint64(21000)

	//get the estimated gasprice with SuggestGasPrice func
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//concert the to Address to byte format
	toAddress := common.HexToAddress(toAddr)


	//Generate the Tx body, the data field is nil for just sending ETH
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//sign the Tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//SendTransaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()

}

//Transfer with ERC20 token
//func TransferERC20(rootDir, node, fromName, password, toAddr string, amount int64, gasLimit uint64) string {
//
//}