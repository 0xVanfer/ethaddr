package ethaddr

import "github.com/0xVanfer/chainId"

// Enjin token: ENJ.
//
// map[network] = address.
var EnjinTokenList = map[string]string{
	chainId.EthereumChainName: "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c", // ENJ
}

// Same as EnjinTokenList.
var ENJList = EnjinTokenList
