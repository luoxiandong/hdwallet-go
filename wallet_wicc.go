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

	return hash, nil
}

/*
 * 部署合约
 * ee Minimum: 110000000 sawi
 */
func SignRegisterContractTx(privateKey string, txParam *wiccwallet.RegisterContractTxParam) (string, error) {
	hash, err := wiccwallet.SignRegisterContractTx(privateKey, txParam)
	if err != nil {
		fmt.Println("SignRegisterContractTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
Asset Create 创建资产

fee Minimum 0.01WICC
make sure account balance >550.01 wicc
发布资产需要扣除550wicc的发布费用
*/
func SignAssetCreateTx(privateKey string, param *wiccwallet.AssetIssueTxParam) (string, error) {
	hash, err := wiccwallet.SignAssetCreateTx(privateKey, param)
	if err != nil {
		fmt.Println("SignAssetIssueTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
Asset Update 更新资产
fee Minimum 0.01WICC
make sure account balance >110.01 wicc
发布资产需要扣除110wicc的发布费用
*/
func SignAssetUpdateTx(privateKey string, param *wiccwallet.AssetUpdateTxParam) (string, error) {
	hash, err := wiccwallet.SignAssetUpdateTx(privateKey, param)
	if err != nil {
		fmt.Println("AssetUpdateTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
* 投票交易
* Voting transaction
* fee Minimum 0.01 wicc
* */
func SignDelegateTx(privateKey string, param *wiccwallet.DelegateTxParam) (string, error) {
	hash, err := wiccwallet.SignDelegateTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDelegateTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 * 创建,追加cdp交易
 * Create or append an  cdp transaction
 * fee Minimum 0.01 wicc
 * */
func SignCdpStakeTx(privateKey string, param *wiccwallet.CdpStakeTxParam) (string, error) {
	hash, err := wiccwallet.SignCdpStakeTx(privateKey, param)
	if err != nil {
		fmt.Println("SignCdpStakeTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 * 赎回cdp交易
 * Redeem cdp transaction
 * fee Minimum 0.01 wicc
 * */
func SignCdpRedeemTx(privateKey string, param *wiccwallet.CdpRedeemTxParam) (string, error) {
	hash, err := wiccwallet.SignCdpRedeemTx(privateKey, param)
	if err != nil {
		fmt.Println("SignCdpRedeemTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 * 清算cdp交易
 * Liquidate cdp transaction
 * fee Minimum 0.01 wicc
 * */
func SignCdpLiquidateTx(privateKey string, param *wiccwallet.CdpLiquidateTxParam) (string, error) {
	hash, err := wiccwallet.SignCdpLiquidateTx(privateKey, param)
	if err != nil {
		fmt.Println("SignCdpLiquidateTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 * Dex 限价买单交易
 * Dex limit price transaction
 * fee Minimum 0.001 wicc
 * */
func SignDexBuyLimitTx(privateKey string, param *wiccwallet.DexLimitTxParam) (string, error) {
	hash, err := wiccwallet.SignDexBuyLimitTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDexBuyLimitTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 * Dex 限价卖单交易
 * Dex limit sell price transaction
 * fee Minimum 0.001 wicc
* */
func SignDexSellLimitTx(privateKey string, param *wiccwallet.DexLimitTxParam) (string, error) {
	hash, err := wiccwallet.SignDexSellLimitTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDexSellLimitTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 *  Dex 市价卖单交易
 * Dex market sell price transaction
 * fee Minimum 0.001 wicc
* */
func SignDexMarketSellTx(privateKey string, param *wiccwallet.DexMarketTxParam) (string, error) {
	hash, err := wiccwallet.SignDexMarketSellTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDexSellLimitTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 *  Dex 市价买单交易
 * Dex market buy price transaction
 * fee Minimum 0.001 wicc
* */
func SignDexMarketBuyTx(privateKey string, param *wiccwallet.DexMarketTxParam) (string, error) {
	hash, err := wiccwallet.SignDexMarketBuyTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDexMarketBuyTx err: ", err)
		return "", err
	}

	return hash, nil
}

/*
 *  Dex 取消挂单交易
 * Dex cancel order tx
 * fee Minimum 0.001 wicc
* */
func SignDexCancelTx(privateKey string, param *wiccwallet.DexCancelTxParam) (string, error) {
	hash, err := wiccwallet.SignDexCancelTx(privateKey, param)
	if err != nil {
		fmt.Println("SignDexCancelTx err: ", err)
		return "", err
	}

	return hash, nil
}
