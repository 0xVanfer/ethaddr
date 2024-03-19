package ethaddr

import "github.com/0xVanfer/chainId"

// Trueusd token: TUSD.
//
// map[network] = address.
var TUSDList = map[string]string{
	chainId.ArbitrumChainName:  "0x4D15a3A2286D883AF0AA1B3f21367843FAc63E07", // TUSD, 0x4d15a3a2286d883af0aa1b3f21367843fac63e07
	chainId.AvalancheChainName: "0x1C20E891Bab6b1727d14Da358FAe2984Ed9B59EB", // TUSD, 0x1c20e891bab6b1727d14da358fae2984ed9b59eb
	chainId.BNBSmartChainName:  "0x40af3827F39D0EAcBF4A168f8D4ee67c121D11c9", // TUSD, 0x40af3827f39d0eacbf4a168f8d4ee67c121d11c9
	chainId.EthereumChainName:  "0x0000000000085d4780B73119b644AE5ecd22b376", // TUSD, 0x0000000000085d4780b73119b644ae5ecd22b376
}

// Old Trueusd token: TUSD.
//
// map[network] = address.
var TUSDOldList = map[string]string{
	chainId.BNBSmartChainName: "0x14016E85a25aeb13065688cAFB43044C2ef86784", // TUSDOLD, 0x14016e85a25aeb13065688cafb43044c2ef86784
}
