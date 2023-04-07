package ethaddr

import "github.com/0xVanfer/chainId"

// USDC.e for avalanche
//
// map[network] = address.
var USDCeList = map[string]string{
	chainId.AvalancheChainName: "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664", // USDC.e
}

// USDT.e for avalanche
//
// map[network] = address.
var USDTeList = map[string]string{
	chainId.AvalancheChainName: "0xc7198437980c041c805a1edcba50c1ce5db95118", // USDT.e
}

// DAI.e
//
// map[network] = address.
var DAIeList = map[string]string{
	chainId.AvalancheChainName: "0xd586e7f844cea2f87f50152665bcbc2c279d8d70", // DAI.e
}

// Wrapped avalanche chain token: WAVAX.
//
// map[network] = address.
var WAVAXList = map[string]string{
	chainId.EthereumChainName:  "0x85f138bfee4ef8e540890cfb48f620571d67eda3", // WAVAX(wormwhole)
	chainId.AvalancheChainName: "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7", // WAVAX
	chainId.PolygonChainName:   "0x7bb11e7f8b10e9e571e5d8eace04735fdfb2358a", // WAVAX(wormwhole)
}

// AVAX on Polygon.
//
// map[network] = address.
var AVAXList = map[string]string{
	chainId.PolygonChainName: "0x2c89bbc92bd86f8075d1decc58c7f4e0107f286b", // AVAX
}

// XAVA
//
// map[network] = address.
var XAVAList = map[string]string{
	chainId.AvalancheChainName: "0xd1c3f94de7e5b45fa4edbba472491a9f4b166fc4", // XAVA
}

// YUSD
//
// map[network] = address.
var YUSDList = map[string]string{
	chainId.AvalancheChainName: "0x111111111111ed1d73f860f57b2798b683f2d325", // YUSD
}
