package ethaddr

import "github.com/0xVanfer/chainId"

// Aave token: AAVE.
//
// map[network] = address.
var AaveTokenList = map[string]string{
	chainId.EthereumChainName:  "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9", // AAVE
	chainId.AvalancheChainName: "0x63a72806098bd3d9520cc43356dd78afe5d386d9", // AAVE.e
	chainId.OptimismChainName:  "0x76fb31fb4af56892a25e32cfc43de717950c9278", // AAVE
	chainId.PolygonChainName:   "0xd6df932a45c0f255f85145f286ea0b292b21c90b", // AAVE
	chainId.ArbitrumChainName:  "0xba5ddd1f9d7f570dc94a51479a000e3bce967196", // AAVE
}

// Same as AaveTokenList.
var AAVEList = AaveTokenList
