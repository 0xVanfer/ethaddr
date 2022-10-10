package ethaddr

import "github.com/0xVanfer/chainId"

// Chainlink token: LINK.
//
// map[network] = address.
var LINKList = map[string]string{
	chainId.EthereumChainName:  "0x514910771af9ca656af840dff83e8264ecf986ca", // LINK
	chainId.AvalancheChainName: "0x5947bb275c521040051d82396192181b413227a3", // LINK.e
	chainId.OptimismChainName:  "0x350a791bfc2c21f9ed5d10980dad2e2638ffa7f6", // LINK
	chainId.PolygonChainName:   "0x53e0bca35ec356bd5dddfebbd1fc0fd03fabad39", // LINK
}
