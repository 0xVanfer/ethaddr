package ethaddr

import "github.com/0xVanfer/chainId"

const ZEROAddress string = "0x0000000000000000000000000000000000000000"
const EtherAddress string = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

// USDC
//
// map[network] = address.
var USDCList = map[string]string{
	chainId.EthereumChainName:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // USDC
	chainId.AvalancheChainName: "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e", // USDC
	chainId.OptimismChainName:  "0x7f5c764cbc14f9669b88837ca1490cca17c31607", // USDC
	chainId.PolygonChainName:   "0x2791bca1f2de4661ed88a30c99a7a9449aa84174", // USDC
	chainId.ArbitrumChainName:  "0xff970a61a04b1ca14834a43f5de4533ebddb5cc8", // USDC
	chainId.BNBSmartChainName:  "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d", // USDC
}

// USDT
//
// map[network] = address.
var USDTList = map[string]string{
	chainId.EthereumChainName:  "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT
	chainId.AvalancheChainName: "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7", // USDt
	chainId.OptimismChainName:  "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58", // USDT
	chainId.PolygonChainName:   "0xc2132d05d31c914a87c6611c10748aeb04b58e8f", // USDT
	chainId.ArbitrumChainName:  "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9", // USDT
	chainId.BNBSmartChainName:  "0x55d398326f99059ff775485246999027b3197955", // USDT
}

// DAI
//
// map[network] = address.
var DAIList = map[string]string{
	chainId.EthereumChainName: "0x6b175474e89094c44da98b954eedeac495271d0f", // DAI
	chainId.OptimismChainName: "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1", // DAI
	chainId.PolygonChainName:  "0x8f3cf7ad23cd3cadbd9735aff958023239c6a063", // DAI
	chainId.ArbitrumChainName: "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1", // DAI
	chainId.BNBSmartChainName: "0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3", // DAI
}

// UST
//
// map[network] = address.
var USTList = map[string]string{
	chainId.EthereumChainName:  "0xa693b19d2931d498c5b318df961919bb4aee87a5", // UST
	chainId.AvalancheChainName: "0x260bbf5698121eb85e7a74f2e45e16ce762ebe11", // UST
	chainId.BNBSmartChainName:  "0x3d4350cd54aef9f9b2c29435e0fa809957b3f30a", // UST
}

// Wrapped eth: WETH.
// Including avalanche WETH.e.
//
// map[network] = address.
var WETHList = map[string]string{
	chainId.AvalancheChainName:    "0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab", // WETH.e
	chainId.ArbitrumChainName:     "0x82af49447d8a07e3bd95bd0d56f35241523fbab1", // WETH
	chainId.BaseChainName:         "0x4200000000000000000000000000000000000006", // WETH
	chainId.EthereumChainName:     "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", // WETH
	chainId.OptimismChainName:     "0x4200000000000000000000000000000000000006", // WETH
	chainId.PolygonChainName:      "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619", // WETH
	chainId.PolygonZkEVMChainName: "0x4F9A0e7FD2Bf6067db6994CF12E4495Df938E6e9", // WETH
	chainId.ScrollChainName:       "0x5300000000000000000000000000000000000004", // WETH
}

// BNB chain ETH is BEP20 token.
//
// map[network] = address.
var ETHList = map[string]string{
	chainId.BNBSmartChainName: "0x2170ed0880ac9a755fd29b2688956bd959f933f8", // ETH
}

// Wrapped btc: WBTC.
// Including avalanche WBTC.e.
//
// map[network] = address.
var WBTCList = map[string]string{
	chainId.EthereumChainName:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", // WBTC
	chainId.AvalancheChainName: "0x50b7545627a5162f82a992c33b87adc75187b218", // WBTC.e
	chainId.OptimismChainName:  "0x68f180fcce6836688e9084f035309e29bf0a2095", // WBTC
	chainId.PolygonChainName:   "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6", // WBTC
	chainId.ArbitrumChainName:  "0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f", // WBTC
}
