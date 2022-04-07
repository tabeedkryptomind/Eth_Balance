package services

import (
	"context"
	"math/big"
	"regexp"
	er "kryptomind-web3-service/app/utils/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BalanceDetails(addr string)(*big.Int, *er.ResErr){
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/83db17056614454792c715a1c4083d8d")
	if err != nil {
		return nil, er.BadRequestError("Cannot Connect of client..")

	}else{
		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
		if re.MatchString(addr){
			account := common.HexToAddress(addr)
			balance , err := client.BalanceAt(context.Background(), account, nil)
			if err != nil {
				return nil, er.BadRequestError("cannot find balance of the address")
			}
			return balance, nil
		}else {
			return nil, er.BadRequestError("in valid address cannot find...!!")
		}
	}

}