package ethaddr

import "github.com/0xVanfer/chainId"

// Enjin token: ENJ.
//
// map[network] = address.
var EnjinTokenList = map[string]string{
	chainId.EthereumChainName: "0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c", // ENJ, 0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c
}

// Same as EnjinTokenList.
var ENJList = EnjinTokenList
