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
	chainId.EthereumChainName:  "0x07d91f5fb9bf7798734c3f606db065549f6893bb",
	chainId.AvalancheChainName: "0xbd0c7aaf0bf082712ebe919a9dd94b2d978f79a9",
}
