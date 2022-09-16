package ethaddr

import "github.com/0xVanfer/chainId"

const ConvexProtocol string = "convex"

// Convex token: CVX.
// map[network] = address.
var ConvexTokenList = map[string]string{
	chainId.EthereumChainName: "0x4e3fbd56cd56c3e72c1403e103b45db9da5b9d2b", // CVX
}
