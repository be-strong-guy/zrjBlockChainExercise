package task1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"zrjBlockChainExercise/golangExercise/goEthereumExperience/task1/count"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**
任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
 具体任务
编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/

func AbigenTest(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {

	countContract, err := count.NewCount(common.HexToAddress("0xa0a0459e77B3F2BC419f17bba4a43C613b01f8c9"), client)
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := countContract.PlusOne(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := countContract.I(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("目前计数器的值是:", valueInContract)
}
