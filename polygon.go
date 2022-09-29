package ethaddr

import "github.com/0xVanfer/chainId"

// Wrapped polygon(matic) chain token: WMATIC.
//
// map[network] = address.
var WMATICList = map[string]string{
	chainId.EthereumChainName: "0x7c9f4c87d911613fe9ca58b579f737911aad2d43", // WMATIC
	chainId.PolygonChainName:  "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270", // WMATIC
}
