package ethaddr

import "github.com/0xVanfer/chainId"

const PancakeswapProtocol string = "pancakeswap"

// Pancakeswap token: CAKE.
//
// map[network] = address.
var PancakeswapTokenList = map[string]string{
	chainId.BinanceSmartChainName: "0x60781c2586d68229fde47564546784ab3faca982", // CAKE
}

// Same as PancakeswapTokenList.
var CAKEList = PancakeswapTokenList

// Pancakeswap factory.
//
// map[network] = address.
var PancakeswapFactoryList = map[string]string{
	chainId.BinanceSmartChainName: "0xca143ce32fe78f1f7019d7d551a6402fc5350c73",
}

// Pancakeswap router02.
//
// map[network] = address.
var PancakeswapRouter02List = map[string]string{
	chainId.BinanceSmartChainName: "0x10ed43c718714eb63d5aa57b78b54704e256024e",
}
