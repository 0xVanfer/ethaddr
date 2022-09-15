package ethaddr

import "github.com/0xVanfer/chainId"

// USDC.e for avalanche
var USDCeList = map[string]string{
	chainId.AvalancheChainName: "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664", // USDC.e
}

// USDT.e for avalanche
var USDTeList = map[string]string{
	chainId.AvalancheChainName: "0xc7198437980c041c805a1edcba50c1ce5db95118", // USDT.e
}

// DAI.e
var DAIeList = map[string]string{
	chainId.AvalancheChainName: "0xd586e7f844cea2f87f50152665bcbc2c279d8d70", // DAI.e
}

// Wrapped avalanche chain token: WAVAX.
var WAVAXList = map[string]string{
	chainId.EthereumChainName:  "0x85f138bfEE4ef8e540890CFb48F620571d67Eda3", // WAVAX
	chainId.AvalancheChainName: "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7", // WAVAX
}

// XAVA
var XAVAList = map[string]string{
	chainId.AvalancheChainName: "0xd1c3f94de7e5b45fa4edbba472491a9f4b166fc4", // XAVA
}

// BTC.b for avalanche
var BTCbList = map[string]string{
	chainId.AvalancheChainName: "0x152b9d0fdc40c096757f570a51e494bd4b943e50", // BTC.b
}

// YUSD
var YUSDList = map[string]string{
	chainId.AvalancheChainName: "0x111111111111ed1d73f860f57b2798b683f2d325", // YUSD
}
