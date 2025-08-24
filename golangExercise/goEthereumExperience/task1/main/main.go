package main

import (
	"crypto/ecdsa"
	"log"
	"zrjBlockChainExercise/golangExercise/goEthereumExperience/task1"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClientAndPrivateKey() (*ethclient.Client, *ecdsa.PrivateKey) {
	client, err := ethclient.Dial("xx")
	privateKey, err := crypto.HexToECDSA("xx")
	if err != nil {
		log.Fatal(err)
	}
	return client, privateKey

}
func main1() {
	client, privateKey := getClientAndPrivateKey()
	task1.QueryBlock(client)
	task1.Transaction(client, privateKey)
}

func main() {
	client, privateKey := getClientAndPrivateKey()
	task1.AbigenTest(client, privateKey)
}
