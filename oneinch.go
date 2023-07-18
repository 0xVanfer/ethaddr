package ethaddr

import "github.com/0xVanfer/chainId"

// 1 Inch token: 1INCH.
//
// map[network] = address.
var OneInchTokenList = map[string]string{
	chainId.EthereumChainName: "0x111111111117dc0aa78b770fa6a738034120c302", // 1INCH
	chainId.PolygonChainName:  "0x9c2c5fd7b07e95ee044ddeba0e97a665f142394f", // 1INCH
}

// Deprecated: Use OneInchRouterListV4 instead.
var OneInchRouterList = OneInchRouterListV4

// 1 Inch router V4.
//
// map[network] = address
var OneInchRouterListV4 = map[string]string{
	chainId.EthereumChainName:  "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.PolygonChainName:   "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.ArbitrumChainName:  "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.AvalancheChainName: "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.BNBSmartChainName:  "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.FantomChainName:    "0x1111111254fb6c44bac0bed2854e76f90643097d",
	chainId.OptimismChainName:  "0x1111111254fb6c44bac0bed2854e76f90643097d",
}

// 1 Inch router V5.
//
// map[network] = address
var OneInchRouterListV5 = map[string]string{
	chainId.EthereumChainName:  "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.PolygonChainName:   "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.ArbitrumChainName:  "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.AvalancheChainName: "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.BNBSmartChainName:  "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.FantomChainName:    "0x1111111254eeb25477b68fb85ed929f73a960582",
	chainId.OptimismChainName:  "0x1111111254eeb25477b68fb85ed929f73a960582",
}

// 1 Inch spot aggregator.
//
// map[network] = address
var OneInchSpotPriceAggregatorList = map[string]string{
	chainId.EthereumChainName:  "0x07d91f5fb9bf7798734c3f606db065549f6893bb",
	chainId.AvalancheChainName: "0xbd0c7aaf0bf082712ebe919a9dd94b2d978f79a9",
	chainId.ArbitrumChainName:  "0x735247fb0a604c0adC6cab38ACE16D0DbA31295F",
	chainId.OptimismChainName:  "0x11DEE30E710B8d4a8630392781Cc3c0046365d4c",
	chainId.PolygonChainName:   "0x7F069df72b7A39bCE9806e3AfaF579E54D8CF2b9",
	chainId.BNBSmartChainName:  "0xfbD61B037C325b959c0F6A7e69D8f37770C2c550",
	chainId.XDaiChainName:      "0x142DB045195CEcaBe415161e1dF1CF0337A4d02E",
}
