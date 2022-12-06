package ethaddr

import "github.com/0xVanfer/chainId"

// 1 Inch token: 1INCH.
//
// map[network] = address.
var OneInchTokenList = map[string]string{
	chainId.EthereumChainName: "0x111111111117dc0aa78b770fa6a738034120c302", // 1INCH
	chainId.PolygonChainName:  "0x9c2c5fd7b07e95ee044ddeba0e97a665f142394f", // 1INCH
}

// 1 Inch router.
//
// map[network] = address
var OneInchRouterList = map[string]string{
	chainId.PolygonChainName: "0x1111111254fb6c44bAC0beD2854e76F90643097d",
}

// 1 Inch spot aggregator.
//
// map[network] = address
var OneInchSpotPriceAggregatorList = map[string]string{
	chainId.EthereumChainName:     "0x07d91f5fb9bf7798734c3f606db065549f6893bb",
	chainId.AvalancheChainName:    "0xbd0c7aaf0bf082712ebe919a9dd94b2d978f79a9",
	chainId.ArbitrumChainName:     "0x735247fb0a604c0adC6cab38ACE16D0DbA31295F",
	chainId.OptimismChainName:     "0x11DEE30E710B8d4a8630392781Cc3c0046365d4c",
	chainId.PolygonChainName:      "0x7F069df72b7A39bCE9806e3AfaF579E54D8CF2b9",
	chainId.BinanceSmartChainName: "0xfbD61B037C325b959c0F6A7e69D8f37770C2c550",
	chainId.XDaiChainName:         "0x142DB045195CEcaBe415161e1dF1CF0337A4d02E",
}
