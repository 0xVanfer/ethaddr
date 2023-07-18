package ethaddr

import "github.com/0xVanfer/chainId"

// BTC.b or BTCB.
//
// map[network] = address.
var BTCbList = map[string]string{
	chainId.AvalancheChainName: "0x152b9d0fdc40c096757f570a51e494bd4b943e50", // BTC.b
	chainId.ArbitrumChainName:  "0x2297aebd383787a160dd0d9f71508148769342e3", // BTC.b
	chainId.BNBSmartChainName:  "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c", // BTCB
}

// Bitcoin cash token.
//
// map[network] = address.
var BCHList = map[string]string{
	chainId.BNBSmartChainName: "0x8ff795a6f4d97e7887c79bea79aba5cc76444adf", // BCH
}
