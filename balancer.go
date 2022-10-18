package ethaddr

import "github.com/0xVanfer/chainId"

const BalancerProtocol string = "balancer"

// Balancer token: BAL.
//
// map[network] = address.
var BalancerTokenList = map[string]string{
	chainId.EthereumChainName: "0xba100000625a3754423978a60c9317c58a424e3d", // BAL
	chainId.PolygonChainName:  "0x9a71012b13ca4d3d0cdc72a177df3ef03b0e76a3", // BAL
}

// Same as BalancerTokenList.
var BALList = BalancerTokenList

// Balancer v2 vault.
//
// map[network] = address.
var BalancerV2VaultList = map[string]string{
	chainId.PolygonChainName: "0xba12222222228d8ba445958a75a0704d566bf2c8",
}

// Balancer helper.
//
// map[network] = address.
var BalancherHelperList = map[string]string{
	chainId.PolygonChainName: "0x239e55f427d44c3cc793f49bfb507ebe76638a2b",
}
