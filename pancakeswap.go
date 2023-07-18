package ethaddr

import "github.com/0xVanfer/chainId"

const PancakeswapProtocol string = "pancakeswap"

// Pancakeswap token: CAKE.
//
// map[network] = address.
var PancakeswapTokenList = map[string]string{
	chainId.BNBSmartChainName: "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82", // Cake
}

// Same as PancakeswapTokenList.
var CakeList = PancakeswapTokenList

// Pancakeswap factory.
//
// map[network] = address.
var PancakeswapFactoryList = map[string]string{
	chainId.BNBSmartChainName: "0xca143ce32fe78f1f7019d7d551a6402fc5350c73",
}

// Pancakeswap router02.
//
// map[network] = address.
var PancakeswapRouter02List = map[string]string{
	chainId.BNBSmartChainName: "0x10ed43c718714eb63d5aa57b78b54704e256024e",
}
