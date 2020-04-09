package main

import (
	"fmt"
	wiccwallet "github.com/WaykiChain/wicc-wallet-utils-go"
	"github.com/WaykiChain/wicc-wallet-utils-go/commons"
	hdwallet "go-hdwallet"
	"io/ioutil"
	"testing"
)

/*
 * 多币种转账交易 ,支持多种币种转账
 * Test nUniversal Coin Transfer Tx
 * fee Minimum 100000 sawi
 * */
func TestSignUCoinTransferTx(t *testing.T) {
	privateKey := "Y54BpYARqtZNNqYT5Zf82B6DWMrof3sk7qPFeMBvrdSzf4FK14jz"
	pubKey, _ := wiccwallet.GetPubKeyFromPrivateKey(privateKey)
	var txParam wiccwallet.UCoinTransferTxParam
	txParam.FeeSymbol = string(commons.WICC) //Fee Type (WICC/WUSD)
	txParam.Fees = 1000000                   // fees for mining
	txParam.ValidHeight = 4385931            // valid height Within the height of the latest block
	txParam.SrcRegId = ""
	txParam.Dests = wiccwallet.NewDestArr()
	dest1 := wiccwallet.Dest{
		string(commons.WICC),                 // From Coin Type
		2000000,                              // the values send to the contract app
		"wQPA4YMLGAKm3uxyzL1iBViUM2tazbqA4b", // To address
	}
	txParam.Dests.Add(&dest1)
	txParam.PubKey = pubKey
	txParam.Memo = ""

	hash, err := hdwallet.SignUCoinTransferTx(privateKey, &txParam)
	if err != nil {
		fmt.Println("UCoinTransferTx err: ", err)
	}
	println(hash)
}

/*
 * 多币种合约调用交易 ,支持多种币种转账
 * Test nUniversal Coin Contract Tx
 * fee Minimum 1000000 sawi
 * */
func TestSignUCoinCallContractTx(t *testing.T) {
	privateKey := "Y54BpYARqtZNNqYT5Zf82B6DWMrof3sk7qPFeMBvrdSzf4FK14jz"

	pubKey, _ := wiccwallet.GetPubKeyFromPrivateKey(privateKey)
	var txParam wiccwallet.UCoinContractTxParam
	txParam.FeeSymbol = string(commons.WICC)  //Fee Type (WICC/WUSD)
	txParam.CoinSymbol = string(commons.WICC) //From Coin Type
	txParam.CoinAmount = 1000000              // the values send to the contract app
	txParam.Fees = 1000000                    // fees for mining
	txParam.ValidHeight = 4508191             // valid height Within the height of the latest block
	txParam.SrcRegId = "4385938-2"            // the reg id of the caller
	txParam.AppId = "4503374-2"               // the reg id of the contract app
	txParam.PubKey = pubKey
	txParam.ContractHex = "4eb744000200" // the command of contract, hex format
	hash, err := hdwallet.SignUCoinCallContractTx(privateKey, &txParam)
	if err != nil {
		t.Error("SignUCoinCallContractTx err: ", err)
	}
	println(hash)
}

/*
   部署合约
   fee Minimum :110000000 sawi
*/
func TestSignDeployContractTx(t *testing.T) {

	privateKey := "Y54BpYARqtZNNqYT5Zf82B6DWMrof3sk7qPFeMBvrdSzf4FK14jz" //"Y9sx4Y8sBAbWDAqAWytYuUnJige3ZPwKDZp1SCDqqRby1YMgRG9c"

	script, err := ioutil.ReadFile("./data/hello.lua")
	if err != nil {
		t.Error("Read contract script file err: ", err)
	}

	var txParam wiccwallet.RegisterContractTxParam
	txParam.ValidHeight = 4503364  //20999
	txParam.SrcRegId = "4385938-2" //"7849-1"
	txParam.Fees = 110000000
	txParam.Script = script
	txParam.Description = "My hello contract!!!"
	hash, err := hdwallet.SignRegisterContractTx(privateKey, &txParam)
	if err != nil {
		t.Error("SignRegisterContractTx err: ", err)
	}
	println(hash)

}
