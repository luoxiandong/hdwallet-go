package hdwallet

import (
	"fmt"
	wiccwallet "github.com/WaykiChain/wicc-wallet-utils-go"
)

func init() {
	coins[WICC] = newWICC
}

type wicc struct {
	*btc
}

func newWICC(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "WaykiChain"
	token.symbol = "WICC"

	return &wicc{btc: token}
}

/*
 * 多币种转账交易 ,支持多种币种转账
 * Test nUniversal Coin Transfer Tx
 * fee Minimum 100000 sawi
 * */
func SignUCoinTransferTx(privateKey string, txParam *wiccwallet.UCoinTransferTxParam) (string, error) {
	hash, err := wiccwallet.SignUCoinTransferTx(privateKey, txParam)
	if err != nil {
		return "", err
	}

	return hash, nil
}

/*
 * 多币种合约调用交易 ,支持多种币种转账
 * Test nUniversal Coin Contract Tx
 * fee Minimum 1000000 sawi
 * */
func SignUCoinCallContractTx(privateKey string, txParam *wiccwallet.UCoinContractTxParam) (string, error) {
	hash, err := wiccwallet.SignUCoinCallContractTx(privateKey, txParam)
	if err != nil {
		fmt.Println("SignUCoinCallContractTx err: ", err)
		return "", err
	}
	//fmt.Println("SignUCoinCallContractTx hash: ", hash)

	return hash, nil
}

/*
 *	多币种部署合约,推荐使用
 *	fee Minimum: 110000000 sawi
 */
func SignUCoinDeployContractTx(privateKey string, txParam *wiccwallet.UCoinRegisterContractTxParam) (string, error) {
	hash, err := wiccwallet.SignUCoinRegisterContractTx(privateKey, txParam)
	if err != nil {
		fmt.Println("SignUCoinDeployContractTx err: ", err)
		return "", err
	}
	//fmt.Println("SignUCoinDeployContractTx hash: ", hash)

	return hash, nil
}
