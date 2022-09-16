package ethaddr

import "github.com/0xVanfer/chainId"

const BalancerProtocol string = "balancer"

// Balancer token: BAL.
// map[network] = address.
var BalancerTokenList = map[string]string{
	chainId.EthereumChainName: "0xba100000625a3754423978a60c9317c58a424e3d", //BAL
}
