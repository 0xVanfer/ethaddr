package ethaddr

import "github.com/0xVanfer/chainId"

const OneInchProtocol string = "1inch"

// 1 Inch token: 1INCH.
//
// map[network] = address.
var OneInchTokenList = map[string]string{
	chainId.EthereumChainName: "0x111111111117dC0aa78b770fA6A738034120C302", // 1INCH, 0x111111111117dc0aa78b770fa6a738034120c302
	chainId.PolygonChainName:  "0x9c2C5fd7b07E95EE044DDeba0E97a665F142394f", // 1INCH, 0x9c2c5fd7b07e95ee044ddeba0e97a665f142394f
}

// 1 Inch token: 1INCH.
var OINCHLIST = OneInchTokenList

// 1 Inch router V4.
//
// map[network] = address
var OneInchRouterListV4 = map[string]string{
	chainId.ArbitrumChainName:  "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.AvalancheChainName: "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.BNBSmartChainName:  "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.EthereumChainName:  "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.FantomChainName:    "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.OptimismChainName:  "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
	chainId.PolygonChainName:   "0x1111111254fb6c44bAC0beD2854e76F90643097d", // 0x1111111254fb6c44bac0bed2854e76f90643097d
}

// 1 Inch router V5.
//
// map[network] = address
var OneInchRouterListV5 = map[string]string{
	chainId.ArbitrumChainName:  "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.AvalancheChainName: "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.BNBSmartChainName:  "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.EthereumChainName:  "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.FantomChainName:    "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.OptimismChainName:  "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
	chainId.PolygonChainName:   "0x1111111254EEB25477B68fb85Ed929f73A960582", // 0x1111111254eeb25477b68fb85ed929f73a960582
}

// 1 Inch spot aggregator.
//
// map[network] = address
var OneInchSpotPriceAggregatorList = map[string]string{
	chainId.ArbitrumChainName:  "0x735247fb0a604c0adC6cab38ACE16D0DbA31295F", // 0x735247fb0a604c0adc6cab38ace16d0dba31295f
	chainId.AvalancheChainName: "0xBd0c7AaF0bF082712EbE919a9dD94b2d978f79A9", // 0xbd0c7aaf0bf082712ebe919a9dd94b2d978f79a9
	chainId.BNBSmartChainName:  "0xfbD61B037C325b959c0F6A7e69D8f37770C2c550", // 0xfbd61b037c325b959c0f6a7e69d8f37770c2c550
	chainId.EthereumChainName:  "0x07D91f5fb9Bf7798734C3f606dB065549F6893bb", // 0x07d91f5fb9bf7798734c3f606db065549f6893bb
	chainId.GnosisChainName:    "0x142DB045195CEcaBe415161e1dF1CF0337A4d02E", // 0x142db045195cecabe415161e1df1cf0337a4d02e
	chainId.OptimismChainName:  "0x11DEE30E710B8d4a8630392781Cc3c0046365d4c", // 0x11dee30e710b8d4a8630392781cc3c0046365d4c
	chainId.PolygonChainName:   "0x7F069df72b7A39bCE9806e3AfaF579E54D8CF2b9", // 0x7f069df72b7a39bce9806e3afaf579e54d8cf2b9
}
