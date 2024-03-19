package ethaddr

import "github.com/0xVanfer/chainId"

const ZEROAddress string = "0x0000000000000000000000000000000000000000"
const EtherAddress string = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

// USDC
//
// map[network] = address.
var USDCList = map[string]string{
	chainId.ArbitrumChainName:  "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8", // USDC, 0xff970a61a04b1ca14834a43f5de4533ebddb5cc8
	chainId.AvalancheChainName: "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E", // USDC, 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e
	chainId.BaseChainName:      "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913", // USDC, 0x833589fcd6edb6e08f4c7c32d4f71b54bda02913
	chainId.BNBSmartChainName:  "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d", // USDC, 0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d
	chainId.EthereumChainName:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC, 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
	chainId.OptimismChainName:  "0x7F5c764cBc14f9669B88837ca1490cCa17c31607", // USDC, 0x7f5c764cbc14f9669b88837ca1490cca17c31607
	chainId.PolygonChainName:   "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", // USDC, 0x2791bca1f2de4661ed88a30c99a7a9449aa84174
}

// USDT
//
// map[network] = address.
var USDTList = map[string]string{
	chainId.ArbitrumChainName:  "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9", // USDT, 0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9
	chainId.AvalancheChainName: "0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7", // USDt, 0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7
	chainId.BNBSmartChainName:  "0x55d398326f99059fF775485246999027B3197955", // USDT, 0x55d398326f99059ff775485246999027b3197955
	chainId.EthereumChainName:  "0xdAC17F958D2ee523a2206206994597C13D831ec7", // USDT, 0xdac17f958d2ee523a2206206994597c13d831ec7
	chainId.OptimismChainName:  "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58", // USDT, 0x94b008aa00579c1307b0ef2c499ad98a8ce58e58
	chainId.PolygonChainName:   "0xc2132D05D31c914a87C6611C10748AEb04B58e8F", // USDT, 0xc2132d05d31c914a87c6611c10748aeb04b58e8f
}

// Wrapped eth: WETH.
// Including avalanche WETH.e.
//
// map[network] = address.
var WETHList = map[string]string{
	chainId.AvalancheChainName:    "0x49D5c2BdFfac6CE2BFdB6640F4F80f226bc10bAB", // WETH.e, 0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab
	chainId.ArbitrumChainName:     "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1", // WETH, 0x82af49447d8a07e3bd95bd0d56f35241523fbab1
	chainId.BaseChainName:         "0x4200000000000000000000000000000000000006", // WETH, 0x4200000000000000000000000000000000000006
	chainId.EthereumChainName:     "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH, 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
	chainId.OptimismChainName:     "0x4200000000000000000000000000000000000006", // WETH, 0x4200000000000000000000000000000000000006
	chainId.PolygonChainName:      "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619", // WETH, 0x7ceb23fd6bc0add59e62ac25578270cff1b9f619
	chainId.PolygonZkEVMChainName: "0x4F9A0e7FD2Bf6067db6994CF12E4495Df938E6e9", // WETH, 0x4f9a0e7fd2bf6067db6994cf12e4495df938e6e9
	chainId.ScrollChainName:       "0x5300000000000000000000000000000000000004", // WETH, 0x5300000000000000000000000000000000000004
}

// BNB chain ETH is BEP20 token.
//
// map[network] = address.
var ETHList = map[string]string{
	chainId.BNBSmartChainName: "0x2170Ed0880ac9A755fd29B2688956BD959F933F8", // ETH, 0x2170ed0880ac9a755fd29b2688956bd959f933f8
}

// Wrapped btc: WBTC.
// Including avalanche WBTC.e.
//
// map[network] = address.
var WBTCList = map[string]string{
	chainId.ArbitrumChainName:  "0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f", // WBTC, 0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f
	chainId.AvalancheChainName: "0x50b7545627a5162F82A992c33b87aDc75187B218", // WBTC.e, 0x50b7545627a5162f82a992c33b87adc75187b218
	chainId.EthereumChainName:  "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", // WBTC, 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599
	chainId.OptimismChainName:  "0x68f180fcCe6836688e9084f035309E29Bf0A2095", // WBTC, 0x68f180fcce6836688e9084f035309e29bf0a2095
	chainId.PolygonChainName:   "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6", // WBTC, 0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6
}
