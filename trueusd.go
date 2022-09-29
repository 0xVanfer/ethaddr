package ethaddr

import "github.com/0xVanfer/chainId"

// Trueusd token: TUSD.
//
// map[network] = address.
var TUSDList = map[string]string{
	chainId.EthereumChainName:  "0x0000000000085d4780b73119b644ae5ecd22b376", // TUSD
	chainId.AvalancheChainName: "0x1c20e891bab6b1727d14da358fae2984ed9b59eb", // TUSD
}
