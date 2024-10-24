package ethaddr

// Uniswap token: UNI.
//
// map[network] = address.
var UniswapTokenList = map[int64]string{
	ChainEthereum:  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", // UNI, 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984
	ChainArbitrum:  "0xFa7F8980b0f1E64A2062791cc3b0871572f1F7f0", // UNI, 0xfa7f8980b0f1e64a2062791cc3b0871572f1f7f0
	ChainPolygon:   "0xb33EaAd8d922B1083446DC23f610c2567fB5180f", // UNI, 0xb33eaad8d922b1083446dc23f610c2567fb5180f
	ChainAvalanche: "0x8eBAf22B6F053dFFeaf46f4Dd9eFA95D89ba8580", // UNI, 0x8ebaf22b6f053dffeaf46f4dd9efa95d89ba8580
}

// Same as UniswapTokenList.
var UNIList = UniswapTokenList
