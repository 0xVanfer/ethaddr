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
