package main

import (
	wiccwallet "github.com/WaykiChain/wicc-wallet-utils-go"
	"github.com/WaykiChain/wicc-wallet-utils-go/commons"
	hdwallet "go-hdwallet"
	"testing"
)

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

	hash, _ := hdwallet.SignUCoinTransferTx(privateKey, &txParam)
	println(hash)
}
