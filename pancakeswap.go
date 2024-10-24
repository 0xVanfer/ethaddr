package ethaddr

// Website: https://pancakeswap.finance/
const PancakeswapProtocol string = "pancakeswap"

// Pancakeswap token: Cake. (symbol is in camel)
//
// map[network] = address.
var PancakeswapTokenList = map[int64]string{
	ChainBSC: "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82", // Cake, 0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82
}

// Same as PancakeswapTokenList.
var CakeList = PancakeswapTokenList

// Pancakeswap factory.
//
// map[network] = address.
var PancakeswapFactoryList = map[int64]string{
	ChainBSC: "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73", // 0xca143ce32fe78f1f7019d7d551a6402fc5350c73
}

// Pancakeswap router02.
//
// map[network] = address.
var PancakeswapRouter02List = map[int64]string{
	ChainBSC: "0x10ED43C718714eb63d5aA57B78B54704E256024E", // 0x10ed43c718714eb63d5aa57b78b54704e256024e
}
