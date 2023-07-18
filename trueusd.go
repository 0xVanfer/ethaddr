package ethaddr

import "github.com/0xVanfer/chainId"

// Trueusd token: TUSD.
//
// map[network] = address.
var TUSDList = map[string]string{
	chainId.EthereumChainName:  "0x0000000000085d4780b73119b644ae5ecd22b376", // TUSD
	chainId.AvalancheChainName: "0x1c20e891bab6b1727d14da358fae2984ed9b59eb", // TUSD
	chainId.ArbitrumChainName:  "0x4d15a3a2286d883af0aa1b3f21367843fac63e07", // TUSD
	chainId.BNBSmartChainName:  "0x40af3827f39d0eacbf4a168f8d4ee67c121d11c9", // TUSD
}

// Old Trueusd token: TUSD.
//
// map[network] = address.
var TUSDOldList = map[string]string{
	chainId.BNBSmartChainName: "0x14016e85a25aeb13065688cafb43044c2ef86784", // TUSDOLD
}
