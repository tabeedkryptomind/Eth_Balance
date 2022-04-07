package main

import "kryptomind-web3-service/app/routes"

//"github.com/chenzhijie/go-web3"

// "context"
// "fmt"
// "log"
// "regexp"
// "github.com/ethereum/go-ethereum/common"
// "github.com/ethereum/go-ethereum/ethclient"
// "github.com/umbracle/ethgo/jsonrpc"
func main() {
	routes.StartApp()
	
}
// client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/83db17056614454792c715a1c4083d8d")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	addr :=  "0x984AF56d8E54bac6AeB1d5F8b941c256E03CFcF2"
// 	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
// 	if re.MatchString(addr) {
// 		account := common.HexToAddress(addr)
// 		balance , err := client.BalanceAt(context.Background(), account,nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(balance)
// 	}

// var rpcProvider = "wss://mainnet.infura.io/ws/v3/83db17056614454792c715a1c4083d8d"
	// web3, err := web3.NewWeb3(rpcProvider)
	// if err != nil {
	// 	panic(err)
	// }
	// blockNumber, err := web3.Eth.GetBlockNumber()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Current block number: ", blockNumber)
	// fee, err := web3.Eth.EstimateFee()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(fee.BaseFee)
	// client, err := jsonrpc.NewClient("wss://mainnet.infura.io/ws/v3/83db17056614454792c715a1c4083d8d")
	// if err != nil {
	// 	panic(err)
	// }
	// eth := client.Eth()
	// balance , err :=  eth.GetBalance("0xcD818F4F74A8170DA28165127C05AA9fD2eF75F3")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(balance)