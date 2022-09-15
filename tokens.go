package ethaddr

import "github.com/0xVanfer/chainId"

const ZEROAddress string = "0x0000000000000000000000000000000000000000"

// USDC
var USDCList = map[string]string{
	chainId.EthereumChainName:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // USDC
	chainId.AvalancheChainName: "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e", // USDC
	chainId.OptimismChainName:  "0x7f5c764cbc14f9669b88837ca1490cca17c31607", // USDC
}

// USDT
var USDTList = map[string]string{
	chainId.EthereumChainName:  "0xdac17f958d2ee523a2206206994597c13d831ec7", // USDT
	chainId.AvalancheChainName: "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7", // USDt
	chainId.OptimismChainName:  "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58", // USDT
}

// DAI
var DAIList = map[string]string{
	chainId.EthereumChainName: "0x6b175474e89094c44da98b954eedeac495271d0f", // DAI
	chainId.OptimismChainName: "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1", // DAI
}

// UST
var USTList = map[string]string{
	chainId.EthereumChainName:  "0xa693b19d2931d498c5b318df961919bb4aee87a5", // UST
	chainId.AvalancheChainName: "0x260bbf5698121eb85e7a74f2e45e16ce762ebe11", // UST
}

// WETH
var WETHList = map[string]string{
	chainId.EthereumChainName:  "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", // WETH
	chainId.AvalancheChainName: "0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab", // WETH.e
	chainId.OptimismChainName:  "0x4200000000000000000000000000000000000006", // WETH
}

// WBTC
var WBTCList = map[string]string{
	chainId.EthereumChainName:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", // WBTC
	chainId.AvalancheChainName: "0x50b7545627a5162f82a992c33b87adc75187b218", // WBTC.e
	chainId.OptimismChainName:  "0x68f180fcce6836688e9084f035309e29bf0a2095", // WBTC
}
