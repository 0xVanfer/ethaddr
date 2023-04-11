package ethaddr

import "github.com/0xVanfer/chainId"

// Rocket pool rETH.
//
// map[network] = address.
var RETHList = map[string]string{
	chainId.EthereumChainName: "0xae78736cd615f374d3085123a210448e74fc6393",
	chainId.ArbitrumChainName: "0xae78736cd615f374d3085123a210448e74fc6393",
}
