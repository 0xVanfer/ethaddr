package ethaddr

import "github.com/0xVanfer/chainId"

// Chainlink token: LINK.
//
// map[network] = address.
var ChainlinkTokenList = map[string]string{
	chainId.ArbitrumChainName:  "0xf97f4df75117a78c1a5a0dbb814af92458539fb4", // LINK
	chainId.AvalancheChainName: "0x5947bb275c521040051d82396192181b413227a3", // LINK.e
	chainId.BNBSmartChainName:  "0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd", // LINK
	chainId.EthereumChainName:  "0x514910771AF9Ca656af840dff83E8264EcF986CA", // LINK, 0x514910771af9ca656af840dff83e8264ecf986ca
	chainId.OptimismChainName:  "0x350a791bfc2c21f9ed5d10980dad2e2638ffa7f6", // LINK
	chainId.PolygonChainName:   "0x53e0bca35ec356bd5dddfebbd1fc0fd03fabad39", // LINK
}

// Same as ChainlinkTokenList.
var LINKList = ChainlinkTokenList
