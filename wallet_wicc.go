package hdwallet

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
